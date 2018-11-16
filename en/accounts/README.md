---
description: Tutorial on how to load an Ethereum account with Go.
---

# Accounts

Accounts on Ethereum are either wallet addresses or smart contract addresses. They look like `0x71c7656ec7ab88b098defb751b7401b5f6d8976f` and they're what you use for sending ETH to another user and also are used for referring to a smart contract on the blockchain when needing to interact with it. They are unique and are derived from a private key. We'll go more in depth into private/public key pairs in later sections.

In order to use account addresses with go-ethereum, you must first convert them to the go-ethereum `common.Address` type.

```go
address := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")

fmt.Println(address.Hex()) // 0x71C7656EC7ab88b098defB751B7401B5f6d8976F
```

Pretty much you'd use this type anywhere you'd pass an ethereum address to methods from go-ethereum. Now that you know the basics of accounts and addresses, let's learn how to retrieve the ETH account balance in the next section.

---

### Full code

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
