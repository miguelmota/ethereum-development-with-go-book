---
概述: 用Go从Swarm下载文件的教程。
---

# 从Swarm下载文件

在[上个章节](../swarm-upload) 我们将一个hello.txt文件上传到swarm，作为返回值，我们得到了一个内容清单哈希。

```go
manifestHash := "f9192507e2e8e118bfedac428c3aa1dec4ae156e954128ec5fb27f63ee67bcac"
```

让我们首先通过调用“DownloadManfest”来下载它，并检查清单的内容。


```go
manifest, isEncrypted, err := client.DownloadManifest(manifestHash)
if err != nil {
  log.Fatal(err)
}
```

我们可以遍历清单条目，看看内容类型，大小和内容哈希是什么。

```go
for _, entry := range manifest.Entries {
  fmt.Println(entry.Hash)        // 42179060941352ba7b400b16c40f1e1290423a826de2a70587034dc14bc4ab2f
  fmt.Println(entry.ContentType) // text/plain; charset=utf-8
  fmt.Println(entry.Path)        // ""
}
```

如果您熟悉swarm url，它们的格式为`bzz：/ <hash> / <path>`，因此为了下载文件，我们指定了清单哈希和路径。 在这个例子里，路径是一个空字符串。 我们将这些数据传递给`Download`函数并返回一个文件对象。


```go
file, err := client.Download(manifestHash, "")
if err != nil {
  log.Fatal(err)
}
```

我们现在可以阅读并打印返回的文件阅读器的内容。

```go
content, err := ioutil.ReadAll(file)
if err != nil {
  log.Fatal(err)
}

fmt.Println(string(content)) // hello world
```

正如预期的那样，它记录了我们原始文件所包含的 *hello world*。

---

### 完整代码

Commands

```bash
geth account new
export BZZKEY=970ef9790b54425bea2c02e25cab01e48cf92573
swarm --bzzaccount $BZZKEY
```

[swarm_download.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/swarm_download.go)

```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"

	bzzclient "github.com/ethereum/go-ethereum/swarm/api/client"
)

func main() {
	client := bzzclient.NewClient("http://127.0.0.1:8500")
	manifestHash := "2e0849490b62e706a5f1cb8e7219db7b01677f2a859bac4b5f522afd2a5f02c0"
	manifest, isEncrypted, err := client.DownloadManifest(manifestHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(isEncrypted) // false

	for _, entry := range manifest.Entries {
		fmt.Println(entry.Hash)        // 42179060941352ba7b400b16c40f1e1290423a826de2a70587034dc14bc4ab2f
		fmt.Println(entry.ContentType) // text/plain; charset=utf-8
		fmt.Println(entry.Size)        // 12
		fmt.Println(entry.Path)        // ""
	}

	file, err := client.Download(manifestHash, "")
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content)) // hello world
}
```
