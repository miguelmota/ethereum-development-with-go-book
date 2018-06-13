---
description: Tutorial on how to transfer ERC-20 tokens to another wallet or smart contract with Go.
---

# Transferring Tokens

This section will walk you through on how to transfer ERC-20 tokens. To learn how to transfer other types tokens that are non-ERC-20 compliant check out the [section on smart contracts](../smart-contracts) to learn how to interact with smart contracts.

Assuming you've already connected a client, loaded your private key, and configured the gas price, the next step is to set the data field of the transaction. If you're not sure about what I just said, check out the [section on transferring ETH](../transfer-eth) first.

Token transfers don't require ETH to be transferred so set the value to `0`.

```go
value := big.NewInt(0)
```

Store the address you'll be sending tokens to in a variable.

```go
toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
```

Now the fun part. We'll need to figure out *data* part of the transaction. This means that we'll need to figure out the signature of the smart contract function we'll be calling, along with the inputs that the function will be receiving. We then take the keccak-256 hash of the function signature to retreive the *method ID* which is the first 8 characters (4 bytes). Afterwards we append the address we're sending to, as well append the amount of tokens we're transferring. These inputs will need to be 256 bits long (32 bytes) and left padded. The method ID is not padded.

For demo purposes I've created a token (HelloToken HTN) using token factory [https://tokenfactory.surge.sh](https://tokenfactory.surge.sh/), and deployed it to the Rinkeby testnet.

Let's assign the token contract address to a variable.

```go
tokenAddress := common.HexToAddress("0x28b149020d2152179873ec60bed6bf7cd705775d")
```

The function signature will be the name of the transfer function, which is `transfer` in the ERC-20 specification, and the argument types. The first argument type is `address` (receiver of the tokens) and the second type is `uint256` (amount of tokens to send). There should be no spaces or argument names. We'll also need it as a byte slice.

```go
transferFnSignature := []byte("transfer(address,uint256)")
```
We'll now import the `crypto/sha3` package from go-ethereum to generate the Keccak256 hash of the function signature. We then take only the first 4 bytes to have the method ID.

```go
hash := sha3.NewKeccak256()
hash.Write(transferFnSignature)
methodID := hash.Sum(nil)[:4]
fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb
```

Next we'll need to left pad 32 bytes the address we're sending tokens to.

```go
paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d
```

Next we determine how many tokens we want to send, in this case it'll be 1,000 tokens which will need to be formatted to wei in a `big.Int`.

```go
amount := new(big.Int)
amount.SetString("1000000000000000000000", 10) // 1000 tokens
```

Left padding to 32 bytes will also be required for the amount.

```go
paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
fmt.Println(hexutil.Encode(paddedAmount))  // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000
```

Now we simply concanate the method ID, padded address, and padded amount to a byte slice that will be our data field.

```go
var data []byte
data = append(data, methodID...)
data = append(data, paddedAddress...)
data = append(data, paddedAmount...)
```

The gas limit will depend on the size of the transaction data and computational steps that the smart contract has to perform. Fortunately the client provides the method `EstimateGas` which is able to esimate the gas for us. This function takes a `CallMsg` struct from the `ethereum` package where we specify the data and to address. It'll return the estimated gas limit units we'll be needing for generating the complete transaction.

```go
gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
  To:   &toAddress,
  Data: data,
})
if err != nil {
  log.Fatal(err)
}

fmt.Println(gasLimit) // 23256
```

Next thing we need to do is generate the transaction type, similar to what you've seen in the transfer ETH section, EXCEPT the *to* field will be the token smart contract address. This is a gotcha that confuses people. We must also include the value field which will be 0 ETH, and the data bytes that we just generated.

```go
tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
```

The next step is to sign the transaction with the private key of the sender.

```go
signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
if err != nil {
  log.Fatal(err)
}
```

And finally broadcast the transaction.

```go
err = client.SendTransaction(context.Background(), signedTx)
if err != nil {
  log.Fatal(err)
}

fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // tx sent: 0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc
```

You can check the progress on Etherscan: [https://rinkeby.etherscan.io/tx/0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc](https://rinkeby.etherscan.io/tx/0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc)

---

### Full code

[transfer_tokens.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/transfer_tokens.go)

```go
package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"
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
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	tokenAddress := common.HexToAddress("0x28b149020d2152179873ec60bed6bf7cd705775d")

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10) // 1000 tokens
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gasLimit) // 23256

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // tx sent: 0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc
}
```
