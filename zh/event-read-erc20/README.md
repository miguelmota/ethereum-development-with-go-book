---
概述: 用Go来读取ERC20代币智能合约的教程。
---

# 读取ERC-20代币的事件日志

首先，创建ERC-20智能合约的事件日志的interface文件 `erc20.sol`:

```solidity
pragma solidity ^0.4.24;

contract ERC20 {
    event Transfer(address indexed from, address indexed to, uint tokens);
    event Approval(address indexed tokenOwner, address indexed spender, uint tokens);
}
```

然后在给定abi使用`abigen`创建Go包

```bash
solc --abi erc20.sol
abigen --abi=erc20_sol_ERC20.abi --pkg=token --out=erc20.go
```

现在在我们的Go应用程序中，让我们创建与ERC-20事件日志签名类型相匹配的结构类型：

```go
type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}

type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}
```

初始化以太坊客户端

```go
client, err := ethclient.Dial("https://mainnet.infura.io")
if err != nil {
  log.Fatal(err)
}
```

按照ERC-20智能合约地址和所需的块范围创建一个“FilterQuery”。这个例子我们会用[ZRX](https://etherscan.io/token/0xe41d2489571d322189246dafa5ebde1f4699f498) 代币:

```go
// 0x Protocol (ZRX) token address
contractAddress := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
query := ethereum.FilterQuery{
  FromBlock: big.NewInt(6383820),
  ToBlock:   big.NewInt(6383840),
  Addresses: []common.Address{
    contractAddress,
  },
}
```

用`FilterLogs`来过滤日志：

```go
logs, err := client.FilterLogs(context.Background(), query)
if err != nil {
  log.Fatal(err)
}
```

接下来我们将解析JSON abi，稍后我们将使用解压缩原始日志数据：

```go
contractAbi, err := abi.JSON(strings.NewReader(string(token.TokenABI)))
if err != nil {
  log.Fatal(err)
}
```

为了按某种日志类型进行过滤，我们需要弄清楚每个事件日志函数签名的keccak256哈希值。 事件日志函数签名哈希始终是`topic [0]`，我们很快就会看到。 以下是使用go-ethereum`crypto`包计算keccak256哈希的方法：


```go
logTransferSig := []byte("Transfer(address,address,uint256)")
LogApprovalSig := []byte("Approval(address,address,uint256)")
logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)
```

现在我们将遍历所有日志并设置switch语句以按事件日志类型进行过滤：

```go
for _, vLog := range logs {
  fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
  fmt.Printf("Log Index: %d\n", vLog.Index)

  switch vLog.Topics[0].Hex() {
  case logTransferSigHash.Hex():
    //
  case logApprovalSigHash.Hex():
    //
  }
}
```

现在要解析`Transfer`事件日志，我们将使用`abi.Unpack`将原始日志数据解析为我们的日志类型结构。 解包不会解析`indexed`事件类型，因为它们存储在`topics`下，所以对于那些我们必须单独解析，如下例所示：

```go
fmt.Printf("Log Name: Transfer\n")

var transferEvent LogTransfer

err := contractAbi.Unpack(&transferEvent, "Transfer", vLog.Data)
if err != nil {
  log.Fatal(err)
}

transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

fmt.Printf("From: %s\n", transferEvent.From.Hex())
fmt.Printf("To: %s\n", transferEvent.To.Hex())
fmt.Printf("Tokens: %s\n", transferEvent.Tokens.String())
```

`Approval` 日志也是类似的方法：

```go
fmt.Printf("Log Name: Approval\n")

var approvalEvent LogApproval

err := contractAbi.Unpack(&approvalEvent, "Approval", vLog.Data)
if err != nil {
  log.Fatal(err)
}

approvalEvent.TokenOwner = common.HexToAddress(vLog.Topics[1].Hex())
approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())

fmt.Printf("Token Owner: %s\n", approvalEvent.TokenOwner.Hex())
fmt.Printf("Spender: %s\n", approvalEvent.Spender.Hex())
fmt.Printf("Tokens: %s\n", approvalEvent.Tokens.String())
```

最后，把所有的步骤放一起：

```bash
Log Block Number: 6383829
Log Index: 20
Log Name: Transfer
From: 0xd03dB9CF89A9b1f856a8E1650cFD78FAF2338eB2
To: 0x924CD9b60F4173DCDd5254ddD38C4F9CAB68FE6b
Tokens: 2804000000000000000000


Log Block Number: 6383831
Log Index: 62
Log Name: Approval
Token Owner: 0xDD3b9186Da521AbE707B48B8f805Fb3Cd5EEe0EE
Spender: 0xCf67d7A481CEEca0a77f658991A00366FED558F7
Tokens: 10000000000000000000000000000000000000000000000000000000000000000


Log Block Number: 6383838
Log Index: 13
Log Name: Transfer
From: 0xBA826fEc90CEFdf6706858E5FbaFcb27A290Fbe0
To: 0x4aEE792A88eDDA29932254099b9d1e06D537883f
Tokens: 2863452144424379687066
```

我们可以把解析的日志与etherscan的数据对比: [https://etherscan.io/tx/0x0c3b6cf604275c7e44dc7db400428c1a39f33f0c6cbc19ff625f6057a5cb32c0#eventlog](https://etherscan.io/tx/0x0c3b6cf604275c7e44dc7db400428c1a39f33f0c6cbc19ff625f6057a5cb32c0#eventlog)

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
    event Transfer(address indexed from, address indexed to, uint tokens);
    event Approval(address indexed tokenOwner, address indexed spender, uint tokens);
}
```

[event_read_erc20.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/event_read_erc20.go)

```go
package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	token "./contracts_erc20" // for demo
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// LogTransfer ..
type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}

// LogApproval ..
type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	// 0x Protocol (ZRX) token address
	contractAddress := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(6383820),
		ToBlock:   big.NewInt(6383840),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(token.TokenABI)))
	if err != nil {
		log.Fatal(err)
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	LogApprovalSig := []byte("Approval(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)

	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)

		switch vLog.Topics[0].Hex() {
		case logTransferSigHash.Hex():
			fmt.Printf("Log Name: Transfer\n")

			var transferEvent LogTransfer

			err := contractAbi.Unpack(&transferEvent, "Transfer", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("From: %s\n", transferEvent.From.Hex())
			fmt.Printf("To: %s\n", transferEvent.To.Hex())
			fmt.Printf("Tokens: %s\n", transferEvent.Tokens.String())

		case logApprovalSigHash.Hex():
			fmt.Printf("Log Name: Approval\n")

			var approvalEvent LogApproval

			err := contractAbi.Unpack(&approvalEvent, "Approval", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			approvalEvent.TokenOwner = common.HexToAddress(vLog.Topics[1].Hex())
			approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("Token Owner: %s\n", approvalEvent.TokenOwner.Hex())
			fmt.Printf("Spender: %s\n", approvalEvent.Spender.Hex())
			fmt.Printf("Tokens: %s\n", approvalEvent.Tokens.String())
		}

		fmt.Printf("\n\n")
	}
}
```

solc version used for these examples

```bash
$ solc --version
0.4.24+commit.e67f0147.Emscripten.clang
```
