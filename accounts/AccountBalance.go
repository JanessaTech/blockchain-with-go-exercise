package accounts

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getClient() *ethclient.Client {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("We have a connection")

	return client
}

func ReadAccountInfo() {
	client := getClient()

	account := common.HexToAddress("0xEb600bE51572beB77B86F9f32BF14E8DbFAb144a")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("balance is ", balance, "for ", account)
	//output:
	// balance is  100000000000000000000 for  0xEb600bE51572beB77B86F9f32BF14E8DbFAb144a

}
