package erc20

import (
	"github.com/YeMiao0715/ether"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type Erc20Contract struct {
	contract common.Address
	erc20    *Erc20
}

func NewErc20WithContract(engine *ether.Engine, contract common.Address) *Erc20Contract {
	return &Erc20Contract{
		contract: contract,
		erc20:    NewErc20(engine),
	}
}

func (e *Erc20Contract) Name() (string, error) {
	return e.erc20.Name(e.contract)
}

func (e *Erc20Contract) Symbol() (string, error) {
	return e.erc20.Symbol(e.contract)
}

func (e *Erc20Contract) Decimals() (uint8, error) {
	return e.erc20.Decimals(e.contract)
}

func (e *Erc20Contract) TotalSupply() (*big.Int, error) {
	return e.erc20.TotalSupply(e.contract)
}

func (e *Erc20Contract) BalanceOf(address common.Address) (*big.Int, error) {
	return e.erc20.BalanceOf(e.contract, address)
}

func (e *Erc20Contract) Approve(owner, spender common.Address, amount *big.Int, privateKey string) (string, *types.Transaction, error) {
	return e.erc20.Approve(e.contract, owner, spender, amount, privateKey)
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
