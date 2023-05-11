package subscribingtonewblocks

import (
	"context"
	"fmt"
	getstarted "hi-supergirl/blockchain-with-go-exercise/get-started"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
)

func Subscribe() {
	client := getstarted.CreateConn()

	headers := make(chan *types.Header)

	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex()) // Print the hash id of the new block
		}
	}

}
