package queryingtransactions

import (
	"context"
	"fmt"
	getstarted "hi-supergirl/blockchain-with-go-exercise/get-started"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// before call this method, make sure ganache is ready and there are at least 2 blocks mined
func QueryTxInfo() {
	client := getstarted.CreateConn()

	blockNumber := big.NewInt(3)
	block, err := client.BlockByNumber(context.Background(), blockNumber)

	if err != nil {
		log.Fatal(err)
	}

	// 1. by going through all transactions in a block
	log.Println("Query transaction info by iterating all txs in a block")
	for _, tx := range block.Transactions() {
		printTxDetail(tx)
	}

	// 2. by calling TransactionInBlock
	blockHash := common.HexToHash("0x7415088cbcab6f5ae47c0fd522eb71fdbe773ce1903360c821b264877c2e0f9e")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("---------------------------------------------------------------")

	log.Println("Query transaction info by calling TransactionInBlock")
	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}
		printTxDetail(tx)
	}
	log.Println("---------------------------------------------------------------")

	log.Println("Query transaction info by calling TransactionByHash")
	txHash := common.HexToHash("0xa99489a09dbeb616e2570beabd7dead4b2a417f3cab4da5696fef2c69c80bded")
	tx, _, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	printTxDetail(tx)

	log.Println("---------------------------------------------------------------")
}

func printTxDetail(tx *types.Transaction) {
	fmt.Println(tx.Hash().Hex())
	fmt.Println(tx.Cost())
	fmt.Println(tx.Gas())
	fmt.Println(tx.GasPrice())
	fmt.Println(tx.To().Hex()) // For the transaction to create a contract, it doesn't have to address
}

func QueryrChainID() {
	client := getstarted.CreateConn()
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Chain id: %d", chainId)
}

func QueryTransactionReceipt() {
	client := getstarted.CreateConn()

	txHash := common.HexToHash("0xa99489a09dbeb616e2570beabd7dead4b2a417f3cab4da5696fef2c69c80bded")
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(receipt.Status)
	fmt.Println(receipt.Logs)

}
