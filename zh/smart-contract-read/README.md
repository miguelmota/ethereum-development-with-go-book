---
概述: Tutorial on how to query a smart contract with Go.
---

# Querying a Smart Contract

这写章节需要了解如何将智能合约的ABI编译成Go的合约文件。如果你还没看， 前先读[上一个章节](../smart-contract-compile) 。

在上个章节我们学习了如何在Go应用程序中初始化合约实例。 现在我们将使用新合约实例提供的方法来阅读智能合约。 如果你还记得我们在部署过程中设置的合约中有一个名为`version`的全局变量。 因为它是公开的，这意味着它们将成为我们自动创建的getter函数。 常量和view函数也接受`bind.CallOpts`作为第一个参数。了解可用的具体选项要看相应类的[文档](https://godoc.org/github.com/ethereum/go-ethereum/accounts/abi/bind#CallOpts) 一般情况下我们可以用 `nil`。

```go
version, err := instance.Version(nil)
if err != nil {
  log.Fatal(err)
}

fmt.Println(version) // "1.0"
```

---

### 完整代码

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
