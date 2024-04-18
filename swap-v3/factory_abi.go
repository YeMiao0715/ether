package swap_v3

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/pkg/errors"
	"strings"
)

var factoryAbiJsonStr = `[{"type":"function","name":"createPool","inputs":[{"name":"tokenA","type":"address","internalType":"address"},{"name":"tokenB","type":"address","internalType":"address"},{"name":"fee","type":"uint24","internalType":"uint24"}],"outputs":[{"name":"pool","type":"address","internalType":"address"}],"stateMutability":"nonpayable"},{"type":"function","name":"enableFeeAmount","inputs":[{"name":"fee","type":"uint24","internalType":"uint24"},{"name":"tickSpacing","type":"int24","internalType":"int24"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"feeAmountTickSpacing","inputs":[{"name":"fee","type":"uint24","internalType":"uint24"}],"outputs":[{"name":"","type":"int24","internalType":"int24"}],"stateMutability":"view"},{"type":"function","name":"getPool","inputs":[{"name":"tokenA","type":"address","internalType":"address"},{"name":"tokenB","type":"address","internalType":"address"},{"name":"fee","type":"uint24","internalType":"uint24"}],"outputs":[{"name":"pool","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"owner","inputs":[],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"setOwner","inputs":[{"name":"_owner","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"event","name":"FeeAmountEnabled","inputs":[{"name":"fee","type":"uint24","indexed":true,"internalType":"uint24"},{"name":"tickSpacing","type":"int24","indexed":true,"internalType":"int24"}],"anonymous":false},{"type":"event","name":"OwnerChanged","inputs":[{"name":"oldOwner","type":"address","indexed":true,"internalType":"address"},{"name":"newOwner","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"PoolCreated","inputs":[{"name":"token0","type":"address","indexed":true,"internalType":"address"},{"name":"token1","type":"address","indexed":true,"internalType":"address"},{"name":"fee","type":"uint24","indexed":true,"internalType":"uint24"},{"name":"tickSpacing","type":"int24","indexed":false,"internalType":"int24"},{"name":"pool","type":"address","indexed":false,"internalType":"address"}],"anonymous":false}]`

var IFactoryAbi = &IFactory{}

type IFactory struct {
	abi *abi.ABI
}

func (i *IFactory) GetAbi() (*abi.ABI, error) {
	if i.abi == nil {
		_erc721Abi, err := abi.JSON(strings.NewReader(factoryAbiJsonStr))
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

func (i *IFactory) MustAbi() *abi.ABI {
	abi, _ := i.GetAbi()
	return abi
}
