package transaction

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TransactionParams struct {
	FromAddress   string
	ToAddress     string
	Amount        int64
	GasLimit      int64
	FromPassword  string
	SignatureFile string
	Network       string
}

func Tx(tx TransactionParams) {

	fmt.Println("Network:", tx.Network)
	client, err := ethclient.Dial(tx.Network)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	a1 := common.HexToAddress(tx.FromAddress)
	a2 := common.HexToAddress(tx.ToAddress)

	b1, err := client.BalanceAt(context.Background(), a1, nil)
	if err != nil {
		log.Fatal(err)
	}

	b2, err := client.BalanceAt(context.Background(), a2, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Balance one:", b1)
	fmt.Println("Balance two:", b2)

	value := big.NewInt(tx.Amount)
	gasLimit := uint64(tx.GasLimit) // in units default 21000
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), a1)

	if err != nil {
		log.Fatal(err)
	}

	var data []byte

	txs := types.NewTransaction(nonce, a2, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())

	fmt.Println(chainID)

	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadFile("./store/" + tx.SignatureFile)

	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, tx.FromPassword)

	if err != nil {
		log.Fatal(err)
	}
	txs, err = types.SignTx(txs, types.NewEIP155Signer(chainID), key.PrivateKey)

	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), txs)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx has sent to: %s", txs.Hash().Hex())
}
