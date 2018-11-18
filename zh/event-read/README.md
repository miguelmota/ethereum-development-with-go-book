---
概述: 用Go读取智能合约事件教程。
---

# 读取事件日志

智能合约可以可选地释放“事件”，其作为交易收据的一部分存储日志。读取这些事件相当简单。首先我们需要构造一个过滤查询。我们从go-ethereum包中导入`FilterQuery`结构体并用过滤选项初始化它。我们告诉它我们想过滤的区块范围并指定从中读取此日志的合约地址。在示例中，我们将从在[智能合约章节]((../smart-contract-compile))创建的智能合约中读取特定区块所有日志。

```go
query := ethereum.FilterQuery{
  FromBlock: big.NewInt(2394201),
  ToBlock:   big.NewInt(2394201),
  Addresses: []common.Address{
    contractAddress,
  },
}
```

下一步是调用ethclient的`FilterLogs`，它接收我们的查询并将返回所有的匹配事件日志。

```go
logs, err := client.FilterLogs(context.Background(), query)
if err != nil {
  log.Fatal(err)
}
```

返回的所有日志将是ABI编码，因此它们本身不会非常易读。为了解码日志，我们需要导入我们智能合约的ABI。为此，我们导入编译好的智能合约Go包，它将包含名称格式为`<Contract>ABI`的外部属性。之后，我们使用go-ethereum中的`accounts/abi`包的`abi.JSON`函数返回一个我们可以在Go应用程序中使用的解析过的ABI接口。

```go
contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
if err != nil {
  log.Fatal(err)
}
```

现在我们可以通过日志进行迭代并将它们解码为我么可以使用的类型。若您回忆起我们的样例合约释放的日志在Solidity中是类型为`bytes32`，那么Go中的等价物将是`[32]byte`。我们可以使用这些类型创建一个匿名结构体，并将指针作为第一个参数传递给解析后的ABI接口的`Unpack`函数，以解码原始的日志数据。第二个参数是我们尝试解码的事件名称，最后一个参数是编码的日志数据。

```go
for _, vLog := range logs {
  event := struct {
    Key   [32]byte
    Value [32]byte
  }{}
  err := contractAbi.Unpack(&event, "ItemSet", vLog.Data)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(string(event.Key[:]))   // foo
  fmt.Println(string(event.Value[:])) // bar
}
```

此外，日志结构体包含附加信息，例如，区块摘要，区块号和交易摘要。

```go
fmt.Println(vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
fmt.Println(vLog.BlockNumber)     // 2394201
fmt.Println(vLog.TxHash.Hex())    // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6
```

### 主题(Topics)

若您的solidity事件包含`indexed`事件类型，那么它们将成为*主题*而不是日志的数据属性的一部分。在solidity中您最多只能有4个主题，但只有3个可索引的事件类型。第一个主题总是事件的签名。我们的示例合约不包含可索引的事件，但如果它确实包含，这是如何读取事件主题。

```go
var topics [4]string
for i := range vLog.Topics {
  topics[i] = vLog.Topics[i].Hex()
}

fmt.Println(topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
```

正如您所见，首个主题只是被哈希过的事件签名。

```go
eventSignature := []byte("ItemSet(bytes32,bytes32)")
hash := crypto.Keccak256Hash(eventSignature)
fmt.Println(hash.Hex()) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
```

这就是阅读和解析日志的全部内容。要学习如何订阅日志，阅读[上个章节]((../event-subscribe))。

### 完整代码

命令

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

[event_read.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/event_read.go)

```go
package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	store "./contracts" // for demo
)

func main() {
	client, err := ethclient.Dial("wss://rinkeby.infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(2394201),
		ToBlock:   big.NewInt(2394201),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		fmt.Println(vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
		fmt.Println(vLog.BlockNumber)     // 2394201
		fmt.Println(vLog.TxHash.Hex())    // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6

		event := struct {
			Key   [32]byte
			Value [32]byte
		}{}
		err := contractAbi.Unpack(&event, "ItemSet", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(event.Key[:]))   // foo
		fmt.Println(string(event.Value[:])) // bar

		var topics [4]string
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}

		fmt.Println(topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
	}

	eventSignature := []byte("ItemSet(bytes32,bytes32)")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Println(hash.Hex()) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
}
```

```bash
$ solc --version
0.4.24+commit.e67f0147.Emscripten.clang
```
