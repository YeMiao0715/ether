package swap_v2

import (
	"github.com/YeMiao0715/ether"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type Factory struct {
	engine   *ether.Engine
	IFactory *IFactory
}

func NewFactory(engine *ether.Engine) *Factory {
	return &Factory{
		engine:   engine,
		IFactory: IFactoryAbi,
	}
}

func (f *Factory) FeeTo(contract common.Address) (common.Address, error) {
	b, err := f.IFactory.FeeTo()
	if err != nil {
		return common.Address{}, err
	}
	resb, err := f.engine.CallContract(contract, b)
	if err != nil {
		return common.Address{}, err
	}

	return f.IFactory.UnpackFeeTo(resb)
}

func (f *Factory) FeeToSetter(contract common.Address) (common.Address, error) {
	b, err := f.IFactory.FeeToSetter()
	if err != nil {
		return common.Address{}, err
	}
	resb, err := f.engine.CallContract(contract, b)
	if err != nil {
		return common.Address{}, err
	}

	return f.IFactory.UnpackFeeToSetter(resb)
}

func (f *Factory) GetPair(contract, tokenA, tokenB common.Address) (common.Address, error) {
	b, err := f.IFactory.GetPair(tokenA, tokenB)
	if err != nil {
		return common.Address{}, err
	}
	resb, err := f.engine.CallContract(contract, b)
	if err != nil {
		return common.Address{}, err
	}

	return f.IFactory.UnpackGetPair(resb)
}

func (f *Factory) AllPairs(contract common.Address, index *big.Int) (common.Address, error) {
	b, err := f.IFactory.AllPairs(index)
	if err != nil {
		return common.Address{}, err
	}
	resb, err := f.engine.CallContract(contract, b)
	if err != nil {
		return common.Address{}, err
	}

	return f.IFactory.UnpackAllPairs(resb)
}

func (f *Factory) AllPairsLength(contract common.Address) (*big.Int, error) {
	b, err := f.IFactory.AllPairsLength()
	if err != nil {
		return nil, err
	}
	resb, err := f.engine.CallContract(contract, b)
	if err != nil {
		return nil, err
	}

	return f.IFactory.UnpackAllPairsLength(resb)
}

func (f *Factory) CreatePair(contract, tokenA, tokenB common.Address, privateKey string) (string, *types.Transaction, error) {
	abiData, err := f.IFactory.CreatePair(tokenA, tokenB)
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
	abiData, err := f.IFactory.SetFeeTo(addr)
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
	abiData, err := f.IFactory.SetFeeToSetter(addr)
	if err != nil {
		return "", nil, err
	}

	tx, err := f.engine.BuildTxByContractWithPrivateKey(contract, abiData, privateKey)
	if err != nil {
		return "", nil, err
	}

	return f.engine.SendTransactionWithPrivateKey(tx, privateKey)
}
