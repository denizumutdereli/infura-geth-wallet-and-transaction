package wallets

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func Create(password string) {
	key := keystore.NewKeyStore("./store", keystore.StandardScryptN, keystore.StandardScryptP)
	a, err := key.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a.Address)
}

func Read(password string, address string) {
	b, err := ioutil.ReadFile("./store/" + address)
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, password)

	if err != nil {
		log.Fatal(err)
	}

	pData := crypto.FromECDSA(key.PrivateKey)
	fmt.Println("Priv:", hexutil.Encode(pData))

	pData = crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Println("Pub:", hexutil.Encode(pData))

	fmt.Println("Add", crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex())

}
