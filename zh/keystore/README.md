---
概述: 用Go创建和导入keystore。
---

# Keystores

keystore是一个包含经过加密了的钱包私钥。go-ethereum中的keystore，每个文件只能包含一个钱包密钥对。要生成keystore，首先您必须调用`NewKeyStore`，给它提供保存keystore的目录路径。然后，您可调用`NewAccount`方法创建新的钱包，并给它传入一个用于加密的口令。您每次调用`NewAccount`，它将在磁盘上生成新的keystore文件。

这是一个完整的生成新的keystore账户的示例。

```go
ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
password := "secret"
account, err := ks.NewAccount(password)
if err != nil {
  log.Fatal(err)
}

fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3
```

现在要导入您的keystore，您基本上像往常一样再次调用`NewKeyStore`，然后调用`Import`方法，该方法接收keystore的JSON数据作为字节。第二个参数是用于加密私钥的口令。第三个参数是指定一个新的加密口令，但我们在示例中使用一样的口令。导入账户将允许您按期访问该账户，但它将生成新keystore文件！有两个相同的事物是没有意义的，所以我们将删除旧的。

这是一个导入keystore和访问账户的示例。

```go
file := "./wallets/UTC--2018-07-04T09-58-30.122808598Z--20f8d42fb0f667f2e53930fed426f225752453b3"
ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
jsonBytes, err := ioutil.ReadFile(file)
if err != nil {
  log.Fatal(err)
}

password := "secret"
account, err := ks.Import(jsonBytes, password, password)
if err != nil {
  log.Fatal(err)
}

fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3

if err := os.Remove(file); err != nil {
  log.Fatal(err)
}
```

----

### 完整代码

[keystore.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/keystore.go)

```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func createKs() {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3
}

func importKs() {
	file := "./tmp/UTC--2018-07-04T09-58-30.122808598Z--20f8d42fb0f667f2e53930fed426f225752453b3"
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	password := "secret"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3

	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}

func main() {
	createKs()
	//importKs()
}
```
