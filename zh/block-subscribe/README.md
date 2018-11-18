---
概述: 用Go订阅以太坊中最新区块的教程。
---

# 订阅新区块 

在本节中，我们将讨论如何设置订阅以便在新区块被开采时获取事件。首先，我们需要一个支持websocket RPC的以太坊服务提供者。在示例中，我们将使用infura 的websocket端点。

```go
client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
if err != nil {
  log.Fatal(err)
}
```

接下来，我们将创建一个新的通道，用于接收最新的区块头。

```go
headers := make(chan *types.Header)
```

现在我们调用客户端的`SubscribeNewHead`方法，它接收我们刚创建的区块头通道，该方法将返回一个订阅对象。

```go
sub, err := client.SubscribeNewHead(context.Background(), headers)
if err != nil {
  log.Fatal(err)
}
```

订阅将推送新的区块头事件到我们的通道，因此我们可以使用一个select语句来监听新消息。订阅对象还包括一个error通道，该通道将在订阅失败时发送消息。

```go
for {
  select {
  case err := <-sub.Err():
    log.Fatal(err)
  case header := <-headers:
    fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
  }
}
```

要获得该区块的完整内容，我们可以将区块头的摘要传递给客户端的`BlockByHash`函数。

```go
block, err := client.BlockByHash(context.Background(), header.Hash())
if err != nil {
  log.Fatal(err)
}

fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
fmt.Println(block.Number().Uint64())   // 3477413
fmt.Println(block.Time().Uint64())     // 1529525947
fmt.Println(block.Nonce())             // 130524141876765836
fmt.Println(len(block.Transactions())) // 7
```

正如您所见，您可以读取整个区块的元数据字段，交易列表等等。

### 完整代码

[block_subscribe.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/block_subscribe.go)

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f

			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println(block.Number().Uint64())   // 3477413
			fmt.Println(block.Time().Uint64())     // 1529525947
			fmt.Println(block.Nonce())             // 130524141876765836
			fmt.Println(len(block.Transactions())) // 7
		}
	}
}
```
