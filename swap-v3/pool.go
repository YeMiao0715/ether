package swap_v3

import (
	"github.com/YeMiao0715/ether"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Pool struct {
	engine *ether.Engine
	IPool  *IPool
}

func NewPool(engine *ether.Engine) *Pool {
	return &Pool{
		engine: engine,
		IPool:  IPoolAbi,
	}
}

func (p *Pool) Slot0(contractAddress common.Address) (*Slot0, error) {
	data, _ := p.IPool.Slot0()
	data, err := p.engine.CallContract(contractAddress, data)
	if err != nil {
		return nil, err
	}

	return p.IPool.UnpackSlot0(data)
}

func (p *Pool) Token0(contract common.Address) (common.Address, error) {
	b, err := p.IPool.Token0()
	if err != nil {
		return common.Address{}, err
	}
	resb, err := p.engine.CallContract(contract, b)
	if err != nil {
		return common.Address{}, err
	}

	return p.IPool.UnpackToken0(resb)
}

func (p *Pool) Token1(contract common.Address) (common.Address, error) {
	b, err := p.IPool.Token1()
	if err != nil {
		return common.Address{}, err
	}
	resb, err := p.engine.CallContract(contract, b)
	if err != nil {
		return common.Address{}, err
	}

	return p.IPool.UnpackToken1(resb)
}

func (p *Pool) Fee(contract common.Address) (*big.Int, error) {
	b, err := p.IPool.Fee()
	if err != nil {
		return nil, err
	}
	resb, err := p.engine.CallContract(contract, b)
	if err != nil {
		return nil, err
	}

	return p.IPool.UnpackFee(resb)
}
