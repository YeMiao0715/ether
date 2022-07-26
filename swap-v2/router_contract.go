package swap_v2

import (
	"github.com/YeMiao0715/ether"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type RouterContract struct {
	contract common.Address
	*Router
}

func NewRouterContract(engine *ether.Engine, contract common.Address) *RouterContract {
	return &RouterContract{
		contract: contract,
		Router:   NewRouter2(engine),
	}
}

func (c *RouterContract) Contract() common.Address {
	return c.contract
}

func (c *RouterContract) Factory() (common.Address, error) {
	return c.Router.Factory(c.contract)
}

func (c *RouterContract) WETH() (common.Address, error) {
	return c.Router.WETH(c.contract)
}

func (c *RouterContract) AddLiquidity(
	tokenA,
	tokenB common.Address,
	amountADesired,
	amountBDesired,
	amountAMin,
	amountBMin *big.Int,
	to common.Address,
	deadline *big.Int,
	privateKey string) (string, *types.Transaction, error) {
	return c.Router.AddLiquidity(c.contract, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline, privateKey)
}

func (c *RouterContract) AddLiquidityETH(
	token common.Address,
	amountTokenDesired,
	amountTokenMin,
	amountETHMin *big.Int,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	return c.Router.AddLiquidityETH(c.contract, token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline, privateKey)
}

func (c *RouterContract) RemoveLiquidity(
	tokenA,
	tokenB common.Address,
	liquidity,
	amountAMin,
	amountBMin *big.Int,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	return c.Router.RemoveLiquidity(c.contract, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, privateKey)
}

func (c *RouterContract) RemoveLiquidityETH(
	token common.Address,
	liquidity,
	amountTokenMin,
	amountETHMin *big.Int,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	return c.Router.RemoveLiquidityETH(c.contract, token, liquidity, amountTokenMin, amountETHMin, to, deadline, privateKey)
}

func (c *RouterContract) RemoveLiquidityWithPermit(
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
	return c.Router.RemoveLiquidityWithPermit(c.contract, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s, privateKey)
}

func (c *RouterContract) RemoveLiquidityETHWithPermit(
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
	return c.Router.RemoveLiquidityETHWithPermit(c.contract, token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s, privateKey)
}

func (c *RouterContract) SwapExactTokensForTokens(
	amountIn,
	amountOutMin *big.Int,
	path []common.Address,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	return c.Router.SwapExactTokensForTokens(c.contract, amountIn, amountOutMin, path, to, deadline, privateKey)
}

func (c *RouterContract) SwapTokensForExactTokens(
	amountOut,
	amountInMax *big.Int,
	path []common.Address,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	return c.Router.SwapTokensForExactTokens(c.contract, amountOut, amountInMax, path, to, deadline, privateKey)
}

func (c *RouterContract) SwapExactETHForTokens(
	amountOutMin *big.Int,
	path []common.Address,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	return c.Router.SwapExactETHForTokens(c.contract, amountOutMin, path, to, deadline, privateKey)
}

func (c *RouterContract) SwapTokensForExactETH(
	amountOut,
	amountInMax *big.Int,
	path []common.Address,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	return c.Router.SwapTokensForExactETH(c.contract, amountOut, amountInMax, path, to, deadline, privateKey)
}

func (c *RouterContract) SwapExactTokensForETH(
	amountIn,
	amountOutMin *big.Int,
	path []common.Address,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	return c.Router.SwapExactTokensForETH(c.contract, amountIn, amountOutMin, path, to, deadline, privateKey)
}

func (c *RouterContract) SwapETHForExactTokens(
	amountOut *big.Int,
	path []common.Address,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	return c.Router.SwapETHForExactTokens(c.contract, amountOut, path, to, deadline, privateKey)
}

func (c *RouterContract) Quote(amountA, reserveA, reserveB *big.Int) (*big.Int, error) {
	return c.Router.Quote(c.contract, amountA, reserveA, reserveB)
}

func (c *RouterContract) GetAmountOut(amountIn, reserveIn, reserveOut *big.Int) (*big.Int, error) {
	return c.Router.GetAmountOut(c.contract, amountIn, reserveIn, reserveOut)
}

func (c *RouterContract) GetAmountIn(amountOut, reserveIn, reserveOut *big.Int) (*big.Int, error) {
	return c.Router.GetAmountIn(c.contract, amountOut, reserveIn, reserveOut)
}

func (c *RouterContract) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return c.Router.GetAmountsOut(c.contract, amountIn, path)
}

func (c *RouterContract) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return c.Router.GetAmountsIn(c.contract, amountOut, path)
}
