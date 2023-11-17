package ether

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"testing"
)

var engine *Engine

func init() {
	logger, _ := zap.NewDevelopment()
	engine = NewEngine(logger, "http://127.0.0.1:8545", "")
	//engine = NewEngine(logger, "https://data-seed-prebsc-1-s1.binance.org:8545/", "")
	//engine.SetGasPrice(decimal.New(10, 9).BigInt())
}

func TestClient_GetEthClient(t *testing.T) {
	ethClient, isWs, err := engine.GetEthClient()
	t.Log(ethClient, isWs, err)
	blockNumber, err := ethClient.BlockNumber(context.Background())
	t.Log(blockNumber, err)
}

func TestEngine_TransactionByHash(t *testing.T) {
	tx, err := engine.TransactionByHash(common.HexToHash("0x8a219af9443ebce5d82853c4357e84b2b3aa7c8b909890595dc61dbe7aeb060e"))
	t.Log(tx, err)

	isContract, code, err := tx.IsContract(engine)
	fmt.Println(isContract, code, err)

	contract := vm.Contract{Code: code}
	for _, b := range code {
		fmt.Println(contract.GetOp(uint64(b)))
	}
}

func TestEngine_TransferEth(t *testing.T) {
	sender, err := engine.PrivateKeyToAddress("0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	balance, err := engine.GetBalance(*sender)
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}

	fmt.Println(balance)

	hash, _, err := engine.TransferEth(common.HexToAddress("0x06eea78c7722d79b5B4B4681cB0E5798146f193d"), decimal.New(1, 18).BigInt(), "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if err != nil {
		fmt.Printf("%+v", err)
	} else {
		t.Log(hash)
	}
}
