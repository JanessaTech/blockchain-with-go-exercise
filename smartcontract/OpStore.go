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

// before running this method, make sure ganache is ready with port 8545
// mnemonic is "lady never blame vintage world talent believe almost apology knee keep scout" when start ganache
func DeployContract() {
	client := getstarted.CreateConn()

	auth := getAuth(client)

	input := "My store deployed by go 1.0"
	address, tx, instance, err := store.DeployStore(auth, client, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex()) // the address of the newly create contract
	// 0xaB78f3724EFF846A0baDdBBDd0c61c5594a3d195
	fmt.Println(tx.Hash().Hex()) // The hash id of the transaction used to deploy contract
	// 0x81aefe49fe24b72968b26400f599c8eaddcdff0ceb4bb117daadb53f1891301f

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

	address := common.HexToAddress("0xaB78f3724EFF846A0baDdBBDd0c61c5594a3d195")
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

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex()) //0x9c604fbfae289d592bd9ee1e4818ed08d98ac2f075d5b449352e3844dc983f1e

	result, err := instance.Items(nil, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(result[:])) // "Janessa"

}
