package swap_v2

import (
	"github.com/YeMiao0715/ether"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"math/big"
	"testing"
	"time"
)

var engine *ether.Engine
var router2 = common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")
var tokenA = common.HexToAddress("0xec8142A785eEA84A7A0db084039bE496CbdB2890")
var tokenB = common.HexToAddress("0xc5CA44fa5efdd9F4F6e10858FFe1665f186bC84B")
var swapRouter2 *Router

var sender = common.HexToAddress("0xadD275a8Ee37acC0E82F6EB1F5ccAa3b2B51E8C0")
var senderPrivateKey = ""

func init() {
	logger, _ := zap.NewDevelopment()
	//engine = ether.NewEngine(logger, "http://8.218.112.195:8545", "")
	//engine = ether.NewEngine(logger, "https://mainnet.infura.io/v3/14e5c24b98634138a9127fc8db299970", "")
	engine = ether.NewEngine(logger, "http://127.0.0.1:8545", "")
	engine.SetGasPrice(decimal.New(10, 9).BigInt())
	swapRouter2 = NewRouter2(engine)
	pair = NewPair(engine)
}

func Test_WETH(t *testing.T) {

	t.Log(swapRouter2.Factory(router2))
	t.Log(swapRouter2.WETH(router2))

	t.Log(swapRouter2.GetAmountsOut(router2, decimal.New(1, 4).BigInt(), []common.Address{
		tokenB, tokenA,
	}))
	t.Log(swapRouter2.GetAmountsIn(router2, decimal.New(1, 6).BigInt(), []common.Address{
		tokenA, tokenB,
	}))
}

func TestSwapRouter2_AddLiquidity(t *testing.T) {
	t.Log(swapRouter2.AddLiquidity(router2,
		tokenA,
		tokenB,
		decimal.New(1, 18).BigInt(),
		decimal.New(1, 6).BigInt(),
		decimal.Zero.BigInt(),
		decimal.Zero.BigInt(),
		common.HexToAddress("0xadD275a8Ee37acC0E82F6EB1F5ccAa3b2B51E8C0"),
		big.NewInt(time.Now().Unix()+600),
		senderPrivateKey,
	))
}

func TestSwapRouter2_RemoveLiquidity(t *testing.T) {
	pairContract := NewPairContract(engine, pairAddress)

	total, err := pairContract.TotalSupply()
	t.Log(total, err)
	lpAmount, err := pairContract.BalanceOf(sender)
	t.Log(lpAmount, err)
	reserves, reserve1, last, err := pairContract.GetReserves()
	t.Log(reserves, reserve1, last, err)

	token0, _ := pairContract.Token0()
	t.Log(token0)
	token1, _ := pairContract.Token1()
	t.Log(token1)

	rate := decimal.NewFromBigInt(lpAmount, 0).Mul(decimal.NewFromFloat(0.1)).Floor()

	tokenAAmount := decimal.NewFromBigInt(reserves, 0).Mul(
		rate.Div(decimal.NewFromBigInt(total, 0)),
	).Floor()

	tokenBAmount := decimal.NewFromBigInt(reserve1, 0).Mul(
		rate.Div(decimal.NewFromBigInt(total, 0)),
	).Floor()

	t.Log(tokenAAmount, tokenBAmount)

	//_, buildTx, err := pairContract.Approve(router2, rate.BigInt(), senderPrivateKey)
	//if err != nil {
	//	panic(err)
	//}
	deadline := big.NewInt(time.Now().Unix() + 600)
	v, r, s, err := pairContract.PermitSign(router2, rate.BigInt(), deadline, senderPrivateKey)
	t.Log(v, r, s, err)

	to32Array := func(bytes []byte) [32]byte {
		res := [32]byte{}
		for i, _ := range res {
			res[i] = bytes[i]
		}
		return res
	}

	t.Log(swapRouter2.RemoveLiquidityWithPermit(router2,
		token0,
		token1,
		rate.BigInt(),
		tokenAAmount.BigInt(),
		tokenBAmount.Sub(decimal.New(1, 18)).BigInt(),
		common.HexToAddress("0xadD275a8Ee37acC0E82F6EB1F5ccAa3b2B51E8C0"),
		deadline,
		false,
		uint8(v.Uint64()),
		to32Array(r.Bytes()),
		to32Array(s.Bytes()),
		senderPrivateKey,
	))
}
