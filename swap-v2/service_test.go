package swap_v2

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestNewServiceWithPairAndRouter(t *testing.T) {
	// 0x4E99615101cCBB83A462dC4DE2bc1362EF1365e5 uni
	//serv := NewServiceWithPairAndRouter(engine, common.HexToAddress("0x4E99615101cCBB83A462dC4DE2bc1362EF1365e5"), router2)
	serv := NewServiceWithPairAndRouter(engine, pairAddress, router2)
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

	factory := NewFactoryContract(engine, common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"))

	t.Log(factory.GetPair(tokenB, common.HexToAddress("0xD92FC79A4A713cEB005FD4D6901D527F6C62112A")))
	t.Log(factory.GetPair(tokenA, tokenB))
	//// 0x4E99615101cCBB83A462dC4DE2bc1362EF1365e5 uni
	serv, err := NewServiceWithFactory(engine,
		factory,
		router2,
		//tokenA,
		tokenB,
		common.HexToAddress("0xD92FC79A4A713cEB005FD4D6901D527F6C62112A"),
	)
	if err != nil {
		fmt.Printf("%+v", err)
	}
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
	factory := NewFactoryContract(engine, common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"))
	serv, err := NewServiceWithFactory(engine,
		factory,
		router2,
		//tokenA,
		tokenB,
		common.HexToAddress("0xD92FC79A4A713cEB005FD4D6901D527F6C62112A"),
	)
	t.Log(serv, err)
	amountA, err := serv.AmountByTokenAFromFloat(10000)
	tx, tokenATx, tokenBTx, err := serv.AddLiquidityWithTokenA(amountA, true, "")
	if err != nil {
		fmt.Printf("%+v", err)
	}
	t.Log(tx, tokenATx, tokenBTx, err)
}

func TestService_RemoveLiquidityWithPermit(t *testing.T) {
	factory := NewFactoryContract(engine, common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"))
	serv, err := NewServiceWithFactory(engine,
		factory,
		router2,
		//tokenA,
		tokenB,
		common.HexToAddress("0xD92FC79A4A713cEB005FD4D6901D527F6C62112A"),
	)
	if err != nil {
		panic(err)
	}
	t.Log(serv.Symbol())

	owner, err := serv.engine.PrivateKeyToAddress("")
	balance, err := serv.pair.BalanceOf(*owner)
	t.Log(balance)

	lpAmount, err := serv.AmountByLpFromFloat(0.001)
	t.Log(lpAmount, err)
	t.Log(serv.RemoveLiquidityWithPermit(lpAmount, ""))
}
