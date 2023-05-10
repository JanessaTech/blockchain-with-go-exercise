package accounts

import (
	accountbalances "hi-supergirl/blockchain-with-go-exercise/accounts/AccountBalances"
	accounttokenbalances "hi-supergirl/blockchain-with-go-exercise/accounts/AccountTokenBalances"
	generatingnewwallets "hi-supergirl/blockchain-with-go-exercise/accounts/GeneratingNewWallets"
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
