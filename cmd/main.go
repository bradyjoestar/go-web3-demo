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

	for i := 58720; i < 58740; i++ {
		block, err := c.Eth().GetBlockByNumber(60000, false)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("block height:%d, block time: %d, total transactions:%d",
			i, block.Timestamp, len(block.Transactions))
	}
}

func handleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
