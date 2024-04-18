package swap_v3

import (
	"context"
	"fmt"
	"github.com/YeMiao0715/ether"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"testing"
)

var engine *ether.Engine
var ipoolAbi = &IPool{}

func init() {
	logger, _ := zap.NewDevelopment()
	engine = ether.NewEngine(logger, "http://127.0.0.1:8545", "")
}

func TestIPoolAbi_GetAbi(t *testing.T) {
	lastBlockNumber, _ := engine.GetBlockNumber()

	ethClient := engine.MustEthClient()
	logs, err := ethClient.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: decimal.NewFromUint64(lastBlockNumber).Sub(decimal.NewFromInt(100)).BigInt(),
		Topics: [][]common.Hash{
			{ipoolAbi.MustAbi().Events["Mint"].ID},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	for _, l := range logs {
		fmt.Println(l.Address)
	}
}

func TestIPoolAbi_Slot0(t *testing.T) {
	//data, _ := ipoolAbi.Slot0()
	//data, err := engine.CallContract(common.HexToAddress("0x11b815efB8f581194ae79006d24E0d814B7697F6"), data)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//ipoolAbi.UnpackSlot0(data)
	pool := NewPoolContract(engine, common.HexToAddress("0x11b815efB8f581194ae79006d24E0d814B7697F6"))
	//slot0, err := pool.Slot0()
	//if err != nil {
	//	t.Fatal(err)
	//}

	t.Log(pool.Price())
	t.Log(pool.Price())

	//token0Address, _ := pool.Token0()
	//token1Address, _ := pool.Token1()
	//token0 := erc20.NewErc20WithContract(engine, token0Address)
	//token1 := erc20.NewErc20WithContract(engine, token1Address)
	//
	//amount0, _ := token0.ToAmount(1)
	//amount1, _ := token1.ToAmount(1)
	//t.Log(pool.Slot0())
	//t.Log(slot0.Token0Price(amount0, amount1))
	//t.Log(slot0.Token0Price(amount1, amount0))
	//t.Log()
	//t.Log(pool.Token1())
}
