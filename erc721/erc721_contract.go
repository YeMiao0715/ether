package erc721

import (
	"github.com/YeMiao0715/ether"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type Erc721Contract struct {
	contract common.Address
	erc721   *Erc721
}

func NewErc721Contract(contract common.Address, engine *ether.Engine) *Erc721Contract {
	return &Erc721Contract{
		contract: contract,
		erc721:   NewErc721(engine),
	}
}

func (e *Erc721Contract) Contract() common.Address {
	return e.contract
}

func (e *Erc721Contract) Name() (string, error) {
	return e.erc721.Name(e.contract)
}
func (e *Erc721Contract) Symbol() (string, error) {
	return e.erc721.Symbol(e.contract)
}
func (e *Erc721Contract) BalanceOf(owner common.Address) (*big.Int, error) {
	return e.erc721.BalanceOf(e.contract, owner)
}
func (e *Erc721Contract) OwnerOf(tokenId *big.Int) (*common.Address, error) {
	return e.erc721.OwnerOf(e.contract, tokenId)
}
func (e *Erc721Contract) TokenURI(tokenId *big.Int) (string, error) {
	return e.erc721.TokenURI(e.contract, tokenId)
}
func (e *Erc721Contract) Approve(to common.Address, tokenId *big.Int, privateKey string) (string, *types.Transaction, error) {
	return e.erc721.Approve(e.contract, to, tokenId, privateKey)
}
func (e *Erc721Contract) GetApproved(tokenId *big.Int) (*common.Address, error) {
	return e.erc721.GetApproved(e.contract, tokenId)
}
func (e *Erc721Contract) SetApprovalForAll(operator common.Address, approved bool, privateKey string) (string, *types.Transaction, error) {
	return e.erc721.SetApprovalForAll(e.contract, operator, approved, privateKey)
}
func (e *Erc721Contract) IsApprovedForAll(owner, operator common.Address) (bool, error) {
	return e.erc721.IsApprovedForAll(e.contract, owner, operator)
}
func (e *Erc721Contract) TransferFrom(from, to common.Address, tokenId *big.Int, privateKey string) (string, *types.Transaction, error) {
	return e.erc721.TransferFrom(e.contract, from, to, tokenId, privateKey)
}
func (e *Erc721Contract) SafeTransferFrom(from, to common.Address, tokenId *big.Int, privateKey string) (string, *types.Transaction, error) {
	return e.erc721.SafeTransferFrom(e.contract, from, to, tokenId, privateKey)
}
