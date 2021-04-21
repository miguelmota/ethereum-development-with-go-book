---
description: Tutorial on how to check if an address is a smart contract or an account with Go.
---

# Address Check

This section will describe how to validate an address and determine if it's a smart contract address.

## Check if Address is Valid

We can use a simple regular expression to check if the ethereum address is valid:

```go
re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

fmt.Printf("is valid: %v\n", re.MatchString("0x323b5d4c32345ced77393b3530b1eed0f346429d")) // is valid: true
fmt.Printf("is valid: %v\n", re.MatchString("0xZYXb5d4c32345ced77393b3530b1eed0f346429d")) // is valid: false
```

## Check if Address is an Account or a Smart Contract

We can determine if an address is a smart contract if there's bytecode stored at that address. Here's an example where we fetch the code for a token smart contract and check the length to verify that it's a smart contract:

```go
// 0x Protocol Token (ZRX) smart contract address
address := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is latest block
if err != nil {
  log.Fatal(err)
}

isContract := len(bytecode) > 0

fmt.Printf("is contract: %v\n", isContract) // is contract: true
```

When there's no bytecode at the address then we know that it's not a smart contract and it's a standard ethereum account:

```go
// a random user account address
address := common.HexToAddress("0x8e215d06ea7ec1fdb4fc5fd21768f4b34ee92ef4")
bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is latest block
if err != nil {
  log.Fatal(err)
}

isContract = len(bytecode) > 0

fmt.Printf("is contract: %v\n", isContract) // is contract: false
```

---

### Full code

[address_check.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/address_check.go)

```go
package main

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	fmt.Printf("is valid: %v\n", re.MatchString("0x323b5d4c32345ced77393b3530b1eed0f346429d")) // is valid: true
	fmt.Printf("is valid: %v\n", re.MatchString("0xZYXb5d4c32345ced77393b3530b1eed0f346429d")) // is valid: false

	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	// 0x Protocol Token (ZRX) smart contract address
	address := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
	bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	isContract := len(bytecode) > 0

	fmt.Printf("is contract: %v\n", isContract) // is contract: true

	// a random user account address
	address = common.HexToAddress("0x8e215d06ea7ec1fdb4fc5fd21768f4b34ee92ef4")
	bytecode, err = client.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	isContract = len(bytecode) > 0

	fmt.Printf("is contract: %v\n", isContract) // is contract: false
}
```
