package erctokendemo

import (
	"context"
	"fmt"
	getstarted "hi-supergirl/blockchain-with-go-exercise/get-started"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// How to run this demo:
//  1. Start up ganache with port 8545, mnemonic is "lady never blame vintage world talent believe almost apology knee keep scout"
//  2. deploy code at exercises\blockchain\get-started\truffle\token\ERC20\demo1  by running "truffle migrate"
//     you could check https://www.notion.so/ERC20-71e673002df14695ab4c22af76927725?pvs=4 to how to run the contracts
//  3. Run command:  solcjs --abi erc20.sol (make sure you installed solc on your pc, in my case, it is windows)
//     file erc20_sol_IERC20.abi will be created
//  4. Run command:  abigen --abi=erc20_sol_IERC20.abi --pkg=token --out=erc20.go
//  5. Copy the content of erc20.go to this folder. (need to change package name to erctokendemo)
func ERC20Testcases() {
	client := getstarted.CreateConn()
	tokenAddress := common.HexToAddress("0x0F1C3B16E0626e0d9d0f910D536AFf75b6e2e353") // the address of METoken
	instance, err := NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	firstAccount := common.HexToAddress("0xEb600bE51572beB77B86F9f32BF14E8DbFAb144a") // this is the address of first account
	bal, err := instance.BalanceOf(&bind.CallOpts{Context: context.Background()}, firstAccount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("token: %s for firstAccount %s\n", bal, firstAccount) // By default, we set 2100000000 token to the first account

	// to fix: no signer to authorize the transaction with
	FaucetAddress := common.HexToAddress("Faucet") // address of Faucet
	transaction, err := instance.Transfer(&bind.TransactOpts{From: firstAccount}, FaucetAddress, big.NewInt(1000))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("transaction hash: %s", transaction.Hash())

	balLeft, err := instance.BalanceOf(&bind.CallOpts{Context: context.Background()}, firstAccount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("token: %s for firstAccount %s\n", balLeft, firstAccount)

}
