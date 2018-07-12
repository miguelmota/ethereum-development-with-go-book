package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	ethtrie "github.com/ethereum/go-ethereum/trie"
)

func main() {
	diskdb := ethdb.NewMemDatabase()
	triedb := ethtrie.NewDatabase(diskdb)
	trie, err := ethtrie.New(common.Hash{}, triedb)
	if err != nil {
		log.Fatal(err)
	}

	trie.Update([]byte("foo"), []byte("bar"))
	value, err := trie.TryGet([]byte("foo"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(value)) // bar
}
