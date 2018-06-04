# Account Balances

Reading the balance of an account is pretty simple; call `ethclient.BalanceAt` passing the account address and optional block number.

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

---

### Full code

[account_balance.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/account_balance.go)

```go
package main

import (
	"context"
	"fmt"
	"log"
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
}
```
