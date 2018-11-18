---
概述: 用Go在whisper中监听/订阅消息的教程。
---

# 监听/订阅Whisper消息

在本节中，我们将订阅websockets上的Whisper消息。 我们首先需要的是一个通道，它将从`whisper/whisperv6`包中的`Message`类型接收Whispe消息。

```go
messages := make(chan *whisperv6.Message)
```

在我们调用订阅之前，我们首先需要确定消息的过滤标准。 从whisperv6包中初始化一个新的`Criteria`对象。 由于我们只对定位到我们的消息感兴趣，因此我们将条件对象上的`PrivateKeyID`属性设置为我们用于加密消息的相同密钥ID。

```go
criteria := whisperv6.Criteria{
  PrivateKeyID: keyID,
}
```

接下来，我们调用客户端的`SubscribeMessages`方法，该方法订阅符合给定条件的消息。 HTTP不支持此方法; 仅支持双向连接，例如websockets和IPC。 最后一个参数是我们之前创建的消息通道。

```go
sub, err := client.SubscribeMessages(context.Background(), criteria, messages)
if err != nil {
  log.Fatal(err)
}
```

现在我们已经订阅了，我们可以使用`select`语句来读取消息，并处理订阅中的错误。 如果您从上一节回忆起来，消息内容在`Payload`属性中作为字节切片，我们可以将其转换回人类可读的字符串。

```go
for {
  select {
  case err := <-sub.Err():
    log.Fatal(err)
  case message := <-messages:
    fmt.Printf(string(message.Payload)) // "Hello"
  }
}
```

查看下面的完整代码，获取完整的栗子。 这就是消息订阅的所有内容。

---

### 完整代码

Commands

```bash
geth --shh --rpc --ws
```

[whisper_subscribe.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/whisper_subscribe.go)

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"

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

	messages := make(chan *whisperv6.Message)
	criteria := whisperv6.Criteria{
		PrivateKeyID: keyID,
	}
	sub, err := client.SubscribeMessages(context.Background(), criteria, messages)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case message := <-messages:
				fmt.Printf(string(message.Payload)) // "Hello"
				os.Exit(0)
			}
		}
	}()

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

	runtime.Goexit() // wait for goroutines to finish
}
```
