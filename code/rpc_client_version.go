package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	client, err := rpc.DialHTTP("https://mainnet.infura.io")
	if err != nil {
		panic(err)
	}

	var res string
	req := struct {
		To   string `json:"to"`
		Data string `json:"data"`
	}{}

	if err := client.Call(&res, "web3_clientVersion", req, "latest"); err != nil {
		panic(err)
	}

	fmt.Println(res) // Geth/v1.8.22-omnibus-260f7fbd/linux-amd64/go1.11.1
}
