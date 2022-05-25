package erc20

import (
	"github.com/YeMiao0715/ether"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type Erc20 struct {
	erc20Abi Erc20Abi
	engine   *ether.Engine
}

func NewErc20(engine *ether.Engine) *Erc20 {
	return &Erc20{
		erc20Abi: Erc20Abi{},
		engine:   engine,
	}
}

func (e *Erc20) Name(contract common.Address) (string, error) {
	b, err := e.erc20Abi.Name()
	if err != nil {
		return "", err
	}

	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return "", err
	}

	return e.erc20Abi.UnpackName(resb)
}

func (e *Erc20) Symbol(contract common.Address) (string, error) {
	b, err := e.erc20Abi.Symbol()
	if err != nil {
		return "", err
	}

	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	return e.erc20Abi.UnpackName(resb)
}

func (e *Erc20) Decimals(contract common.Address) (uint8, error) {
	b, err := e.erc20Abi.Decimals()
	if err != nil {
		return 0, err
	}

	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return 0, err
	}

	return e.erc20Abi.UnpackDecimals(resb)
}

func (e *Erc20) TotalSupply(contract common.Address) (*big.Int, error) {
	b, err := e.erc20Abi.TotalSupply()
	if err != nil {
		return nil, err
	}

	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return nil, err
	}

	return e.erc20Abi.UnpackTotalSupply(resb)
}

func (e *Erc20) BalanceOf(contract, address common.Address) (*big.Int, error) {
	b, err := e.erc20Abi.BalanceOf(address)
	if err != nil {
		return nil, err
	}

	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return nil, err
	}

	return e.erc20Abi.UnpackBalanceOf(resb)
}

func (e *Erc20) Approve(contract, owner, spender common.Address, amount *big.Int, privateKey string) (string, *types.Transaction, error) {
	b, err := e.erc20Abi.Approve(spender, amount)
	if err != nil {
		return "", nil, err
	}
	buildTx, err := e.engine.BuildTxByContract(owner, contract, b)
	if err != nil {
		return "", nil, err
	}
	hash, buildTx, err := e.engine.SendTransactionWithPrivateKey(buildTx, privateKey)
	if err != nil {
		return "", nil, err
	}
	return hash, buildTx, nil
}

func (e *Erc20) Allowance(contract, owner, spender common.Address) (*big.Int, error) {
	b, err := e.erc20Abi.Allowance(owner, spender)
	if err != nil {
		return nil, err
	}

	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return nil, err
	}

	return e.erc20Abi.UnpackAllowance(resb)
}

func (e *Erc20) Transfer(contract, recipient common.Address, amount *big.Int, privateKey string) (string, *types.Transaction, error) {
	b, err := e.erc20Abi.Transfer(recipient, amount)
	if err != nil {
		return "", nil, err
	}
	buildTx, err := e.engine.BuildTxByContractWithPrivateKey(contract, b, privateKey)
	if err != nil {
		return "", nil, err
	}
	hash, buildTx, err := e.engine.SendTransactionWithPrivateKey(buildTx, privateKey)
	if err != nil {
		return "", nil, err
	}
	return hash, buildTx, nil
}

func (e *Erc20) TransferFrom(contract, from, recipient common.Address, amount *big.Int, privateKey string) (string, *types.Transaction, error) {
	b, err := e.erc20Abi.TransferFrom(from, recipient, amount)
	if err != nil {
		return "", nil, err
	}
	buildTx, err := e.engine.BuildTxByContractWithPrivateKey(contract, b, privateKey)
	if err != nil {
		return "", nil, err
	}
	hash, buildTx, err := e.engine.SendTransactionWithPrivateKey(buildTx, privateKey)
	if err != nil {
		return "", nil, err
	}
	return hash, buildTx, nil
}
