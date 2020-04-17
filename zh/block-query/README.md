---
概述: 用Go查询以太坊区块。
---

# 查询区块 

正如我们所见，您可以有两种方式查询区块信息。

#### 区块头 

您可以调用客户端的`HeaderByNumber`来返回有关一个区块的头信息。若您传入`nil`，它将返回最新的区块头。

```go
header, err := client.HeaderByNumber(context.Background(), nil)
if err != nil {
  log.Fatal(err)
}

fmt.Println(header.Number.String()) // 5671744
```

#### 完整区块 

调用客户端的`BlockByNumber`方法来获得完整区块。您可以读取该区块的所有内容和元数据，例如，区块号，区块时间戳，区块摘要，区块难度以及交易列表等等。

```go
blockNumber := big.NewInt(5671744)
block, err := client.BlockByNumber(context.Background(), blockNumber)
if err != nil {
  log.Fatal(err)
}

fmt.Println(block.Number().Uint64())     // 5671744
fmt.Println(block.Time().Uint64())       // 1527211625
fmt.Println(block.Difficulty().Uint64()) // 3217000136609065
fmt.Println(block.Hash().Hex())          // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
fmt.Println(len(block.Transactions()))   // 144
```

调用`Transaction`只返回一个区块的交易数目。

```go
count, err := client.TransactionCount(context.Background(), block.Hash())
if err != nil {
  log.Fatal(err)
}

fmt.Println(count) // 144
```

在下个章节，我们将学习查询区块中的交易。

### 完整代码

[blocks.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/blocks.go)

```go
package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(header.Number.String()) // 5671744

	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64())     // 5671744
	fmt.Println(block.Time().Uint64())       // 1527211625
	fmt.Println(block.Difficulty().Uint64()) // 3217000136609065
	fmt.Println(block.Hash().Hex())          // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	fmt.Println(len(block.Transactions()))   // 144

	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count) // 144
}
```
