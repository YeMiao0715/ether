package swap_v2

import (
	"github.com/YeMiao0715/ether"
	"github.com/YeMiao0715/ether/erc20"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"
	"math/big"
)

type PairContract struct {
	contract common.Address
	*Pair
	*erc20.Erc20Contract

	cacheToken0Address *common.Address
	cacheToken1Address *common.Address
	cacheToken0Erc20   *erc20.Erc20Contract
	cacheToken1Erc20   *erc20.Erc20Contract
}

func NewPairContract(engine *ether.Engine, contract common.Address) *PairContract {
	return &PairContract{
		contract:      contract,
		Pair:          NewPair(engine),
		Erc20Contract: erc20.NewErc20WithContract(engine, contract),
	}
}

func (p *PairContract) Contract() common.Address {
	return p.contract
}

func (p *PairContract) Factory() (common.Address, error) {
	return p.Pair.Factory(p.contract)
}

func (p *PairContract) Token0() (common.Address, error) {
	if p.cacheToken0Address == nil {
		token0, err := p.Pair.Token0(p.contract)
		if err != nil {
			return common.Address{}, err
		}
		p.cacheToken0Address = &token0
	}
	return *p.cacheToken0Address, nil
}
func (p *PairContract) Token0Contract() (*erc20.Erc20Contract, error) {
	if p.cacheToken0Erc20 == nil {
		token0, err := p.Token0()
		if err != nil {
			return nil, err
		}
		p.cacheToken0Erc20 = erc20.NewErc20WithContract(p.engine, token0)
	}
	return p.cacheToken0Erc20, nil
}

func (p *PairContract) Token1() (common.Address, error) {
	if p.cacheToken1Address == nil {
		token1, err := p.Pair.Token1(p.contract)
		if err != nil {
			return common.Address{}, err
		}
		p.cacheToken1Address = &token1
	}
	return *p.cacheToken1Address, nil
}
func (p *PairContract) Token1Contract() (*erc20.Erc20Contract, error) {
	if p.cacheToken1Erc20 == nil {
		token1, err := p.Token1()
		if err != nil {
			return nil, err
		}
		p.cacheToken1Erc20 = erc20.NewErc20WithContract(p.engine, token1)
	}
	return p.cacheToken1Erc20, nil
}

func (p *PairContract) Price0CumulativeLast() (*big.Int, error) {
	return p.Pair.Price0CumulativeLast(p.contract)
}

func (p *PairContract) Price1CumulativeLast() (*big.Int, error) {
	return p.Pair.Price1CumulativeLast(p.contract)
}

func (p *PairContract) KLast() (*big.Int, error) {
	return p.Pair.KLast(p.contract)
}

func (p *PairContract) GetReserves() (reserve0, reserve1 *big.Int, blockTimestampLast uint32, err error) {
	return p.Pair.GetReserves(p.contract)
}

func (p *PairContract) DOMAIN_SEPARATOR() (common.Hash, error) {
	return p.Pair.DOMAIN_SEPARATOR(p.contract)
}

func (p *PairContract) PERMIT_TYPEHASH() (common.Hash, error) {
	return p.Pair.PERMIT_TYPEHASH(p.contract)
}

func (p *PairContract) Nonces(addr common.Address) (*big.Int, error) {
	return p.Pair.Nonces(p.contract, addr)
}

func (p *PairContract) PermitSign(sender common.Address, amount *big.Int, deadline *big.Int, privateKey string) (v, r, s *big.Int, err error) {
	PERMIT_TYPEHASH, err := p.Pair.PERMIT_TYPEHASH(p.contract)
	if err != nil {
		return
	}
	DOMAIN_SEPARATOR, err := p.Pair.DOMAIN_SEPARATOR(p.contract)
	if err != nil {
		return
	}
	bytes32, _ := abi.NewType("bytes32", "", nil)
	address, _ := abi.NewType("address", "", nil)
	uint256, _ := abi.NewType("uint256", "", nil)
	args := abi.Arguments{
		{Name: "PERMIT_TYPEHASH", Type: bytes32},
		{Name: "owner", Type: address},
		{Name: "spender", Type: address},
		{Name: "permit_value", Type: uint256},
		{Name: "nonce", Type: uint256},
		{Name: "deadline", Type: uint256},
	}

	owner, err := p.engine.PrivateKeyToAddress(privateKey)
	if err != nil {
		return
	}
	nonce, err := p.Nonces(*owner)
	if err != nil {
		return
	}
	b, err := args.Pack(PERMIT_TYPEHASH, *owner, sender, amount, nonce, deadline)
	if err != nil {
		return
	}

	DETAIL_HASH := crypto.Keccak256(b)
	//bytes, _ := abi.NewType("bytes", "", nil)
	//args = abi.Arguments{
	//	{Name: "", Type: bytes},
	//	{Name: "DOMAIN_SEPARATOR", Type: bytes},
	//	{Name: "DETAIL_HASH", Type: bytes},
	//}
	//fmt.Println(args)
	//b, err = args.Pack([]byte("\\x19\\x01"), DOMAIN_SEPARATOR[:], DETAIL_HASH)
	//if err != nil {
	//	return
	//}
	hash := crypto.Keccak256Hash([]byte("\x19\x01"), DOMAIN_SEPARATOR[:], DETAIL_HASH)

	priv, err := p.engine.HexToEcdsaPrivateKey(privateKey)
	if err != nil {
		return
	}

	sig, err := crypto.Sign(hash.Bytes(), priv)
	if err != nil {
		return
	}

	r = new(big.Int).SetBytes(sig[:32])
	s = new(big.Int).SetBytes(sig[32:64])
	v = new(big.Int).SetBytes([]byte{sig[64] + 27})

	return
}

// Quote 实现：https://github.com/Uniswap/v2-periphery/blob/master/contracts/libraries/UniswapV2Library.sol
func (p *PairContract) Quote(amountA decimal.Decimal, tokenA common.Address) (amountB *big.Int, err error) {
	reserve0, reserve1, _, err := p.GetReserves()
	if err != nil {
		return
	}
	token0, err := p.Token0()
	if err != nil {
		return
	}
	if tokenA != token0 {
		reserve0, reserve1 = reserve1, reserve0
	}
	reserveA := decimal.NewFromBigInt(reserve0, 0)
	reserveB := decimal.NewFromBigInt(reserve1, 0)
	amountB = amountA.Mul(reserveB).Div(reserveA).BigInt()
	return
}

// GetAmountOut 实现：https://github.com/Uniswap/v2-periphery/blob/master/contracts/libraries/UniswapV2Library.sol
func (p *PairContract) GetAmountOut(amountInBigInt *big.Int, tokenA common.Address) (amountOut *big.Int, err error) {
	reserve0, reserve1, _, err := p.GetReserves()
	if err != nil {
		return
	}
	token0, err := p.Token0()
	if err != nil {
		return
	}
	if tokenA != token0 {
		reserve0, reserve1 = reserve1, reserve0
	}

	amountIn := decimal.NewFromBigInt(amountInBigInt, 0)
	reserveIn := decimal.NewFromBigInt(reserve0, 0)
	reserveOut := decimal.NewFromBigInt(reserve1, 0)
	amountInWithFee := amountIn.Mul(decimal.NewFromInt(997))
	numerator := amountInWithFee.Mul(reserveOut)
	denominator := reserveIn.Mul(decimal.NewFromInt(1000)).Add(amountInWithFee)
	amountOut = numerator.Div(denominator).Truncate(0).BigInt()
	return
}

// GetAmountIn 实现：https://github.com/Uniswap/v2-periphery/blob/master/contracts/libraries/UniswapV2Library.sol
func (p *PairContract) GetAmountIn(amountOutBigInt *big.Int, tokenA common.Address) (amountIn *big.Int, err error) {
	reserve0, reserve1, _, err := p.GetReserves()
	if err != nil {
		return
	}
	token0, err := p.Token0()
	if err != nil {
		return
	}
	if tokenA != token0 {
		reserve0, reserve1 = reserve1, reserve0
	}
	reserveIn := decimal.NewFromBigInt(reserve0, 0)
	reserveOut := decimal.NewFromBigInt(reserve1, 0)
	amountOut := decimal.NewFromBigInt(amountOutBigInt, 0)
	numerator := reserveIn.Mul(amountOut).Mul(decimal.NewFromInt(1000))
	denominator := reserveOut.Sub(amountOut).Mul(decimal.NewFromInt(997))
	amountIn = numerator.Div(denominator).Truncate(0).Add(decimal.NewFromInt(1)).BigInt()
	return
}
