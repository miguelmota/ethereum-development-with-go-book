package main

import (
	"fmt"
	"log"

	bzzclient "github.com/ethereum/go-ethereum/swarm/api/client"
)

func main() {
	client := bzzclient.NewClient("http://127.0.0.1:8500")

	file, err := bzzclient.Open("hello.txt")
	if err != nil {
		log.Fatal(err)
	}

	manifestHash, err := client.Upload(file, "", false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(manifestHash) // 2e0849490b62e706a5f1cb8e7219db7b01677f2a859bac4b5f522afd2a5f02c0
}
