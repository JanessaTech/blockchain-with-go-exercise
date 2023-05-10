package transferringeth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	getstarted "hi-supergirl/blockchain-with-go-exercise/get-started"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

//before call this method, make sure ganache is ready
//mnemonic is "lady never blame vintage world talent believe almost apology knee keep scout" when start ganache

func SendEther() {
	// 1. get a connect to network saying ganache
	client := getstarted.CreateConn()

	// 2. get a private key
	privateKey, err := crypto.HexToECDSA("daf1e04560cd49a5f240d541680e0710d1b2617da21032d193c6a789f392488f") // the private key of the first account in ganache
	//privateKey := crypto.ToECDSAUnsafe(common.FromHex("0xdaf1e04560cd49a5f240d541680e0710d1b2617da21032d193c6a789f392488f"))
	if err != nil {
		log.Fatal(err)
	}

	// 3. get public key from the private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey) // the syntax is called type assertion: publicKey.(<expectedType>)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//4. get next nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress) // sometimes, it method is not accurate if we want to get the next nonce
	if err != nil {
		log.Fatal(err)
	}

	// 5. we start to construct a tx
	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)                // in units
	//gasPrice := big.NewInt(30000000000)      // in wei (30 gwei)
	gasPrice, err := client.SuggestGasPrice(context.Background()) // we can either hardcode gasPrice or let client suggest one
	if err != nil {
		log.Fatal(err)
	}
	_ = gasPrice

	toAddress := common.HexToAddress("0x5d0B82b53F61A304940A6e3DEd6AC881Dd6774a5") // this is the second account

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil) // tx is ready to be sent

	// 6. sign the tx
	//chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	/*tipCap, _ := client.SuggestGasTipCap(context.Background())
	feeCap, _ := client.SuggestGasPrice(context.Background())

	tx := types.NewTx(
		&types.DynamicFeeTx{
			ChainID:   chainID,
			Nonce:     nonce,
			GasTipCap: tipCap,
			GasFeeCap: feeCap,
			Gas:       gasLimit,
			To:        &toAddress,
			Value:     value,
			Data:      nil,
		})
	*/
	//signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)
	//signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	//signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// There is an exeception thrown here:  transaction type not supported

	// 7.finally, we send the signed tx
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())

}

// this solution comes from ：https://medium.com/@akshay_111meher/creating-offline-raw-transactions-with-go-ethereum-8d6cc8174c5d
