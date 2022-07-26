package swap_v2

import (
	"github.com/YeMiao0715/ether"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type Factory struct {
	engine   *ether.Engine
	iFactory *IFactory
}

func NewFactory(engine *ether.Engine) *Factory {
	return &Factory{
		engine:   engine,
		iFactory: &IFactory{},
	}
}

func (f *Factory) FeeTo(contract common.Address) (common.Address, error) {
	b, err := f.iFactory.FeeTo()
	if err != nil {
		return common.Address{}, err
	}
	resb, err := f.engine.CallContract(contract, b)
	if err != nil {
		return common.Address{}, err
	}

	return f.iFactory.UnpackFeeTo(resb)
}

func (f *Factory) FeeToSetter(contract common.Address) (common.Address, error) {
	b, err := f.iFactory.FeeToSetter()
	if err != nil {
		return common.Address{}, err
	}
	resb, err := f.engine.CallContract(contract, b)
	if err != nil {
		return common.Address{}, err
	}

	return f.iFactory.UnpackFeeToSetter(resb)
}

func (f *Factory) GetPair(contract, tokenA, tokenB common.Address) (common.Address, error) {
	b, err := f.iFactory.GetPair(tokenA, tokenB)
	if err != nil {
		return common.Address{}, err
	}
	resb, err := f.engine.CallContract(contract, b)
	if err != nil {
		return common.Address{}, err
	}

	return f.iFactory.UnpackGetPair(resb)
}

func (f *Factory) AllPairs(contract common.Address, index *big.Int) (common.Address, error) {
	b, err := f.iFactory.AllPairs(index)
	if err != nil {
		return common.Address{}, err
	}
	resb, err := f.engine.CallContract(contract, b)
	if err != nil {
		return common.Address{}, err
	}

	return f.iFactory.UnpackAllPairs(resb)
}

func (f *Factory) AllPairsLength(contract common.Address) (*big.Int, error) {
	b, err := f.iFactory.AllPairsLength()
	if err != nil {
		return nil, err
	}
	resb, err := f.engine.CallContract(contract, b)
	if err != nil {
		return nil, err
	}

	return f.iFactory.UnpackAllPairsLength(resb)
}

func (f *Factory) CreatePair(contract, tokenA, tokenB common.Address, privateKey string) (string, *types.Transaction, error) {
	abiData, err := f.iFactory.CreatePair(tokenA, tokenB)
	if err != nil {
		return "", nil, err
	}

	tx, err := f.engine.BuildTxByContractWithPrivateKey(contract, abiData, privateKey)
	if err != nil {
		return "", nil, err
	}

	return f.engine.SendTransactionWithPrivateKey(tx, privateKey)
}

func (f *Factory) SetFeeTo(contract, addr common.Address, privateKey string) (string, *types.Transaction, error) {
	abiData, err := f.iFactory.SetFeeTo(addr)
	if err != nil {
		return "", nil, err
	}

	tx, err := f.engine.BuildTxByContractWithPrivateKey(contract, abiData, privateKey)
	if err != nil {
		return "", nil, err
	}

	return f.engine.SendTransactionWithPrivateKey(tx, privateKey)
}

func (f *Factory) SetFeeToSetter(contract, addr common.Address, privateKey string) (string, *types.Transaction, error) {
	abiData, err := f.iFactory.SetFeeToSetter(addr)
	if err != nil {
		return "", nil, err
	}

	tx, err := f.engine.BuildTxByContractWithPrivateKey(contract, abiData, privateKey)
	if err != nil {
		return "", nil, err
	}

	return f.engine.SendTransactionWithPrivateKey(tx, privateKey)
}
