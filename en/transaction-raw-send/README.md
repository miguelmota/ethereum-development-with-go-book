---
description: Tutorial on how to send a raw Ethereum transaction with Go.
---

# Send Raw Transaction

In the [previous section](../transaction-raw-create) we learned how to create a raw transaction. Now we'll learn how to broadcast it to the Ethereum network in order for it to get processed and mined.

First decode the raw transaction hex to bytes format.

```go
rawTx := "f86d8202b28477359400825208944592d8f8d7b001e72cb26a73e4fa1806a51ac79d880de0b6b3a7640000802ca05924bde7ef10aa88db9c66dd4f5fb16b46dff2319b9968be983118b57bb50562a001b24b31010004f13d9a26b320845257a6cfc2bf819a3d55e3fc86263c5f0772"

rawTxBytes, err := hex.DecodeString(rawTx)
```

Now initialize a new `types.Transaction` pointer and call `DecodeBytes` from the go-ethereum `rlp` package passing it the raw transaction bytes and the pointer to the ethereum transaction type. RLP is an encoding method used by Ethereum to serialized and derialized data.

```go
tx := new(types.Transaction)
rlp.DecodeBytes(rawTxBytes, &tx)
```

Now we can easily broadcast the transaction with our ethereum client.

```go
err := client.SendTransaction(context.Background(), tx)
if err != nil {
  log.Fatal(err)
}

fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0xc429e5f128387d224ba8bed6885e86525e14bfdc2eb24b5e9c3351a1176fd81f
```

You can see the transaction on etherscan: [https://rinkeby.etherscan.io/tx/0xc429e5f128387d224ba8bed6885e86525e14bfdc2eb24b5e9c3351a1176fd81f](https://rinkeby.etherscan.io/tx/0xc429e5f128387d224ba8bed6885e86525e14bfdc2eb24b5e9c3351a1176fd81f)

---

### Full code

[transaction_raw_sendreate.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/transaction_raw_send.go)

```go
package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	rawTx := "f86d8202b28477359400825208944592d8f8d7b001e72cb26a73e4fa1806a51ac79d880de0b6b3a7640000802ca05924bde7ef10aa88db9c66dd4f5fb16b46dff2319b9968be983118b57bb50562a001b24b31010004f13d9a26b320845257a6cfc2bf819a3d55e3fc86263c5f0772"

	rawTxBytes, err := hex.DecodeString(rawTx)

	tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0xc429e5f128387d224ba8bed6885e86525e14bfdc2eb24b5e9c3351a1176fd81f
}
```
