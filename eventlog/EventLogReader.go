package eventlog

import (
	"context"
	"fmt"
	accounttokenbalances "hi-supergirl/blockchain-with-go-exercise/accounts/AccountTokenBalances"
	getstarted "hi-supergirl/blockchain-with-go-exercise/get-started"
	"hi-supergirl/blockchain-with-go-exercise/smartcontract/store"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type LogTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
}

// LogApproval ..
type LogApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
}

// before run this method, make sure:
// - Ganache is ready with port 8545.
// - mnemonic is "lady never blame vintage world talent believe almost apology knee keep scout" when start ganache
// - You have deployed Store.sol by running all methods defined in smartcontract.OpStore.go
// - When you run methods defined in smartcontract.OpStore.go make sure : in order:
//    -- run methods in this order: 1. DeployContract 2. WriteContract
//    -- Once you done DeployContract, modify line89 with the laest address of Store.sol

func ReadEventLogs() {
	client := getstarted.CreateConn()

	contractAddress := common.HexToAddress("0x0F1C3B16E0626e0d9d0f910D536AFf75b6e2e353")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   big.NewInt(4),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		event := struct {
			Key   [32]byte
			Value [32]byte
		}{}

		err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(event.Key[:]))   // me
		fmt.Println(string(event.Value[:])) // Janessa
		fmt.Println(vLog.BlockHash.Hex())   // block hash
		fmt.Println(vLog.BlockNumber)       // block id
		fmt.Println(vLog.TxHash.Hex())      // tx hash id
	}
}

// before run this method, make sure:
// - Ganache is ready with port 8545.
// - mnemonic is "lady never blame vintage world talent believe almost apology knee keep scout" when start ganache
// - run truffle migrate to deploy contracts. See code: https://github.com/hi-supergirl/exercises/tree/master/blockchain/get-started/truffle/token/ERC20/demo1
// - Once deployment is finished, you will see an event Transfer is existing in ganache
// - Check the address of METoken, update the value with the lastest address of METoken
func Reading_ERC_20_Token_Event_Logs() {
	client := getstarted.CreateConn()

	contractAddress := common.HexToAddress("0x0F1C3B16E0626e0d9d0f910D536AFf75b6e2e353") // the lastest address of METoken
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   big.NewInt(10),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(accounttokenbalances.TokenABI)))
	if err != nil {
		log.Fatal(err)
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	LogApprovalSig := []byte("Approval(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)

	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)

		switch vLog.Topics[0].Hex() {
		case logTransferSigHash.Hex():
			fmt.Printf("Log Name: Transfer\n")

			var transferEvent LogTransfer

			err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("From: %s\n", transferEvent.From.Hex())
			fmt.Printf("To: %s\n", transferEvent.To.Hex())
			fmt.Printf("Value: %s\n", transferEvent.Value.String())
		case logApprovalSigHash.Hex():
			fmt.Printf("Log Name: Approval\n")

			var approvalEvent LogApproval

			err := contractAbi.UnpackIntoInterface(&approvalEvent, "Approval", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			approvalEvent.Owner = common.HexToAddress(vLog.Topics[1].Hex())
			approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("Token Owner: %s\n", approvalEvent.Owner.Hex())
			fmt.Printf("Spender: %s\n", approvalEvent.Spender.Hex())
			fmt.Printf("Value: %s\n", approvalEvent.Value.String())
		}
	}

}
