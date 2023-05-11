package transferringtokenserc20

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	getstarted "hi-supergirl/blockchain-with-go-exercise/get-started"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

// Some explanations about this demo:
//   - This demo is doing exactly the same as what accounttokenbalances.ERC20Testcases is doing
//   - The difference between this demo and  accounttokenbalances.ERC20Testcases is
//     this demo uses raw transaction to call transfer() by mannually constructing data part
//     whereras accounttokenbalances.ERC20Testcases uses go binding with the help of abigen to call transfer()
//
// How to run this demo:
//  1. Start up ganache with port 8545, mnemonic is "lady never blame vintage world talent believe almost apology knee keep scout"
//  2. deploy code at exercises\blockchain\get-started\truffle\token\ERC20\demo1  by running "truffle migrate"
//     you could check https://www.notion.so/ERC20-71e673002df14695ab4c22af76927725?pvs=4 to how to run the contracts
//  3. check link below to learn more about how to run the equivalent steps in truffle console :
//     https://www.notion.so/ERC20-71e673002df14695ab4c22af76927725?pvs=4
func TransferERC20() {
	client := getstarted.CreateConn()

	privateKey, err := crypto.HexToECDSA("daf1e04560cd49a5f240d541680e0710d1b2617da21032d193c6a789f392488f") // the private key of the first account
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey) // the syntax is called type assertion: publicKey.(<expectedType>)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gasPrice", gasPrice)

	toAddress := common.HexToAddress("0xeA2fDe9a249A058356544FaD2810fF2a51B2cd3A")    //Faucet.address
	tokenAddress := common.HexToAddress("0x0F1C3B16E0626e0d9d0f910D536AFf75b6e2e353") //METoken.address

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println("function selector", hexutil.Encode(methodID)) // function selector: 0xa9059cbb

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)    // first parameter
	fmt.Println("first parameter:", hexutil.Encode(paddedAddress)) // 0x000000000000000000000000eA2fDe9a249A058356544FaD2810fF2a51B2cd3A

	amount := new(big.Int)
	amount.SetString("1000", 10) // sets the value to 1000 tokens, in the token denomination
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println("second parameter:", hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000000000000000000003e8

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	/*
		gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
			To:   &tokenAddress,
			Data: data,
		})
		if err != nil {
			log.Fatal(err)
		}*/
	gasLimit := uint64(90000)
	fmt.Println("gasLimit", gasLimit) // 90000

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// finally, we send the signed tx
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
	// once successfully, you will see not only a tx hash printed here but also a transfer envent in ganache GUI
}
