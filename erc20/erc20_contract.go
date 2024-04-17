package erc20

import (
	"github.com/YeMiao0715/ether"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
	"math/big"
)

type Erc20Contract struct {
	contract      common.Address
	erc20         *Erc20
	cacheName     *string
	cacheSymbol   *string
	cacheDecimals *uint8
}

func NewErc20WithContract(engine *ether.Engine, contract common.Address) *Erc20Contract {
	return &Erc20Contract{
		contract: contract,
		erc20:    NewErc20(engine),
	}
}

func (e *Erc20Contract) Contract() common.Address {
	return e.contract
}

func (e *Erc20Contract) Name() (string, error) {
	if e.cacheName == nil {
		name, err := e.erc20.Name(e.contract)
		if err != nil {
			return "", err
		}
		e.cacheName = &name
	}
	return *e.cacheName, nil
}

func (e *Erc20Contract) Symbol() (string, error) {
	if e.cacheSymbol == nil {
		symbol, err := e.erc20.Symbol(e.contract)
		if err != nil {
			return "", err
		}
		e.cacheSymbol = &symbol
	}
	return *e.cacheSymbol, nil
}

func (e *Erc20Contract) Decimals() (uint8, error) {
	if e.cacheDecimals == nil {
		decimals, err := e.erc20.Decimals(e.contract)
		if err != nil {
			return 0, err
		}
		e.cacheDecimals = &decimals
	}
	return *e.cacheDecimals, nil
}

func (e *Erc20Contract) TotalSupply() (*big.Int, error) {
	return e.erc20.TotalSupply(e.contract)
}

func (e *Erc20Contract) BalanceOf(address common.Address) (*big.Int, error) {
	return e.erc20.BalanceOf(e.contract, address)
}

func (e *Erc20Contract) Approve(spender common.Address, amount *big.Int, privateKey string) (string, *types.Transaction, error) {
	owner, err := e.erc20.engine.PrivateKeyToAddress(privateKey)
	if err != nil {
		return "", nil, err
	}
	return e.erc20.Approve(e.contract, *owner, spender, amount, privateKey)
}

func (e *Erc20Contract) Allowance(owner, spender common.Address) (*big.Int, error) {
	return e.erc20.Allowance(e.contract, owner, spender)
}

func (e *Erc20Contract) Transfer(recipient common.Address, amount *big.Int, privateKey string) (string, *types.Transaction, error) {
	return e.erc20.Transfer(e.contract, recipient, amount, privateKey)
}

func (e *Erc20Contract) TransferFrom(from, recipient common.Address, amount *big.Int, privateKey string) (string, *types.Transaction, error) {
	return e.erc20.TransferFrom(e.contract, from, recipient, amount, privateKey)
}

type Erc20Info struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals uint8  `json:"decimals"`
}

func (e *Erc20Contract) Info() (*Erc20Info, error) {
	name, err := e.Name()
	if err != nil {
		return nil, err
	}
	symbol, err := e.Symbol()
	if err != nil {
		return nil, err
	}
	decimals, err := e.Decimals()
	if err != nil {
		return nil, err
	}

	return &Erc20Info{
		Name:     name,
		Symbol:   symbol,
		Decimals: decimals,
	}, nil
}

func (e *Erc20Contract) ToAmount(amount float64) (decimal.Decimal, error) {
	decimals, err := e.Decimals()
	if err != nil {
		return decimal.Zero, nil
	}
	return decimal.NewFromFloat(amount).Mul(decimal.New(1, int32(decimals))), nil
}
