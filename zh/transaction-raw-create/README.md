---
概述: 用Go构建以太坊原始交易的教程。
---

# 构建原始交易（Raw Transaction）

如果你看过[上个章节](../transfer-eth), 那么你知道如何加载你的私钥来签名交易。 我们现在假设你知道如何做到这一点，现在你想让原始交易数据能够在以后广播它。

首先构造事务对象并对其进行签名，例如：

```go
tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
if err != nil {
  log.Fatal(err)
}
```

现在，在我们以原始字节格式获取事务之前，我们需要初始化一个`types.Transactions`类型，并将签名后的交易作为第一个值。

```go
ts := types.Transactions{signedTx}
```

这样做的原因是因为`Transactions`类型提供了一个`GetRlp`方法，用于以RLP编码格式返回事务。 RLP是以太坊用于序列化对象的特殊编码方法。 结果是原始字节。

```go
rawTxBytes := ts.GetRlp(0)
```

最后，我们可以非常轻松地将原始字节转换为十六进制字符串。

```go
rawTxHex := hex.EncodeToString(rawTxBytes)

fmt.Printf(rawTxHex)
// f86d8202b38477359400825208944592d8f8d7b001e72cb26a73e4fa1806a51ac79d880de0b6b3a7640000802ba0699ff162205967ccbabae13e07cdd4284258d46ec1051a70a51be51ec2bc69f3a04e6944d508244ea54a62ebf9a72683eeadacb73ad7c373ee542f1998147b220e
```

接下来，你就可以广播原始交易数据。在[下一章](../transaction-raw-send) 我们将学习如何广播一个原始交易。

---

### 完整代码

[transaction_raw_create.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/transaction_raw_create.go)

```go
package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
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

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	ts := types.Transactions{signedTx}
	rawTxBytes := ts.GetRlp(0)
	rawTxHex := hex.EncodeToString(rawTxBytes)

	fmt.Printf(rawTxHex) // f86...772
}
```
