---
description: Tutorial on downloadig files from swarm in Go.
---

# Downloading Files from Swarm

In the [previous section](../swarm-upload) we uploaded a hello.txt file to swarm and in return we got a manifest hash.

```go
manifestHash := "f9192507e2e8e118bfedac428c3aa1dec4ae156e954128ec5fb27f63ee67bcac"
```

Let's inspect the manifest by downloading it first by calling `DownloadManfest`.

```go
manifest, isEncrypted, err := client.DownloadManifest(manifestHash)
if err != nil {
  log.Fatal(err)
}
```

We can iterate over the manifest entries and see what the content-type, size, and content hash are.

```go
for _, entry := range manifest.Entries {
  fmt.Println(entry.Hash)        // 42179060941352ba7b400b16c40f1e1290423a826de2a70587034dc14bc4ab2f
  fmt.Println(entry.ContentType) // text/plain; charset=utf-8
  fmt.Println(entry.Path)        // ""
}
```

If you're familiar with swarm urls, they're in the format `bzz:/<hash>/<path>`, so in order to download the file we specify the manifest hash and path. The path in this case is an empty string. We pass this data to the `Download` function and get back a file object.

```go
file, err := client.Download(manifestHash, "")
if err != nil {
  log.Fatal(err)
}
```

We may now read and print the contents of the returned file reader.

```go
content, err := ioutil.ReadAll(file)
if err != nil {
  log.Fatal(err)
}

fmt.Println(string(content)) // hello world
```

As expected, it logs *hello world* which what our original file contained.

---

### Full code

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

	bzzclient "github.com/ethersphere/swarm/api/client"
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
