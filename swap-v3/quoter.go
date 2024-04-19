package swap_v3

import (
	"github.com/YeMiao0715/ether"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Quoter struct {
	engine  *ether.Engine
	IQuoter *IQuoter
}

func NewQuoter(engine *ether.Engine) *Quoter {
	return &Quoter{
		engine:  engine,
		IQuoter: IQuoterAbi,
	}
}

func (q *Quoter) QuoteExactInputSingleWithPool(contractAddress common.Address, params QuoteExactInputSingleWithPoolParams) (amountOut, sqrtPriceX96After *big.Int, initializedTicksCrossed uint32, gasEstimate *big.Int, err error) {
	data, err := q.IQuoter.QuoteExactInputSingleWithPool(params)
	if err != nil {
		return
	}
	res, err := q.engine.CallContract(contractAddress, data)
	if err != nil {
		return
	}

	return q.IQuoter.UnpackQuoteExactInputSingleWithPool(res)
}

func (q *Quoter) QuoteExactInputSingle(contractAddress common.Address, params QuoteExactInputSingleParams) (amountOut, sqrtPriceX96After *big.Int, initializedTicksCrossed uint32, gasEstimate *big.Int, err error) {
	data, err := q.IQuoter.QuoteExactInputSingle(params)
	if err != nil {
		return
	}

	res, err := q.engine.CallContract(contractAddress, data)
	if err != nil {
		return
	}

	return q.IQuoter.UnpackQuoteExactInputSingle(res)
}

func (q *Quoter) QuoteExactOutputSingleWithPool(contractAddress common.Address, params QuoteExactOutputSingleWithPoolParams) (amountIn *big.Int, sqrtPriceX96After *big.Int, initializedTicksCrossed uint32, gasEstimate *big.Int, err error) {
	data, err := q.IQuoter.QuoteExactOutputSingleWithPool(params)
	if err != nil {
		return
	}
	res, err := q.engine.CallContract(contractAddress, data)
	if err != nil {
		return
	}

	return q.IQuoter.UnpackQuoteExactOutputSingleWithPool(res)
}

func (q *Quoter) QuoteExactOutputSingle(contractAddress common.Address, params QuoteExactOutputSingleParams) (amountIn *big.Int, sqrtPriceX96After *big.Int, initializedTicksCrossed uint32, gasEstimate *big.Int, err error) {
	data, err := q.IQuoter.QuoteExactOutputSingle(params)
	if err != nil {
		return
	}
	res, err := q.engine.CallContract(contractAddress, data)
	if err != nil {
		return
	}

	return q.IQuoter.UnpackQuoteExactOutputSingle(res)
}

func (q *Quoter) QuoteExactInput(contractAddress common.Address, paths []common.Address, amountIn *big.Int) (amountOut *big.Int, sqrtPriceX96AfterList []*big.Int, initializedTicksCrossedList []*big.Int, gasEstimate *big.Int, err error) {
	data, err := q.IQuoter.QuoteExactInput(paths, amountIn)
	if err != nil {
		return
	}
	res, err := q.engine.CallContract(contractAddress, data)
	if err != nil {
		return
	}
	return q.IQuoter.UnpackQuoteExactInput(res)
}

func (q *Quoter) QuoteExactOutput(contractAddress common.Address, paths []common.Address, amountOut *big.Int) (amountIn *big.Int, sqrtPriceX96AfterList []*big.Int, initializedTicksCrossedList []*big.Int, gasEstimate *big.Int, err error) {
	data, err := q.IQuoter.QuoteExactOutput(paths, amountOut)
	if err != nil {
		return
	}
	res, err := q.engine.CallContract(contractAddress, data)
	if err != nil {
		return
	}

	return q.IQuoter.UnpackQuoteExactOutput(res)
}
