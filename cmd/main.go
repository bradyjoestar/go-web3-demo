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

	for i := 60910; i < 60940; i++ {
		block, err := c.Eth().GetBlockByNumber(ethgo.BlockNumber(i), true)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("block height:%d, block time: %d, total transactions:%d\n",
			i, block.Timestamp, len(block.Transactions))

		if len(block.Transactions) != 0 {
			if receipt, err := c.Eth().GetTransactionReceipt(block.Transactions[0].Hash); err == nil {
				fmt.Printf("block height:%d, receipt result is %d\n", i, receipt.Status)
			}
		}
	}
}

func handleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
