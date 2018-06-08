# Sending Messages on Whisper

```bash
geth --shh --rpc
```

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	shh "github.com/ethereum/go-ethereum/whisper/shhclient"
	whisper6 "github.com/ethereum/go-ethereum/whisper/whisperv6"
)

func main() {
	client, err := shh.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	privateKeyID, err := client.NewKeyPair(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(privateKeyID) // 0ec5cfe4e215239756054992dbc2e10f011db1cdfc88b9ba6301e2f9ea1b58d2

	filterID, err := client.NewMessageFilter(context.Background(), whisper6.Criteria{PrivateKeyID: privateKeyID})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(filterID) // 21171f8b4e7ac0d7a1ce0d121b647ce10d4f0293b95d8fba69c5b4e9d0f235a6

	publicKey, err := client.PublicKey(context.Background(), privateKeyID)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(hexutil.Encode(publicKey)) // 0x04f17356fd52b0d13e5ede84f998d26276f1fc9d08d9e73dcac6ded5f3553405db38c2f257c956f32a0c1fca4c3ff6a38a2c277c1751e59a574aecae26d3bf5d1d

	messageHash, err := client.Post(context.Background(), whisper6.NewMessage{
		TTL:       60,
		PowTime:   2,
		PowTarget: 2.5,
		Payload:   []byte("Hello"),
		PublicKey: publicKey,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messageHash) // 0xdbfc815d3d122a90d7fb44d1fc6a46f3d76ec752f3f3d04230fe5f1b97d2209a

	messages, err := client.FilterMessages(context.Background(), filterID)
	if err != nil {
		log.Fatal(err)
	}
	for _, message := range messages {
		fmt.Printf(string(message.Payload)) // Hello
	}
}
```
