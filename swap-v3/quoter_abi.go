package swap_v3

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"math/big"
	"strings"
)

var iQuoterAbiJsonStr = `[{"type":"constructor","inputs":[{"name":"_factory","type":"address","internalType":"address"}],"stateMutability":"nonpayable"},{"type":"function","name":"factory","inputs":[],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"quoteExactInput","inputs":[{"name":"path","type":"bytes","internalType":"bytes"},{"name":"amountIn","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"amountOut","type":"uint256","internalType":"uint256"},{"name":"sqrtPriceX96AfterList","type":"uint160[]","internalType":"uint160[]"},{"name":"initializedTicksCrossedList","type":"uint32[]","internalType":"uint32[]"},{"name":"gasEstimate","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"quoteExactInputSingle","inputs":[{"name":"params","type":"tuple","internalType":"struct IQuoter.QuoteExactInputSingleParams","components":[{"name":"tokenIn","type":"address","internalType":"address"},{"name":"tokenOut","type":"address","internalType":"address"},{"name":"amountIn","type":"uint256","internalType":"uint256"},{"name":"fee","type":"uint24","internalType":"uint24"},{"name":"sqrtPriceLimitX96","type":"uint160","internalType":"uint160"}]}],"outputs":[{"name":"amountReceived","type":"uint256","internalType":"uint256"},{"name":"sqrtPriceX96After","type":"uint160","internalType":"uint160"},{"name":"initializedTicksCrossed","type":"uint32","internalType":"uint32"},{"name":"gasEstimate","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"quoteExactInputSingleWithPool","inputs":[{"name":"params","type":"tuple","internalType":"struct IQuoter.QuoteExactInputSingleWithPoolParams","components":[{"name":"tokenIn","type":"address","internalType":"address"},{"name":"tokenOut","type":"address","internalType":"address"},{"name":"amountIn","type":"uint256","internalType":"uint256"},{"name":"pool","type":"address","internalType":"address"},{"name":"fee","type":"uint24","internalType":"uint24"},{"name":"sqrtPriceLimitX96","type":"uint160","internalType":"uint160"}]}],"outputs":[{"name":"amountReceived","type":"uint256","internalType":"uint256"},{"name":"sqrtPriceX96After","type":"uint160","internalType":"uint160"},{"name":"initializedTicksCrossed","type":"uint32","internalType":"uint32"},{"name":"gasEstimate","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"quoteExactOutput","inputs":[{"name":"path","type":"bytes","internalType":"bytes"},{"name":"amountOut","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"amountIn","type":"uint256","internalType":"uint256"},{"name":"sqrtPriceX96AfterList","type":"uint160[]","internalType":"uint160[]"},{"name":"initializedTicksCrossedList","type":"uint32[]","internalType":"uint32[]"},{"name":"gasEstimate","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"quoteExactOutputSingle","inputs":[{"name":"params","type":"tuple","internalType":"struct IQuoter.QuoteExactOutputSingleParams","components":[{"name":"tokenIn","type":"address","internalType":"address"},{"name":"tokenOut","type":"address","internalType":"address"},{"name":"amount","type":"uint256","internalType":"uint256"},{"name":"fee","type":"uint24","internalType":"uint24"},{"name":"sqrtPriceLimitX96","type":"uint160","internalType":"uint160"}]}],"outputs":[{"name":"amountIn","type":"uint256","internalType":"uint256"},{"name":"sqrtPriceX96After","type":"uint160","internalType":"uint160"},{"name":"initializedTicksCrossed","type":"uint32","internalType":"uint32"},{"name":"gasEstimate","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"quoteExactOutputSingleWithPool","inputs":[{"name":"params","type":"tuple","internalType":"struct IQuoter.QuoteExactOutputSingleWithPoolParams","components":[{"name":"tokenIn","type":"address","internalType":"address"},{"name":"tokenOut","type":"address","internalType":"address"},{"name":"amount","type":"uint256","internalType":"uint256"},{"name":"fee","type":"uint24","internalType":"uint24"},{"name":"pool","type":"address","internalType":"address"},{"name":"sqrtPriceLimitX96","type":"uint160","internalType":"uint160"}]}],"outputs":[{"name":"amountIn","type":"uint256","internalType":"uint256"},{"name":"sqrtPriceX96After","type":"uint160","internalType":"uint160"},{"name":"initializedTicksCrossed","type":"uint32","internalType":"uint32"},{"name":"gasEstimate","type":"uint256","internalType":"uint256"}],"stateMutability":"view"}]`

var IQuoterAbi = &IQuoter{}

// IQuoter https://github.com/Uniswap/view-quoter-v3/blob/master/contracts/Quoter.sol
// https://etherscan.io/address/0x4752ba5DBc23f44D87826276BF6Fd6b1C372aD24#code
// https://github.com/Uniswap/view-quoter-v3
type IQuoter struct {
	abi *abi.ABI
}

func (i *IQuoter) GetAbi() (*abi.ABI, error) {
	if i.abi == nil {
		_erc721Abi, err := abi.JSON(strings.NewReader(iQuoterAbiJsonStr))
		if err != nil {
			return nil, err
		}
		i.abi = &_erc721Abi
	}
	return i.abi, nil
}

func (i *IQuoter) Method(fn string, param ...interface{}) ([]byte, error) {
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

func (i *IQuoter) Unpack(fn string, data []byte) ([]interface{}, error) {
	contractAbi, err := i.GetAbi()
	if err != nil {
		return nil, err
	}

	return contractAbi.Unpack(fn, data)
}

func (i *IQuoter) MustAbi() *abi.ABI {
	targetAbi, _ := i.GetAbi()
	return targetAbi
}

func (i *IQuoter) UnpackFactory(data []byte) (common.Address, error) {
	result, err := i.Unpack("factory", data)
	if err != nil {
		return common.Address{}, err
	}

	return result[0].(common.Address), nil
}

type QuoteExactInputSingleWithPoolParams struct {
	TokenIn           common.Address `json:"tokenIn"`
	TokenOut          common.Address `json:"tokenOut"`
	AmountIn          *big.Int       `json:"amountIn"`
	Fee               *big.Int       `json:"fee"`
	Pool              common.Address `json:"pool"`
	SqrtPriceLimitX96 *big.Int       `json:"sqrtPriceLimitX96"`
}

func (i *IQuoter) QuoteExactInputSingleWithPool(params QuoteExactInputSingleWithPoolParams) ([]byte, error) {
	return i.Method("quoteExactInputSingleWithPool", params)
}

func (i *IQuoter) UnpackQuoteExactInputSingleWithPool(data []byte) (amountOut, sqrtPriceX96After *big.Int, initializedTicksCrossed uint32, gasEstimate *big.Int, err error) {
	result, err := i.Unpack("quoteExactInputSingleWithPool", data)
	if err != nil {
		return
	}

	amountOut = result[0].(*big.Int)
	sqrtPriceX96After = result[1].(*big.Int)
	initializedTicksCrossed = result[2].(uint32)
	gasEstimate = result[3].(*big.Int)
	return
}

type QuoteExactInputSingleParams struct {
	TokenIn           common.Address `json:"tokenIn"`
	TokenOut          common.Address `json:"tokenOut"`
	AmountIn          *big.Int       `json:"amountIn"`
	Fee               *big.Int       `json:"fee"`
	SqrtPriceLimitX96 *big.Int       `json:"sqrtPriceLimitX96"`
}

func (i *IQuoter) QuoteExactInputSingle(params QuoteExactInputSingleParams) ([]byte, error) {
	return i.Method("quoteExactInputSingle", params)
}

func (i *IQuoter) UnpackQuoteExactInputSingle(data []byte) (amountOut, sqrtPriceX96After *big.Int, initializedTicksCrossed uint32, gasEstimate *big.Int, err error) {
	result, err := i.Unpack("quoteExactInputSingle", data)
	if err != nil {
		return
	}
	amountOut = result[0].(*big.Int)
	sqrtPriceX96After = result[1].(*big.Int)
	initializedTicksCrossed = result[2].(uint32)
	gasEstimate = result[3].(*big.Int)
	return
}

type QuoteExactOutputSingleWithPoolParams struct {
	TokenIn           common.Address `json:"tokenIn"`
	TokenOut          common.Address `json:"tokenOut"`
	Amount            *big.Int       `json:"amount"`
	Fee               *big.Int       `json:"fee"`
	Pool              common.Address `json:"pool"`
	SqrtPriceLimitX96 *big.Int       `json:"sqrtPriceLimitX96"`
}

func (i *IQuoter) QuoteExactOutputSingleWithPool(params QuoteExactOutputSingleWithPoolParams) ([]byte, error) {
	return i.Method("quoteExactOutputSingleWithPool", params)
}

func (i *IQuoter) UnpackQuoteExactOutputSingleWithPool(data []byte) (amountIn *big.Int, sqrtPriceX96After *big.Int, initializedTicksCrossed uint32, gasEstimate *big.Int, err error) {
	result, err := i.Unpack("quoteExactOutputSingleWithPool", data)
	if err != nil {
		return
	}

	amountIn = result[0].(*big.Int)
	sqrtPriceX96After = result[1].(*big.Int)
	initializedTicksCrossed = result[2].(uint32)
	gasEstimate = result[3].(*big.Int)
	return
}

type QuoteExactOutputSingleParams struct {
	TokenIn           common.Address `json:"tokenIn"`
	TokenOut          common.Address `json:"tokenOut"`
	Amount            *big.Int       `json:"amount"`
	Fee               *big.Int       `json:"fee"`
	SqrtPriceLimitX96 *big.Int       `json:"sqrtPriceLimitX96"`
}

func (i *IQuoter) QuoteExactOutputSingle(params QuoteExactOutputSingleParams) ([]byte, error) {
	return i.Method("quoteExactOutputSingle", params)
}

func (i *IQuoter) UnpackQuoteExactOutputSingle(data []byte) (amountIn *big.Int, sqrtPriceX96After *big.Int, initializedTicksCrossed uint32, gasEstimate *big.Int, err error) {
	result, err := i.Unpack("quoteExactOutputSingle", data)
	if err != nil {
		return
	}

	amountIn = result[0].(*big.Int)
	sqrtPriceX96After = result[1].(*big.Int)
	initializedTicksCrossed = result[2].(uint32)
	gasEstimate = result[3].(*big.Int)
	return
}

func (i *IQuoter) QuoteExactInput(paths []common.Address, amountIn *big.Int) ([]byte, error) {
	path := make([]byte, 0)
	for _, item := range paths {
		path = append(path, common.HexToHash(item.Hex()).Bytes()...)
	}
	return i.Method("quoteExactInput", path, amountIn)
}

func (i *IQuoter) UnpackQuoteExactInput(data []byte) (amountOut *big.Int, sqrtPriceX96AfterList []*big.Int, initializedTicksCrossedList []*big.Int, gasEstimate *big.Int, err error) {
	result, err := i.Unpack("quoteExactInput", data)
	if err != nil {
		return
	}

	amountOut = result[0].(*big.Int)
	sqrtPriceX96AfterList = result[1].([]*big.Int)
	initializedTicksCrossedList = result[2].([]*big.Int)
	gasEstimate = result[3].(*big.Int)
	return
}

func (i *IQuoter) QuoteExactOutput(paths []common.Address, amountOut *big.Int) ([]byte, error) {
	path := make([]byte, 0)
	for _, item := range paths {
		path = append(path, item.Bytes()...)
	}
	return i.Method("quoteExactOutput", path, amountOut)
}

func (i *IQuoter) UnpackQuoteExactOutput(data []byte) (amountIn *big.Int, sqrtPriceX96AfterList []*big.Int, initializedTicksCrossedList []*big.Int, gasEstimate *big.Int, err error) {
	result, err := i.Unpack("quoteExactOutput", data)
	if err != nil {
		return
	}

	amountIn = result[0].(*big.Int)
	sqrtPriceX96AfterList = result[1].([]*big.Int)
	initializedTicksCrossedList = result[2].([]*big.Int)
	gasEstimate = result[3].(*big.Int)
	return
}
