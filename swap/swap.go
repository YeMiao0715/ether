package swap

import (
	"github.com/YeMiao0715/ether"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type SwapRouter2 struct {
	engine       *ether.Engine
	iSwapRouter2 *ISwapRouter2
}

func NewSwapRouter2(engine *ether.Engine) *SwapRouter2 {
	return &SwapRouter2{
		engine:       engine,
		iSwapRouter2: &ISwapRouter2{},
	}
}

func (e *SwapRouter2) Factory(contract common.Address) (common.Address, error) {
	b, err := e.iSwapRouter2.Factory()
	if err != nil {
		return common.Address{}, err
	}
	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return common.Address{}, err
	}

	return e.iSwapRouter2.UnpackFactory(resb)
}

func (e *SwapRouter2) WETH(contract common.Address) (common.Address, error) {
	b, err := e.iSwapRouter2.WETH()
	if err != nil {
		return common.Address{}, err
	}
	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return common.Address{}, err
	}

	return e.iSwapRouter2.UnpackWETH(resb)
}

func (e *SwapRouter2) AddLiquidity(
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
	abiData, err := e.iSwapRouter2.AddLiquidity(tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
	if err != nil {
		return "", nil, err
	}

	tx, err := e.engine.BuildTxByContractWithPrivateKey(contract, abiData, privateKey)
	if err != nil {
		return "", nil, err
	}

	return e.engine.SendTransactionWithPrivateKey(tx, privateKey)
}

func (e *SwapRouter2) RemoveLiquidity(
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
	abiData, err := e.iSwapRouter2.RemoveLiquidity(tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
	if err != nil {
		return "", nil, err
	}

	tx, err := e.engine.BuildTxByContractWithPrivateKey(contract, abiData, privateKey)
	if err != nil {
		return "", nil, err
	}

	return e.engine.SendTransactionWithPrivateKey(tx, privateKey)
}

func (e *SwapRouter2) RemoveLiquidityWithPermit(
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
	abiData, err := e.iSwapRouter2.RemoveLiquidityWithPermit(tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
	if err != nil {
		return "", nil, err
	}

	tx, err := e.engine.BuildTxByContractWithPrivateKey(contract, abiData, privateKey)
	if err != nil {
		return "", nil, err
	}

	return e.engine.SendTransactionWithPrivateKey(tx, privateKey)
}

func (e *SwapRouter2) GetAmountOut(contract common.Address, amountIn, reserveIn, reserveOut *big.Int) (*big.Int, error) {
	b, err := e.iSwapRouter2.GetAmountOut(amountIn, reserveIn, reserveOut)
	if err != nil {
		return nil, err
	}
	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return nil, err
	}

	return e.iSwapRouter2.UnpackGetAmountOut(resb)
}

func (e *SwapRouter2) GetAmountIn(contract common.Address, amountIn, reserveIn, reserveOut *big.Int) (*big.Int, error) {
	b, err := e.iSwapRouter2.GetAmountIn(amountIn, reserveIn, reserveOut)
	if err != nil {
		return nil, err
	}
	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return nil, err
	}

	return e.iSwapRouter2.UnpackGetAmountIn(resb)
}

func (e *SwapRouter2) GetAmountsOut(contract common.Address, amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	b, err := e.iSwapRouter2.GetAmountsOut(amountIn, path)
	if err != nil {
		return nil, err
	}
	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return nil, err
	}

	return e.iSwapRouter2.UnpackGetAmountsOut(resb)
}

func (e *SwapRouter2) GetAmountsIn(contract common.Address, amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	b, err := e.iSwapRouter2.GetAmountsIn(amountOut, path)
	if err != nil {
		return nil, err
	}
	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return nil, err
	}

	return e.iSwapRouter2.UnpackGetAmountsIn(resb)
}
