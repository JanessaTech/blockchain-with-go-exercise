package queryingblocks

import (
	"context"
	"fmt"
	getstarted "hi-supergirl/blockchain-with-go-exercise/get-started"
	"log"
	"math/big"
)

// before call this method, make sure ganache is ready and there are at least 2 blocks mined
func QueryBlockHeaderInfo() {
	client := getstarted.CreateConn()

	header, err := client.HeaderByNumber(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(header.Number.String())
}

// before call this method, make sure ganache is ready and there are at least 2 blocks mined
func QueryFullBlock() {
	client := getstarted.CreateConn()
	blockNumber := big.NewInt(2)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Int64())
	fmt.Println(block.Time())                // 1683688348
	fmt.Println(block.Difficulty().Uint64()) // 0
	fmt.Println(block.Hash().Hex())          // 0x41fc3367de044a7aa3e7b5a5ceadd756aa7db3c1cd4606ce7ddb00765893c1e2
	fmt.Println(len(block.Transactions()))   // 1
}
