package erc20

import (
	"fmt"
	"github.com/YeMiao0715/ether"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"testing"
)

var engine *ether.Engine

func init() {
	logger, _ := zap.NewDevelopment()
	engine = ether.NewEngine(logger, "https://data-seed-prebsc-1-s1.binance.org:8545/", "")
	engine.SetGasPrice(decimal.New(10, 9).BigInt())
}

func TestErc20Abi_Method(t *testing.T) {

	//b, _ := Erc20Abi.BalanceOf(common.HexToAddress("0x06eea78c7722d79b5B4B4681cB0E5798146f193d"))
	//t.Log(hexutil.Encode(b))

	_, err := IErc20Abi.abi.EventByID(common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"))

	if err != nil {
		panic(err)
	}

	//fmt.Println(event.Inputs)

}

func TestErc20_Name(t *testing.T) {
	e := NewErc20(engine)
	t.Log(e.Name(common.HexToAddress("0x762f7957714CfCfd271a9078e8f446692627126f")))
	t.Log(e.Symbol(common.HexToAddress("0x762f7957714CfCfd271a9078e8f446692627126f")))
	t.Log(e.Decimals(common.HexToAddress("0x762f7957714CfCfd271a9078e8f446692627126f")))
	t.Log(e.Allowance(
		common.HexToAddress("0x78867BbEeF44f2326bF8DDd1941a4439382EF2A7"),
		common.HexToAddress("0x06eea78c7722d79b5B4B4681cB0E5798146f193d"),
		common.HexToAddress("0xa1be04c6F760D887Fd83570734c8B06F77B8826e"),
	))
}

func TestErc20_BalanceOf(t *testing.T) {
	e := NewErc20(engine)
	t.Log(e.BalanceOf(common.HexToAddress("0x762f7957714CfCfd271a9078e8f446692627126f"), common.HexToAddress("0x06eea78c7722d79b5B4B4681cB0E5798146f193d")))
}

func TestErc20_Approve(t *testing.T) {
	e := NewErc20(engine)
	t.Log(e.BalanceOf(common.HexToAddress("0x762f7957714CfCfd271a9078e8f446692627126f"), common.HexToAddress("0x06eea78c7722d79b5B4B4681cB0E5798146f193d")))

	hash, buildTx, err := e.Approve(
		common.HexToAddress("0x78867BbEeF44f2326bF8DDd1941a4439382EF2A7"),
		common.HexToAddress("0x06eea78c7722d79b5B4B4681cB0E5798146f193d"),
		common.HexToAddress("0x762f7957714CfCfd271a9078e8f446692627126f"),
		decimal.New(2, 18).BigInt(),
		"3a78981e002260660e21e27b66a16a56a8f3ada60d9fc59e67410c1f3576ea61",
	)
	t.Log(hash, buildTx, err)
}

func TestErc20_Transfer(t *testing.T) {
	e := NewErc20(engine)
	hash, buildTx, err := e.Transfer(
		common.HexToAddress("0x762f7957714CfCfd271a9078e8f446692627126f"),
		common.HexToAddress("0xa1be04c6F760D887Fd83570734c8B06F77B8826e"),
		decimal.New(100, 18).BigInt(),
		"3a78981e002260660e21e27b66a16a56a8f3ada60d9fc59e67410c1f3576ea61",
	)

	if err != nil {
		fmt.Printf("%+v", err)
	}

	t.Log(hash, buildTx)
}

func TestErrors(t *testing.T) {

	err := errors.New("a err")
	new := errors.WithStack(err)

	panic(new)
}
