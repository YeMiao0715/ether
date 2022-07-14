package erc721

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"math/big"
	"strings"
)

const erc721AbiJson = `[{"inputs":[{"internalType":"string","name":"name_","type":"string"},{"internalType":"string","name":"symbol_","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"approved","type":"address"},{"indexed":true,"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"operator","type":"address"},{"indexed":false,"internalType":"bool","name":"approved","type":"bool"}],"name":"ApprovalForAll","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":true,"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"Transfer","type":"event"},{"inputs":[{"internalType":"bytes4","name":"interfaceId","type":"bytes4"}],"name":"supportsInterface","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"owner","type":"address"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"ownerOf","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"name","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"symbol","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"tokenURI","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"approve","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"getApproved","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"operator","type":"address"},{"internalType":"bool","name":"approved","type":"bool"}],"name":"setApprovalForAll","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"address","name":"operator","type":"address"}],"name":"isApprovedForAll","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"transferFrom","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"safeTransferFrom","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"},{"internalType":"bytes","name":"_data","type":"bytes"}],"name":"safeTransferFrom","outputs":[],"stateMutability":"nonpayable","type":"function"}]`

var IErc721Abi = &Erc721Abi{}

type Erc721Abi struct {
	abi *abi.ABI
}

func (e *Erc721Abi) GetAbi() (*abi.ABI, error) {
	if e.abi == nil {
		_erc721Abi, err := abi.JSON(strings.NewReader(erc721AbiJson))
		if err != nil {
			return nil, err
		}
		e.abi = &_erc721Abi
	}
	return e.abi, nil
}

func (e *Erc721Abi) Method(fn string, param ...interface{}) ([]byte, error) {
	erc721Abi, err := e.GetAbi()
	if err != nil {
		return nil, err
	}

	b, err := erc721Abi.Pack(fn, param...)
	if err != nil {
		return nil, errors.Wrap(err, "abi pack err")
	}

	return b, nil
}

func (e *Erc721Abi) Name() ([]byte, error) {
	return e.Method("name")
}

func (e *Erc721Abi) UnpackName(data []byte) (string, error) {
	result, err := e.abi.Unpack("name", data)
	if err != nil {
		return "", errors.Wrap(err, "unpack name error")
	}
	return result[0].(string), err
}

func (e *Erc721Abi) Symbol() ([]byte, error) {
	return e.Method("symbol")
}

func (e *Erc721Abi) UnpackSymbol(data []byte) (string, error) {
	result, err := e.abi.Unpack("symbol", data)
	if err != nil {
		return "", errors.Wrap(err, "unpack symbol error")
	}
	return result[0].(string), err
}

func (e *Erc721Abi) BalanceOf(owner common.Address) ([]byte, error) {
	return e.Method("balanceOf", owner)
}

func (e *Erc721Abi) UnpackBalanceOf(data []byte) (*big.Int, error) {
	result, err := e.abi.Unpack("balanceOf", data)
	if err != nil {
		return nil, errors.Wrap(err, "unpack balanceOf error")
	}
	return result[0].(*big.Int), err
}

func (e *Erc721Abi) TokenURI(tokenId *big.Int) ([]byte, error) {
	return e.Method("tokenURI", tokenId)
}

func (e *Erc721Abi) UnpackTokenURI(data []byte) (string, error) {
	result, err := e.abi.Unpack("tokenURI", data)
	if err != nil {
		return "", errors.Wrap(err, "unpack tokenURI error")
	}
	return result[0].(string), err
}

func (e *Erc721Abi) OwnerOf(tokenId *big.Int) ([]byte, error) {
	return e.Method("ownerOf", tokenId)
}

func (e *Erc721Abi) UnpackOwnerOf(data []byte) (*common.Address, error) {
	result, err := e.abi.Unpack("ownerOf", data)
	if err != nil {
		return nil, errors.Wrap(err, "unpack ownerOf error")
	}
	addr := result[0].(common.Address)
	return &addr, err
}

func (e *Erc721Abi) TransferFrom(from, to common.Address, tokenId *big.Int) ([]byte, error) {
	return e.Method("transferFrom", from, to, tokenId)
}

func (e *Erc721Abi) Approve(to common.Address, tokenId *big.Int) ([]byte, error) {
	return e.Method("approve", to, tokenId)
}

func (e *Erc721Abi) GetApproved(tokenId *big.Int) ([]byte, error) {
	return e.Method("getApproved", tokenId)
}

func (e *Erc721Abi) UnpackGetApproved(data []byte) (*common.Address, error) {
	result, err := e.abi.Unpack("getApproved", data)
	if err != nil {
		return nil, errors.Wrap(err, "unpack getApproved error")
	}
	addr := result[0].(common.Address)
	return &addr, err
}

func (e *Erc721Abi) SetApprovalForAll(operator common.Address, approved bool) ([]byte, error) {
	return e.Method("setApprovalForAll", operator, approved)
}

func (e *Erc721Abi) IsApprovedForAll(owner, operator common.Address) ([]byte, error) {
	return e.Method("isApprovedForAll", owner, operator)
}

func (e *Erc721Abi) UnpackIsApprovedForAll(data []byte) (bool, error) {
	result, err := e.abi.Unpack("isApprovedForAll", data)
	if err != nil {
		return false, errors.Wrap(err, "unpack isApprovedForAll error")
	}
	return result[0].(bool), err
}

func (e *Erc721Abi) SafeTransferFrom(from, to common.Address, tokenId *big.Int) ([]byte, error) {
	return e.Method("safeTransferFrom", from, to, tokenId)
}
