package transactions

import (
	queryingblocks "hi-supergirl/blockchain-with-go-exercise/Transactions/QueryingBlocks"
	queryingtransactions "hi-supergirl/blockchain-with-go-exercise/Transactions/QueryingTransactions"
	subscribingtonewblocks "hi-supergirl/blockchain-with-go-exercise/Transactions/SubscribingToNewBlocks"
	transferringeth "hi-supergirl/blockchain-with-go-exercise/Transactions/TransferringETH"
	transferringtokenserc20 "hi-supergirl/blockchain-with-go-exercise/Transactions/TransferringTokensERC20"
)

func QueryBlockHeaderInfo() {
	queryingblocks.QueryBlockHeaderInfo()
}

func QueryFullBlock() {
	queryingblocks.QueryFullBlock()
}

func QueryTxInfo() {
	queryingtransactions.QueryTxInfo()
}

func QueryrChainID() {
	queryingtransactions.QueryrChainID()
}

func QueryTransactionReceipt() {
	queryingtransactions.QueryTransactionReceipt()
}

func SendEther() {
	transferringeth.SendEther()
}

func TransferERC20() {
	transferringtokenserc20.TransferERC20()
}

func MonitorNewBlock() {
	subscribingtonewblocks.Subscribe()
}
