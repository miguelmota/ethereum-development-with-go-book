---
description: Tutorial on creating and importing a keystore with Go.
---

# Keystores

A keystore is a file containing an encrypted wallet private key. Keystores in go-ethereum can only contain one wallet key pair per file. To generate keystores first you must invoke `NewKeyStore` giving it the directory path to save the keystores. After that, you may generate a new wallet by calling the method `NewAccount` passing it a password for encryption. Every time you call `NewAccount` it will generate a new keystore file on disk.

Here's a full example of generating a new keystore account.

```go
ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
password := "secret"
account, err := ks.NewAccount(password)
if err != nil {
  log.Fatal(err)
}

fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3
```

Now to import your keystore you basically need to invoke `NewKeyStore` again as usual and then call the `Import` method which accepts the keystore JSON data as bytes. The second argument is the password used to encrypt it in order to decrypt it. The third argument is to specify a new encryption password but we'll use the same one in the example. Importing the account will give you access to the account as expected but it'll generate a new keystore file! There's no point in having two of the same thing so we'll delete the old one.

Here's an example of importing a keystore and accessing the account.

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

### Full code

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
