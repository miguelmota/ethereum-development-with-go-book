---
概述: 用Go来向另外一个钱包地址或者合约转账以太币的教程。
---

# 转账以太币ETH

在本课程中，您将学习如何将ETH从一个帐户转移到另一个帐户。如果您已熟悉以太坊，那么您就知道如何交易包括您打算转账的以太币数量量，燃气限额，燃气价格，一个随机数(nonce)，接收地址以及可选择性的添加的数据。 在广告发送到网络之前，必须使用发送方的私钥对该交易进行签名。

假设您已经连接了客户端，下一步就是加载您的私钥。

```go
privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
if err != nil {
  log.Fatal(err)
}
```

之后我们需要获得帐户的随机数(nonce)。 每笔交易都需要一个nonce。 根据定义，nonce是仅使用一次的数字。 如果是发送交易的新帐户，则该随机数将为“0”。 来自帐户的每个新事务都必须具有前一个nonce增加1的nonce。很难对所有nonce进行手动跟踪，于是ethereum客户端提供一个帮助方法`PendingNonceAt`，它将返回你应该使用的下一个nonce。

该函数需要我们发送的帐户的公共地址 - 这个我们可以从私钥派生。


```go
publicKey := privateKey.Public()
publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
if !ok {
  log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
}

fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
```

接下来我们可以读取我们应该用于帐户交易的随机数。

```go
nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
if err != nil {
  log.Fatal(err)
}
```

下一步是设置我们将要转移的ETH数量。 但是我们必须将ETH以太转换为wei，因为这是以太坊区块链所使用的。 以太网支持最多18个小数位，因此1个ETH为1加18个零。 这里有一个小工具可以帮助您在ETH和wei之间进行转换: [https://etherconverter.netlify.com](https://etherconverter.netlify.com)

```go
value := big.NewInt(1000000000000000000) // in wei (1 eth)
```

ETH转账的燃气应设上限为“21000”单位。

```go
gasLimit := uint64(21000) // in units
```

燃气价格必须以wei为单位设定。 在撰写本文时，将在一个区块中比较快的打包交易的燃气价格为30 gwei。

```go
gasPrice := big.NewInt(30000000000) // in wei (30 gwei)
```

然而，燃气价格总是根据市场需求和用户愿意支付的价格而波动的，因此对燃气价格进行硬编码有时并不理想。 go-ethereum客户端提供`SuggestGasPrice`函数，用于根据'x'个先前块来获得平均燃气价格。

```go
gasPrice, err := client.SuggestGasPrice(context.Background())
if err != nil {
  log.Fatal(err)
}
```

接下来我们弄清楚我们将ETH发送给谁。

```go
toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
```

现在我们最终可以通过导入go-ethereum`core/types`包并调用`NewTransaction`来生成我们的未签名以太坊事务，这个函数需要接收nonce，地址，值，燃气上限值，燃气价格和可选发的数据。 发送ETH的数据字段为“nil”。 在与智能合约进行交互时，我们将使用数据字段，仅仅转账以太币是不需要数据字段的。

```go
tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)
```

下一步是使用发件人的私钥对事务进行签名。 为此，我们调用`SignTx`方法，该方法接受一个未签名的事务和我们之前构造的私钥。 `SignTx`方法需要EIP155签名者，这个也需要我们先从客户端拿到链ID。


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

现在我们终于准备通过在客户端上调用“SendTransaction”来将已签名的事务广播到整个网络。

```go
err = client.SendTransaction(context.Background(), signedTx)
if err != nil {
  log.Fatal(err)
}

fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // tx sent: 0x77006fcb3938f648e2cc65bafd27dec30b9bfbe9df41f78498b9c8b7322a249e
```

然后你可以去Etherscan看交易的确认过程:  [https://rinkeby.etherscan.io/tx/0x77006fcb3938f648e2cc65bafd27dec30b9bfbe9df41f78498b9c8b7322a249e](https://rinkeby.etherscan.io/tx/0x77006fcb3938f648e2cc65bafd27dec30b9bfbe9df41f78498b9c8b7322a249e)

---

### 完整代码

[transfer_eth.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/transfer_eth.go)

```go
package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
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

	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)                // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

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

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
```
