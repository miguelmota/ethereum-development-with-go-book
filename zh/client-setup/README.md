---
概述: 用Go初始化客户端以连接以太坊的教程
---

# 初始化客户端

用Go初始化以太坊客户端是和区块链交互所需的基本步骤。首先，导入go-etherem的`ethclient`包并通过调用接收区块链服务提供者URL的`Dial`来初始化它。

若您没有现有以太坊客户端，您可以连接到infura网关。Infura管理着一批安全，可靠，可扩展的以太坊[geth和parity]节点，并且在接入以太坊网络时降低了新人的入门门槛。

```go
client, err := ethclient.Dial("https://mainnet.infura.io")
```

若您运行了本地geth实例，您还可以将路径传递给IPC端点文件。

```go
client, err := ethclient.Dial("/home/user/.ethereum/geth.ipc")
```

对每个Go以太坊项目，使用ethclient是您开始的必要事项，您将在本书中非常多的看到这一步骤。

## 使用Ganache

[Ganache](https://github.com/trufflesuite/ganache-cli)(正式名称为testrpc)是一个用Node.js编写的以太坊实现，用于在本地开发去中心化应用程序时进行测试。现在我们将带着您完成安装并连接到它。

首先通过[NPM](https://www.npmjs.com/package/ganache-cli)安装ganache。

```bash
npm install -g ganache-cli
```

然后运行ganache cli客户端。

```bash
ganache-cli
```

现在连到`http://localhost:8584`上的ganache RPC主机。

```go
client, err := ethclient.Dial("http://localhost:8545")
if err != nil {
  log.Fatal(err)
}
```

在启动ganache时，您还可以使用相同的助记词来生成相同序列的公开地址。

```bash
ganache-cli -m "much repair shock carbon improve miss forget sock include bullet interest solution"
```

我强烈推荐您通过阅读其[文档](http://truffleframework.com/ganache/)熟悉ganache。

---

### 完整代码

[client.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/client.go)

```go
package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections
}
```
