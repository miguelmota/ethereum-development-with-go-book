---
description: Tutorial on how to subscribe to messages on whisper with Go.
---

# Subscribing to Whisper Messages

In this section we'll be subscribing to whisper messages over websockets. First thing we need is a channel that will be receiving whisper messages in the `Message` type from the `whisper/whisperv6` package.

```go
messages := make(chan *whisperv6.Message)
```

Before we invoke a subscription, we first need to determine the criteria. From the whisperv6 package initialize a new `Criteria` object. Since we're only interested in messages targeted to us, we'll set the `PrivateKeyID` property on the criteria object to the same key ID we used for encrypting messages.

```go
criteria := whisperv6.Criteria{
  PrivateKeyID: keyID,
}
```

Next we invoke the client's `SubscribeMessages` method which subscribes to messages that match the given criteria. This method is not supported over HTTP; only supported on bi-directional connections such as websockets and IPC. The last argument is the messages channel we created earlier.

```go
sub, err := client.SubscribeMessages(context.Background(), criteria, messages)
if err != nil {
  log.Fatal(err)
}
```

Now that we have our subscription, we can use a `select` statement to read messages as they come in and also to handle errors from the subscription. If you recall from the previous section, the message content is in the `Payload` property as a byte slice which we can convert back to a human readable string.

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

Check out the full code below for a complete working example. That's all there is to whisper message subscriptions.

---

### Full code

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
