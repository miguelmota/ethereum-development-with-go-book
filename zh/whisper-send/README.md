---
概述: 用Go在whisper上发送消息的教程。
---

# 在Whisper上发送消息


在我们能够创建消息之前，我们必须首先使用公钥来加密消息。在[上个章节](../whisper-keys)中，我们学习了如何使用`NewKeyPair`函数生成公钥和私钥对，该函数返回了引用该密钥对的密钥ID。 我们现在必须调用`PublicKey`函数以字节格式读取密钥对的公钥，我们将使用它来加密消息。

```go
publicKey, err := client.PublicKey(context.Background(), keyID)
if err != nil {
  log.Print(err)
}

fmt.Println(hexutil.Encode(publicKey)) // 0x04f17356fd52b0d13e5ede84f998d26276f1fc9d08d9e73dcac6ded5f3553405db38c2f257c956f32a0c1fca4c3ff6a38a2c277c1751e59a574aecae26d3bf5d1d
```

现在我们将通过从go-ethereum`whisper/whisperv6`包中初始化`NewMessage`结构来构造我们的私语消息，这需要以下属性：

- `Payload` 字节格式的消息内容
- `PublicKey` 加密的公钥
- `TTL` 消息的活跃时间
- `PowTime` 做工证明的时间上限
- `PowTarget` 做工证明的时间下限

```go
message := whisperv6.NewMessage{
  Payload:   []byte("Hello"),
  PublicKey: publicKey,
  TTL:       60,
  PowTime:   2,
  PowTarget: 2.5,
}
```

我们现在可以通过调用客户端的`Post`函数向网络广播，给它消息，它是否会返回消息的哈希值。

```go
messageHash, err := client.Post(context.Background(), message)
if err != nil {
  log.Fatal(err)
}

fmt.Println(messageHash) // 0xdbfc815d3d122a90d7fb44d1fc6a46f3d76ec752f3f3d04230fe5f1b97d2209a
```

在[下个章节](../whisper-subscribe)中我们将看到如何创建消息订阅以便能够实时接收消息。

---

### 完整代码

Commands

```bash
geth --shh --rpc --ws
```

[whisper_send.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/whisper_send.go)

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/whisper/shhclient"
	"github.com/ethereum/go-ethereum/whisper/whisperv6"
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

	publicKey, err := client.PublicKey(context.Background(), keyID)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(hexutil.Encode(publicKey)) // 0x04f17356fd52b0d13e5ede84f998d26276f1fc9d08d9e73dcac6ded5f3553405db38c2f257c956f32a0c1fca4c3ff6a38a2c277c1751e59a574aecae26d3bf5d1d

	message := whisperv6.NewMessage{
		Payload:   []byte("Hello"),
		PublicKey: publicKey,
		TTL:       60,
		PowTime:   2,
		PowTarget: 2.5,
	}
	messageHash, err := client.Post(context.Background(), message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messageHash) // 0xdbfc815d3d122a90d7fb44d1fc6a46f3d76ec752f3f3d04230fe5f1b97d2209a
}
```
