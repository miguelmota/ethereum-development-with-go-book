---
概述: 用Go从区块链读取账户余额教程。
---

# 账户余额 

读取一个账户的余额相当简单。调用客户端的`BalanceAt`方法，给它传递账户地址和可选的区块号。将区块号设置为`nil`将返回最新的余额。

```go
account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
balance, err := client.BalanceAt(context.Background(), account, nil)
if err != nil {
  log.Fatal(err)
}

fmt.Println(balance) // 25893180161173005034
```

传区块号能让您读取该区块时的账户余额。区块号必须是`big.Int`类型。

```go
blockNumber := big.NewInt(5532993)
balance, err := client.BalanceAt(context.Background(), account, blockNumber)
if err != nil {
  log.Fatal(err)
}

fmt.Println(balance) // 25729324269165216042
```

以太坊中的数字是使用尽可能小的单位来处理的，因为它们是定点精度，在ETH中它是*wei*。要读取ETH值，您必须做计算`wei/10^18`。因为我们正在处理大数，我们得导入原生的Go`math`和`math/big`包。这是您做的转换。

```go
fbalance := new(big.Float)
fbalance.SetString(balance.String())
ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

fmt.Println(ethValue) // 25.729324269165216041
```

#### 待处理的余额

有时您想知道待处理的账户余额是多少，例如，在提交或等待交易确认后。客户端提供了类似`BalanceAt`的方法，名为`PendingBalanceAt`，它接收账户地址作为参数。

```go
pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
fmt.Println(pendingBalance) // 25729324269165216042
```

---

### 完整代码

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
