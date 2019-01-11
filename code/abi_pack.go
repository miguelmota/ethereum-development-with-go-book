package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	store "./contracts" // for demo
)

func main() {
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	abiContract, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
	if err != nil {
		panic(err)
	}

	var key [32]byte
	copy(key[:], []byte("foo"))
	iargs := []interface{}{key}

	packed, err := abiContract.Pack("items", iargs...)
	if err != nil {
		panic(err)
	}

	fmt.Println(packed) // [72 243 67 243 102 111 111 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]

	contractAddress := common.HexToAddress("0x147b8eb97fd247d06c4006d269c90c1908fb5d54")

	msg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: packed,
	}

	output, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		panic(err)
	}

	var item [32]byte
	err = abiContract.Unpack(&item, "items", output)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(item[:])) // "bar"
}
