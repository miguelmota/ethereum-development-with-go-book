---
概述: 用Go生成whisper密钥对的。
---

# 生成 Whisper 密匙对

在Whisper中，消息必须使用对称或非对称密钥加密，以防止除预期接收者以外的任何人读取消息。

在连接到Whisper客户端后，您需要调用客户端的`NewKeyPair`方法来生成该节点将管理的新公共和私有对。 此函数的结果将是一个唯一的ID，它引用我们将在接下来的几节中用于加密和解密消息的密钥对。


```go
keyID, err := client.NewKeyPair(context.Background())
if err != nil {
  log.Fatal(err)
}

fmt.Println(keyID) // 0ec5cfe4e215239756054992dbc2e10f011db1cdfc88b9ba6301e2f9ea1b58d2
```

在[下一章节](../whisper-send) 让我们学习如何发送一个加密的消息。

---

### 完整代码

Commands

```bash
geth --rpc --shh --ws
```

[whisper_keypair.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/whisper_keypair.go)

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/whisper/shhclient"
)

func main() {
	client, err := shhclient.Dial("ws://127.0.0.1:8546")
	if err != nil {
		log.Fatal(err)
	}

	keyID, err := client.NewKeyPair(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(keyID) // 0ec5cfe4e215239756054992dbc2e10f011db1cdfc88b9ba6301e2f9ea1b58d2
}
```
