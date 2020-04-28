---
description: Tutorial on how to subscribe to latest blocks in Ethereum with Go.
---

# Subscribing to New Blocks

In this section we'll go over how to set up a subscription to get events when their is a new block mined. First thing is we need an Ethereum provider that supports RPC over websockets. In this example we'll use the infura websocket endpoint.

```go
client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
if err != nil {
  log.Fatal(err)
}
```

Next we'll create a new channel that will be receiving the latest block headers.

```go
headers := make(chan *types.Header)
```

Now we call the client's `SubscribeNewHead` method which takes in the headers channel we just created, which will return a subscription object.

```go
sub, err := client.SubscribeNewHead(context.Background(), headers)
if err != nil {
  log.Fatal(err)
}
```

The subscription will push new block headers to our channel so we'll use a select statement to listen for new messages. The subscription object also contains an error channel that will send a message in case of a failure with the subscription.

```go
for {
  select {
  case err := <-sub.Err():
    log.Fatal(err)
  case header := <-headers:
    fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
  }
}
```

To get the full contents of the block, we can pass the block header hash to the client's `BlockByHash` function.

```go
block, err := client.BlockByHash(context.Background(), header.Hash())
if err != nil {
  log.Fatal(err)
}

fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
fmt.Println(block.Number().Uint64())   // 3477413
fmt.Println(block.Time())              // 1529525947
fmt.Println(block.Nonce())             // 130524141876765836
fmt.Println(len(block.Transactions())) // 7
```

As you can see, you can read the entire block's metadata fields, list of transactions, and much more.

### Full code

[block_subscribe.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/block_subscribe.go)

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f

			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println(block.Number().Uint64())   // 3477413
			fmt.Println(block.Time())              // 1529525947
			fmt.Println(block.Nonce())             // 130524141876765836
			fmt.Println(len(block.Transactions())) // 7
		}
	}
}
```
