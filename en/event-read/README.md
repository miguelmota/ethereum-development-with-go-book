---
description: Tutorial on how to read smart contract events with Go.
---

# Reading Event Logs

A smart contract may optionally emit "events" which get stored as logs as part of the transaction receipt. Reading these events is pretty simple. First we need to construct a filter query. We import the `FilterQuery` struct from the go-ethereum package and initialize it with filter options. We tell it the range of blocks that we want to filter through and specify the contract address to read these logs from. In this example we'll be reading all the logs from a particular block, from the smart contract we created in the [smart contract sections](../smart-contract-compile).

```go
query := ethereum.FilterQuery{
  FromBlock: big.NewInt(2394201),
  ToBlock:   big.NewInt(2394201),
  Addresses: []common.Address{
    contractAddress,
  },
}
```

The next is step is to call `FilterLogs` from the ethclient that takes in our query and will return all the matching event logs.

```go
logs, err := client.FilterLogs(context.Background(), query)
if err != nil {
  log.Fatal(err)
}
```

All the logs returned will be ABI encoded so by themselves they won't be very readable. In order to decode the logs we'll need to import our smart contract ABI. To do that, we import our compiled smart contract Go package which will contain an external property in the name format `<ContractName>ABI` containing our ABI. Afterwards we use the `abi.JSON` function from the go-ethereum `accounts/abi` go-ethereum package to return a parsed ABI interface that we can use in our Go application.

```go
contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
if err != nil {
  log.Fatal(err)
}
```

Now we can interate through the logs and decode them into a type we can use. If you recall the logs that our sample contract emitted were of type `bytes32` in Solidity, so the equivalent in Go would be `[32]byte`. We can create an anonymous struct with these types and pass a pointer as the first argument to the `Unpack` function of the parsed ABI interface to decode the raw log data. The second argument is the name of the event we're trying to decode and the last argument is the encoded log data.

```go
for _, vLog := range logs {
  event := struct {
    Key   [32]byte
    Value [32]byte
  }{}
  err := contractAbi.Unpack(&event, "ItemSet", vLog.Data)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(string(event.Key[:]))   // foo
  fmt.Println(string(event.Value[:])) // bar
}
```

Also, the log struct contains additional information such as the block hash, block number, and transaction hash.

```go
fmt.Println(vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
fmt.Println(vLog.BlockNumber)     // 2394201
fmt.Println(vLog.TxHash.Hex())    // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6
```

### Topics

If your solidity event contains `indexed` event types, then they become a *topic* rather than part of the data property of the log. In solidity you may only have up to 4 topics but only 3 indexed event types. The first topic is *always* the signature of the event. Our example contract didn't contain indexed events, but if it did this is how to read the event topics.

```go
var topics [4]string
for i := range vLog.Topics {
  topics[i] = vLog.Topics[i].Hex()
}

fmt.Println(topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
```

As you can see here the first topic is just the hashed event signature.

```go
eventSignature := []byte("ItemSet(bytes32,bytes32)")
hash := crypto.Keccak256Hash(eventSignature)
fmt.Println(hash.Hex()) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
```

That's all there is to reading and parsing logs. To learn how to subscribe to logs, read the [previous section](../event-subscribe).

---

### Full code

Commands

```bash
solc --abi Store.sol
solc --bin Store.sol
abigen --bin=Store_sol_Store.bin --abi=Store_sol_Store.abi --pkg=store --out=Store.go
```

[Store.sol](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/contracts/Store.sol)

```solidity
pragma solidity ^0.4.24;

contract Store {
  event ItemSet(bytes32 key, bytes32 value);

  string public version;
  mapping (bytes32 => bytes32) public items;

  constructor(string _version) public {
    version = _version;
  }

  function setItem(bytes32 key, bytes32 value) external {
    items[key] = value;
    emit ItemSet(key, value);
  }
}
```

[event_read.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/event_read.go)

```go
package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	store "./contracts" // for demo
)

func main() {
	client, err := ethclient.Dial("wss://rinkeby.infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(2394201),
		ToBlock:   big.NewInt(2394201),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		fmt.Println(vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
		fmt.Println(vLog.BlockNumber)     // 2394201
		fmt.Println(vLog.TxHash.Hex())    // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6

		event := struct {
			Key   [32]byte
			Value [32]byte
		}{}
		err := contractAbi.Unpack(&event, "ItemSet", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(event.Key[:]))   // foo
		fmt.Println(string(event.Value[:])) // bar

		var topics [4]string
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}

		fmt.Println(topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
	}

	eventSignature := []byte("ItemSet(bytes32,bytes32)")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Println(hash.Hex()) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
}
```

```bash
$ solc --version
0.4.24+commit.e67f0147.Emscripten.clang
```
