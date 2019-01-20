---
description: Tutorial on how to transfer ERC-20 tokens to another wallet or smart contract with Go.
---

# Transferring Tokens (ERC-20)

This section will walk you through on how to transfer ERC-20 tokens. To learn how to transfer other types of tokens that are non-ERC-20 compliant check out the [section on smart contracts](../smart-contracts) to learn how to interact with smart contracts.

To transfer ERC-20 tokens, we'll need to broadcast a transaction to the blockchain just like before, but with a few changed parameters:

- Instead of setting a `value` for the broadcasted transaction, we'll need to embed the value of tokens to transfer in the `data` send in the transaction.
- Construct a contract function call and embed it in the `data` field of the transaction we're broadcasting to the blockchain.

We'll assume that you've already completed the previous [section on transferring ETH](../transfer-eth), and have a Go application that has:

1. Connected a client.
1. Loaded your account private key.
1. Configured the gas price to use for your transaction.

## Creating a Token for testing

You can create a token using the Token Factory [https://tokenfactory.surge.sh](https://tokenfactory.surge.sh/), a website for conveniently deploying ERC-20 token contracts, to follow the examples in this guide.

When you create your ERC-20 Token, be sure to note down the **address of the token contract**.

For demonstration purposes, I've created a token (HelloToken HTN) using the Token Factory and deployed it to the Rinkeby testnet at the token contract address `0x28b149020d2152179873ec60bed6bf7cd705775d`.

You can check it out with a Web3-enabled browser here (make sure to be connected to the Rinkeby testnet in MetaMask): [https://tokenfactory.surge.sh/#/token/0x28b149020d2152179873ec60bed6bf7cd705775d](https://tokenfactory.surge.sh/#/token/0x28b149020d2152179873ec60bed6bf7cd705775d)

## ETH value and destination address

First, we'll set a few variables.

Set the `value` of the transaction to 0.

```go
value := big.NewInt(0)
```

This `value` is the amount of ETH to be transferred for this transaction, which should be `0` since we're transferring ERC-20 Tokens and not ETH. We'll set the value of Tokens to be transferred in the `data` field later.

Then, store the address you'll be sending tokens to in a variable.

```go
toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
```

## Forming the data field

Now the fun part. We'll need to figure out what goes into the `data` field of the transaction. This is the message that we broadcast to the blockchain as part of the transaction.

To make a token transfer, we need to use this data field to invoke a function on the smart contract. For more information on the functions available on an ERC-20 token contract, see the [ERC-20 Token Standard specification](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-20.md).

To transfer tokens from our active account to another, we need to invoke the `transfer()` function in our ERC-20 token in our transactions data field. We do this by doing the following:

1. Figure out the function signature of the `transfer()` smart contract function we'll be calling.
1. Figure out the inputs for the function — the `address` of the token recipients, and the `value` of tokens to be transferred.
1. Get the first 8 characters (4 bytes) of the Keccak256 hash of that function signature. This is the *method ID* of the contract function we're invoking.
1. Zero-pad (on the left) the inputs of our function call — the `address` and `value`. These input values need to be 256-bits (32 bytes) long.

First, let's assign the token contract address to a variable.

```go
tokenAddress := common.HexToAddress("0x28b149020d2152179873ec60bed6bf7cd705775d")
```

Next, we need to form the smart contract function call. The signature of the function we'll be calling is the `transfer()` function in the ERC-20 specification, and the types of the argument we'll be passing to it. The first argument type is `address` (the address to which we're sending tokens), and the second argument's type is `uint256` (the amount of tokens to send). The result is the string `transfer(address,uint256)` (no spaces!).

We need this function signature as a byte slice, which we assign to `transferFnSignature`:

```go
transferFnSignature := []byte("transfer(address,uint256)") // do not include spaces in the string
```

We then need to get the `methodID` of our function. To do this, we'll import the `crypto/sha3` to generate the Keccak256 hash of the function signature. The first 4 bytes of the resulting hash is the `methodID`:

```go
hash := sha3.NewLegacyKeccak256()
hash.Write(transferFnSignature)
methodID := hash.Sum(nil)[:4]
fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb
```

Next we'll zero pad (to the left) the account address we're sending tokens. The resulting byte slice must be 32 bytes long:

```go
paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d
```

Next we'll set the value tokens to send as a `*big.Int` number. Note that the denomination used here is determined by the token contract that you're interacting with, and **not** in ETH or wei.

For example, if we were working with TokenA where 1 token is set as the smallest unit of TokenA (i.e. the `decimal()` value of the token contract is `0`; for more information, see the [ERC-20 Token Standard specification](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-20.md)), then `amount := big.NewInt(1000)` would set `amount` to `1000` units of TokenA.

The example token we're using, HelloToken, uses 18 decimals which is standard practice for ERC-20 tokens. This means that in order to represent 1 token we have to do the calculation _amount * 10^18_. In this example we'll use 1,000 tokens so we'll need to calculate _1000 * 10^18_ which is *1e+21* or *1000000000000000000000*. This is the value the smart contract understands as 1,000 tokens from a user representation.

```go
amount := new(big.Int)
amount.SetString("1000000000000000000000", 10) // sets the value to 1000 tokens, in the token denomination
```

There are utility functions available in the [utils](../util-go) section to easily do these conversions.

Left padding to 32 bytes will also be required for the amount since the EVM use 32 byte wide data structures.

```go
paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
fmt.Println(hexutil.Encode(paddedAmount))  // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000
```

Now we concanate the method ID, padded address, and padded amount into a byte slice that will be our data field.

```go
var data []byte
data = append(data, methodID...)
data = append(data, paddedAddress...)
data = append(data, paddedAmount...)
```

## Set gas limit

The gas limit will depend on the size of the transaction data and computational steps that the smart contract has to perform. Fortunately the client provides the `EstimateGas` method which is able to esimate the gas for us based on the most recent state of the blockchain. This function takes a `CallMsg` struct from the `ethereum` package where we specify the data and the address of the token contract to which we're sending the function call message. It'll return the estimated gas limit units we'll use to generate the complete transaction.

```go
gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
  To:   &tokenAddress,
  Data: data,
})
if err != nil {
  log.Fatal(err)
}

fmt.Println(gasLimit) // 23256
```

**NOTE**: The gas limit set by the `EstimateGas()` method is based on the current state of the blockchain, and is just an *estimate*. If your transactions are constantly failing, or if you prefer to have full control over the amount of gas your application spends, you may want to set this value manually.

## Create transaction

Now we have all the information we need to generate the transaction.

We'll create a transaction similar the one we used in [section on transferring ETH](../transfer-eth), EXCEPT that the *to* field should contain the token smart contract address, and the value field should be set to `0` since we're not transferring ETH. This is a gotcha that confuses people.

```go
tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
```

The next step is to sign the transaction with the private key of the sender. The `SignTx` method requires the EIP155 signer, which we derive the chain ID from the client.

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

And finally, broadcast the transaction:

```go
err = client.SendTransaction(context.Background(), signedTx)
if err != nil {
  log.Fatal(err)
}

fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // tx sent: 0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc
```

You can check the progress on Etherscan: [https://rinkeby.etherscan.io/tx/0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc](https://rinkeby.etherscan.io/tx/0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc)

To learn how to load and interact with an ERC20 smart contract, check out the [section on ERC20 token smart contracts](../smart-contract-read-erc20).

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

	"golang.org/x/crypto/sha3"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	tokenAddress := common.HexToAddress("0x28b149020d2152179873ec60bed6bf7cd705775d")

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10) // sets the value to 1000 tokens, in the token denomination

	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gasLimit) // 23256

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

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

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // tx sent: 0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc
}
```
