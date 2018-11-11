---
description: Tutorial on how to send a message on whisper with Go.
---

# Sending Messages on Whisper

Before we're able to create a message, we must first have a public key to encrypt the message. In the [previous section](../whisper-keys) we learned how to generate a public and private key pair using the `NewKeyPair` function which returned a key ID that references this key pair. We now have to call the `PublicKey` function to read the key pair's public key in bytes format which we'll be using to encrypt the message.

```go
publicKey, err := client.PublicKey(context.Background(), keyID)
if err != nil {
  log.Print(err)
}

fmt.Println(hexutil.Encode(publicKey)) // 0x04f17356fd52b0d13e5ede84f998d26276f1fc9d08d9e73dcac6ded5f3553405db38c2f257c956f32a0c1fca4c3ff6a38a2c277c1751e59a574aecae26d3bf5d1d
```

Now we'll construct our whisper message by initializing the `NewMessage` struct from the go-ethereum `whisper/whisperv6` package, which requires the following properties:

- `Payload` as the message content in bytes format
- `PublicKey` as the key we'll use for encryption
- `TTL` as the time-to-live in seconds for the message
- `PowTime` as maximal time in seconds to be spent on proof of work.
- `PowTarget` as the minimal PoW target required for this message.

```go
message := whisperv6.NewMessage{
  Payload:   []byte("Hello"),
  PublicKey: publicKey,
  TTL:       60,
  PowTime:   2,
  PowTarget: 2.5,
}
```

We can now broadcast to the network by invoking the client's `Post` function giving it the message, will it'll return a hash of the message.

```go
messageHash, err := client.Post(context.Background(), message)
if err != nil {
  log.Fatal(err)
}

fmt.Println(messageHash) // 0xdbfc815d3d122a90d7fb44d1fc6a46f3d76ec752f3f3d04230fe5f1b97d2209a
```

In the [next section](../whisper-subscribe) we'll see how we can create a message subscription to be able to receive the messages in real time.

---

### Full code

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
