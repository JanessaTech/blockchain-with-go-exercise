package transactions

import (
	queryingblocks "hi-supergirl/blockchain-with-go-exercise/Transactions/QueryingBlocks"
	queryingtransactions "hi-supergirl/blockchain-with-go-exercise/Transactions/QueryingTransactions"
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
