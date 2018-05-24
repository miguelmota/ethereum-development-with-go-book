# Setting Up Client

Setting up the Ethereum client in Go is very simply. Import the `ethclient` go-ethereum package and initialize it by calling `Dial` which accepts a provider URL.

```go
client, err := ethclient.Dial("https://mainnet.infura.io")
```

You may also pass the path to the IPC endpoint file if you have a local instance of geth running.

```go
client, err := ethclient.Dial("/home/user/.ethereum/geth.ipc")
```

Full code

```go
package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client // we'll use this in the next section
}
```
