---
概述: 用Go来上传文件到Swarm的教程。
---

# 上传文件到Swarm

在[上个章节](../swarm-setup) 我们在端口“8500”上运行了一个作为背景进程的swarm节点。 接下来就导入swarm包go-ethereum`swearm/api/client`。 我将把包装别名为`bzzclient`。


```go
import (
  bzzclient "github.com/ethereum/go-ethereum/swarm/api/client"
)
```

调用`NewClient`函数向它传递swarm背景程序的url。

```go
client := bzzclient.NewClient("http://127.0.0.1:8500")
```

用内容 *hello world* 创建示例文本文件`hello.txt`。 我们将会把这个文件上传到swarm。

```txt
hello world
```

在我们的Go应用程序中，我们将使用Swarm客户端软件包中的“Open”打开我们刚刚创建的文件。 该函数将返回一个`File`类型，它表示swarm清单中的文件，用于上传和下载swarm内容。

```go
file, err := bzzclient.Open("hello.txt")
if err != nil {
  log.Fatal(err)
}
```

现在我们可以从客户端实例调用`Upload`函数，为它提供文件对象。 第二个参数是一个可选添的现有内容清单字符串，用于添加文件，否则它将为我们创建。 第三个参数是我们是否希望我们的数据被加密。

返回的哈希值是文件的内容清单的哈希值，其中包含hello.txt文件作为其唯一条目。 默认情况下，主要内容和清单都会上传。 清单确保您可以使用正确的mime类型检索文件。


```go
manifestHash, err := client.Upload(file, "", false)
if err != nil {
  log.Fatal(err)
}

fmt.Println(manifestHash) // 2e0849490b62e706a5f1cb8e7219db7b01677f2a859bac4b5f522afd2a5f02c0
```

然后我们就可以在这里查看上传的文件 `bzz://2e0849490b62e706a5f1cb8e7219db7b01677f2a859bac4b5f522afd2a5f02c0`，具体如何下载，我们会在[下个章节](../swarm-download)介绍。

---

### 完整代码

Commands

```bash
geth account new
export BZZKEY=970ef9790b54425bea2c02e25cab01e48cf92573
swarm --bzzaccount $BZZKEY
```

[hello.txt](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/hello.txt)

```txt
hello world
```

[swarm_upload.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/swarm_upload.go)

```go
package main

import (
	"fmt"
	"log"

	bzzclient "github.com/ethereum/go-ethereum/swarm/api/client"
)

func main() {
	client := bzzclient.NewClient("http://127.0.0.1:8500")

	file, err := bzzclient.Open("hello.txt")
	if err != nil {
		log.Fatal(err)
	}

	manifestHash, err := client.Upload(file, "", false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(manifestHash) // 2e0849490b62e706a5f1cb8e7219db7b01677f2a859bac4b5f522afd2a5f02c0
}
```
