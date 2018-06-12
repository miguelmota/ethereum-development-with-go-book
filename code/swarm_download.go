package main

import (
	"fmt"
	"log"

	bzzclient "github.com/ethereum/go-ethereum/swarm/api/client"
)

func main() {
	client := bzzclient.NewClient("http://127.0.0.1:8500")

	manifestHash := "2e0849490b62e706a5f1cb8e7219db7b01677f2a859bac4b5f522afd2a5f02c0"

	file, err := client.Download(manifest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(file)
}
