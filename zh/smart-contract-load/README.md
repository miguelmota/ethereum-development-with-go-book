---
概述: 用Go加载和初始化智能合约的教程。
---

# 加载智能合约

这写章节需要了解如何将智能合约的ABI编译成Go的合约文件。如果你还没看， 前先读[上一个章节](../smart-contract-compile) 。

一旦使用`abigen`工具将智能合约的ABI编译为Go包，下一步就是调用“New”方法，其格式为“New<ContractName>”，所以在我们的例子中如果你 回想一下它将是*NewStore*。 此初始化方法接收智能合约的地址，并返回可以开始与之交互的合约实例。

```go
address := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
instance, err := store.NewStore(address, client)
if err != nil {
  log.Fatal(err)
}

_ = instance // we'll be using this in the 下个章节
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

[contract_load.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/contract_load.go)

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

	fmt.Println("contract is loaded")
	_ = instance
}
```

solc version used for these examples

```bash
$ solc --version
0.4.24+commit.e67f0147.Emscripten.clang
```
