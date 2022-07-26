package swap_v2

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"math/big"
	"strings"
)

const FactoryAbiJson = `[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"token0","type":"address"},{"indexed":true,"internalType":"address","name":"token1","type":"address"},{"indexed":false,"internalType":"address","name":"pair","type":"address"},{"indexed":false,"internalType":"uint256","name":"","type":"uint256"}],"name":"PairCreated","type":"event"},{"constant":true,"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"allPairs","outputs":[{"internalType":"address","name":"pair","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"allPairsLength","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"tokenA","type":"address"},{"internalType":"address","name":"tokenB","type":"address"}],"name":"createPair","outputs":[{"internalType":"address","name":"pair","type":"address"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"feeTo","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"feeToSetter","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"address","name":"tokenA","type":"address"},{"internalType":"address","name":"tokenB","type":"address"}],"name":"getPair","outputs":[{"internalType":"address","name":"pair","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"","type":"address"}],"name":"setFeeTo","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"","type":"address"}],"name":"setFeeToSetter","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`

type IFactory struct {
	abi *abi.ABI
}

func (i *IFactory) GetAbi() (*abi.ABI, error) {
	if i.abi == nil {
		_erc721Abi, err := abi.JSON(strings.NewReader(FactoryAbiJson))
		if err != nil {
			return nil, err
		}
		i.abi = &_erc721Abi
	}
	return i.abi, nil
}

func (i *IFactory) Method(fn string, param ...interface{}) ([]byte, error) {
	contractAbi, err := i.GetAbi()
	if err != nil {
		return nil, err
	}

	b, err := contractAbi.Pack(fn, param...)
	if err != nil {
		return nil, errors.Wrap(err, "abi pack err")
	}

	return b, nil
}

func (i *IFactory) FeeTo() ([]byte, error) {
	return i.Method("feeTo")
}

func (i *IFactory) UnpackFeeTo(data []byte) (common.Address, error) {
	result, err := i.abi.Unpack("feeTo", data)
	if err != nil {
		return common.Address{}, errors.WithStack(err)
	}
	return result[0].(common.Address), err
}

func (i *IFactory) FeeToSetter() ([]byte, error) {
	return i.Method("feeToSetter")
}

func (i *IFactory) UnpackFeeToSetter(data []byte) (common.Address, error) {
	result, err := i.abi.Unpack("feeToSetter", data)
	if err != nil {
		return common.Address{}, errors.WithStack(err)
	}
	return result[0].(common.Address), err
}

func (i *IFactory) GetPair(tokenA, tokenB common.Address) ([]byte, error) {
	return i.Method("getPair", tokenA, tokenB)
}

func (i *IFactory) UnpackGetPair(data []byte) (common.Address, error) {
	result, err := i.abi.Unpack("getPair", data)
	if err != nil {
		return common.Address{}, errors.WithStack(err)
	}
	return result[0].(common.Address), err
}

func (i *IFactory) AllPairs(index *big.Int) ([]byte, error) {
	return i.Method("allPairs", index)
}

func (i *IFactory) UnpackAllPairs(data []byte) (common.Address, error) {
	result, err := i.abi.Unpack("allPairs", data)
	if err != nil {
		return common.Address{}, errors.WithStack(err)
	}
	return result[0].(common.Address), err
}

func (i *IFactory) AllPairsLength() ([]byte, error) {
	return i.Method("allPairsLength")
}

func (i *IFactory) UnpackAllPairsLength(data []byte) (*big.Int, error) {
	result, err := i.abi.Unpack("allPairsLength", data)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return result[0].(*big.Int), err
}

func (i *IFactory) CreatePair(tokenA, tokenB common.Address) ([]byte, error) {
	return i.Method("createPair", tokenA, tokenB)
}

func (i *IFactory) UnpackCreatePair(data []byte) (common.Address, error) {
	result, err := i.abi.Unpack("createPair", data)
	if err != nil {
		return common.Address{}, errors.WithStack(err)
	}
	return result[0].(common.Address), err
}

func (i *IFactory) SetFeeTo(addr common.Address) ([]byte, error) {
	return i.Method("setFeeTo", addr)
}

func (i *IFactory) SetFeeToSetter(addr common.Address) ([]byte, error) {
	return i.Method("setFeeToSetter", addr)
}
