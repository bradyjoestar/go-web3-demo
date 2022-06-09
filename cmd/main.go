package main

import (
	"fmt"
	"github.com/umbracle/ethgo/jsonrpc"
)

var (
	walletFile = "wallet.json"
	password   = "123"
)

// call a contract
func main() {

	c, err := jsonrpc.NewClient("http://172.17.0.1:8545")
	handleErr(err)

	block, err := c.Eth().GetBlockByNumber(60000, false)
	fmt.Println(block.Timestamp)
	fmt.Println("transactions")
	fmt.Println(len(block.Transactions))
}

func handleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
