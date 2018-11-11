---
description: Tutorial on generating signatures with Go.
---

# Generating a Signature

The components for generating a signature are: the signers private key, and the hash of the data that will be signed. Any hashing algorithm may be used as long as the output is 32 bytes. We'll be using Keccak-256 as the hashing algorithm which is what Ethereum prefers to use.

First we'll load private key.

```go
privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
if err != nil {
  log.Fatal(err)
}
```

Next we'll take the Keccak-256 of the data that we wish to sign, in this case it'll be the word *hello*. The go-ethereum `crypto` package provides a handy `Keccak256Hash` method for doing this.

```go
data := []byte("hello")
hash := crypto.Keccak256Hash(data)
fmt.Println(hash.Hex()) // 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8
```

Finally we sign the hash with our private, which gives us the signature.

```go
signature, err := crypto.Sign(hash.Bytes(), privateKey)
if err != nil {
  log.Fatal(err)
}

fmt.Println(hexutil.Encode(signature)) // 0x789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c62621578113ddbb62d509bf6049b8fb544ab06d36f916685a2eb8e57ffadde02301
```

Now that we have successfully generated the signature, in the next section we'll learn how to verify that the signature indeed was signed by the holder of that private key.

---

### Full code

[signature_generate.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/signature_generate.go)

```go
package main

import (
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

	data := []byte("hello")
	hash := crypto.Keccak256Hash(data)
	fmt.Println(hash.Hex()) // 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hexutil.Encode(signature)) // 0x789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c62621578113ddbb62d509bf6049b8fb544ab06d36f916685a2eb8e57ffadde02301
}
```
