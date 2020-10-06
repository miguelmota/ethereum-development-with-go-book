---
概述: 用Go生成以太坊钱包的教程。
---

# 生成新钱包

要首先生成一个新的钱包，我们需要导入go-ethereum`crypto`包，该包提供用于生成随机私钥的`GenerateKey`方法。

```go
privateKey, err := crypto.GenerateKey()
if err != nil {
  log.Fatal(err)
}
```

然后我们可以通过导入golang`crypto/ecdsa`包并使用`FromECDSA`方法将其转换为字节。

```go
privateKeyBytes := crypto.FromECDSA(privateKey)
```

我们现在可以使用go-ethereum`hexutil`包将它转换为十六进制字符串，该包提供了一个带有字节切片的`Encode`方法。 然后我们在十六进制编码之后删除“0x”。

```go
fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19
```

这就是用于签署交易的私钥，将被视为密码，永远不应该被共享给别人，因为谁拥有它可以访问你的所有资产。

由于公钥是从私钥派生的，因此go-ethereum的加密私钥具有一个返回公钥的`Public`方法。

```go
publicKey := privateKey.Public()
```

将其转换为十六进制的过程与我们使用转化私钥的过程类似。 我们剥离了`0x`和前2个字符`04`，它始终是EC前缀，不是必需的。

```go
publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
if !ok {
  log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
}

publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
fmt.Println(hexutil.Encode(publicKeyBytes)[4:]) // 9a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05
```

现在我们拥有公钥，就可以轻松生成你经常看到的公共地址。 为了做到这一点，go-ethereum加密包有一个`PubkeyToAddress`方法，它接受一个ECDSA公钥，并返回公共地址。

```go
address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
fmt.Println(address) // 0x96216849c49358B10257cb55b28eA603c874b05E
```

公共地址其实就是公钥的Keccak-256哈希，然后我们取最后40个字符（20个字节）并用“0x”作为前缀。 以下是使用 `golang.org/x/crypto/sha3` 的 Keccak256函数手动完成的方法。

```go
hash := sha3.NewLegacyKeccak256()
hash.Write(publicKeyBytes[1:])
fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 0x96216849c49358b10257cb55b28ea603c874b05e
```

---

### 完整代码

[wallet_generate.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/wallet_generate.go)

```go
package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func main() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // 0xfad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:]) // 0x049a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address) // 0x96216849c49358B10257cb55b28eA603c874b05E

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 0x96216849c49358b10257cb55b28ea603c874b05e
}
```
