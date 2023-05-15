package main

import testing "hi-supergirl/blockchain-with-go-exercise/Testing"

func main() {
	//getstarted.CreateConn()

	//accounts.ReadBalanceInfo()
	//accounts.ERC20Testcases()
	//accounts.CreateWallet()
	//accounts.CreateKeyStore()
	//accounts.ImportKeyStore()
	//accounts.CheckPublicAddress()
	//accounts.CheckAddressIsSmartContract()

	//transactions.QueryBlockHeaderInfo()
	//transactions.QueryFullBlock()
	//transactions.QueryTxInfo()
	//transactions.QueryrChainID()
	//transactions.QueryTransactionReceipt()
	//transactions.SendEther()
	//transactions.TransferERC20()
	//transactions.MonitorNewBlock() //notifications not supported

	// make sure run the following methods in order one by one, otherwise there will some problems
	//smartcontract.DeployContract()
	//smartcontract.LoadContract()
	//smartcontract.QueryContract()
	//smartcontract.WriteContract()

	//eventlog.ReadEventLogs()
	//eventlog.Reading_ERC_20_Token_Event_Logs()

	//signatures.GnerateAndVerifySignature()

	testing.MockTest()
}
