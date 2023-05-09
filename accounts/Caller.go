package accounts

import (
	accountbalances "hi-supergirl/blockchain-with-go-exercise/accounts/AccountBalances"
	accounttokenbalances "hi-supergirl/blockchain-with-go-exercise/accounts/AccountTokenBalances"
)

func ReadBalanceInfo() {
	accountbalances.ReadBalanceInfo()
}

func ERC20Testcases() {
	accounttokenbalances.ERC20Testcases()
}
