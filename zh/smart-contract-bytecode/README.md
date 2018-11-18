---
概述: 用Go来读取已经部署的智能合约的字节码的教程。
---

# 读取智能合约的字节码

有时您需要读取已部署的智能合约的字节码。 由于所有智能合约字节码都存在于区块链中，因此我们可以轻松获取它。

首先设置客户端和要读取的字节码的智能合约地址。

```go
client, err := ethclient.Dial("https://rinkeby.infura.io")
if err != nil {
  log.Fatal(err)
}

contractAddress := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
```

现在你需要调用客户端的`codeAt`方法。 `codeAt`方法接受智能合约地址和可选的块编号，并以字节格式返回字节码。

```go
bytecode, err := client.CodeAt(context.Background(), contractAddress, nil) // nil is latest block
if err != nil {
  log.Fatal(err)
}

fmt.Println(hex.EncodeToString(bytecode)) // 60806...10029
```


你也可以在etherscan上查询16进制格式的字节码 [https://rinkeby.etherscan.io/address/0x147b8eb97fd247d06c4006d269c90c1908fb5d54#code](https://rinkeby.etherscan.io/address/0x147b8eb97fd247d06c4006d269c90c1908fb5d54#code)

---

### 完整代码

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
