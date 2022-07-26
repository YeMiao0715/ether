package swap_v2

import (
	"github.com/YeMiao0715/ether"
	"github.com/YeMiao0715/ether/erc20"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Pair struct {
	engine *ether.Engine
	iPair  *IPair
	*erc20.Erc20
}

func NewPair(engine *ether.Engine) *Pair {
	return &Pair{
		engine: engine,
		iPair:  &IPair{},
		Erc20:  erc20.NewErc20(engine),
	}
}

func (p *Pair) Factory(contract common.Address) (common.Address, error) {
	b, err := p.iPair.Factory()
	if err != nil {
		return common.Address{}, err
	}
	resb, err := p.engine.CallContract(contract, b)
	if err != nil {
		return common.Address{}, err
	}

	return p.iPair.UnpackFactory(resb)
}

func (p *Pair) Token0(contract common.Address) (common.Address, error) {
	b, err := p.iPair.Token0()
	if err != nil {
		return common.Address{}, err
	}
	resb, err := p.engine.CallContract(contract, b)
	if err != nil {
		return common.Address{}, err
	}

	return p.iPair.UnpackToken0(resb)
}

func (p *Pair) Token1(contract common.Address) (common.Address, error) {
	b, err := p.iPair.Token1()
	if err != nil {
		return common.Address{}, err
	}
	resb, err := p.engine.CallContract(contract, b)
	if err != nil {
		return common.Address{}, err
	}

	return p.iPair.UnpackToken1(resb)
}

func (p *Pair) Price0CumulativeLast(contract common.Address) (*big.Int, error) {
	b, err := p.iPair.Price0CumulativeLast()
	if err != nil {
		return nil, err
	}
	resb, err := p.engine.CallContract(contract, b)
	if err != nil {
		return nil, err
	}

	return p.iPair.UnpackPrice0CumulativeLast(resb)
}

func (p *Pair) Price1CumulativeLast(contract common.Address) (*big.Int, error) {
	b, err := p.iPair.Price1CumulativeLast()
	if err != nil {
		return nil, err
	}
	resb, err := p.engine.CallContract(contract, b)
	if err != nil {
		return nil, err
	}

	return p.iPair.UnpackPrice1CumulativeLast(resb)
}

func (p *Pair) KLast(contract common.Address) (*big.Int, error) {
	b, err := p.iPair.KLast()
	if err != nil {
		return nil, err
	}
	resb, err := p.engine.CallContract(contract, b)
	if err != nil {
		return nil, err
	}

	return p.iPair.UnpackKLast(resb)
}

func (p *Pair) GetReserves(contract common.Address) (reserve0, reserve1 *big.Int, blockTimestampLast uint32, err error) {
	b, err := p.iPair.GetReserves()
	if err != nil {
		return
	}
	resb, err := p.engine.CallContract(contract, b)
	if err != nil {
		return
	}

	return p.iPair.UnpackGetReserves(resb)
}

func (p *Pair) DOMAIN_SEPARATOR(contract common.Address) (common.Hash, error) {
	b, err := p.iPair.DOMAIN_SEPARATOR()
	if err != nil {
		return common.Hash{}, err
	}
	resb, err := p.engine.CallContract(contract, b)
	if err != nil {
		return common.Hash{}, err
	}
	return p.iPair.UnpackDOMAIN_SEPARATOR(resb)
}

func (p *Pair) PERMIT_TYPEHASH(contract common.Address) (common.Hash, error) {
	b, err := p.iPair.PERMIT_TYPEHASH()
	if err != nil {
		return common.Hash{}, err
	}
	resb, err := p.engine.CallContract(contract, b)
	if err != nil {
		return common.Hash{}, err
	}
	return p.iPair.UnpackPERMIT_TYPEHASH(resb)
}

func (p *Pair) Nonces(contract, sender common.Address) (*big.Int, error) {
	b, err := p.iPair.Nonces(sender)
	if err != nil {
		return nil, err
	}
	resb, err := p.engine.CallContract(contract, b)
	if err != nil {
		return nil, err
	}
	return p.iPair.UnpackNonces(resb)
}
