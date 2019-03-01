package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	rawTx := "f86d8202b28477359400825208944592d8f8d7b001e72cb26a73e4fa1806a51ac79d880de0b6b3a7640000802ca05924bde7ef10aa88db9c66dd4f5fb16b46dff2319b9968be983118b57bb50562a001b24b31010004f13d9a26b320845257a6cfc2bf819a3d55e3fc86263c5f0772"

	tx := new(types.Transaction)
	rawTxBytes, err := hex.DecodeString(rawTx)
	rlp.DecodeBytes(rawTxBytes, &tx)

	fmt.Println(tx.Hash().Hex())

	msg, err := tx.AsMessage(types.NewEIP155Signer(tx.ChainId()))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg.From().Hex())
	// 0x96216849c49358B10257cb55b28eA603c874b05E
}
