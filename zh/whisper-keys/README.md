---
description: Tutorial on how to generate whisper key pairs with Go.
---

# Generating Whisper Key Pair

In whisper, messages have to be encrypted with either a symmetric or an asymmetric key to prevent them from being read by anyone other than the intended recipient.

After you've connected to the whisper client you'll need to call the client's `NewKeyPair` method to generate a new public and private pair that the node will manage. The result of this function will be a unique ID that references the key pair which we'll be using for encrypting and decrypting the message in the next few sections.

```go
keyID, err := client.NewKeyPair(context.Background())
if err != nil {
  log.Fatal(err)
}

fmt.Println(keyID) // 0ec5cfe4e215239756054992dbc2e10f011db1cdfc88b9ba6301e2f9ea1b58d2
```

Let's learn how to send an encrypted message in the [next section](../whisper-send).

---

### Full code

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
