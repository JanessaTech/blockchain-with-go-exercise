package accounts

import (
	accountbalances "hi-supergirl/blockchain-with-go-exercise/accounts/AccountBalances"
	accounttokenbalances "hi-supergirl/blockchain-with-go-exercise/accounts/AccountTokenBalances"
	generatingnewwallets "hi-supergirl/blockchain-with-go-exercise/accounts/GeneratingNewWallets"
	keystores "hi-supergirl/blockchain-with-go-exercise/accounts/Keystores"
)

func ReadBalanceInfo() {
	accountbalances.ReadBalanceInfo()
}

func ERC20Testcases() {
	accounttokenbalances.ERC20Testcases()
}

func CreateWallet() {
	generatingnewwallets.CreateWallet()
}

func CreateKeyStore() {
	keystores.CreateKeyStore()
}

func ImportKeyStore() {
	keystores.ImportKeyStore()
}
