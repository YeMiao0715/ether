package main

import (
	"context"
	"fmt"
	"github.com/YeMiao0715/ether"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()
	engine := ether.NewEngine(logger, "http://8.218.112.195:8545", "")
	engine.SetGasPrice(decimal.New(1, 9).BigInt())
	client, _, err := engine.GetEthClient()
	if err != nil {
		panic(err)
	}

	//if isWs {
	//	newHeader := make(chan *types.Header)
	//	subscribe, err := client.SubscribeNewHead(context.Background(), newHeader)
	//	if err != nil {
	//		subscribe.Unsubscribe()
	//		panic(err)
	//	}
	//	for {
	//		select {
	//		case head := <-newHeader:
	//			fmt.Println(head.Number.String())
	//		case err := <-subscribe.Err():
	//			fmt.Println(err)
	//		}
	//	}
	//}

	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash("0x77bf3df5dd63f91675c5107122ced5061dc69838127cc6ba769bcaef0c414f2f"))
	if err != nil {
		panic(err)
	}
	fmt.Println(receipt.ContractAddress)
	for _, log := range receipt.Logs {
		fmt.Println(log)
		//log.
	}
	b, err := receipt.MarshalJSON()
	fmt.Println(string(b))
}
