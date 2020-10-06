---
description: Tutorial on how to generate Ethereum wallets with Go.
---

# Generating New Wallets

To generate a new wallet first we need to import the go-ethereum `crypto` package that provides the `GenerateKey` method for generating a random private key.

```go
privateKey, err := crypto.GenerateKey()
if err != nil {
  log.Fatal(err)
}
```

Then we can convert it to bytes by importing the golang `crypto/ecdsa` package and using the `FromECDSA` method.

```go
privateKeyBytes := crypto.FromECDSA(privateKey)
```

We can now convert it to a hexadecimal string by using the go-ethereum `hexutil` package which provides the `Encode` method which takes a byte slice. Then we strip off the `0x` after it's hex encoded.

```go
fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19
```

This is the private key which is used for signing transactions and is to be treated like a password and never be shared, since who ever is in possesion of it will have access to all your funds.

Since the public key is derived from the private key, go-ethereum's crypto private key has a `Public` method that will return the public key.

```go
publicKey := privateKey.Public()
```

Converting it to hex is a similar process that we went through with the private key. We strip off the `0x` and the first 2 characters `04` which is always the EC prefix and is not required.

```go
publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
if !ok {
  log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
}

publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
fmt.Println(hexutil.Encode(publicKeyBytes)[4:]) // 9a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05
```

Now that we have the public key we can easily generate the public address which is what you're used to seeing. In order to do that, the go-ethereum crypto package has a `PubkeyToAddress` method which accepts an ECDSA public key, and returns the public address.

```go
address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
fmt.Println(address) // 0x96216849c49358B10257cb55b28eA603c874b05E
```

The public address is simply the Keccak-256 hash of the public key, and then we take the last 40 characters (20 bytes) and prefix it with `0x`. Here's how you can do it manually using the `crypto/sha3` keccak256 function.

```go
hash := sha3.NewLegacyKeccak256()
hash.Write(publicKeyBytes[1:])
fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 0x96216849c49358b10257cb55b28ea603c874b05e
```

---

### Full code

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
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:]) // 9a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address) // 0x96216849c49358B10257cb55b28eA603c874b05E

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 0x96216849c49358b10257cb55b28ea603c874b05e
}
```
