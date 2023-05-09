package main

import (
	"hi-supergirl/blockchain-with-go-exercise/accounts"
	getstarted "hi-supergirl/blockchain-with-go-exercise/get-started"
)

func getStarted() {
	getstarted.CreateConn()
}

func main() {
	//getStarted()
	accounts.ReadAccountInfo()
}
