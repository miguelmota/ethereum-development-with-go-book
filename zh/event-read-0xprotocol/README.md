---
概述: 用Go读取0x Protocol智能合约事件日志的教程
---

# 读取0x Protocol事件日志

要读取[0x Protocol](https://0xproject.com/)事件日志，我们必须首先将solidity智能合约编译为一个Go包。

安装solc版本`0.4.11`

```bash
npm i -g solc@0.4.11
```

为例如`Exchange.sol`的事件日志创建0x Protocol交易所智能合约接口:

Create the 0x protocol exchange smart contract interface for event logs as `Exchange.sol`:

```solidity
pragma solidity 0.4.11;

contract Exchange {
    event LogFill(
        address indexed maker,
        address taker,
        address indexed feeRecipient,
        address makerToken,
        address takerToken,
        uint filledMakerTokenAmount,
        uint filledTakerTokenAmount,
        uint paidMakerFee,
        uint paidTakerFee,
        bytes32 indexed tokens, // keccak256(makerToken, takerToken), allows subscribing to a token pair
        bytes32 orderHash
    );

    event LogCancel(
        address indexed maker,
        address indexed feeRecipient,
        address makerToken,
        address takerToken,
        uint cancelledMakerTokenAmount,
        uint cancelledTakerTokenAmount,
        bytes32 indexed tokens,
        bytes32 orderHash
    );

    event LogError(uint8 indexed errorId, bytes32 indexed orderHash);
}
```

接着给定abi，使用`abigen`来创建Go`exchange`包：

Then use `abigen` to create the Go `exchange` package given the abi:

```bash
solc --abi Exchange.sol
abigen --abi="Exchange.sol:Exchange.abi" --pkg=exchange --out=Exchange.go
```

现在在我们的Go应用程序中，让我们创建与0xProtocol事件日志签名类型匹配的结构体类型：

```go
type LogFill struct {
	Maker                  common.Address
	Taker                  common.Address
	FeeRecipient           common.Address
	MakerToken             common.Address
	TakerToken             common.Address
	FilledMakerTokenAmount *big.Int
	FilledTakerTokenAmount *big.Int
	PaidMakerFee           *big.Int
	PaidTakerFee           *big.Int
	Tokens                 [32]byte
	OrderHash              [32]byte
}

type LogCancel struct {
	Maker                     common.Address
	FeeRecipient              common.Address
	MakerToken                common.Address
	TakerToken                common.Address
	CancelledMakerTokenAmount *big.Int
	CancelledTakerTokenAmount *big.Int
	Tokens                    [32]byte
	OrderHash                 [32]byte
}

type LogError struct {
	ErrorID   uint8
	OrderHash [32]byte
}
```

初始化以太坊客户端：

```go
client, err := ethclient.Dial("https://mainnet.infura.io")
if err != nil {
  log.Fatal(err)
}
```

创建一个`FilterQuery`，并为其传递0x Protocol智能合约地址和所需的区块范围：

```go
// 0x Protocol Exchange smart contract address
contractAddress := common.HexToAddress("0x12459C951127e0c374FF9105DdA097662A027093")
query := ethereum.FilterQuery{
  FromBlock: big.NewInt(6383482),
  ToBlock:   big.NewInt(6383488),
  Addresses: []common.Address{
    contractAddress,
  },
}
```

用`FilterLogs`查询日志：

```go
logs, err := client.FilterLogs(context.Background(), query)
if err != nil {
  log.Fatal(err)
}
```

接下来我们将解析JSON abi，我们后续将使用解压缩原始日志数据：

```go
contractAbi, err := abi.JSON(strings.NewReader(string(exchange.ExchangeABI)))
if err != nil {
  log.Fatal(err)
}
```

为了按某种日志类型过滤，我们需要知晓每个事件日志函数签名的keccak256摘要。正如我们很快所见到的那样，事件日志函数签名摘要总是`topic[0]`：

```go
// NOTE: keccak256("LogFill(address,address,address,address,address,uint256,uint256,uint256,uint256,bytes32,bytes32)")
logFillEvent := common.HexToHash("0d0b9391970d9a25552f37d436d2aae2925e2bfe1b2a923754bada030c498cb3")

// NOTE: keccak256("LogCancel(address,address,address,address,uint256,uint256,bytes32,bytes32)")
logCancelEvent := common.HexToHash("67d66f160bc93d925d05dae1794c90d2d6d6688b29b84ff069398a9b04587131")

// NOTE: keccak256("LogError(uint8,bytes32)")
logErrorEvent := common.HexToHash("36d86c59e00bd73dc19ba3adfe068e4b64ac7e92be35546adeddf1b956a87e90")
```

现在我们迭代所有的日志并设置一个switch语句来按事件日志类型过滤：

```go
for _, vLog := range logs {
  fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
  fmt.Printf("Log Index: %d\n", vLog.Index)

  switch vLog.Topics[0].Hex() {
  case logFillEvent.Hex():
    //
  case logCancelEvent.Hex():
    //
  case logErrorEvent.Hex():
    //
  }
}
```

现在要解析`LogFill`，我们将使用`abi.Unpack`将原始数据类型解析为我们自定义的日志类型结构体。Unpack不会解析`indexed`事件类型，因为这些它们存储在`topics`下，所以对于那些我们必须单独解析，如下例所示：

```go
fmt.Printf("Log Name: LogFill\n")

var fillEvent LogFill

err := contractAbi.Unpack(&fillEvent, "LogFill", vLog.Data)
if err != nil {
  log.Fatal(err)
}

fillEvent.Maker = common.HexToAddress(vLog.Topics[1].Hex())
fillEvent.FeeRecipient = common.HexToAddress(vLog.Topics[2].Hex())
fillEvent.Tokens = vLog.Topics[3]

fmt.Printf("Maker: %s\n", fillEvent.Maker.Hex())
fmt.Printf("Taker: %s\n", fillEvent.Taker.Hex())
fmt.Printf("Fee Recipient: %s\n", fillEvent.FeeRecipient.Hex())
fmt.Printf("Maker Token: %s\n", fillEvent.MakerToken.Hex())
fmt.Printf("Taker Token: %s\n", fillEvent.TakerToken.Hex())
fmt.Printf("Filled Maker Token Amount: %s\n", fillEvent.FilledMakerTokenAmount.String())
fmt.Printf("Filled Taker Token Amount: %s\n", fillEvent.FilledTakerTokenAmount.String())
fmt.Printf("Paid Maker Fee: %s\n", fillEvent.PaidMakerFee.String())
fmt.Printf("Paid Taker Fee: %s\n", fillEvent.PaidTakerFee.String())
fmt.Printf("Tokens: %s\n", hexutil.Encode(fillEvent.Tokens[:]))
fmt.Printf("Order Hash: %s\n", hexutil.Encode(fillEvent.OrderHash[:]))
```

对于`LogCancel`类似：

```go
fmt.Printf("Log Name: LogCancel\n")

var cancelEvent LogCancel

err := contractAbi.Unpack(&cancelEvent, "LogCancel", vLog.Data)
if err != nil {
  log.Fatal(err)
}

cancelEvent.Maker = common.HexToAddress(vLog.Topics[1].Hex())
cancelEvent.FeeRecipient = common.HexToAddress(vLog.Topics[2].Hex())
cancelEvent.Tokens = vLog.Topics[3]

fmt.Printf("Maker: %s\n", cancelEvent.Maker.Hex())
fmt.Printf("Fee Recipient: %s\n", cancelEvent.FeeRecipient.Hex())
fmt.Printf("Maker Token: %s\n", cancelEvent.MakerToken.Hex())
fmt.Printf("Taker Token: %s\n", cancelEvent.TakerToken.Hex())
fmt.Printf("Cancelled Maker Token Amount: %s\n", cancelEvent.CancelledMakerTokenAmount.String())
fmt.Printf("Cancelled Taker Token Amount: %s\n", cancelEvent.CancelledTakerTokenAmount.String())
fmt.Printf("Tokens: %s\n", hexutil.Encode(cancelEvent.Tokens[:]))
fmt.Printf("Order Hash: %s\n", hexutil.Encode(cancelEvent.OrderHash[:]))
```

最后是`LogError`：

```go
fmt.Printf("Log Name: LogError\n")

errorID, err := strconv.ParseInt(vLog.Topics[1].Hex(), 16, 64)
if err != nil {
  log.Fatal(err)
}

errorEvent := &LogError{
  ErrorID:   uint8(errorID),
  OrderHash: vLog.Topics[2],
}

fmt.Printf("Error ID: %d\n", errorEvent.ErrorID)
fmt.Printf("Order Hash: %s\n", hexutil.Encode(errorEvent.OrderHash[:]))
```

将它们放在一起并运行我们将看到以下输出：

```bash
Log Block Number: 6383482
Log Index: 35
Log Name: LogFill
Maker: 0x8dd688660ec0BaBD0B8a2f2DE3232645F73cC5eb
Taker: 0xe269E891A2Ec8585a378882fFA531141205e92E9
Fee Recipient: 0xe269E891A2Ec8585a378882fFA531141205e92E9
Maker Token: 0xD7732e3783b0047aa251928960063f863AD022D8
Taker Token: 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2
Filled Maker Token Amount: 240000000000000000000000
Filled Taker Token Amount: 6930282000000000000
Paid Maker Fee: 0
Paid Taker Fee: 0
Tokens: 0xf08499c9e419ea8c08c4b991f88632593fb36baf4124c62758acb21898711088
Order Hash: 0x306a9a7ecbd9446559a2c650b4cfc16d1fb615aa2b3f4f63078da6d021268440


Log Block Number: 6383482
Log Index: 38
Log Name: LogFill
Maker: 0x04aa059b2e31B5898fAB5aB24761e67E8a196AB8
Taker: 0xe269E891A2Ec8585a378882fFA531141205e92E9
Fee Recipient: 0xe269E891A2Ec8585a378882fFA531141205e92E9
Maker Token: 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2
Taker Token: 0xD7732e3783b0047aa251928960063f863AD022D8
Filled Maker Token Amount: 6941718000000000000
Filled Taker Token Amount: 240000000000000000000000
Paid Maker Fee: 0
Paid Taker Fee: 0
Tokens: 0x97ef123f2b566f36ab1e6f5d462a8079fbe34fa667b4eae67194b3f9cce60f2a
Order Hash: 0xac270e88ce27b6bb78ee5b68ebaef666a77195020a6ab8922834f07bc9e0d524


Log Block Number: 6383488
Log Index: 43
Log Name: LogCancel
Maker: 0x0004E79C978B95974dCa16F56B516bE0c50CC652
Fee Recipient: 0xA258b39954ceF5cB142fd567A46cDdB31a670124
Maker Token: 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2
Taker Token: 0x89d24A6b4CcB1B6fAA2625fE562bDD9a23260359
Cancelled Maker Token Amount: 30000000000000000000
Cancelled Taker Token Amount: 7274848425000000000000
Tokens: 0x9dd48110dcc444fdc242510c09bbbbe21a5975cac061d82f7b843bce061ba391
Order Hash: 0xe43eff38dc27af046bfbd431926926c072bbc7a509d56f6f1a7ae1f5ad7efe4f
```

将解析后的日志输出与etherscan上的内容进行比较：[https://etherscan.io/tx/0xb73a4492c5db1f67930b25ce3869c1e6b9bdbccb239a23b6454925a5bc0e03c5](https://etherscan.io/tx/0xb73a4492c5db1f67930b25ce3869c1e6b9bdbccb239a23b6454925a5bc0e03c5)

---
### 完整代码

命令

```bash
solc --abi Exchange.sol
abigen --abi="Exchange.sol:Exchange.abi" --pkg=exchange --out=Exchange.go
```

[Exchange.sol](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/contracts_0xprotocol/Exchange.sol)

```solidity
pragma solidity 0.4.11;

contract Exchange {
    event LogFill(
        address indexed maker,
        address taker,
        address indexed feeRecipient,
        address makerToken,
        address takerToken,
        uint filledMakerTokenAmount,
        uint filledTakerTokenAmount,
        uint paidMakerFee,
        uint paidTakerFee,
        bytes32 indexed tokens, // keccak256(makerToken, takerToken), allows subscribing to a token pair
        bytes32 orderHash
    );

    event LogCancel(
        address indexed maker,
        address indexed feeRecipient,
        address makerToken,
        address takerToken,
        uint cancelledMakerTokenAmount,
        uint cancelledTakerTokenAmount,
        bytes32 indexed tokens,
        bytes32 orderHash
    );

    event LogError(uint8 indexed errorId, bytes32 indexed orderHash);
}
```

[event_read_0xprotocol.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/event_read.go)

```go
package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"

	exchange "./contracts_0xprotocol" // for demo
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
)

// LogFill ...
type LogFill struct {
	Maker                  common.Address
	Taker                  common.Address
	FeeRecipient           common.Address
	MakerToken             common.Address
	TakerToken             common.Address
	FilledMakerTokenAmount *big.Int
	FilledTakerTokenAmount *big.Int
	PaidMakerFee           *big.Int
	PaidTakerFee           *big.Int
	Tokens                 [32]byte
	OrderHash              [32]byte
}

// LogCancel ...
type LogCancel struct {
	Maker                     common.Address
	FeeRecipient              common.Address
	MakerToken                common.Address
	TakerToken                common.Address
	CancelledMakerTokenAmount *big.Int
	CancelledTakerTokenAmount *big.Int
	Tokens                    [32]byte
	OrderHash                 [32]byte
}

// LogError ...
type LogError struct {
	ErrorID   uint8
	OrderHash [32]byte
}

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	// 0x Protocol Exchange smart contract address
	contractAddress := common.HexToAddress("0x12459C951127e0c374FF9105DdA097662A027093")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(6383482),
		ToBlock:   big.NewInt(6383488),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(exchange.ExchangeABI)))
	if err != nil {
		log.Fatal(err)
	}

	// NOTE: keccak256("LogFill(address,address,address,address,address,uint256,uint256,uint256,uint256,bytes32,bytes32)")
	logFillEvent := common.HexToHash("0d0b9391970d9a25552f37d436d2aae2925e2bfe1b2a923754bada030c498cb3")

	// NOTE: keccak256("LogCancel(address,address,address,address,uint256,uint256,bytes32,bytes32)")
	logCancelEvent := common.HexToHash("67d66f160bc93d925d05dae1794c90d2d6d6688b29b84ff069398a9b04587131")

	// NOTE: keccak256("LogError(uint8,bytes32)")
	logErrorEvent := common.HexToHash("36d86c59e00bd73dc19ba3adfe068e4b64ac7e92be35546adeddf1b956a87e90")

	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)

		switch vLog.Topics[0].Hex() {
		case logFillEvent.Hex():
			fmt.Printf("Log Name: LogFill\n")

			var fillEvent LogFill

			err := contractAbi.Unpack(&fillEvent, "LogFill", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			fillEvent.Maker = common.HexToAddress(vLog.Topics[1].Hex())
			fillEvent.FeeRecipient = common.HexToAddress(vLog.Topics[2].Hex())
			fillEvent.Tokens = vLog.Topics[3]

			fmt.Printf("Maker: %s\n", fillEvent.Maker.Hex())
			fmt.Printf("Taker: %s\n", fillEvent.Taker.Hex())
			fmt.Printf("Fee Recipient: %s\n", fillEvent.FeeRecipient.Hex())
			fmt.Printf("Maker Token: %s\n", fillEvent.MakerToken.Hex())
			fmt.Printf("Taker Token: %s\n", fillEvent.TakerToken.Hex())
			fmt.Printf("Filled Maker Token Amount: %s\n", fillEvent.FilledMakerTokenAmount.String())
			fmt.Printf("Filled Taker Token Amount: %s\n", fillEvent.FilledTakerTokenAmount.String())
			fmt.Printf("Paid Maker Fee: %s\n", fillEvent.PaidMakerFee.String())
			fmt.Printf("Paid Taker Fee: %s\n", fillEvent.PaidTakerFee.String())
			fmt.Printf("Tokens: %s\n", hexutil.Encode(fillEvent.Tokens[:]))
			fmt.Printf("Order Hash: %s\n", hexutil.Encode(fillEvent.OrderHash[:]))

		case logCancelEvent.Hex():
			fmt.Printf("Log Name: LogCancel\n")

			var cancelEvent LogCancel

			err := contractAbi.Unpack(&cancelEvent, "LogCancel", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			cancelEvent.Maker = common.HexToAddress(vLog.Topics[1].Hex())
			cancelEvent.FeeRecipient = common.HexToAddress(vLog.Topics[2].Hex())
			cancelEvent.Tokens = vLog.Topics[3]

			fmt.Printf("Maker: %s\n", cancelEvent.Maker.Hex())
			fmt.Printf("Fee Recipient: %s\n", cancelEvent.FeeRecipient.Hex())
			fmt.Printf("Maker Token: %s\n", cancelEvent.MakerToken.Hex())
			fmt.Printf("Taker Token: %s\n", cancelEvent.TakerToken.Hex())
			fmt.Printf("Cancelled Maker Token Amount: %s\n", cancelEvent.CancelledMakerTokenAmount.String())
			fmt.Printf("Cancelled Taker Token Amount: %s\n", cancelEvent.CancelledTakerTokenAmount.String())
			fmt.Printf("Tokens: %s\n", hexutil.Encode(cancelEvent.Tokens[:]))
			fmt.Printf("Order Hash: %s\n", hexutil.Encode(cancelEvent.OrderHash[:]))

		case logErrorEvent.Hex():
			fmt.Printf("Log Name: LogError\n")

			errorID, err := strconv.ParseInt(vLog.Topics[1].Hex(), 16, 64)
			if err != nil {
				log.Fatal(err)
			}

			errorEvent := &LogError{
				ErrorID:   uint8(errorID),
				OrderHash: vLog.Topics[2],
			}

			fmt.Printf("Error ID: %d\n", errorEvent.ErrorID)
			fmt.Printf("Order Hash: %s\n", hexutil.Encode(errorEvent.OrderHash[:]))
		}

		fmt.Printf("\n\n")
	}
}
```

这些示例使用的solc版本

```bash
$ solc --version
0.4.11+commit.68ef5810.Emscripten.clang
```
