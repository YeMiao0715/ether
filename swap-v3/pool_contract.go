package swap_v3

import (
	"github.com/YeMiao0715/ether"
	"github.com/YeMiao0715/ether/erc20"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"math/big"
)

type PoolContract struct {
	engine *ether.Engine
	*Pool
	contractAddress common.Address

	cacheToken0 *common.Address
	cacheToken1 *common.Address
}

func NewPoolContract(engine *ether.Engine, contractAddress common.Address) *PoolContract {
	return &PoolContract{
		engine:          engine,
		Pool:            NewPool(engine),
		contractAddress: contractAddress,
	}
}

func (p *PoolContract) Slot0() (*Slot0, error) {
	return p.Pool.Slot0(p.contractAddress)
}

func (p *PoolContract) Token0() (*common.Address, error) {
	if p.cacheToken0 == nil {
		token0, err := p.Pool.Token0(p.contractAddress)
		if err != nil {
			return nil, err
		}
		p.cacheToken0 = &token0
	}
	return p.cacheToken0, nil
}

func (p *PoolContract) Token1() (*common.Address, error) {
	if p.cacheToken1 == nil {
		token1, err := p.Pool.Token1(p.contractAddress)
		if err != nil {
			return nil, err
		}
		p.cacheToken1 = &token1
	}
	return p.cacheToken1, nil
}

func (p *PoolContract) Fee() (*big.Int, error) {
	return p.Pool.Fee(p.contractAddress)
}

//func (p *PoolContract) GetPriceInput(amountIn decimal.Decimal, token common.Address) (decimal.Decimal, error) {
//
//}

func (p *PoolContract) Price() (decimal.Decimal, error) {
	slot0, err := p.Slot0()
	if err != nil {
		return decimal.Decimal{}, err
	}
	token0Address, err := p.Token0()
	if err != nil {
		return decimal.Decimal{}, err
	}
	token0 := erc20.NewErc20WithContract(p.engine, *token0Address)
	token1Address, err := p.Token1()
	if err != nil {
		return decimal.Decimal{}, err
	}
	token1 := erc20.NewErc20WithContract(p.engine, *token1Address)
	amount0, err := token0.ToAmount(1)
	if err != nil {
		return decimal.Decimal{}, err
	}
	amount1, err := token1.ToAmount(1)
	if err != nil {
		return decimal.Decimal{}, err
	}

	return slot0.Price(amount0, amount1), nil
}
