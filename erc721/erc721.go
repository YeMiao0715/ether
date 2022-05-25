package erc721

import (
	"github.com/YeMiao0715/ether"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type Erc721 struct {
	engine    *ether.Engine
	erc721Abi *Erc721Abi
}

func NewErc721(engine *ether.Engine) *Erc721 {
	return &Erc721{
		engine:    engine,
		erc721Abi: &Erc721Abi{},
	}
}

func (e *Erc721) Name(contract common.Address) (string, error) {
	b, err := e.erc721Abi.Name()
	if err != nil {
		return "", err
	}
	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return "", err
	}

	return e.erc721Abi.UnpackName(resb)
}
func (e *Erc721) Symbol(contract common.Address) (string, error) {
	b, err := e.erc721Abi.Symbol()
	if err != nil {
		return "", err
	}
	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return "", err
	}

	return e.erc721Abi.UnpackSymbol(resb)
}
func (e *Erc721) BalanceOf(contract common.Address, owner common.Address) (*big.Int, error) {
	b, err := e.erc721Abi.BalanceOf(owner)
	if err != nil {
		return nil, err
	}
	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return nil, err
	}
	return e.erc721Abi.UnpackBalanceOf(resb)
}
func (e *Erc721) OwnerOf(contract common.Address, tokenId *big.Int) (*common.Address, error) {
	b, err := e.erc721Abi.OwnerOf(tokenId)
	if err != nil {
		return nil, err
	}
	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return nil, err
	}
	return e.erc721Abi.UnpackOwnerOf(resb)
}
func (e *Erc721) TokenURI(contract common.Address, tokenId *big.Int) (string, error) {
	b, err := e.erc721Abi.TokenURI(tokenId)
	if err != nil {
		return "", err
	}
	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return "", err
	}
	return e.erc721Abi.UnpackTokenURI(resb)
}
func (e *Erc721) Approve(contract, to common.Address, tokenId *big.Int, privateKey string) (string, *types.Transaction, error) {
	data, err := e.erc721Abi.Approve(to, tokenId)
	if err != nil {
		return "", nil, err
	}
	buildTx, err := e.engine.BuildTxByContractWithPrivateKey(contract, data, privateKey)
	if err != nil {
		return "", nil, err
	}
	hash, buildTx, err := e.engine.SendTransactionWithPrivateKey(buildTx, privateKey)
	if err != nil {
		return "", nil, err
	}
	return hash, buildTx, err
}
func (e *Erc721) GetApproved(contract common.Address, tokenId *big.Int) (*common.Address, error) {
	b, err := e.erc721Abi.GetApproved(tokenId)
	if err != nil {
		return nil, err
	}
	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return nil, err
	}
	return e.erc721Abi.UnpackGetApproved(resb)
}
func (e *Erc721) SetApprovalForAll(contract, operator common.Address, approved bool, privateKey string) (string, *types.Transaction, error) {
	data, err := e.erc721Abi.SetApprovalForAll(operator, approved)
	if err != nil {
		return "", nil, err
	}
	buildTx, err := e.engine.BuildTxByContractWithPrivateKey(contract, data, privateKey)
	if err != nil {
		return "", nil, err
	}
	hash, buildTx, err := e.engine.SendTransactionWithPrivateKey(buildTx, privateKey)
	if err != nil {
		return "", nil, err
	}
	return hash, buildTx, err
}
func (e *Erc721) IsApprovedForAll(contract, owner, operator common.Address) (bool, error) {
	b, err := e.erc721Abi.IsApprovedForAll(owner, operator)
	if err != nil {
		return false, err
	}
	resb, err := e.engine.CallContract(contract, b)
	if err != nil {
		return false, err
	}
	return e.erc721Abi.UnpackIsApprovedForAll(resb)
}
func (e *Erc721) TransferFrom(contract, from, to common.Address, tokenId *big.Int, privateKey string) (string, *types.Transaction, error) {
	data, err := e.erc721Abi.TransferFrom(from, to, tokenId)
	if err != nil {
		return "", nil, err
	}
	buildTx, err := e.engine.BuildTxByContractWithPrivateKey(contract, data, privateKey)
	if err != nil {
		return "", nil, err
	}
	hash, buildTx, err := e.engine.SendTransactionWithPrivateKey(buildTx, privateKey)
	if err != nil {
		return "", nil, err
	}
	return hash, buildTx, err
}
func (e *Erc721) SafeTransferFrom(contract, from, to common.Address, tokenId *big.Int, privateKey string) (string, *types.Transaction, error) {
	data, err := e.erc721Abi.SafeTransferFrom(from, to, tokenId)
	if err != nil {
		return "", nil, err
	}
	buildTx, err := e.engine.BuildTxByContractWithPrivateKey(contract, data, privateKey)
	if err != nil {
		return "", nil, err
	}
	hash, buildTx, err := e.engine.SendTransactionWithPrivateKey(buildTx, privateKey)
	if err != nil {
		return "", nil, err
	}
	return hash, buildTx, err
}
