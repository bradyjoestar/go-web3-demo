package main

import (
	"fmt"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/jsonrpc"
	"github.com/umbracle/ethgo/wallet"
)

var (
	walletFile = "wallet.json"
	password   = "123"
)

// call a contract
func main() {
	key, _ := importWallet()
	fmt.Println(key.Address())

	c, err := jsonrpc.NewClient("http://localhost:8545")
	handleErr(err)
	found, err := c.Eth().GetBalance(key.Address(), ethgo.Latest)
	handleErr(err)

	fmt.Println(found.String())
}

func importWallet() (ethgo.Key, error) {
	key, err := wallet.NewJSONWalletFromFile(walletFile, password)
	handleErr(err)
	return key, nil
}

func handleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
