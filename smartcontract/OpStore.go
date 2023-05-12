package smartcontract

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	getstarted "hi-supergirl/blockchain-with-go-exercise/get-started"
	"hi-supergirl/blockchain-with-go-exercise/smartcontract/store"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// before running this demo, make sure ganache is ready with port 8545
// mnemonic is "lady never blame vintage world talent believe almost apology knee keep scout" when start ganache
// Here are commands we used to generate Store.go. I assume you have solc(solcjs on windows) and abigen installed on your local
// solcjs --abi Store.sol -o build
// solcjs --bin Store.sol -o build
// abigen --abi=./build/Store_sol_Store.abi --pkg=store --out=Store.go
// abigen --abi ./build/Store_sol_Store.abi --pkg store --type Store --out Store.go --bin Store_sol_Store.bin
// Copy Store.go to ./store folder revative to the current file
func DeployContract() {
	client := getstarted.CreateConn()

	auth := getAuth(client)

	input := "My store deployed by go 1.0"
	address, tx, instance, err := store.DeployStore(auth, client, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex()) // the address of the newly create contract
	// 0x0F1C3B16E0626e0d9d0f910D536AFf75b6e2e353   . The address is different every time you deploy the contract
	fmt.Println(tx.Hash().Hex()) // The hash id of the transaction used to deploy contract
	// 0xdb92a4fc056c351f8c45dfe0ef520c86716102c53661523d59effdcaa05d4d54    . hash id is different every time you deploy the contract

	_ = instance // we'll be using this in the next section

}

func getAuth(client *ethclient.Client) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA("daf1e04560cd49a5f240d541680e0710d1b2617da21032d193c6a789f392488f")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	_ = gasPrice
	if err != nil {
		log.Fatal(err)
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	//auth.GasLimit = uint64(300000) // in units
	//auth.GasPrice = gasPrice
	return auth
}

func LoadContract() *store.Store {
	client := getstarted.CreateConn()

	address := common.HexToAddress("0x0F1C3B16E0626e0d9d0f910D536AFf75b6e2e353") // the address of the newly create contract deployed in DeployContract()

	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	return instance // we'll be using this in the next section
}

func QueryContract() {
	instance := LoadContract()
	version, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version) // "My store deployed by go 1.0"
}

func WriteContract() {
	client := getstarted.CreateConn()
	instance := LoadContract()

	auth := getAuth(client)

	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("me"))
	copy(value[:], []byte("Janessa"))

	tx, err := instance.SetItem(auth, key, value)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex()) //0xb1572a86742936f172a735f999f8a11697144f0a3e4a0a81723ef231b17889ab  , the tx hash id. it is different everytime you write the contract

	result, err := instance.Items(nil, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(result[:])) // "Janessa"

}
