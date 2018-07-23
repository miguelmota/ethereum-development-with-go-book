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
	foo := &simplestruct{
		A: 123,
		B: "hello",
	}

	// encode

	serialized, err := rlp.EncodeToBytes(&foo)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(fmt.Sprintf("%x", serialized)) // c77b8568656c6c6f

	// decode

	bar := new(simplestruct)

	b, err := hex.DecodeString("c77b8568656c6c6f")
	if err != nil {
		log.Fatal(err)
	}

	r := bytes.NewReader(b)
	err = rlp.Decode(r, &bar)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bar.A) // 123
	fmt.Println(bar.B) // "hello"

	// look at tags "-" "tail" "nil"
}
