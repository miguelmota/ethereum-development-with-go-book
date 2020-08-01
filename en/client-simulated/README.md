---
description: Tutorial on how to set up a simulated backend as your client to test your Ethereum application with Go.
---

# Using a Simulated Client

You can use a simulated client for testing your transactions locally quickly and easily, ideal for unit tests. In order to get started we're going to need an account with some initial ETH in it. To do that first generate an account private key.

```go
privateKey, err := crypto.GenerateKey()
if err != nil {
  log.Fatal(err)
}
```

Then create a `NewKeyedTransactor` from the `accounts/abi/bind` package passing the private key.

```go
auth := bind.NewKeyedTransactor(privateKey)
```

The next step is to create a genesis account and assign it an initial balance. We'll be using the `GenesisAccount` type from the `core` package.

```go
balance := new(big.Int)
balance.SetString("10000000000000000000", 10) // 10 eth in wei

address := auth.From
genesisAlloc := map[common.Address]core.GenesisAccount{
  address: {
    Balance: balance,
  },
}
```

Now we pass the genesis allocation struct and a configured block gas limit to the `NewSimulatedBackend` method from the `accounts/abi/bind/backends` package which will return a new simulated ethereum client.

```go
blockGasLimit := uint64(4712388)
client := backends.NewSimulatedBackend(genesisAlloc, blockGasLimit)
```

You can use this client as you'd normally would. As an example, we'll construct a new transaction and broadcast it.

```go
fromAddress := auth.From
nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
if err != nil {
  log.Fatal(err)
}

value := big.NewInt(1000000000000000000) // in wei (1 eth)
gasLimit := uint64(21000)                // in units
gasPrice, err := client.SuggestGasPrice(context.Background())
if err != nil {
  log.Fatal(err)
}

toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
var data []byte
tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
if err != nil {
  log.Fatal(err)
}

err = client.SendTransaction(context.Background(), signedTx)
if err != nil {
  log.Fatal(err)
}

fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex()) // tx sent: 0xec3ceb05642c61d33fa6c951b54080d1953ac8227be81e7b5e4e2cfed69eeb51
```

By now you're probably wondering when will the transaction actually get mined. Well in order to "mine" it, there's one additional important thing you must do; call `Commit` on the client to commit a new mined block.

```go
client.Commit()
```

Now you can fetch the transaction receipt and see that it was processed.

```go
receipt, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
if err != nil {
  log.Fatal(err)
}
if receipt == nil {
  log.Fatal("receipt is nil. Forgot to commit?")
}

fmt.Printf("status: %v\n", receipt.Status) // status: 1
```

So remember that the simulated client allows you to manually mine blocks at your command using the simulated client's `Commit` method.

----

### Full code

[client_simulated.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/client_simulated.go)

```go
package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)

	balance := new(big.Int)
	balance.SetString("10000000000000000000", 10) // 10 eth in wei

	address := auth.From
	genesisAlloc := map[common.Address]core.GenesisAccount{
		address: {
			Balance: balance,
		},
	}

	blockGasLimit := uint64(4712388)
	client := backends.NewSimulatedBackend(genesisAlloc, blockGasLimit)

	fromAddress := auth.From
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)                // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex()) // tx sent: 0xec3ceb05642c61d33fa6c951b54080d1953ac8227be81e7b5e4e2cfed69eeb51

	client.Commit()

	receipt, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	if receipt == nil {
		log.Fatal("receipt is nil. Forgot to commit?")
	}

	fmt.Printf("status: %v\n", receipt.Status) // status: 1
}
```
