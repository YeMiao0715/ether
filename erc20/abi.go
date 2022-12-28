package erc20

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"math/big"
	"strings"
)

const erc20ApiJson = `[{"inputs":[{"internalType":"string","name":"name_","type":"string"},{"internalType":"string","name":"symbol_","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"spender","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"previousOwner","type":"address"},{"indexed":true,"internalType":"address","name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"Transfer","type":"event"},{"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"address","name":"spender","type":"address"}],"name":"allowance","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"approve","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"account","type":"address"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"decimals","outputs":[{"internalType":"uint8","name":"","type":"uint8"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"subtractedValue","type":"uint256"}],"name":"decreaseAllowance","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"getOwner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"addedValue","type":"uint256"}],"name":"increaseAllowance","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"msgSender","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"name","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"renounceOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"symbol","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"totalSupply","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"recipient","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"transfer","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"sender","type":"address"},{"internalType":"address","name":"recipient","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"transferFrom","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"}]`

var IErc20Abi = Erc20Abi{}

type Erc20Abi struct {
	abi *abi.ABI
}

func (e *Erc20Abi) GetAbi() (*abi.ABI, error) {
	if e.abi == nil {
		_erc20Abi, err := abi.JSON(strings.NewReader(erc20ApiJson))
		if err != nil {
			return nil, err
		}
		e.abi = &_erc20Abi
	}
	return e.abi, nil
}

func (e *Erc20Abi) MustAbi() *abi.ABI {
	abi, _ := e.GetAbi()
	return abi
}

func (e *Erc20Abi) Method(fn string, param ...interface{}) ([]byte, error) {
	erc20Abi, err := e.GetAbi()
	if err != nil {
		return nil, err
	}

	b, err := erc20Abi.Pack(fn, param...)
	if err != nil {
		return nil, errors.Wrap(err, "erc20Abi pack err")
	}

	return b, nil
}

func (e *Erc20Abi) Name() ([]byte, error) {
	return e.Method("name")
}

func (e *Erc20Abi) UnpackName(data []byte) (string, error) {
	result, err := e.abi.Unpack("name", data)
	if err != nil {
		return "", errors.Wrap(err, "unpack name error")
	}
	return result[0].(string), err
}

func (e *Erc20Abi) Symbol() ([]byte, error) {
	return e.Method("symbol")
}

func (e *Erc20Abi) UnpackSymbol(data []byte) (string, error) {
	result, err := e.abi.Unpack("symbol", data)
	if err != nil {
		return "", errors.Wrap(err, "unpack symbol error")
	}
	return result[0].(string), err
}

func (e *Erc20Abi) Decimals() ([]byte, error) {
	return e.Method("decimals")
}

func (e *Erc20Abi) UnpackDecimals(data []byte) (uint8, error) {
	result, err := e.abi.Unpack("decimals", data)
	if err != nil {
		return 0, errors.Wrap(err, "unpack decimals error")
	}
	return result[0].(uint8), err
}

func (e *Erc20Abi) TotalSupply() ([]byte, error) {
	return e.Method("totalSupply")
}

func (e *Erc20Abi) UnpackTotalSupply(data []byte) (*big.Int, error) {
	result, err := e.abi.Unpack("totalSupply", data)
	if err != nil {
		return nil, errors.Wrap(err, "unpack totalSupply error")
	}
	return result[0].(*big.Int), err
}

func (e *Erc20Abi) BalanceOf(address common.Address) ([]byte, error) {
	return e.Method("balanceOf", address)
}

func (e *Erc20Abi) UnpackBalanceOf(data []byte) (*big.Int, error) {
	result, err := e.abi.Unpack("balanceOf", data)
	if err != nil {
		return nil, errors.Wrap(err, "unpack balanceOf error")
	}
	return result[0].(*big.Int), err
}

func (e *Erc20Abi) Approve(spender common.Address, amount *big.Int) ([]byte, error) {
	return e.Method("approve", spender, amount)
}

func (e *Erc20Abi) Transfer(recipient common.Address, amount *big.Int) ([]byte, error) {
	return e.Method("transfer", recipient, amount)
}

func (e *Erc20Abi) Allowance(owner, spender common.Address) ([]byte, error) {
	return e.Method("allowance", owner, spender)
}

func (e *Erc20Abi) UnpackAllowance(data []byte) (*big.Int, error) {
	result, err := e.abi.Unpack("allowance", data)
	if err != nil {
		return nil, errors.Wrap(err, "unpack allowance error")
	}
	return result[0].(*big.Int), err
}

func (e *Erc20Abi) TransferFrom(form, to common.Address, amount *big.Int) ([]byte, error) {
	return e.Method("transferFrom", form, to, amount)
}
