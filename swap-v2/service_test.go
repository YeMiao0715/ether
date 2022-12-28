package swap_v2

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

var serv *Service

func init() {
	_serv, err := NewServiceWithRouter(engine, router2)
	if err != nil {
		panic(err)
	}

	//factory := NewFactoryContract(engine, common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"))
	//_serv, err := NewServiceWithFactory(engine,
	//	factory,
	//	router2,
	//	//tokenA,
	//	tokenB,
	//	common.HexToAddress("0xD92FC79A4A713cEB005FD4D6901D527F6C62112A"),
	//)
	//if err != nil {
	//	panic(err)
	//}
	serv = _serv
}

func TestService_NewServiceForTokenAndWETH(t *testing.T) {
	// 0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45 goerli swapRouter2
	tokenA := common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7")
	//tokenB := common.HexToAddress("")
	//serv.MustFactory().GetPair(tokenA, tokenB)
	//serv.MustFactory().

	weth2Usdt, err := serv.NewServiceForTokenAndWETH(tokenA)
	if err != nil {
		panic(err)
	}
	fmt.Println(weth2Usdt.Symbol())
	fmt.Println(weth2Usdt.tokenA.Contract())
	fmt.Println(weth2Usdt.tokenB.Contract())
	fmt.Println(weth2Usdt.Price())

}

func TestNewServiceWithPairAndRouter(t *testing.T) {
	// 0x4E99615101cCBB83A462dC4DE2bc1362EF1365e5 uni
	//serv := NewServiceWithPairAndRouter(engine, common.HexToAddress("0x4E99615101cCBB83A462dC4DE2bc1362EF1365e5"), router2)
	t.Log(serv.Symbol())
	_price, err := serv.Price()
	t.Log(_price.ToDecimal().String(), err)
	t.Log(_price.ToString(), err)
	inAmount, err := serv.AmountByTokenAFromFloat(1)
	t.Log(inAmount, err)
	t.Log(serv.GetAmountsOut(inAmount))
	//factoryContract, _ := serv.Factory()
	//tokenb, err := serv.TokenB()
	//if err != nil {
	//	fmt.Printf("%+v", err)
	//}
	//fmt.Println(tokenb.Contract())
	//t.Log(factoryContract.Contract())
	//t.Log(factoryContract.FeeTo())
}

func TestNewServiceWithFactory(t *testing.T) {
	////serv := NewServiceWithPairAndRouter2(engine, pairAddress, router2)
	symbol, err := serv.Symbol()
	if err != nil {
		fmt.Printf("%+v", err)
	}
	t.Log(symbol)
	factoryContract, _ := serv.Factory()
	t.Log(factoryContract.FeeTo())
}

func TestService_AddLiquidityByTokenA(t *testing.T) {
	amountA, err := serv.AmountByTokenAFromFloat(100)
	tx, tokenATx, tokenBTx, err := serv.AddLiquidityWithTokenA(amountA, "")
	if err != nil {
		fmt.Printf("%+v", err)
	}
	t.Log(tx, tokenATx, tokenBTx, err)
}

func TestService_RemoveLiquidityWithPermit(t *testing.T) {
	owner, err := serv.engine.PrivateKeyToAddress("")
	balance, err := serv.pair.BalanceOf(*owner)
	t.Log(balance)

	lpAmount, err := serv.AmountByLpFromFloat(0.001)
	t.Log(lpAmount, err)
	t.Log(serv.RemoveLiquidity(lpAmount, ""))
}

func TestService_RemoveLiquidityWithTokenA(t *testing.T) {
	amountB, err := serv.AmountByTokenBFromFloat(10000)
	t.Log(amountB, err)
	t.Log(serv.RemoveLiquidityWithTokenB(amountB, ""))
}

func TestService_SwapWithTokenA(t *testing.T) {
	amountA, _ := serv.AmountByTokenAFromFloat(1)
	swapTx, _, err := serv.SwapTokenA2TokenBWithSupportingFee(amountA, 0.07, "")
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	t.Log(swapTx.Hash().String())
}

func TestService_SwapWithTokenB(t *testing.T) {
	//_tokenA, _ := serv.TokenA()
	//_tokenB, _ := serv.TokenB()
	//owner, _ := serv.engine.PrivateKeyToAddress("")
	//_pair, _ := serv.Pair()
	//t.Log(_tokenA.Approve(_pair.contract, big.NewInt(0), ""))
	//t.Log(_tokenB.Approve(_pair.contract, big.NewInt(0), ""))
	//t.Log(_tokenA.Approve(serv.Router().contract, big.NewInt(0), ""))
	//t.Log(_tokenB.Approve(serv.Router().contract, big.NewInt(0), ""))
	//t.Log(_tokenA.Allowance(*owner, _pair.contract))
	//t.Log(_tokenB.Allowance(*owner, _pair.contract))
	//t.Log(_tokenA.Allowance(*owner, serv.Router().contract))
	//t.Log(_tokenB.Allowance(*owner, serv.Router().contract))
	amountB, err := serv.AmountByTokenBFromFloat(100)
	swapTx, _, err := serv.SwapTokenB2TokenAWithSupportingFee(amountB, 0.07, "")
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	t.Log(swapTx.Hash().String())
}

func TestService_GetAmountsIn(t *testing.T) {
	_pair, _ := serv.Pair()
	amountB, _ := serv.AmountByTokenBFromFloat(1000)
	amounts := [2]*big.Int{}
	amounts[len(amounts)-1] = amountB
	for i := 1; i > 0; i-- {
		reserve0, reserve1, _, _ := _pair.GetReserves()
		t.Log(serv.Router().GetAmountIn(amounts[i], reserve0, reserve1))
	}

	fmt.Println(amounts)
}
