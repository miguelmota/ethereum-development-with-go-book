---
概述: 用Go使用whisper客户端的教程。
---

# 连接Whisper客户端


要使用连接Whisper客户端，我们必须首先连接到运行whisper的以太坊节点。 不幸的是，诸如infura之类的公共网关不支持whisper，因为没有金钱动力免费处理这些消息。 Infura可能会在不久的将来支持whisper，但现在我们必须运行我们自己的`geth`节点。一旦你[安装 geth](https://geth.ethereum.org/downloads/), 运行geth的时候加 `--shh` flag来支持whisper协议, 并且加 `--ws`flag和 `--rpc`，来支持websocket来接收实时信息，

```bash
geth --rpc --shh --ws
```

现在在我们的Go应用程序中，我们将导入在`whisper/shhclient`中找到的go-ethereum whisper客户端软件包并初始化客户端，使用默认的websocket端口“8546”通过websockets连接我们的本地geth节点。

```go
client, err := shhclient.Dial("ws://127.0.0.1:8546")
if err != nil {
  log.Fatal(err)
}

_ = client // we'll be using this in the 下个章节
```

现在我们已经拨打了，让我们创建一个密钥对来加密消息，然后再发送消息 [在下一章节](../whisper-keys).

---

### 完整代码

Commands

```bash
geth --rpc --shh --ws
```

[whisper_client.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/whisper_client.go)

```go
package main

import (
	"log"

	"github.com/ethereum/go-ethereum/whisper/shhclient"
)

func main() {
	client, err := shhclient.Dial("ws://127.0.0.1:8546")
	if err != nil {
		log.Fatal(err)
	}

	_ = client // we'll be using this in the 下个章节
	fmt.Println("we have a whisper connection")
}
```
