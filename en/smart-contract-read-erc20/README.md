---
description: Tutorial on how to query an ERC20 token smart contract with Go.
---

# Querying an ERC20 Token Smart Contract

First create an ERC20 smart contract interface. This is just a contract with the function definitions of the functions that you can call.

```solidity
pragma solidity ^0.4.24;

contract ERC20 {
    string public constant name = "";
    string public constant symbol = "";
    uint8 public constant decimals = 0;

    function totalSupply() public constant returns (uint);
    function balanceOf(address tokenOwner) public constant returns (uint balance);
    function allowance(address tokenOwner, address spender) public constant returns (uint remaining);
    function transfer(address to, uint tokens) public returns (bool success);
    function approve(address spender, uint tokens) public returns (bool success);
    function transferFrom(address from, address to, uint tokens) public returns (bool success);

    event Transfer(address indexed from, address indexed to, uint tokens);
    event Approval(address indexed tokenOwner, address indexed spender, uint tokens);
}
```

Then compile the smart contract to the JSON ABI, and create a Go token package out of the ABI using `abigen`.

```
solc --abi erc20.sol
abigen --abi=erc20_sol_ERC20.abi --pkg=token --out=erc20.go
```

Assuming we already have Ethereum client set up as usual, we can now import the new *token* package into our application and instantiate it. In this example we'll be using the [Golem](https://etherscan.io/address/0xa74476443119a942de498590fe1f2454d7d4ac0d) token.

```go
tokenAddress := common.HexToAddress("0xa74476443119A942dE498590Fe1f2454d7D4aC0d")
instance, err := token.NewToken(tokenAddress, client)
if err != nil {
  log.Fatal(err)
}
```

We may now call any ERC20 method that we like. For example, we can query the token balance of a user.

```
address := common.HexToAddress("0x0536806df512d6cdde913cf95c9886f65b1d3462")
bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
if err != nil {
  log.Fatal(err)
}

fmt.Printf("wei: %s\n", bal) // "wei: 74605500647408739782407023"
```

We can also read the public variables of the ERC20 smart contract.

```go
name, err := instance.Name(&bind.CallOpts{})
if err != nil {
  log.Fatal(err)
}

symbol, err := instance.Symbol(&bind.CallOpts{})
if err != nil {
  log.Fatal(err)
}

decimals, err := instance.Decimals(&bind.CallOpts{})
if err != nil {
  log.Fatal(err)
}

fmt.Printf("name: %s\n", name)         // "name: Golem Network"
fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"
```

We can do some simple math to convert the balance into a human readable decimal format.

```go
fbal := new(big.Float)
fbal.SetString(bal.String())
value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
```

See the same information on etherscan: [https://etherscan.io/token/0xa74476443119a942de498590fe1f2454d7d4ac0d?a=0x0536806df512d6cdde913cf95c9886f65b1d3462](https://etherscan.io/token/0xa74476443119a942de498590fe1f2454d7d4ac0d?a=0x0536806df512d6cdde913cf95c9886f65b1d3462)

---

### Full code

Commands

```bash
solc --abi erc20.sol
abigen --abi=erc20_sol_ERC20.abi --pkg=token --out=erc20.go
```

[erc20.sol](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/contracts_erc20/erc20.sol)

```solidity
pragma solidity ^0.4.24;

contract ERC20 {
    string public constant name = "";
    string public constant symbol = "";
    uint8 public constant decimals = 0;

    function totalSupply() public constant returns (uint);
    function balanceOf(address tokenOwner) public constant returns (uint balance);
    function allowance(address tokenOwner, address spender) public constant returns (uint remaining);
    function transfer(address to, uint tokens) public returns (bool success);
    function approve(address spender, uint tokens) public returns (bool success);
    function transferFrom(address from, address to, uint tokens) public returns (bool success);

    event Transfer(address indexed from, address indexed to, uint tokens);
    event Approval(address indexed tokenOwner, address indexed spender, uint tokens);
}
```

[contract_read_erc20.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/contract_read_erc20.go)

```go
package main

import (
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	token "./contracts_erc20" // for demo
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	// Golem (GNT) Address
	tokenAddress := common.HexToAddress("0xa74476443119A942dE498590Fe1f2454d7D4aC0d")
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x0536806df512d6cdde913cf95c9886f65b1d3462")
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}

	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name: %s\n", name)         // "name: Golem Network"
	fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
	fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"

	fmt.Printf("wei: %s\n", bal) // "wei: 74605500647408739782407023"

	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

	fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
}
```

solc version used for these examples

```bash
$ solc --version
0.4.24+commit.e67f0147.Emscripten.clang
```
