package swap_v2

import (
	"github.com/YeMiao0715/ether"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type Router struct {
	engine  *ether.Engine
	iRouter *IRouter
}

func NewRouter2(engine *ether.Engine) *Router {
	return &Router{
		engine:  engine,
		iRouter: &IRouter{},
	}
}

func (e *Router) Factory(contract common.Address) (common.Address, error) {
	b, err := e.iRouter.Factory()
	if err != nil {
		return common.Address{}, err
	}
	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return common.Address{}, err
	}

	return e.iRouter.UnpackFactory(resb)
}

func (e *Router) WETH(contract common.Address) (common.Address, error) {
	b, err := e.iRouter.WETH()
	if err != nil {
		return common.Address{}, err
	}
	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return common.Address{}, err
	}

	return e.iRouter.UnpackWETH(resb)
}

func (e *Router) AddLiquidity(
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
	abiData, err := e.iRouter.AddLiquidity(tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
	if err != nil {
		return "", nil, err
	}

	tx, err := e.engine.BuildTxByContractWithPrivateKey(contract, abiData, privateKey)
	if err != nil {
		return "", nil, err
	}

	return e.engine.SendTransactionWithPrivateKey(tx, privateKey)
}

func (e *Router) AddLiquidityETH(
	contract common.Address,
	token common.Address,
	amountTokenDesired,
	amountTokenMin,
	amountETHMin *big.Int,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	abiData, err := e.iRouter.AddLiquidityETH(token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
	if err != nil {
		return "", nil, err
	}

	tx, err := e.engine.BuildTxByContractWithPrivateKey(contract, abiData, privateKey)
	if err != nil {
		return "", nil, err
	}

	return e.engine.SendTransactionWithPrivateKey(tx, privateKey)
}

func (e *Router) RemoveLiquidity(
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
	abiData, err := e.iRouter.RemoveLiquidity(tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
	if err != nil {
		return "", nil, err
	}

	tx, err := e.engine.BuildTxByContractWithPrivateKey(contract, abiData, privateKey)
	if err != nil {
		return "", nil, err
	}

	return e.engine.SendTransactionWithPrivateKey(tx, privateKey)
}

func (e *Router) RemoveLiquidityETH(
	contract common.Address,
	token common.Address,
	liquidity,
	amountTokenMin,
	amountETHMin *big.Int,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	abiData, err := e.iRouter.RemoveLiquidityETH(token, liquidity, amountTokenMin, amountETHMin, to, deadline)
	if err != nil {
		return "", nil, err
	}

	tx, err := e.engine.BuildTxByContractWithPrivateKey(contract, abiData, privateKey)
	if err != nil {
		return "", nil, err
	}

	return e.engine.SendTransactionWithPrivateKey(tx, privateKey)
}

func (e *Router) RemoveLiquidityWithPermit(
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
	abiData, err := e.iRouter.RemoveLiquidityWithPermit(tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
	if err != nil {
		return "", nil, err
	}

	tx, err := e.engine.BuildTxByContractWithPrivateKey(contract, abiData, privateKey)
	if err != nil {
		return "", nil, err
	}

	return e.engine.SendTransactionWithPrivateKey(tx, privateKey)
}

func (e *Router) RemoveLiquidityETHWithPermit(
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
	abiData, err := e.iRouter.RemoveLiquidityETHWithPermit(token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
	if err != nil {
		return "", nil, err
	}

	tx, err := e.engine.BuildTxByContractWithPrivateKey(contract, abiData, privateKey)
	if err != nil {
		return "", nil, err
	}

	return e.engine.SendTransactionWithPrivateKey(tx, privateKey)
}

func (e *Router) SwapExactTokensForTokens(
	contract common.Address,
	amountIn,
	amountOutMin *big.Int,
	path []common.Address,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	abiData, err := e.iRouter.SwapExactTokensForTokens(amountIn, amountOutMin, path, to, deadline)
	if err != nil {
		return "", nil, err
	}

	tx, err := e.engine.BuildTxByContractWithPrivateKey(contract, abiData, privateKey)
	if err != nil {
		return "", nil, err
	}

	return e.engine.SendTransactionWithPrivateKey(tx, privateKey)
}

func (e *Router) SwapTokensForExactTokens(
	contract common.Address,
	amountOut,
	amountInMax *big.Int,
	path []common.Address,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	abiData, err := e.iRouter.SwapTokensForExactTokens(amountOut, amountInMax, path, to, deadline)
	if err != nil {
		return "", nil, err
	}

	tx, err := e.engine.BuildTxByContractWithPrivateKey(contract, abiData, privateKey)
	if err != nil {
		return "", nil, err
	}

	return e.engine.SendTransactionWithPrivateKey(tx, privateKey)
}

func (e *Router) SwapExactETHForTokens(
	contract common.Address,
	amountOutMin *big.Int,
	path []common.Address,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	abiData, err := e.iRouter.SwapExactETHForTokens(amountOutMin, path, to, deadline)
	if err != nil {
		return "", nil, err
	}

	tx, err := e.engine.BuildTxByContractWithPrivateKey(contract, abiData, privateKey)
	if err != nil {
		return "", nil, err
	}

	return e.engine.SendTransactionWithPrivateKey(tx, privateKey)
}

func (e *Router) SwapTokensForExactETH(
	contract common.Address,
	amountOut,
	amountInMax *big.Int,
	path []common.Address,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	abiData, err := e.iRouter.SwapTokensForExactETH(amountOut, amountInMax, path, to, deadline)
	if err != nil {
		return "", nil, err
	}

	tx, err := e.engine.BuildTxByContractWithPrivateKey(contract, abiData, privateKey)
	if err != nil {
		return "", nil, err
	}

	return e.engine.SendTransactionWithPrivateKey(tx, privateKey)
}

func (e *Router) SwapExactTokensForETH(
	contract common.Address,
	amountIn,
	amountOutMin *big.Int,
	path []common.Address,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	abiData, err := e.iRouter.SwapExactTokensForETH(amountIn, amountOutMin, path, to, deadline)
	if err != nil {
		return "", nil, err
	}

	tx, err := e.engine.BuildTxByContractWithPrivateKey(contract, abiData, privateKey)
	if err != nil {
		return "", nil, err
	}

	return e.engine.SendTransactionWithPrivateKey(tx, privateKey)
}

func (e *Router) SwapETHForExactTokens(
	contract common.Address,
	amountOut *big.Int,
	path []common.Address,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	abiData, err := e.iRouter.SwapETHForExactTokens(amountOut, path, to, deadline)
	if err != nil {
		return "", nil, err
	}

	tx, err := e.engine.BuildTxByContractWithPrivateKey(contract, abiData, privateKey)
	if err != nil {
		return "", nil, err
	}

	return e.engine.SendTransactionWithPrivateKey(tx, privateKey)
}

func (e *Router) SwapExactTokensForTokensSupportingFeeOnTransferTokens(
	contract common.Address,
	amountIn,
	amountOutMin *big.Int,
	path []common.Address,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	abiData, err := e.iRouter.SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn, amountOutMin, path, to, deadline)
	if err != nil {
		return "", nil, err
	}

	tx, err := e.engine.BuildTxByContractWithPrivateKey(contract, abiData, privateKey)
	if err != nil {
		return "", nil, err
	}

	return e.engine.SendTransactionWithPrivateKey(tx, privateKey)
}

func (e *Router) Quote(contract common.Address, amountA, reserveA, reserveB *big.Int) (*big.Int, error) {
	b, err := e.iRouter.Quote(amountA, reserveA, reserveB)
	if err != nil {
		return nil, err
	}
	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return nil, err
	}

	return e.iRouter.UnpackGetAmountOut(resb)
}

func (e *Router) GetAmountOut(contract common.Address, amountIn, reserveIn, reserveOut *big.Int) (*big.Int, error) {
	b, err := e.iRouter.GetAmountOut(amountIn, reserveIn, reserveOut)
	if err != nil {
		return nil, err
	}
	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return nil, err
	}

	return e.iRouter.UnpackGetAmountOut(resb)
}

func (e *Router) GetAmountIn(contract common.Address, amountOut, reserveIn, reserveOut *big.Int) (*big.Int, error) {
	b, err := e.iRouter.GetAmountIn(amountOut, reserveIn, reserveOut)
	if err != nil {
		return nil, err
	}
	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return nil, err
	}

	return e.iRouter.UnpackGetAmountIn(resb)
}

func (e *Router) GetAmountsOut(contract common.Address, amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	b, err := e.iRouter.GetAmountsOut(amountIn, path)
	if err != nil {
		return nil, err
	}
	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return nil, err
	}

	return e.iRouter.UnpackGetAmountsOut(resb)
}

func (e *Router) GetAmountsIn(contract common.Address, amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	b, err := e.iRouter.GetAmountsIn(amountOut, path)
	if err != nil {
		return nil, err
	}
	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return nil, err
	}

	return e.iRouter.UnpackGetAmountsIn(resb)
}
