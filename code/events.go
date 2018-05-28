package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/miguelmota/go-web3-example/greeter"
)

func main() {
	client, err := ethclient.Dial("wss://rinkeby.infura.io/ws")

	if err != nil {
		log.Fatal(err)
	}

	greeterAddress := "a7b2eb1b9fff7c9625373a6a6d180e36b552fc4c"
	priv := "abcdbcf6bdc3a8e57f311a2b4f513c25b20e3ad4606486d7a927d8074872cefg"

	key, err := crypto.HexToECDSA(priv)

	contractAddress := common.HexToAddress(greeterAddress)
	greeterClient, err := greeter.NewGreeter(contractAddress, client)

	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(key)

	// not sure why I have to set this when using testrpc
	// var nonce int64 = 0
	// auth.Nonce = big.NewInt(nonce)

	tx, err := greeterClient.Greet(auth, "hello")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Pending TX: 0x%x\n", tx.Hash())

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	var ch = make(chan types.Log)
	ctx := context.Background()

	sub, err := client.SubscribeFilterLogs(ctx, query, ch)

	if err != nil {
		log.Println("Subscribe:", err)
		return
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case log := <-ch:
			fmt.Println("Log:", log)
		}
	}

}
