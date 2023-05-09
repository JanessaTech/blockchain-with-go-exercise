package accounts

import (
	"context"
	"fmt"
	getstarted "hi-supergirl/blockchain-with-go-exercise/get-started"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func ReadBalanceInfo(blockNum int64) {
	client := getstarted.CreateConn()
	var blockNumber *big.Int
	if blockNum >= 0 {
		blockNumber = big.NewInt(blockNum)
	} else {
		blockNumber = nil
	}

	account := common.HexToAddress("0xEb600bE51572beB77B86F9f32BF14E8DbFAb144a")
	balance, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("balance is ", balance, "for ", account, " in block ", blockNum)
	//output, eg:
	//balance is  99994904880749389936 for  0xEb600bE51572beB77B86F9f32BF14E8DbFAb144a  in block  2

}
