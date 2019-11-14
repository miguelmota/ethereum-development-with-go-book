---
概述: 用Go加载以太坊账户的教程。
---

# 账户

以太坊上的账户要么是钱包地址要么是智能合约地址。它们看起来像是`0x71c7656ec7ab88b098defb751b7401b5f6d8976f`，它们用于将ETH发送到另一个用户，并且还用于在需要和区块链交互时指一个智能合约。它们是唯一的，且是从私钥导出的。我们将在后面的章节更深入地介绍公私钥对。

要使用go-ethereum的账户地址，您必须先将它们转化为go-ethereum中的`common.Address`类型。

```go
address := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")

fmt.Println(address.Hex()) // 0x71C7656EC7ab88b098defB751B7401B5f6d8976F
```

您可以在几乎任何地方使用这种类型，您可以将以太坊地址传递给go-ethereum的方法。既然您已经了解账户和地址的基础知识，那么让我们在下一节中学习如何检索ETH账户余额。

---

### 完整代码

[address.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/address.go)

```go
package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	address := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")

	fmt.Println(address.Hex())        // 0x71C7656EC7ab88b098defB751B7401B5f6d8976F
	fmt.Println(address.Hash().Hex()) // 0x00000000000000000000000071c7656ec7ab88b098defb751b7401b5f6d8976f
	fmt.Println(address.Bytes())      // [113 199 101 110 199 171 136 176 152 222 251 117 27 116 1 181 246 216 151 111]
}
```
