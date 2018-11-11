---
description: Tutorial on how to query a smart contract with Go.
---

# Querying a Smart Contract

These section requires knowledge of how to compile a smart contract's ABI to a Go contract file. If you haven't already gone through it, please [read the section](../smart-contract-compile) first.

In the previous section we learned how to initialize a contract instance in our Go application. Now we're going to read the smart contract using the provided methods by the new contract instance. If you recall we had a global variable named `version` in our contract that was set during deployment. Because it's public that means that they'll be a getter function automatically created for us. Constant and view functions also accept `bind.CallOpts` as the first argument. To learn about what options you can pass checkout the type's [documentation](https://godoc.org/github.com/ethereum/go-ethereum/accounts/abi/bind#CallOpts) but usually this is set to `nil`.

```go
version, err := instance.Version(nil)
if err != nil {
  log.Fatal(err)
}

fmt.Println(version) // "1.0"
```

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

[contract_read.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/contract_read.go)

```go
package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	store "./contracts" // for demo
)

func main() {
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	version, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version) // "1.0"
}
```

solc version used for these examples

```bash
$ solc --version
0.4.24+commit.e67f0147.Emscripten.clang
```
