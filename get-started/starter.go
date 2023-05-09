package getstarted

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

// mnemonic is "lady never blame vintage world talent believe almost apology knee keep scout" when start ganache
func CreateConn() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("We have a connection")

	_ = client // we'll use this in the upcoming sections
}
