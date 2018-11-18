---
概述: 用Go编译智能合约并读取ABI的教程。
---

# 智能合约的编译与ABI

与智能合约交互，我们要先生成相应智能合约的应用二进制接口ABI(application binary interface)，并把ABI编译成我们可以在Go应用中调用的格式。

第一步是安装 [Solidity编译器](https://solidity.readthedocs.io/en/latest/installing-solidity.html) (`solc`).

Solc 在Ubuntu上有snapcraft包。

```bash
sudo snap install solc --edge
```

Solc在macOS上有Homebrew的包。

```bash
brew update
brew tap ethereum/ethereum
brew install solidity
```

其他的平台或者从源码编译的教程请查阅官方solidity文档[install guide](https://solidity.readthedocs.io/en/latest/installing-solidity.html#building-from-source).

我们还得安装一个叫`abigen`的工具，来从solidity智能合约生成ABI。

假设您已经在计算机上设置了Go，只需运行以下命令即可安装`abigen`工具。

```bash
go get -u github.com/ethereum/go-ethereum
cd $GOPATH/src/github.com/ethereum/go-ethereum/
make
make devtools
```

我们将创建一个简单的智能合约来测试。 学习更复杂的智能合约，或者智能合约的开发的内容则超出了本书的范围。 我强烈建议您查看[truffle framework](http://truffleframework.com/) 来学习开发和测试智能合约。

这里只是一个简单的合约，就是一个键/值存储，只有一个外部方法来设置任何人的键/值对。 我们还在设置值后添加了要发出的事件。

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

虽然这个智能合约很简单，但它将适用于这个例子。

现在我们可以从一个solidity文件生成ABI。

```bash
solc --abi Store.sol
```

它会将其写入名为“Store_sol_Store.abi”的文件中

现在让我们用`abigen`将ABI转换为我们可以导入的Go文件。 这个新文件将包含我们可以用来与Go应用程序中的智能合约进行交互的所有可用方法。

```bash
abigen --abi=Store_sol_Store.abi --pkg=store --out=Store.go
```

为了从Go部署智能合约，我们还需要将solidity智能合约编译为EVM字节码。 EVM字节码将在事务的数据字段中发送。 在Go文件上生成部署方法需要bin文件。


```bash
solc --bin Store.sol
```

现在我们编译Go合约文件，其中包括deploy方法，因为我们包含了bin文件。

```bash
abigen --bin=Store_sol_Store.bin --abi=Store_sol_Store.abi --pkg=store --out=Store.go
```

在接下来的课程中，我们将学习如何部署智能合约，然后与之交互。


---

### 完整代码

Commands

```bash
go get -u github.com/ethereum/go-ethereum
cd $GOPATH/src/github.com/ethereum/go-ethereum/
make
make devtools

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

solc version used for these examples

```bash
$ solc --version
0.4.24+commit.e67f0147.Emscripten.clang
```
