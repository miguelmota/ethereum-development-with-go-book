---
概述: 用Go发送以太坊原始交易事务的教程。
---

# 发送原始交易事务

在[上个章节中](../transaction-raw-create) 我们学会了如何创建原始事务。 现在，我们将学习如何将其广播到以太坊网络，以便最终被处理和被矿工打包到区块。

首先将原始事务十六进制解码为字节格式。

```go
rawTx := "f86d8202b28477359400825208944592d8f8d7b001e72cb26a73e4fa1806a51ac79d880de0b6b3a7640000802ca05924bde7ef10aa88db9c66dd4f5fb16b46dff2319b9968be983118b57bb50562a001b24b31010004f13d9a26b320845257a6cfc2bf819a3d55e3fc86263c5f0772"

rawTxBytes, err := hex.DecodeString(rawTx)
```

接下来初始化一个新的`types.Transaction`指针并从go-ethereum`rlp`包中调用`DecodeBytes`，将原始事务字节和指针传递给以太坊事务类型。 RLP是以太坊用于序列化和反序列化数据的编码方法。


```go
tx := new(types.Transaction)
rlp.DecodeBytes(rawTxBytes, &tx)
```

现在，我们可以使用我们的以太坊客户端轻松地广播交易。

```go
err := client.SendTransaction(context.Background(), tx)
if err != nil {
  log.Fatal(err)
}

fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0xc429e5f128387d224ba8bed6885e86525e14bfdc2eb24b5e9c3351a1176fd81f
```

然后你可以去Etherscan看交易的确认过程: [https://rinkeby.etherscan.io/tx/0xc429e5f128387d224ba8bed6885e86525e14bfdc2eb24b5e9c3351a1176fd81f](https://rinkeby.etherscan.io/tx/0xc429e5f128387d224ba8bed6885e86525e14bfdc2eb24b5e9c3351a1176fd81f)

---

### 完整代码

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
