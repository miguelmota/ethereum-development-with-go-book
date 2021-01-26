package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	store "github.com/miguelmota/ethereum-development-with-go-book/code/contracts"
)

func main() {
	client, err := ethclient.Dial("http://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")

	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("foo"))
	copy(value[:], []byte("bar"))

	parsed, err := abi.JSON(strings.NewReader(store.StoreABI))
	if err != nil {
		log.Fatal(err)
	}

	encodedData, err := parsed.Pack("setItem", key, value)
	if err != nil {
		log.Fatal(err)
	}

	estimatedGas, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: fromAddress,
		To:   &address,
		Data: encodedData,
		GasPrice: gasPrice,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(estimatedGas)
}
