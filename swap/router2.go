package swap

import (
	"github.com/YeMiao0715/ether"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type Router2 struct {
	engine   *ether.Engine
	iRouter2 *IRouter2
}

func NewRouter2(engine *ether.Engine) *Router2 {
	return &Router2{
		engine:   engine,
		iRouter2: &IRouter2{},
	}
}

func (e *Router2) Factory(contract common.Address) (common.Address, error) {
	b, err := e.iRouter2.Factory()
	if err != nil {
		return common.Address{}, err
	}
	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return common.Address{}, err
	}

	return e.iRouter2.UnpackFactory(resb)
}

func (e *Router2) WETH(contract common.Address) (common.Address, error) {
	b, err := e.iRouter2.WETH()
	if err != nil {
		return common.Address{}, err
	}
	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return common.Address{}, err
	}

	return e.iRouter2.UnpackWETH(resb)
}

func (e *Router2) AddLiquidity(
	contract common.Address,
	tokenA,
	tokenB common.Address,
	amountADesired,
	amountBDesired,
	amountAMin,
	amountBMin *big.Int,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	abiData, err := e.iRouter2.AddLiquidity(tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
	if err != nil {
		return "", nil, err
	}

	tx, err := e.engine.BuildTxByContractWithPrivateKey(contract, abiData, privateKey)
	if err != nil {
		return "", nil, err
	}

	return e.engine.SendTransactionWithPrivateKey(tx, privateKey)
}

func (e *Router2) AddLiquidityETH(
	contract common.Address,
	token common.Address,
	amountTokenDesired,
	amountTokenMin,
	amountETHMin *big.Int,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	abiData, err := e.iRouter2.AddLiquidityETH(token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
	if err != nil {
		return "", nil, err
	}

	tx, err := e.engine.BuildTxByContractWithPrivateKey(contract, abiData, privateKey)
	if err != nil {
		return "", nil, err
	}

	return e.engine.SendTransactionWithPrivateKey(tx, privateKey)
}

func (e *Router2) RemoveLiquidity(
	contract common.Address,
	tokenA,
	tokenB common.Address,
	liquidity,
	amountAMin,
	amountBMin *big.Int,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	abiData, err := e.iRouter2.RemoveLiquidity(tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
	if err != nil {
		return "", nil, err
	}

	tx, err := e.engine.BuildTxByContractWithPrivateKey(contract, abiData, privateKey)
	if err != nil {
		return "", nil, err
	}

	return e.engine.SendTransactionWithPrivateKey(tx, privateKey)
}

func (e *Router2) RemoveLiquidityETH(
	contract common.Address,
	token common.Address,
	liquidity,
	amountTokenMin,
	amountETHMin *big.Int,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	abiData, err := e.iRouter2.RemoveLiquidityETH(token, liquidity, amountTokenMin, amountETHMin, to, deadline)
	if err != nil {
		return "", nil, err
	}

	tx, err := e.engine.BuildTxByContractWithPrivateKey(contract, abiData, privateKey)
	if err != nil {
		return "", nil, err
	}

	return e.engine.SendTransactionWithPrivateKey(tx, privateKey)
}

func (e *Router2) RemoveLiquidityWithPermit(
	contract common.Address,
	tokenA,
	tokenB common.Address,
	liquidity,
	amountAMin,
	amountBMin *big.Int,
	to common.Address,
	deadline *big.Int,
	approveMax bool,
	v uint8,
	r [32]byte,
	s [32]byte,
	privateKey string,
) (string, *types.Transaction, error) {
	abiData, err := e.iRouter2.RemoveLiquidityWithPermit(tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
	if err != nil {
		return "", nil, err
	}

	tx, err := e.engine.BuildTxByContractWithPrivateKey(contract, abiData, privateKey)
	if err != nil {
		return "", nil, err
	}

	return e.engine.SendTransactionWithPrivateKey(tx, privateKey)
}

func (e *Router2) RemoveLiquidityETHWithPermit(
	contract common.Address,
	token common.Address,
	liquidity,
	amountTokenMin,
	amountETHMin *big.Int,
	to common.Address,
	deadline *big.Int,
	approveMax bool,
	v uint8,
	r [32]byte,
	s [32]byte,
	privateKey string,
) (string, *types.Transaction, error) {
	abiData, err := e.iRouter2.RemoveLiquidityETHWithPermit(token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
	if err != nil {
		return "", nil, err
	}

	tx, err := e.engine.BuildTxByContractWithPrivateKey(contract, abiData, privateKey)
	if err != nil {
		return "", nil, err
	}

	return e.engine.SendTransactionWithPrivateKey(tx, privateKey)
}

func (e *Router2) SwapExactTokensForTokens(
	contract common.Address,
	amountIn,
	amountOutMin *big.Int,
	path []common.Address,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	abiData, err := e.iRouter2.SwapExactTokensForTokens(amountIn, amountOutMin, path, to, deadline)
	if err != nil {
		return "", nil, err
	}

	tx, err := e.engine.BuildTxByContractWithPrivateKey(contract, abiData, privateKey)
	if err != nil {
		return "", nil, err
	}

	return e.engine.SendTransactionWithPrivateKey(tx, privateKey)
}

func (e *Router2) SwapTokensForExactTokens(
	contract common.Address,
	amountOut,
	amountInMax *big.Int,
	path []common.Address,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	abiData, err := e.iRouter2.SwapTokensForExactTokens(amountOut, amountInMax, path, to, deadline)
	if err != nil {
		return "", nil, err
	}

	tx, err := e.engine.BuildTxByContractWithPrivateKey(contract, abiData, privateKey)
	if err != nil {
		return "", nil, err
	}

	return e.engine.SendTransactionWithPrivateKey(tx, privateKey)
}

func (e *Router2) SwapExactETHForTokens(
	contract common.Address,
	amountOutMin *big.Int,
	path []common.Address,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	abiData, err := e.iRouter2.SwapExactETHForTokens(amountOutMin, path, to, deadline)
	if err != nil {
		return "", nil, err
	}

	tx, err := e.engine.BuildTxByContractWithPrivateKey(contract, abiData, privateKey)
	if err != nil {
		return "", nil, err
	}

	return e.engine.SendTransactionWithPrivateKey(tx, privateKey)
}

func (e *Router2) SwapTokensForExactETH(
	contract common.Address,
	amountOut,
	amountInMax *big.Int,
	path []common.Address,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	abiData, err := e.iRouter2.SwapTokensForExactETH(amountOut, amountInMax, path, to, deadline)
	if err != nil {
		return "", nil, err
	}

	tx, err := e.engine.BuildTxByContractWithPrivateKey(contract, abiData, privateKey)
	if err != nil {
		return "", nil, err
	}

	return e.engine.SendTransactionWithPrivateKey(tx, privateKey)
}

func (e *Router2) SwapExactTokensForETH(
	contract common.Address,
	amountIn,
	amountOutMin *big.Int,
	path []common.Address,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	abiData, err := e.iRouter2.SwapExactTokensForETH(amountIn, amountOutMin, path, to, deadline)
	if err != nil {
		return "", nil, err
	}

	tx, err := e.engine.BuildTxByContractWithPrivateKey(contract, abiData, privateKey)
	if err != nil {
		return "", nil, err
	}

	return e.engine.SendTransactionWithPrivateKey(tx, privateKey)
}

func (e *Router2) SwapETHForExactTokens(
	contract common.Address,
	amountOut *big.Int,
	path []common.Address,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	abiData, err := e.iRouter2.SwapETHForExactTokens(amountOut, path, to, deadline)
	if err != nil {
		return "", nil, err
	}

	tx, err := e.engine.BuildTxByContractWithPrivateKey(contract, abiData, privateKey)
	if err != nil {
		return "", nil, err
	}

	return e.engine.SendTransactionWithPrivateKey(tx, privateKey)
}

func (e *Router2) Quote(contract common.Address, amountA, reserveA, reserveB *big.Int) (*big.Int, error) {
	b, err := e.iRouter2.Quote(amountA, reserveA, reserveB)
	if err != nil {
		return nil, err
	}
	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return nil, err
	}

	return e.iRouter2.UnpackGetAmountOut(resb)
}

func (e *Router2) GetAmountOut(contract common.Address, amountIn, reserveIn, reserveOut *big.Int) (*big.Int, error) {
	b, err := e.iRouter2.GetAmountOut(amountIn, reserveIn, reserveOut)
	if err != nil {
		return nil, err
	}
	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return nil, err
	}

	return e.iRouter2.UnpackGetAmountOut(resb)
}

func (e *Router2) GetAmountIn(contract common.Address, amountOut, reserveIn, reserveOut *big.Int) (*big.Int, error) {
	b, err := e.iRouter2.GetAmountIn(amountOut, reserveIn, reserveOut)
	if err != nil {
		return nil, err
	}
	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return nil, err
	}

	return e.iRouter2.UnpackGetAmountIn(resb)
}

func (e *Router2) GetAmountsOut(contract common.Address, amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	b, err := e.iRouter2.GetAmountsOut(amountIn, path)
	if err != nil {
		return nil, err
	}
	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return nil, err
	}

	return e.iRouter2.UnpackGetAmountsOut(resb)
}

func (e *Router2) GetAmountsIn(contract common.Address, amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	b, err := e.iRouter2.GetAmountsIn(amountOut, path)
	if err != nil {
		return nil, err
	}
	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return nil, err
	}

	return e.iRouter2.UnpackGetAmountsIn(resb)
}
