---
description: Tutorial on how to create a raw Ethereum transaction with Go.
---

# Create Raw Transaction

If you've read the [previous sections](../transfer-eth), then you know how to load your private key to sign transactions. We'll assume you know how to do that by now and now you want to get the raw transaction data to be able to broadcast it at a later time.

First construct the transaction object and sign it, for example:

```go
tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
if err != nil {
  log.Fatal(err)
}
```

Now before we can get the transaction in raw bytes format we'll need to initialize a `types.Transactions` type with the signed transaction as the first value.

```go
ts := types.Transactions{signedTx}
```

The reason for doing this is because the `Transactions` type provides a `GetRlp` method for returning the transaction in RLP encoded format. RLP is a special encoding method Ethereum uses for serializing objects. The result of this is raw bytes.

```go
rawTxBytes := ts.GetRlp(0)
```

Finally we can very easily turn the raw bytes into a hex string.

```go
rawTxHex := hex.EncodeToString(rawTxBytes)

fmt.Printf(rawTxHex)
// f86d8202b38477359400825208944592d8f8d7b001e72cb26a73e4fa1806a51ac79d880de0b6b3a7640000802ba0699ff162205967ccbabae13e07cdd4284258d46ec1051a70a51be51ec2bc69f3a04e6944d508244ea54a62ebf9a72683eeadacb73ad7c373ee542f1998147b220e
```

And now you have the raw transaction data which you can use to broadcast at a future date. In the [next section](../transaction-raw-send) we'll learn how to broadcast a raw transaction.

---

### Full code

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
