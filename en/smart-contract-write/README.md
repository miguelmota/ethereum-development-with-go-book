---
description: Tutorial on how to write to a smart contract with Go.
---

# Writing to a Smart Contract

These section requires knowledge of how to compile a smart contract's ABI to a Go contract file. If you haven't already gone through it, please [read the section](../smart-contract-compile) first.

Writing to a smart contract requires us to sign the sign transaction with our private key.

```go
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
```

We'll also need to figure the nonce and gas price.

```go
nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
if err != nil {
  log.Fatal(err)
}

gasPrice, err := client.SuggestGasPrice(context.Background())
if err != nil {
  log.Fatal(err)
}
```

Next we create a new keyed transactor which takes in the private key.

```go
auth := bind.NewKeyedTransactor(privateKey)
```

Then we need to set the standard transaction options attached to the keyed transactor.

```go
auth.Nonce = big.NewInt(int64(nonce))
auth.Value = big.NewInt(0)     // in wei
auth.GasLimit = uint64(300000) // in units
auth.GasPrice = gasPrice
```

Now we load an instance of the smart contract. If you recall in the [previous sections](../smart-contract-compile) we create a contract called *Store* and generated a Go package file using the `abigen` tool. To initialize it we just invoke the *New* method of the contract package and give the smart contract address and the ethclient, which returns a contract instance that we can use.


```go
address := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
instance, err := store.NewStore(address, client)
if err != nil {
  log.Fatal(err)
}
```

The smart contract that we created has an external method called *SetItem* which takes in two arguments (key, value) in the from of solidity `bytes32`. This means that the Go contract package requires us to pass a byte array of length 32 bytes. Invoking the *SetItem* method requires us to pass the `auth` object we created earlier. Behind the scenes this method will encode this function call with it's arguments, set it as the `data` property of the transaction, and sign it with the private key. The result will be a signed transaction object.

```go
key := [32]byte{}
value := [32]byte{}
copy(key[:], []byte("foo"))
copy(value[:], []byte("bar"))

tx, err := instance.SetItem(auth, key, value)
if err != nil {
  log.Fatal(err)
}

fmt.Printf("tx sent: %s\n", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
```

We can see now that the transaction has been successfully sent on the network: [https://rinkeby.etherscan.io/tx/0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870](https://rinkeby.etherscan.io/tx/0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870)

To verify that the key/value was set, we read the smart contract mapping value.

```go
result, err := instance.Items(nil, key)
if err != nil {
  log.Fatal(err)
}

fmt.Println(string(result[:])) // "bar"
```

There you have it.

---

### Full code

Commands

```bash
solc --abi Store.sol
solc --bin Store.sol
abigen --bin=Store_sol_Store.bin --abi=Store_sol_Store.abi --pkg=store --out=Store.go
```

[Store.sol](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/contracts/Store.sol)

```solidity
pragma solidity ^0.4.24;

contract Store {
  event ItemSet(bytes32 key, bytes32 value);

  string public version;
  mapping (bytes32 => bytes32) public items;

  constructor(string _version) public {
    version = _version;
  }

  function setItem(bytes32 key, bytes32 value) external {
    items[key] = value;
    emit ItemSet(key, value);
  }
}
```

[contract_write.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/contract_write.go)

```go
package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	store "./contracts" // for demo
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

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("foo"))
	copy(value[:], []byte("bar"))

	tx, err := instance.SetItem(auth, key, value)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

	result, err := instance.Items(nil, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(result[:])) // "bar"
}
```

solc version used for these examples

```bash
$ solc --version
0.4.24+commit.e67f0147.Emscripten.clang
```
