package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	txID := common.HexToHash("0xe330601b05c54116da3b06dd17cf483ab3106fb969989458c8587aac1c34fbf3")
	receipt, err := client.TransactionReceipt(context.Background(), txID)
	if err != nil {
		log.Fatal(err)
	}

	logID := "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	for _, vLog := range receipt.Logs {
		if vLog.Topics[0].Hex() == logID {
			if len(vLog.Topics) > 2 {
				id := new(big.Int)
				id.SetBytes(vLog.Topics[3].Bytes())

				fmt.Println(id.Uint64()) // 1133
			}
		}
	}
}
