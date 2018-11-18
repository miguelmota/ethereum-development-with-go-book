---
概述: 用Go来查询ERC20代币智能合约的教程。
---

# 查询ERC20代币智能合约

首先创建一个ERC20智能合约interface。 这只是与您可以调用的函数的函数定义的契约。


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

然后将interface智能合约编译为JSON ABI，并使用`abigen`从ABI创建Go包。

```
solc --abi erc20.sol
abigen --abi=erc20_sol_ERC20.abi --pkg=token --out=erc20.go
```

假设我们已经像往常一样设置了以太坊客户端，我们现在可以将新的*token*包导入我们的应用程序并实例化它。这个例子里我们用[Golem](https://etherscan.io/address/0xa74476443119a942de498590fe1f2454d7d4ac0d) 代币的地址.

```go
tokenAddress := common.HexToAddress("0xa74476443119A942dE498590Fe1f2454d7D4aC0d")
instance, err := token.NewToken(tokenAddress, client)
if err != nil {
  log.Fatal(err)
}
```

我们现在可以调用任何ERC20的方法。 例如，我们可以查询用户的代币余额。

```
address := common.HexToAddress("0x0536806df512d6cdde913cf95c9886f65b1d3462")
bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
if err != nil {
  log.Fatal(err)
}

fmt.Printf("wei: %s\n", bal) // "wei: 74605500647408739782407023"
```

我们还可以读ERC20智能合约的公共变量。

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

我们可以做一些简单的数学运算将余额转换为可读的十进制格式。

```go
fbal := new(big.Float)
fbal.SetString(bal.String())
value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
```

同样的信息也可以在etherscan上查询: [https://etherscan.io/token/0xa74476443119a942de498590fe1f2454d7d4ac0d?a=0x0536806df512d6cdde913cf95c9886f65b1d3462](https://etherscan.io/token/0xa74476443119a942de498590fe1f2454d7d4ac0d?a=0x0536806df512d6cdde913cf95c9886f65b1d3462)

---

### 完整代码

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
