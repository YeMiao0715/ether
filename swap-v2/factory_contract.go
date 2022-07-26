package swap_v2

import (
	"github.com/YeMiao0715/ether"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type FactoryContract struct {
	contract common.Address
	*Factory
}

func NewFactoryContract(engine *ether.Engine, contract common.Address) *FactoryContract {
	return &FactoryContract{
		contract: contract,
		Factory:  NewFactory(engine),
	}
}

func (f *FactoryContract) Contract() common.Address {
	return f.contract
}

func (f *FactoryContract) FeeTo() (common.Address, error) {
	return f.Factory.FeeTo(f.contract)
}

func (f *FactoryContract) FeeToSetter() (common.Address, error) {
	return f.Factory.FeeToSetter(f.contract)
}

func (f *FactoryContract) GetPair(tokenA, tokenB common.Address) (common.Address, error) {
	return f.Factory.GetPair(f.contract, tokenA, tokenB)
}

func (f *FactoryContract) AllPairs(index *big.Int) (common.Address, error) {
	return f.Factory.AllPairs(f.contract, index)
}

func (f *FactoryContract) AllPairsLength() (*big.Int, error) {
	return f.Factory.AllPairsLength(f.contract)
}

func (f *FactoryContract) CreatePair(tokenA, tokenB common.Address, privateKey string) (string, *types.Transaction, error) {
	return f.Factory.CreatePair(f.contract, tokenA, tokenB, privateKey)
}

func (f *FactoryContract) SetFeeTo(fee common.Address, privateKey string) (string, *types.Transaction, error) {
	return f.Factory.SetFeeTo(f.contract, fee, privateKey)
}

func (f *FactoryContract) SetFeeToSetter(feeToSetter common.Address, privateKey string) (string, *types.Transaction, error) {
	return f.Factory.SetFeeToSetter(f.contract, feeToSetter, privateKey)
}
