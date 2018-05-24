# Setting Up Client

Setting up the Ethereum client in Go is very simply. Import the `ethclient` go-ethereum package and initialize it by calling `Dial` which accepts a provider URL.

You can connect to the infura gateway if you don't have an existing client. Infura manages a bunch of Ethereum [geth and parity] nodes that are trusted and reliable and lowers the entry to barrier for newcomers when it comes to plugging into the Ethereum network.

```go
client, err := ethclient.Dial("https://mainnet.infura.io")
```

You may also pass the path to the IPC endpoint file if you have a local instance of geth running.

```go
client, err := ethclient.Dial("/home/user/.ethereum/geth.ipc")
```

Using the ethclient is a fundamental thing you'll need to start with for every Go Ethereum project.

**Full code**

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
