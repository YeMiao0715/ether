package swap_v2

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"math/big"
	"strings"
)

const routerAbiJson = `[{"inputs":[],"name":"WETH","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"pure","type":"function"},{"inputs":[{"internalType":"address","name":"tokenA","type":"address"},{"internalType":"address","name":"tokenB","type":"address"},{"internalType":"uint256","name":"amountADesired","type":"uint256"},{"internalType":"uint256","name":"amountBDesired","type":"uint256"},{"internalType":"uint256","name":"amountAMin","type":"uint256"},{"internalType":"uint256","name":"amountBMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"addLiquidity","outputs":[{"internalType":"uint256","name":"amountA","type":"uint256"},{"internalType":"uint256","name":"amountB","type":"uint256"},{"internalType":"uint256","name":"liquidity","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"token","type":"address"},{"internalType":"uint256","name":"amountTokenDesired","type":"uint256"},{"internalType":"uint256","name":"amountTokenMin","type":"uint256"},{"internalType":"uint256","name":"amountETHMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"addLiquidityETH","outputs":[{"internalType":"uint256","name":"amountToken","type":"uint256"},{"internalType":"uint256","name":"amountETH","type":"uint256"},{"internalType":"uint256","name":"liquidity","type":"uint256"}],"stateMutability":"payable","type":"function"},{"inputs":[],"name":"factory","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"pure","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountOut","type":"uint256"},{"internalType":"uint256","name":"reserveIn","type":"uint256"},{"internalType":"uint256","name":"reserveOut","type":"uint256"}],"name":"getAmountIn","outputs":[{"internalType":"uint256","name":"amountIn","type":"uint256"}],"stateMutability":"pure","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountIn","type":"uint256"},{"internalType":"uint256","name":"reserveIn","type":"uint256"},{"internalType":"uint256","name":"reserveOut","type":"uint256"}],"name":"getAmountOut","outputs":[{"internalType":"uint256","name":"amountOut","type":"uint256"}],"stateMutability":"pure","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountOut","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"}],"name":"getAmountsIn","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountIn","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"}],"name":"getAmountsOut","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountA","type":"uint256"},{"internalType":"uint256","name":"reserveA","type":"uint256"},{"internalType":"uint256","name":"reserveB","type":"uint256"}],"name":"quote","outputs":[{"internalType":"uint256","name":"amountB","type":"uint256"}],"stateMutability":"pure","type":"function"},{"inputs":[{"internalType":"address","name":"tokenA","type":"address"},{"internalType":"address","name":"tokenB","type":"address"},{"internalType":"uint256","name":"liquidity","type":"uint256"},{"internalType":"uint256","name":"amountAMin","type":"uint256"},{"internalType":"uint256","name":"amountBMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"removeLiquidity","outputs":[{"internalType":"uint256","name":"amountA","type":"uint256"},{"internalType":"uint256","name":"amountB","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"token","type":"address"},{"internalType":"uint256","name":"liquidity","type":"uint256"},{"internalType":"uint256","name":"amountTokenMin","type":"uint256"},{"internalType":"uint256","name":"amountETHMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"removeLiquidityETH","outputs":[{"internalType":"uint256","name":"amountToken","type":"uint256"},{"internalType":"uint256","name":"amountETH","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"token","type":"address"},{"internalType":"uint256","name":"liquidity","type":"uint256"},{"internalType":"uint256","name":"amountTokenMin","type":"uint256"},{"internalType":"uint256","name":"amountETHMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"removeLiquidityETHSupportingFeeOnTransferTokens","outputs":[{"internalType":"uint256","name":"amountETH","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"token","type":"address"},{"internalType":"uint256","name":"liquidity","type":"uint256"},{"internalType":"uint256","name":"amountTokenMin","type":"uint256"},{"internalType":"uint256","name":"amountETHMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"},{"internalType":"bool","name":"approveMax","type":"bool"},{"internalType":"uint8","name":"v","type":"uint8"},{"internalType":"bytes32","name":"r","type":"bytes32"},{"internalType":"bytes32","name":"s","type":"bytes32"}],"name":"removeLiquidityETHWithPermit","outputs":[{"internalType":"uint256","name":"amountToken","type":"uint256"},{"internalType":"uint256","name":"amountETH","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"token","type":"address"},{"internalType":"uint256","name":"liquidity","type":"uint256"},{"internalType":"uint256","name":"amountTokenMin","type":"uint256"},{"internalType":"uint256","name":"amountETHMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"},{"internalType":"bool","name":"approveMax","type":"bool"},{"internalType":"uint8","name":"v","type":"uint8"},{"internalType":"bytes32","name":"r","type":"bytes32"},{"internalType":"bytes32","name":"s","type":"bytes32"}],"name":"removeLiquidityETHWithPermitSupportingFeeOnTransferTokens","outputs":[{"internalType":"uint256","name":"amountETH","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"tokenA","type":"address"},{"internalType":"address","name":"tokenB","type":"address"},{"internalType":"uint256","name":"liquidity","type":"uint256"},{"internalType":"uint256","name":"amountAMin","type":"uint256"},{"internalType":"uint256","name":"amountBMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"},{"internalType":"bool","name":"approveMax","type":"bool"},{"internalType":"uint8","name":"v","type":"uint8"},{"internalType":"bytes32","name":"r","type":"bytes32"},{"internalType":"bytes32","name":"s","type":"bytes32"}],"name":"removeLiquidityWithPermit","outputs":[{"internalType":"uint256","name":"amountA","type":"uint256"},{"internalType":"uint256","name":"amountB","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountOut","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapETHForExactTokens","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"payable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountOutMin","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapExactETHForTokens","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"payable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountOutMin","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapExactETHForTokensSupportingFeeOnTransferTokens","outputs":[],"stateMutability":"payable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountIn","type":"uint256"},{"internalType":"uint256","name":"amountOutMin","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapExactTokensForETH","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountIn","type":"uint256"},{"internalType":"uint256","name":"amountOutMin","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapExactTokensForETHSupportingFeeOnTransferTokens","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountIn","type":"uint256"},{"internalType":"uint256","name":"amountOutMin","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapExactTokensForTokens","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountIn","type":"uint256"},{"internalType":"uint256","name":"amountOutMin","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapExactTokensForTokensSupportingFeeOnTransferTokens","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountOut","type":"uint256"},{"internalType":"uint256","name":"amountInMax","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapTokensForExactETH","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountOut","type":"uint256"},{"internalType":"uint256","name":"amountInMax","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapTokensForExactTokens","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"nonpayable","type":"function"}]`

type IRouter struct {
	abi *abi.ABI
}

func (i *IRouter) GetAbi() (*abi.ABI, error) {
	if i.abi == nil {
		_abi, err := abi.JSON(strings.NewReader(routerAbiJson))
		if err != nil {
			return nil, err
		}
		i.abi = &_abi
	}
	return i.abi, nil
}

func (i *IRouter) Method(fn string, param ...interface{}) ([]byte, error) {
	contractAbi, err := i.GetAbi()
	if err != nil {
		return nil, err
	}

	b, err := contractAbi.Pack(fn, param...)
	if err != nil {
		return nil, errors.Wrap(err, "abi pack err")
	}

	return b, nil
}

func (i *IRouter) WETH() ([]byte, error) {
	return i.Method("WETH")
}

func (i *IRouter) UnpackWETH(data []byte) (common.Address, error) {
	result, err := i.abi.Unpack("WETH", data)
	if err != nil {
		return common.Address{}, errors.WithStack(err)
	}
	return result[0].(common.Address), err
}

func (i *IRouter) Factory() ([]byte, error) {
	return i.Method("factory")
}

func (i *IRouter) UnpackFactory(data []byte) (common.Address, error) {
	result, err := i.abi.Unpack("factory", data)
	if err != nil {
		return common.Address{}, errors.WithStack(err)
	}
	return result[0].(common.Address), err
}

func (i *IRouter) AddLiquidity(
	tokenA,
	tokenB common.Address,
	amountADesired,
	amountBDesired,
	amountAMin,
	amountBMin *big.Int,
	to common.Address,
	deadline *big.Int,
) ([]byte, error) {
	return i.Method("addLiquidity", tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

func (i *IRouter) AddLiquidityETH(
	token common.Address,
	amountTokenDesired,
	amountTokenMin,
	amountETHMin *big.Int,
	to common.Address,
	deadline *big.Int,
) ([]byte, error) {
	return i.Method("addLiquidityETH", token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

func (i *IRouter) RemoveLiquidity(
	tokenA,
	tokenB common.Address,
	liquidity,
	amountAMin,
	amountBMin *big.Int,
	to common.Address,
	deadline *big.Int,
) ([]byte, error) {
	return i.Method("removeLiquidity", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

func (i *IRouter) RemoveLiquidityETH(
	token common.Address,
	liquidity,
	amountTokenMin,
	amountETHMin *big.Int,
	to common.Address,
	deadline *big.Int,
) ([]byte, error) {
	return i.Method("removeLiquidityETH", token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

func (i *IRouter) RemoveLiquidityWithPermit(
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
) ([]byte, error) {
	return i.Method("removeLiquidityWithPermit", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

func (i *IRouter) RemoveLiquidityETHWithPermit(
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
) ([]byte, error) {
	return i.Method("removeLiquidityETHWithPermit", token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

func (i *IRouter) SwapExactTokensForTokens(
	amountIn,
	amountOutMin *big.Int,
	path []common.Address,
	to common.Address,
	deadline *big.Int,
) ([]byte, error) {
	return i.Method("swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline)
}

func (i *IRouter) SwapTokensForExactTokens(
	amountOut,
	amountInMax *big.Int,
	path []common.Address,
	to common.Address,
	deadline *big.Int,
) ([]byte, error) {
	return i.Method("swapExactTokensForTokens", amountOut, amountInMax, path, to, deadline)
}

func (i *IRouter) SwapExactETHForTokens(
	amountOutMin *big.Int,
	path []common.Address,
	to common.Address,
	deadline *big.Int,
) ([]byte, error) {
	return i.Method("swapExactETHForTokens", amountOutMin, path, to, deadline)
}

func (i *IRouter) SwapTokensForExactETH(
	amountOut,
	amountInMax *big.Int,
	path []common.Address,
	to common.Address,
	deadline *big.Int,
) ([]byte, error) {
	return i.Method("swapTokensForExactETH", amountOut, amountInMax, path, to, deadline)
}

func (i *IRouter) SwapExactTokensForETH(
	amountIn,
	amountOutMin *big.Int,
	path []common.Address,
	to common.Address,
	deadline *big.Int,
) ([]byte, error) {
	return i.Method("swapExactTokensForETH", amountIn, amountOutMin, path, to, deadline)
}

func (i *IRouter) SwapETHForExactTokens(
	amountOut *big.Int,
	path []common.Address,
	to common.Address,
	deadline *big.Int,
) ([]byte, error) {
	return i.Method("swapETHForExactTokens", amountOut, path, to, deadline)
}

func (i *IRouter) SwapExactTokensForTokensSupportingFeeOnTransferTokens(
	amountIn,
	amountOutMin *big.Int,
	path []common.Address,
	to common.Address,
	deadline *big.Int,
) ([]byte, error) {
	return i.Method("swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
}

func (i *IRouter) Quote(amountA, reserveA, reserveB *big.Int) ([]byte, error) {
	return i.Method("quote", amountA, reserveA, reserveB)
}

func (i *IRouter) UnpackQuote(data []byte) (*big.Int, error) {
	result, err := i.abi.Unpack("quote", data)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return result[0].(*big.Int), err
}

func (i *IRouter) GetAmountOut(amountIn, reserveIn, reserveOut *big.Int) ([]byte, error) {
	return i.Method("getAmountOut", amountIn, reserveIn, reserveOut)
}

func (i *IRouter) UnpackGetAmountOut(data []byte) (*big.Int, error) {
	result, err := i.abi.Unpack("getAmountOut", data)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return result[0].(*big.Int), err
}

func (i *IRouter) GetAmountIn(amountOut, reserveIn, reserveOut *big.Int) ([]byte, error) {
	return i.Method("getAmountIn", amountOut, reserveIn, reserveOut)
}

func (i *IRouter) UnpackGetAmountIn(data []byte) (*big.Int, error) {
	result, err := i.abi.Unpack("getAmountIn", data)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return result[0].(*big.Int), err
}

func (i *IRouter) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]byte, error) {
	return i.Method("getAmountsOut", amountIn, path)
}

func (i *IRouter) UnpackGetAmountsOut(data []byte) (amounts []*big.Int, err error) {
	amounts = make([]*big.Int, 0)
	result, err := i.abi.Unpack("getAmountsOut", data)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	amounts = result[0].([]*big.Int)
	return
}

func (i *IRouter) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]byte, error) {
	return i.Method("getAmountsIn", amountOut, path)
}

func (i *IRouter) UnpackGetAmountsIn(data []byte) (amounts []*big.Int, err error) {
	amounts = make([]*big.Int, 0)
	result, err := i.abi.Unpack("getAmountsIn", data)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	amounts = result[0].([]*big.Int)
	return
}
