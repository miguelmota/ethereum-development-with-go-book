---
description: Tutorial on uploading files to swarm with go.
---

# Uploading Files to Swarm

In the [previous section](../swarm-setup) we setup a swarm node running as a daemon on port `8500`. Now import the swarm package ethersphere `swarm/api/client`. I'll be aliasing the package to `bzzclient`.

```go
import (
  bzzclient "github.com/ethersphere/swarm/api/client"
)
```

Invoke `NewClient` function passing it the swarm daemon url.

```go
client := bzzclient.NewClient("http://127.0.0.1:8500")
```

Create an example text file `hello.txt` with the content *hello world*. We'll be uploading this to swarm.

```txt
hello world
```

In our Go application we'll open the file we just created using `Open` from the client package. This function will return a `File` type which represents a file in a swarm manifest and is used for uploading and downloading content to and from swarm.

```go
file, err := bzzclient.Open("hello.txt")
if err != nil {
  log.Fatal(err)
}
```

Now we can invoke the `Upload` function from our client instance giving it the file object. The second argument is an optional existing manifest string to add the file to, otherwise it'll create on for us. The third argument is if we want our data to be encrypted.

The hash returned is the swarm hash of a manifest that contains the hello.txt file as its only entry. So by default both the primary content and the manifest is uploaded. The manifest makes sure you could retrieve the file with the correct mime type.

```go
manifestHash, err := client.Upload(file, "", false)
if err != nil {
  log.Fatal(err)
}

fmt.Println(manifestHash) // 2e0849490b62e706a5f1cb8e7219db7b01677f2a859bac4b5f522afd2a5f02c0
```

Now we can access our file at `bzz://2e0849490b62e706a5f1cb8e7219db7b01677f2a859bac4b5f522afd2a5f02c0` which learn how to do in the [next section](../swarm-download).

---

### Full code

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

	bzzclient "github.com/ethersphere/swarm/api/client"
)

func main() {
	client := bzzclient.NewClient("http://127.0.0.1:8500")

	file, err := bzzclient.Open("hello.txt")
	if err != nil {
		log.Fatal(err)
	}

	manifestHash, err := client.Upload(file, "", false, false, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(manifestHash) // 2e0849490b62e706a5f1cb8e7219db7b01677f2a859bac4b5f522afd2a5f02c0
}
```
