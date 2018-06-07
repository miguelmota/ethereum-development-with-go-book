package main

import (
	"fmt"
	"math/big"

	util "./util" // for demo
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func main() {
	valid := util.IsValidAddress("0x323b5d4c32345ced77393b3530b1eed0f346429d")
	fmt.Println(valid) // true

	zeroed := util.IsZeroAddress("0x0")
	fmt.Println(zeroed) // true

	wei := util.ToWei(0.02, 18)
	fmt.Println(wei) // 20000000000000000

	wei = new(big.Int)
	wei.SetString("20000000000000000", 10)
	eth := util.ToDecimal(wei, 18)
	fmt.Println(eth) // 0.02

	gasLimit := uint64(21000)
	gasPrice := new(big.Int)
	gasPrice.SetString("2000000000", 10)
	gasCost := util.CalcGasCost(gasLimit, gasPrice)
	fmt.Println(gasCost) // 42000000000000

	sig := "0x789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c62621578113ddbb62d509bf6049b8fb544ab06d36f916685a2eb8e57ffadde02301"
	r, s, v := util.SigRSV(sig)
	fmt.Println(hexutil.Encode(r[:])[2:]) // 789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c6
	fmt.Println(hexutil.Encode(s[:])[2:]) // 2621578113ddbb62d509bf6049b8fb544ab06d36f916685a2eb8e57ffadde023
	fmt.Println(v)                        // 28
}
