package main

import (
	"fmt"
	"github.com/umbracle/ethgo"
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

	for i := 60910; i < 60960; i++ {
		block, err := c.Eth().GetBlockByNumber(ethgo.BlockNumber(i), false)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("block height:%d, block time: %d, total transactions:%d\n",
			i, block.Timestamp, len(block.Transactions))
	}
}

func handleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
