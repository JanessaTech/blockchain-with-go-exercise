package accounts

import (
	"context"
	"fmt"
	getstarted "hi-supergirl/blockchain-with-go-exercise/get-started"
	"log"

	"github.com/ethereum/go-ethereum/common"
)

func ReadAccountInfo() {
	client := getstarted.CreateConn()

	account := common.HexToAddress("0xEb600bE51572beB77B86F9f32BF14E8DbFAb144a")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("balance is ", balance, "for ", account)
	//output:
	// balance is  100000000000000000000 for  0xEb600bE51572beB77B86F9f32BF14E8DbFAb144a

}
