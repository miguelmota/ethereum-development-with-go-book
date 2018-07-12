package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/rlp"
)

type simplestruct struct {
	A uint
	B string
}

func main() {
	s := new(simplestruct)

	b, err := hex.DecodeString("C50583343434")
	if err != nil {
		log.Fatal(err)
	}

	r := bytes.NewReader(b)
	err = rlp.Decode(r, &s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s.A) // 55
	fmt.Println(s.B) // "444"
}
