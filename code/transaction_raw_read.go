package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	txHash := common.HexToHash("0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2")
	tx, _, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	v, r, s := tx.RawSignatureValues()
	R := r.Bytes()
	S := s.Bytes()
	V := byte(v.Uint64() - 27)

	sig := make([]byte, 65)
	copy(sig[32-len(R):32], R)
	copy(sig[64-len(S):64], S)
	sig[64] = V

	hs := types.HomesteadSigner{}
	hash := hs.Hash(tx)

	rawTx, err := rlp.EncodeToBytes(tx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%x\n", rawTx) // f86f8301b0348517bfac7c0083019a289455fe59d8ad77035154ddd0ad0388d09dd4047a8e872386f26fc10000801ba0b5781624f25a362537e79721417af6ac2cbc84cf3b6717cc0a90c2ddc8d5e1eaa07d4860d5690ca93c6d5ff31def3cd08c47d78e446bdf8c6db4d00b629ae86afa

	publicKeyBytes, err := crypto.Ecrecover(hash.Bytes(), sig)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x\n", publicKeyBytes) // 0408e5466792a4710242ec90cdd9a5e228f0c6e6212a2b571bf1eb6fbe344813938f4e1436f4348d71d32294f6e0fd36b876de0b01c4e411404a57a47b661a6f84

	publicKeyECDSA, err := crypto.UnmarshalPubkey(publicKeyBytes)
	if err != nil {
		log.Fatal(err)
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
}
