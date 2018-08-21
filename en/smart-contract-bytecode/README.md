---
description: Tutorial on how to read the bytecode of a deployed smart contract with Go.
---

# Reading Smart Contract Bytecode

Sometimes you'll need to read the bytecode of a deployed smart contract. Since all the smart contract bytecode lives on the blockchain, we can easily fetch it.

First set up the client and the smart contract address you want to read the bytecode of.

```go
client, err := ethclient.Dial("https://rinkeby.infura.io")
if err != nil {
  log.Fatal(err)
}

contractAddress := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
```

Now all you have to is call the `codeAt` method of the client. The `codeAt` method accepts a smart contract address and an optional block number, and returns the bytecode in bytes format.

```go
bytecode, err := client.CodeAt(context.Background(), contractAddress, nil) // nil is latest block
if err != nil {
  log.Fatal(err)
}

fmt.Println(hex.EncodeToString(bytecode)) // 60806...10029
```


See the same bytecode hex on etherscan [https://rinkeby.etherscan.io/address/0x147b8eb97fd247d06c4006d269c90c1908fb5d54#code](https://rinkeby.etherscan.io/address/0x147b8eb97fd247d06c4006d269c90c1908fb5d54#code)

---

### Full code

[contract_bytecode.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/contract_bytecode.go)

```go
package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	bytecode, err := client.CodeAt(context.Background(), contractAddress, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hex.EncodeToString(bytecode)) // 60806...10029
}
```
