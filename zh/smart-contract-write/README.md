---
概述: 用Go向智能合约写入数据的教程。
---

# 写入智能合约

这写章节需要了解如何将智能合约的ABI编译成Go的合约文件。如果你还没看， 前先读[上一个章节](../smart-contract-compile) 。

写入智能合约需要我们用私钥来对交易事务进行签名。

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

我们还需要先查到nonce和燃气价格。


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

接下来，我们创建一个新的keyed transactor，它接收私钥。

```go
auth := bind.NewKeyedTransactor(privateKey)
```

然后我们需要设置keyed transactor的标准交易选项。

```go
auth.Nonce = big.NewInt(int64(nonce))
auth.Value = big.NewInt(0)     // in wei
auth.GasLimit = uint64(300000) // in units
auth.GasPrice = gasPrice
```

现在我们加载一个智能合约的实例。如果你还记得[上个章节](../smart-contract-compile) 我们创建一个名为*Store*的合约，并使用`abigen`工具生成一个Go文件。 要初始化它，我们只需调用合约包的*New*方法，并提供智能合约地址和ethclient，它返回我们可以使用的合约实例。

```go
address := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
instance, err := store.NewStore(address, client)
if err != nil {
  log.Fatal(err)
}
```

我们创建的智能合约有一个名为*SetItem*的外部方法，它接受solidity“bytes32”格式的两个参数（key，value）。 这意味着Go合约包要求我们传递一个长度为32个字节的字节数组。 调用*SetItem*方法需要我们传递我们之前创建的`auth`对象（keyed transactor）。 在幕后，此方法将使用它的参数对此函数调用进行编码，将其设置为事务的`data`属性，并使用私钥对其进行签名。 结果将是一个已签名的事务对象。


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

现在我就可以看到交易已经成功被发送到了以太坊网络了: [https://rinkeby.etherscan.io/tx/0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870](https://rinkeby.etherscan.io/tx/0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870)

要验证键/值是否已设置，我们可以读取智能合约中的值。

```go
result, err := instance.Items(nil, key)
if err != nil {
  log.Fatal(err)
}

fmt.Println(string(result[:])) // "bar"
```

搞定！

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
