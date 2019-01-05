---
概述: 用Go来验证签名的教程。
---

# 验证签名

在上个章节中，我们学习了如何使用私钥对一段数据进行签名以生成签名。 现在我们将学习如何验证签名的真实性。

我们需要有3件事来验证签名：签名，原始数据的哈希以及签名者的公钥。 利用该信息，我们可以确定公钥对的私钥持有者是否确实签署了该消息。

首先，我们需要以字节格式的公钥。


```go
publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
```

接下来我们将需要原始数据哈希。 在上一课中，我们使用Keccak-256生成哈希，因此我们将执行相同的操作以验证签名。

```go
data := []byte("hello")
hash := crypto.Keccak256Hash(data)
fmt.Println(hash.Hex()) // 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8
```

现在假设我们有字节格式的签名，我们可以从go-ethereum`crypto`包调用`Ecrecover`（椭圆曲线签名恢复）来检索签名者的公钥。 此函数采用字节格式的哈希和签名。

```go
sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
if err != nil {
  log.Fatal(err)
}
```

为了验证我们现在必须将签名的公钥与期望的公钥进行比较，如果它们匹配，那么预期的公钥持有者确实是原始消息的签名者。

```go
matches := bytes.Equal(sigPublicKey, publicKeyBytes)
fmt.Println(matches) // true
```

还有`SigToPub`方法做同样的事情，区别是它将返回ECDSA类型中的签名公钥。

```go
sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signature)
if err != nil {
  log.Fatal(err)
}

sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
matches = bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
fmt.Println(matches) // true
```

为方便起见，`go-ethereum/crypto`包提供了`VerifySignature`函数，该函数接收原始数据的签名，哈希值和字节格式的公钥。 它返回一个布尔值，如果公钥与签名的签名者匹配，则为true。 一个重要的问题是我们必须首先删除signture的最后一个字节，因为它是ECDSA恢复ID，不能包含它。

```go
signatureNoRecoverID := signature[:len(signature)-1] // remove recovery ID
verified := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
fmt.Println(verified) // true
```

这些就是使用go-ethereum软件包生成和验证ECDSA签名的基础知识。

---

### 完整代码

[signature_verify.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/signature_verify.go)

```go
package main

import (
	"bytes"
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	data := []byte("hello")
	hash := crypto.Keccak256Hash(data)
	fmt.Println(hash.Hex()) // 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hexutil.Encode(signature)) // 0x789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c62621578113ddbb62d509bf6049b8fb544ab06d36f916685a2eb8e57ffadde02301

	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}

	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
	fmt.Println(matches) // true

	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}

	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
	matches = bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
	fmt.Println(matches) // true

	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	verified := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
	fmt.Println(verified) // true
}
```
