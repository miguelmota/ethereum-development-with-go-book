---
description: Tutorial on how to read account balances from the blockchain with Go.
---

# Account Balances

Reading the balance of an account is pretty simple; call the `BalanceAt` method of the client passing it the account address and optional block number. Setting `nil` as the block number will return the latest balance.

```go
account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
balance, err := client.BalanceAt(context.Background(), account, nil)
if err != nil {
  log.Fatal(err)
}

fmt.Println(balance) // 25893180161173005034
```

Passing the block number let's you read the account balance at the time of that block. The block number must be a `big.Int`.

```go
blockNumber := big.NewInt(5532993)
balance, err := client.BalanceAt(context.Background(), account, blockNumber)
if err != nil {
  log.Fatal(err)
}

fmt.Println(balance) // 25729324269165216042
```

Numbers in ethereum are dealt using the smallest possible unit because they're fixed-point precision, which in the case of ETH it's *wei*. To read the ETH value you must do the calculation `wei / 10^18`. Because we're dealing with big numbers we'll have to import the native Go `math` and `math/big` packages. Here's how'd you do the conversion.

```go
fbalance := new(big.Float)
fbalance.SetString(balance.String())
ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

fmt.Println(ethValue) // 25.729324269165216041
```

#### Pending balance

Sometimes you'll want to know what the pending account balance is, for example after submitting or waiting for a transaction to be confirmed. The client provides a similar method to `BalanceAt` called `PendingBalanceAt` which accepts the account address as a parameter.

```go
pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
fmt.Println(pendingBalance) // 25729324269165216042
```

---

### Full code

[account_balance.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/account_balance.go)

```go
package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance) // 25893180161173005034

	blockNumber := big.NewInt(5532993)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balanceAt) // 25729324269165216042

	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue) // 25.729324269165216041

	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println(pendingBalance) // 25729324269165216042
}
```
