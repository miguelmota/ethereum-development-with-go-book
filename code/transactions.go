package main

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io")

	if err != nil {
		log.Fatal(err)
	}
}

/*
	for idx := uint(0); idx < count; idx++ {
		tx, err := s.client.TransactionInBlock(context.Background(), block.ParentHash(), idx)
		if err != nil {
			log.Debugf("error TransactionInBlock: %s", err)
			continue
		}
		if tx == nil {
			log.Error("TransactionInBlock tx is nil")
			continue
		}
		receipt, err := s.client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Error(err)
			continue
		}
	}
*/
