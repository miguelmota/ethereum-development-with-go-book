---
description: Tutorial on how to transfer ETH to another wallet or smart contract with Go.
---

# Transferring ETH

In this lesson you'll learn how to transfer ETH from one account to another account. If you're already familar with Ethereum then you know that a transaction consists of the amount of ether you're transferring, the gas limit, the gas price, a nonce, the receiving address, and optionally data. The transaction must be signed with the private key of the sender before it's broadcasted to the network.

Assuming you've already connected a client, the next step is to load your private key.

```go
privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
if err != nil {
  log.Fatal(err)
}
```

Afterwards we need to get the account nonce. Every transaction requires a nonce. A nonce by definition is a number that is only used once. If it's a new account sending out a transaction then the nonce will be `0`. Every new transaction from an account must have a nonce that the previous nonce incremented by 1. It's hard to keep manual track of all the nonces so the ethereum client provides a helper method `PendingNonceAt` that will return the next nonce you should use.

The function requires the public address of the account we're sending from -- which we can derive from the private key.

```go
publicKey := privateKey.Public()
publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
if !ok {
  log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
}

fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
```

Here, `privateKey.Public()` returns an interface that contains our public key. We perform a type assertion with `publicKey.(<expectedType>)` to explictly set the type of our `publicKey` variable, and assign it to `publicKeyECDSA`. This allows us to use it where our program expects an input of type `*ecdsa.PublicKey`.

Now we can read the nonce that we should use for the account's transaction.

```go
nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
if err != nil {
  log.Fatal(err)
}
```

The next step is to set the amount of ETH that we'll be transferring. However we must convert ether to wei since that's what the Ethereum blockchain uses. Ether supports up to 18 decimal places so 1 ETH is 1 plus 18 zeros. Here's a little tool to help you convert between ETH and wei: [https://etherconverter.netlify.com](https://etherconverter.netlify.com)

```go
value := big.NewInt(1000000000000000000) // in wei (1 eth)
```

The gas limit for a standard ETH transfer is `21000` units.

```go
gasLimit := uint64(21000) // in units
```

The gas price must be set in wei. At the time of this writing, a gas price that will get your transaction included pretty fast in a block is 30 gwei.

```go
gasPrice := big.NewInt(30000000000) // in wei (30 gwei)
```

However, gas prices are always fluctuating based on market demand and what users are willing to pay, so hardcoding a gas price is sometimes not ideal. The go-ethereum client provides the `SuggestGasPrice` function for getting the average gas price based on `x` number of previous blocks.

```go
gasPrice, err := client.SuggestGasPrice(context.Background())
if err != nil {
  log.Fatal(err)
}
```

We figure out who we're sending the ETH to.

```go
toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
```

Now we can finally generate our unsigned ethereum transaction by importing the go-ethereum `core/types` package and invoking `NewTransaction` which takes in the nonce, to address, value, gas limit, gas price, and optional data. The data field is `nil` for just sending ETH. We'll be using the data field when it comes to interacting with smart contracts.

```go
tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)
```

The next step is to sign the transaction with the private key of the sender. To do this we call the `SignTx` method that takes in the unsigned transaction and the private key that we constructed earlier. The `SignTx` method requires the EIP155 signer, which we derive the chain ID from the client.

```go
chainID, err := client.NetworkID(context.Background())
if err != nil {
  log.Fatal(err)
}

signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
if err != nil {
  log.Fatal(err)
}
```

Now we are finally ready to broadcast the transaction to the entire network by calling `SendTransaction` on the client which takes in the signed transaction.

```go
err = client.SendTransaction(context.Background(), signedTx)
if err != nil {
  log.Fatal(err)
}

fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // tx sent: 0x77006fcb3938f648e2cc65bafd27dec30b9bfbe9df41f78498b9c8b7322a249e
```

Afterwards you can check the progress on a block explorer such as Etherscan: [https://rinkeby.etherscan.io/tx/0x77006fcb3938f648e2cc65bafd27dec30b9bfbe9df41f78498b9c8b7322a249e](https://rinkeby.etherscan.io/tx/0x77006fcb3938f648e2cc65bafd27dec30b9bfbe9df41f78498b9c8b7322a249e)

---

### Full code

[transfer_eth.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/transfer_eth.go)

```go
package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)                // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
```
