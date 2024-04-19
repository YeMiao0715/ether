package swap_v3

import (
	"github.com/YeMiao0715/ether"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type QuoterContract struct {
	engine *ether.Engine
	*Quoter
	contractAddress common.Address
}

func NewQuoterContract(engine *ether.Engine, contractAddress common.Address) *QuoterContract {
	return &QuoterContract{
		engine:          engine,
		Quoter:          NewQuoter(engine),
		contractAddress: contractAddress,
	}
}

func (q *QuoterContract) QuoteExactInputSingleWithPool(params QuoteExactInputSingleWithPoolParams) (amountOut, sqrtPriceX96After *big.Int, initializedTicksCrossed uint32, gasEstimate *big.Int, err error) {
	return q.Quoter.QuoteExactInputSingleWithPool(q.contractAddress, params)
}

func (q *QuoterContract) QuoteExactInputSingle(params QuoteExactInputSingleParams) (amountOut, sqrtPriceX96After *big.Int, initializedTicksCrossed uint32, gasEstimate *big.Int, err error) {
	return q.Quoter.QuoteExactInputSingle(q.contractAddress, params)
}

func (q *QuoterContract) QuoteExactOutputSingleWithPool(params QuoteExactOutputSingleWithPoolParams) (amountIn *big.Int, sqrtPriceX96After *big.Int, initializedTicksCrossed uint32, gasEstimate *big.Int, err error) {
	return q.Quoter.QuoteExactOutputSingleWithPool(q.contractAddress, params)
}

func (q *QuoterContract) QuoteExactOutputSingle(params QuoteExactOutputSingleParams) (amountIn *big.Int, sqrtPriceX96After *big.Int, initializedTicksCrossed uint32, gasEstimate *big.Int, err error) {
	return q.Quoter.QuoteExactOutputSingle(q.contractAddress, params)
}

func (q *QuoterContract) QuoteExactInput(paths []common.Address, amountIn *big.Int) (amountOut *big.Int, sqrtPriceX96AfterList []*big.Int, initializedTicksCrossedList []*big.Int, gasEstimate *big.Int, err error) {
	return q.Quoter.QuoteExactInput(q.contractAddress, paths, amountIn)
}

func (q *QuoterContract) QuoteExactOutput(paths []common.Address, amountOut *big.Int) (amountIn *big.Int, sqrtPriceX96AfterList []*big.Int, initializedTicksCrossedList []*big.Int, gasEstimate *big.Int, err error) {
	return q.Quoter.QuoteExactOutput(q.contractAddress, paths, amountOut)
}
