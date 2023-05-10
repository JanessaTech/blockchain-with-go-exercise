package keystores

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func CreateKeyStore() {
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(account.Address.Hex())
}

func ImportKeyStore() {
	file := "./wallets/UTC--2023-05-10T02-42-29.945016700Z--04535e97043b20052381ca854b65ef9e6963831d"
	// the value of file will be different every time you run CreateKeyStore()
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	password := "secret"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) //

	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}
