package main

import (
	"wallet_mng/transaction"
	"wallet_mng/wallets"
)

var (
	network *string
	mainurl = "https://mainnet.infura.io/v3/baab6a1c72174240b4d74e208fb2481f"
	rinkbey = "https://rinkeby.infura.io/v3/baab6a1c72174240b4d74e208fb2481f"
	bsc     = "https://data-seed-prebsc-1-s1.binance.org:8545"
)

func main() {

	network := &rinkbey

	//create
	wallets.Create("password")

	//read
	wallets.Read("password", "UTC--2022-04-12T22-16-04.986110400Z--b7f4a6ba8e996b924fb933278e17fab64ed7dd60")

	transaction.Tx(transaction.TransactionParams{
		FromAddress:   "73292ACA2EC4d4E7Ab724cf8b6Ab91e9d15e8a3D",
		ToAddress:     "C8d99d0687c9bDB1b93190207Cc3e7117753e7db",
		Amount:        10000,
		GasLimit:      21000,
		FromPassword:  "password",
		SignatureFile: "UTC--2022-04-12T22-16-04.986110400Z--b7f4a6ba8e996b924fb933278e17fab64ed7dd60",
		Network:       *network,
	})

}
