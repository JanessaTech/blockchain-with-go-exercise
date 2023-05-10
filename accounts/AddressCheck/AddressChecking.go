package addresscheck

import (
	"context"
	"fmt"
	getstarted "hi-supergirl/blockchain-with-go-exercise/get-started"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
)

func CheckPublicAddress() {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	addr1 := "0x323b5d4c32345ced77393b3530b1eed0f346429d"
	addr2 := "0xZYXb5d4c32345ced77393b3530b1eed0f346429d"

	fmt.Printf("%s is valid: %v\n", addr1, re.MatchString(addr1))
	//0x323b5d4c32345ced77393b3530b1eed0f346429d is valid: true
	fmt.Printf("%s is valid: %v\n", addr2, re.MatchString(addr2))
	//0xZYXb5d4c32345ced77393b3530b1eed0f346429d is valid: false
}

// check if an address is an account or a smart contract
// Before run this method, make sure ganache is ready, and you have deployed at least one contract
func CheckAddressIsSmartContract() {
	client := getstarted.CreateConn()
	address := common.HexToAddress("0x0F1C3B16E0626e0d9d0f910D536AFf75b6e2e353") // this is the address of one contract

	bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is latest block

	if err != nil {
		log.Fatal(err)
	}

	isContract := len(bytecode) > 0
	fmt.Printf("is contract: %v\n", isContract)

}
