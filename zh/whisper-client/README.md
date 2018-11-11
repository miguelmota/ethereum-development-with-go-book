---
description: Tutorial on whisper with Go.
---

# Connecting Whisper Client

To use whisper, we must first connect to an Ethereum node running whisper. Unfortunately, public gateways such as infura don't support whisper because there is no incentive for processing the messages for free. Infura might support whisper in the near future but for now we must run our own `geth` node. Once you [install geth](https://geth.ethereum.org/downloads/), run it with the `--shh` flag on to enable the whisper protocol, as well as the `--ws` flag to enable websocket support in order to receive messages in real time, and also enable the `--rpc` flag because we'll be communicating over RPC.

```bash
geth --rpc --shh --ws
```

Now in our Go application we'll import the go-ethereum whisper client package found at `whisper/shhclient` and initialize the client to connect our local geth node over websockets using the default websocket port `8546`.

```go
client, err := shhclient.Dial("ws://127.0.0.1:8546")
if err != nil {
  log.Fatal(err)
}

_ = client // we'll be using this in the next section
```

Now that we're dialed in let's create a key pair for encrypting the message before we send it in the [next section](../whisper-keys).

---

### Full code

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

	_ = client // we'll be using this in the next section
	fmt.Println("we have a whisper connection")
}
```
