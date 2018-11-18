---
概述: 搭建swarm节点的教程。
---

# 搭建 Swarm 节点

要运行swarm，首先需要安装`geth`和`bzzd`，这是swarm背景进程。

```go
go get -d github.com/ethereum/go-ethereum
go install github.com/ethereum/go-ethereum/cmd/geth
go install github.com/ethereum/go-ethereum/cmd/swarm
```

然后我们将生成一个新的geth帐户。


```bash
$ geth account new

Your new account is locked with a password. Please give a password. Do not forget this password.
Passphrase:
Repeat passphrase:
Address: {970ef9790b54425bea2c02e25cab01e48cf92573}
```

将环境变量`BZZKEY`导出，并设定为我们刚刚生成的geth帐户地址。

```bash
export BZZKEY=970ef9790b54425bea2c02e25cab01e48cf92573
```

然后使用设定的帐户运行swarm，并作为我们的swarm帐户。 默认情况下，Swarm将在端口“8500”上运行。

```bash
$ swarm --bzzaccount $BZZKEY
Unlocking swarm account 0x970EF9790B54425BEA2C02e25cAb01E48CF92573 [1/3]
Passphrase:
WARN [06-12|13:11:41] Starting Swarm service
```

现在swarm进程已经可以运行了，那么我们会在[下个章节](../swarm-upload)学习如何上传文件。

---

### 完整代码

Commands

```bash
go get -d github.com/ethereum/go-ethereum
go install github.com/ethereum/go-ethereum/cmd/geth
go install github.com/ethereum/go-ethereum/cmd/swarm
geth account new
export BZZKEY=970ef9790b54425bea2c02e25cab01e48cf92573
swarm --bzzaccount $BZZKEY
```
