package swap_v2

import (
	"github.com/YeMiao0715/ether/erc20"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"math/big"
	"strings"
)

const pairAbiJson = `[{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"spender","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"sender","type":"address"},{"indexed":false,"internalType":"uint256","name":"amount0","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"amount1","type":"uint256"},{"indexed":true,"internalType":"address","name":"to","type":"address"}],"name":"Burn","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"sender","type":"address"},{"indexed":false,"internalType":"uint256","name":"amount0","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"amount1","type":"uint256"}],"name":"Mint","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"sender","type":"address"},{"indexed":false,"internalType":"uint256","name":"amount0In","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"amount1In","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"amount0Out","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"amount1Out","type":"uint256"},{"indexed":true,"internalType":"address","name":"to","type":"address"}],"name":"Swap","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"uint112","name":"reserve0","type":"uint112"},{"indexed":false,"internalType":"uint112","name":"reserve1","type":"uint112"}],"name":"Sync","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"Transfer","type":"event"},{"constant":true,"inputs":[],"name":"DOMAIN_SEPARATOR","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"MINIMUM_LIQUIDITY","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"PERMIT_TYPEHASH","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"address","name":"","type":"address"},{"internalType":"address","name":"","type":"address"}],"name":"allowance","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"value","type":"uint256"}],"name":"approve","outputs":[{"internalType":"bool","name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"internalType":"address","name":"","type":"address"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"to","type":"address"}],"name":"burn","outputs":[{"internalType":"uint256","name":"amount0","type":"uint256"},{"internalType":"uint256","name":"amount1","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"internalType":"uint8","name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"factory","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"getReserves","outputs":[{"internalType":"uint112","name":"_reserve0","type":"uint112"},{"internalType":"uint112","name":"_reserve1","type":"uint112"},{"internalType":"uint32","name":"_blockTimestampLast","type":"uint32"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"_token0","type":"address"},{"internalType":"address","name":"_token1","type":"address"}],"name":"initialize","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"kLast","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"to","type":"address"}],"name":"mint","outputs":[{"internalType":"uint256","name":"liquidity","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"name","outputs":[{"internalType":"string","name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"address","name":"","type":"address"}],"name":"nonces","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"value","type":"uint256"},{"internalType":"uint256","name":"deadline","type":"uint256"},{"internalType":"uint8","name":"v","type":"uint8"},{"internalType":"bytes32","name":"r","type":"bytes32"},{"internalType":"bytes32","name":"s","type":"bytes32"}],"name":"permit","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"price0CumulativeLast","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"price1CumulativeLast","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"to","type":"address"}],"name":"skim","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"internalType":"uint256","name":"amount0Out","type":"uint256"},{"internalType":"uint256","name":"amount1Out","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"swap","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"internalType":"string","name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[],"name":"sync","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"token0","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"token1","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"value","type":"uint256"}],"name":"transfer","outputs":[{"internalType":"bool","name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"value","type":"uint256"}],"name":"transferFrom","outputs":[{"internalType":"bool","name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"}]`

type IPair struct {
	abi *abi.ABI
	erc20.Erc20Abi
}

func (i *IPair) GetAbi() (*abi.ABI, error) {
	if i.abi == nil {
		_erc721Abi, err := abi.JSON(strings.NewReader(pairAbiJson))
		if err != nil {
			return nil, err
		}
		i.abi = &_erc721Abi
	}
	return i.abi, nil
}

func (i *IPair) Method(fn string, param ...interface{}) ([]byte, error) {
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

func (e *IPair) MustAbi() *abi.ABI {
	abi, _ := e.GetAbi()
	return abi
}

func (i *IPair) Factory() ([]byte, error) {
	return i.Method("factory")
}

func (i *IPair) UnpackFactory(data []byte) (common.Address, error) {
	result, err := i.abi.Unpack("factory", data)
	if err != nil {
		return common.Address{}, errors.WithStack(err)
	}
	return result[0].(common.Address), err
}

func (i *IPair) Token0() ([]byte, error) {
	return i.Method("token0")
}

func (i *IPair) UnpackToken0(data []byte) (common.Address, error) {
	result, err := i.abi.Unpack("token0", data)
	if err != nil {
		return common.Address{}, errors.WithStack(err)
	}
	return result[0].(common.Address), err
}

func (i *IPair) Token1() ([]byte, error) {
	return i.Method("token1")
}

func (i *IPair) UnpackToken1(data []byte) (common.Address, error) {
	result, err := i.abi.Unpack("token1", data)
	if err != nil {
		return common.Address{}, errors.WithStack(err)
	}
	return result[0].(common.Address), err
}

func (i *IPair) Price0CumulativeLast() ([]byte, error) {
	return i.Method("price0CumulativeLast")
}

func (i *IPair) UnpackPrice0CumulativeLast(data []byte) (*big.Int, error) {
	result, err := i.abi.Unpack("price0CumulativeLast", data)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return result[0].(*big.Int), err
}

func (i *IPair) Price1CumulativeLast() ([]byte, error) {
	return i.Method("price1CumulativeLast")
}

func (i *IPair) UnpackPrice1CumulativeLast(data []byte) (*big.Int, error) {
	result, err := i.abi.Unpack("price1CumulativeLast", data)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return result[0].(*big.Int), err
}

func (i *IPair) KLast() ([]byte, error) {
	return i.Method("kLast")
}

func (i *IPair) UnpackKLast(data []byte) (*big.Int, error) {
	result, err := i.abi.Unpack("kLast", data)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return result[0].(*big.Int), err
}

func (i *IPair) GetReserves() ([]byte, error) {
	return i.Method("getReserves")
}

func (i *IPair) UnpackGetReserves(data []byte) (reserve0, reserve1 *big.Int, blockTimestampLast uint32, err error) {
	result, err := i.abi.Unpack("getReserves", data)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	reserve0 = result[0].(*big.Int)
	reserve1 = result[1].(*big.Int)
	blockTimestampLast = result[2].(uint32)
	return
}

func (i *IPair) DOMAIN_SEPARATOR() ([]byte, error) {
	return i.Method("DOMAIN_SEPARATOR")
}

func (i *IPair) UnpackDOMAIN_SEPARATOR(data []byte) (common.Hash, error) {
	result, err := i.abi.Unpack("DOMAIN_SEPARATOR", data)
	if err != nil {
		err = errors.WithStack(err)
		return common.Hash{}, err
	}
	hash := result[0].([32]byte)
	return hash, nil
}

func (i *IPair) PERMIT_TYPEHASH() ([]byte, error) {
	return i.Method("PERMIT_TYPEHASH")
}

func (i *IPair) UnpackPERMIT_TYPEHASH(data []byte) (common.Hash, error) {
	result, err := i.abi.Unpack("PERMIT_TYPEHASH", data)
	if err != nil {
		err = errors.WithStack(err)
		return common.Hash{}, err
	}
	hash := result[0].([32]byte)
	return hash, nil
}

func (i *IPair) Nonces(addr common.Address) ([]byte, error) {
	return i.Method("nonces", addr)
}

func (i *IPair) UnpackNonces(data []byte) (*big.Int, error) {
	result, err := i.abi.Unpack("nonces", data)
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}
	return result[0].(*big.Int), nil
}
