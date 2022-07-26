package swap

import (
	"github.com/YeMiao0715/ether"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type Router2Contract struct {
	contract common.Address
	*Router2
}

func NewRouter2Contract(engine *ether.Engine, contract common.Address) *Router2Contract {
	return &Router2Contract{
		contract: contract,
		Router2:  NewRouter2(engine),
	}
}

func (c *Router2Contract) Contract() common.Address {
	return c.contract
}

func (c *Router2Contract) Factory() (common.Address, error) {
	return c.Router2.Factory(c.contract)
}

func (c *Router2Contract) WETH() (common.Address, error) {
	return c.Router2.WETH(c.contract)
}

func (c *Router2Contract) AddLiquidity(
	tokenA,
	tokenB common.Address,
	amountADesired,
	amountBDesired,
	amountAMin,
	amountBMin *big.Int,
	to common.Address,
	deadline *big.Int,
	privateKey string) (string, *types.Transaction, error) {
	return c.Router2.AddLiquidity(c.contract, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline, privateKey)
}

func (c *Router2Contract) AddLiquidityETH(
	token common.Address,
	amountTokenDesired,
	amountTokenMin,
	amountETHMin *big.Int,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	return c.Router2.AddLiquidityETH(c.contract, token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline, privateKey)
}

func (c *Router2Contract) RemoveLiquidity(
	tokenA,
	tokenB common.Address,
	liquidity,
	amountAMin,
	amountBMin *big.Int,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	return c.Router2.RemoveLiquidity(c.contract, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, privateKey)
}

func (c *Router2Contract) RemoveLiquidityETH(
	token common.Address,
	liquidity,
	amountTokenMin,
	amountETHMin *big.Int,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	return c.Router2.RemoveLiquidityETH(c.contract, token, liquidity, amountTokenMin, amountETHMin, to, deadline, privateKey)
}

func (c *Router2Contract) RemoveLiquidityWithPermit(
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
	return c.Router2.RemoveLiquidityWithPermit(c.contract, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s, privateKey)
}

func (c *Router2Contract) RemoveLiquidityETHWithPermit(
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
	return c.Router2.RemoveLiquidityETHWithPermit(c.contract, token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s, privateKey)
}

func (c *Router2Contract) SwapExactTokensForTokens(
	amountIn,
	amountOutMin *big.Int,
	path []common.Address,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	return c.Router2.SwapExactTokensForTokens(c.contract, amountIn, amountOutMin, path, to, deadline, privateKey)
}

func (c *Router2Contract) SwapTokensForExactTokens(
	amountOut,
	amountInMax *big.Int,
	path []common.Address,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	return c.Router2.SwapTokensForExactTokens(c.contract, amountOut, amountInMax, path, to, deadline, privateKey)
}

func (c *Router2Contract) SwapExactETHForTokens(
	amountOutMin *big.Int,
	path []common.Address,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	return c.Router2.SwapExactETHForTokens(c.contract, amountOutMin, path, to, deadline, privateKey)
}

func (c *Router2Contract) SwapTokensForExactETH(
	amountOut,
	amountInMax *big.Int,
	path []common.Address,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	return c.Router2.SwapTokensForExactETH(c.contract, amountOut, amountInMax, path, to, deadline, privateKey)
}

func (c *Router2Contract) SwapExactTokensForETH(
	amountIn,
	amountOutMin *big.Int,
	path []common.Address,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	return c.Router2.SwapExactTokensForETH(c.contract, amountIn, amountOutMin, path, to, deadline, privateKey)
}

func (c *Router2Contract) SwapETHForExactTokens(
	amountOut *big.Int,
	path []common.Address,
	to common.Address,
	deadline *big.Int,
	privateKey string,
) (string, *types.Transaction, error) {
	return c.Router2.SwapETHForExactTokens(c.contract, amountOut, path, to, deadline, privateKey)
}

func (c *Router2Contract) Quote(amountA, reserveA, reserveB *big.Int) (*big.Int, error) {
	return c.Router2.Quote(c.contract, amountA, reserveA, reserveB)
}

func (c *Router2Contract) GetAmountOut(amountIn, reserveIn, reserveOut *big.Int) (*big.Int, error) {
	return c.Router2.GetAmountOut(c.contract, amountIn, reserveIn, reserveOut)
}

func (c *Router2Contract) GetAmountIn(amountOut, reserveIn, reserveOut *big.Int) (*big.Int, error) {
	return c.Router2.GetAmountIn(c.contract, amountOut, reserveIn, reserveOut)
}

func (c *Router2Contract) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return c.Router2.GetAmountsOut(c.contract, amountIn, path)
}

func (c *Router2Contract) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return c.Router2.GetAmountsIn(c.contract, amountOut, path)
}
