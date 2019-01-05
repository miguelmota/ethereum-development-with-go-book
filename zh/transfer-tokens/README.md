---
概述: 用Go向其他钱包账号或者智能合约发送ERC-20代币的教程。
---

# 代币的转账

本节将向你介绍如何转移ERC-20代币。了解如何转移非ERC-20兼容的其他类型的代币请查阅[智能合约的章节](../smart-contracts) 来了解如何与智能合约交互。

假设您已连接客户端，加载私钥并配置燃气价格，下一步是设置具体的交易数据字段。 如果你完全不明白我刚讲的这些，请先复习 [以太币转账的章节](../transfer-eth)。

代币传输不需要传输ETH，因此将交易“值”设置为“0”。

```go
value := big.NewInt(0)
```

先将您要发送代币的地址存储在变量中。


```go
toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
```

现在轮到有趣的部分。 我们需要弄清楚交易的 *data* 部分。 这意味着我们需要找出我们将要调用的智能合约函数名，以及函数将接收的输入。 然后我们使用函数名的keccak-256哈希来检索 *方法ID*，它是前8个字符（4个字节）。 然后，我们附加我们发送的地址，并附加我们打算转账的代币数量。 这些输入需要256位长（32字节）并填充左侧。 方法ID不需填充。

为了演示，我创造了一个新的代币(HelloToken HTN)，这个可以用代币工厂服务来完成[https://tokenfactory.surge.sh](https://tokenfactory.surge.sh/), 代币我部署到了Rinkeby测试网。

让我们将代币合约地址分配给变量。

```go
tokenAddress := common.HexToAddress("0x28b149020d2152179873ec60bed6bf7cd705775d")
```

函数名将是传递函数的名称，即ERC-20规范中的`transfer`和参数类型。 第一个参数类型是`address`（令牌的接收者），第二个类型是`uint256`（要发送的代币数量）。 不需要没有空格和参数名称。 我们还需要用字节切片格式。

```go
transferFnSignature := []byte("transfer(address,uint256)")
```
我们现在将从go-ethereum导入`crypto/sha3`包以生成函数签名的Keccak256哈希。 然后我们只使用前4个字节来获取方法ID。

```go
hash := sha3.NewKeccak256()
hash.Write(transferFnSignature)
methodID := hash.Sum(nil)[:4]
fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb
```

接下来，我们需要将给我们发送代币的地址左填充到32字节。

```go
paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d
```

接下来我们确定要发送多少个代币，在这个例子里是1,000个，并且我们需要在`big.Int`中格式化为wei。

```go
amount := new(big.Int)
amount.SetString("1000000000000000000000", 10) // 1000 tokens
```

代币量也需要左填充到32个字节。

```go
paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
fmt.Println(hexutil.Encode(paddedAmount))  // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000
```

接下来我们只需将方法ID，填充后的地址和填后的转账量，接到将成为我们数据字段的字节片。


```go
var data []byte
data = append(data, methodID...)
data = append(data, paddedAddress...)
data = append(data, paddedAmount...)
```

燃气上限制将取决于交易数据的大小和智能合约必须执行的计算步骤。 幸运的是，客户端提供了`EstimateGas`方法，它可以为我们估算所需的燃气量。 这个函数从`ethereum`包中获取`CallMsg`结构，我们在其中指定数据和地址。 它将返回我们估算的完成交易所需的估计燃气上限。


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

接下来我们需要做的是构建交易事务类型，这类似于您在ETH转账部分中看到的，除了*to*字段将是代币智能合约地址。 这个常让人困惑。我们还必须在调用中包含0 ETH的值字段和刚刚生成的数据字节。

```go
tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
```

下一步是使用发件人的私钥对事务进行签名。 `SignTx`方法需要EIP155签名器(EIP155 signer)，这需要我们从客户端拿到链ID。

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

最后广播交易。

```go
err = client.SendTransaction(context.Background(), signedTx)
if err != nil {
  log.Fatal(err)
}

fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // tx sent: 0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc
```

你可以去Etherscan看交易的确认过程: [https://rinkeby.etherscan.io/tx/0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc](https://rinkeby.etherscan.io/tx/0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc)

要了解更多如何加载ERC20智能合约并与之互动的内容，可以查看[ERC20代币的智能合约章节](../smart-contract-read-erc20).

---

### 完整代码

[transfer_tokens.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/transfer_tokens.go)

```go
package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
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
