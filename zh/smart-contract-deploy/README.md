---
概述: 用Go部署智能合约的教程。
---

# 部署智能合约

如果你还没看之前的章节，请先学习[编译智能合约的章节](../smart-contract-compile)因为这节内容，需要先了解如何将智能合约编译为Go文件。


假设你已经导入从`abigen`生成的新创建的Go包文件，并设置ethclient，加载您的私钥，下一步是创建一个有配置密匙的交易发送器(tansactor)。 首先从go-ethereum导入`accounts/abi/bind`包，然后调用传入私钥的`NewKeyedTransactor`。 然后设置通常的属性，如nonce，燃气价格，燃气上线限制和ETH值。

```go
auth := bind.NewKeyedTransactor(privateKey)
auth.Nonce = big.NewInt(int64(nonce))
auth.Value = big.NewInt(0)     // in wei
auth.GasLimit = uint64(300000) // in units
auth.GasPrice = gasPrice
```

如果你还记得上个章节的内容, 我们创建了一个非常简单的“Store”合约，用于设置和存储键/值对。 生成的Go合约文件提供了部署方法。 部署方法名称始终以单词*Deploy*开头，后跟合约名称，在本例中为*Store*。

deploy函数接受有密匙的事务处理器，ethclient，以及智能合约构造函数可能接受的任何输入参数。我们测试的智能合约接受一个版本号的字符串参数。 此函数将返回新部署的合约地址，事务对象，我们可以交互的合约实例，还有错误（如果有）。

```go
input := "1.0"
address, tx, instance, err := store.DeployStore(auth, client, input)
if err != nil {
  log.Fatal(err)
}

fmt.Println(address.Hex())   // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
fmt.Println(tx.Hash().Hex()) // 0xdae8ba5444eefdc99f4d45cd0c4f24056cba6a02cefbf78066ef9f4188ff7dc0

_ = instance // will be using the instance in the 下个章节
```

就这么简单：）你可以用事务哈希来在Etherscan上查询合约的部署状态: [https://rinkeby.etherscan.io/tx/0xdae8ba5444eefdc99f4d45cd0c4f24056cba6a02cefbf78066ef9f4188ff7dc0](https://rinkeby.etherscan.io/tx/0xdae8ba5444eefdc99f4d45cd0c4f24056cba6a02cefbf78066ef9f4188ff7dc0)

---

### 完整代码

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

[contract_deploy.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/contract_deploy.go)

```go
package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
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

	input := "1.0"
	address, tx, instance, err := store.DeployStore(auth, client, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())   // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
	fmt.Println(tx.Hash().Hex()) // 0xdae8ba5444eefdc99f4d45cd0c4f24056cba6a02cefbf78066ef9f4188ff7dc0

	_ = instance
}
```

solc version used for these examples

```bash
$ solc --version
0.4.24+commit.e67f0147.Emscripten.clang
```
