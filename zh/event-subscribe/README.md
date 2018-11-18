---
概述: 用Go订阅智能合约事件日志的教程。
---

# 订阅事件日志

为了订阅事件日志，我们需要做的第一件事就是拨打启用websocket的以太坊客户端。 幸运的是，Infura支持websockets。

```go
client, err := ethclient.Dial("wss://rinkeby.infura.io/ws")
if err != nil {
  log.Fatal(err)
}
```

下一步是创建筛选查询。 在这个例子中，我们将阅读来自我们在之前课程中创建的示例合约中的所有事件。

```go
contractAddress := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
query := ethereum.FilterQuery{
  Addresses: []common.Address{contractAddress},
}
```

我们接收事件的方式是通过Go channel。 让我们从go-ethereum`core/types`包创建一个类型为`Log`的channel。

```go
logs := make(chan types.Log)
```

现在我们所要做的就是通过从客户端调用`SubscribeFilterLogs`来订阅，它接收查询选项和输出通道。 这将返回包含unsubscribe和error方法的订阅结构。


```go
sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
if err != nil {
  log.Fatal(err)
}
```

最后，我们要做的就是使用select语句设置一个连续循环来读入新的日志事件或订阅错误。

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

我们会在[下个章节](../event-read)介绍如何解析日志。

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
