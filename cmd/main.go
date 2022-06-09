package main

import (
	"fmt"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/jsonrpc"
	"time"
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
		fmt.Println(time.Unix(int64(block.Timestamp), 0))

		if len(block.Transactions) != 0 {
			successNumber := 0
			for j := 0; j < len(block.Transactions); j++ {
				if receipt, err := c.Eth().GetTransactionReceipt(block.Transactions[j].Hash); err == nil {
					if receipt.Status == 1 {
						successNumber++
					}
				}
			}
			fmt.Printf("block height:%d, success number is %d\n", i, successNumber)
		}
	}
}

func handleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
