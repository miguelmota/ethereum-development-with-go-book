package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

// https://stackoverflow.com/questions/53237759/how-to-correctly-send-rpc-call-using-golang-to-get-smart-contract-owner/53260846#53260846

func main() {
	client, err := rpc.DialHTTP("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	type request struct {
		To   string `json:"to"`
		Data string `json:"data"`
	}

	var result string

	req := request{"0xcc13fc627effd6e35d2d2706ea3c4d7396c610ea", "0x8da5cb5b"}
	if err := client.Call(&result, "eth_call", req, "latest"); err != nil {
		log.Fatal(err)
	}

	owner := common.HexToAddress(result)
	fmt.Printf("%s\n", owner.Hex()) // 0x281017b4E914b79371d62518b17693B36c7a221e
}
