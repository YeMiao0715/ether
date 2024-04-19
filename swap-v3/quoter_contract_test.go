package swap_v3

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"math/big"
	"testing"
)

var quoter = common.HexToAddress("0x4752ba5DBc23f44D87826276BF6Fd6b1C372aD24")

func TestQuoterContract_GetPool(t *testing.T) {
	WETH := common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	USDT := common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7")
	quoterContract := NewQuoterContract(engine, quoter)
	t.Log(quoterContract.QuoteExactInput([]common.Address{WETH, USDT}, decimal.New(1, 18).BigInt()))
	t.Log(quoterContract.QuoteExactInputSingle(QuoteExactInputSingleParams{
		TokenIn:           WETH,
		TokenOut:          USDT,
		AmountIn:          decimal.New(1, 18).BigInt(),
		Fee:               big.NewInt(500),
		SqrtPriceLimitX96: big.NewInt(0),
	}))
	t.Log(quoterContract.QuoteExactInputSingleWithPool(QuoteExactInputSingleWithPoolParams{
		TokenIn:           WETH,
		TokenOut:          USDT,
		AmountIn:          decimal.New(1, 18).BigInt(),
		Fee:               big.NewInt(500),
		Pool:              common.HexToAddress("0x11b815efB8f581194ae79006d24E0d814B7697F6"),
		SqrtPriceLimitX96: big.NewInt(0),
	}))
	t.Log(quoterContract.QuoteExactOutputSingleWithPool(QuoteExactOutputSingleWithPoolParams{
		TokenIn:           WETH,
		TokenOut:          USDT,
		Amount:            decimal.New(1, 18).BigInt(),
		Fee:               big.NewInt(500),
		Pool:              common.HexToAddress("0x11b815efB8f581194ae79006d24E0d814B7697F6"),
		SqrtPriceLimitX96: big.NewInt(0),
	}))
	t.Log(quoterContract.QuoteExactOutputSingle(QuoteExactOutputSingleParams{
		TokenIn:           WETH,
		TokenOut:          USDT,
		Amount:            decimal.New(1, 18).BigInt(),
		Fee:               big.NewInt(500),
		SqrtPriceLimitX96: big.NewInt(0),
	}))
}
