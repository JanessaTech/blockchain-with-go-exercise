package accountbalances

import (
	"context"
	"fmt"
	getstarted "hi-supergirl/blockchain-with-go-exercise/get-started"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func ReadBalanceInfo() {
	client := getstarted.CreateConn()

	account := common.HexToAddress("0xEb600bE51572beB77B86F9f32BF14E8DbFAb144a")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("balance is ", balance, "for ", account, " in latest block ")

	blockNumber := big.NewInt(2)

	balance1, err1 := client.BalanceAt(context.Background(), account, blockNumber)
	if err1 != nil {
		log.Fatal(err)
	}
	fmt.Println("balance is ", balance1, "for ", account, " in block ", 2)

	fbalance := new(big.Float)
	fbalance.SetString(balance1.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("ethValue:", ethValue)

}
