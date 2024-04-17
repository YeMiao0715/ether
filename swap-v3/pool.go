package swap_v3

import (
	"github.com/YeMiao0715/ether"
	"github.com/ethereum/go-ethereum/common"
)

type Pool struct {
	engine   *ether.Engine
	IPoolAbi IPoolAbi
}

func NewPool(engine *ether.Engine) *Pool {
	return &Pool{
		engine:   engine,
		IPoolAbi: IPoolAbi{},
	}
}

func (p *Pool) Slot0(contractAddress common.Address) (*Slot0, error) {
	data, _ := p.IPoolAbi.Slot0()
	data, err := p.engine.CallContract(contractAddress, data)
	if err != nil {
		return nil, err
	}

	return p.IPoolAbi.UnpackSlot0(data)
}

func (p *Pool) Token0(contract common.Address) (common.Address, error) {
	b, err := p.IPoolAbi.Token0()
	if err != nil {
		return common.Address{}, err
	}
	resb, err := p.engine.CallContract(contract, b)
	if err != nil {
		return common.Address{}, err
	}

	return p.IPoolAbi.UnpackToken0(resb)
}

func (p *Pool) Token1(contract common.Address) (common.Address, error) {
	b, err := p.IPoolAbi.Token1()
	if err != nil {
		return common.Address{}, err
	}
	resb, err := p.engine.CallContract(contract, b)
	if err != nil {
		return common.Address{}, err
	}

	return p.IPoolAbi.UnpackToken1(resb)
}
