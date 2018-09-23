---
description: Tutorial on how to subscribe to smart contract events with Go.
---

# Subscribing to Event Logs

First thing we need to do in order to subscribe to event logs is dial to a websocket enabled Ethereum client. Fortunately for us, Infura supports websockets.

```go
client, err := ethclient.Dial("wss://rinkeby.infura.io/ws")
if err != nil {
  log.Fatal(err)
}
```

The next step is to create a filter query. In this example we'll be reading all events coming from the example contract that we've created in the previous lessons.

```go
contractAddress := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
query := ethereum.FilterQuery{
  Addresses: []common.Address{contractAddress},
}
```

The way we'll be receiving events is through a Go channel. Let's create one with type of `Log` from the go-ethereum `core/types` package.

```go
logs := make(chan types.Log)
```

Now all we have to do is subscribe by calling `SubscribeFilterLogs` from the client, which takes in the query options and the output channel. This will return a subscription struct containing unsubscribe and error methods.

```go
sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
if err != nil {
  log.Fatal(err)
}
```

Finally all we have to do is setup an continous loop with a select statement to read in either new log events or the subscription error.

```go
for {
  select {
  case err := <-sub.Err():
    log.Fatal(err)
  case vLog := <-logs:
    fmt.Println(vLog) // pointer to event log
  }
}
```

You'll have to parse the log entries, which we'll learn how to do in the [next section](../event-read).

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

[event_subscribe.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/event_subscribe.go)

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("wss://rinkeby.infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}
```

```bash
$ solc --version
0.4.24+commit.e67f0147.Emscripten.clang
```
