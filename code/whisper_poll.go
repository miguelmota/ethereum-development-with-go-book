package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/whisper/shhclient"
	"github.com/ethereum/go-ethereum/whisper/whisperv6"
)

func main() {
	client, err := shhclient.Dial("ws://127.0.0.1:8546")
	if err != nil {
		log.Fatal(err)
	}

	filterID, err := client.NewMessageFilter(context.Background(), whisperv6.Criteria{
		PrivateKeyID: privateKeyID,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(filterID) // 21171f8b4e7ac0d7a1ce0d121b647ce10d4f0293b95d8fba69c5b4e9d0f235a6

	messages := make(chan *whisperv6.Message)
	sub, err := client.SubscribeMessages(context.Background(), whisper6.Criteria{
		PrivateKeyID: privateKeyID,
	}, messages)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case message := <-messages:
				fmt.Printf(string(message.Payload)) // "Hello"
			}
		}
	}()

	/*
		messages, err := client.FilterMessages(context.Background(), filterID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(messages)
		for _, message := range messages {
			fmt.Printf(string(message.Payload)) // "Hello"
		}
	*/
}
