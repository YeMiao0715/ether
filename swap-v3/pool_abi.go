package swap_v3

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"math/big"
	"strings"
)

const poolAbiJsonStr = `[{"type":"function","name":"burn","inputs":[{"name":"tickLower","type":"int24","internalType":"int24"},{"name":"tickUpper","type":"int24","internalType":"int24"},{"name":"amount","type":"uint128","internalType":"uint128"}],"outputs":[{"name":"amount0","type":"uint256","internalType":"uint256"},{"name":"amount1","type":"uint256","internalType":"uint256"}],"stateMutability":"nonpayable"},{"type":"function","name":"collect","inputs":[{"name":"recipient","type":"address","internalType":"address"},{"name":"tickLower","type":"int24","internalType":"int24"},{"name":"tickUpper","type":"int24","internalType":"int24"},{"name":"amount0Requested","type":"uint128","internalType":"uint128"},{"name":"amount1Requested","type":"uint128","internalType":"uint128"}],"outputs":[{"name":"amount0","type":"uint128","internalType":"uint128"},{"name":"amount1","type":"uint128","internalType":"uint128"}],"stateMutability":"nonpayable"},{"type":"function","name":"collectProtocol","inputs":[{"name":"recipient","type":"address","internalType":"address"},{"name":"amount0Requested","type":"uint128","internalType":"uint128"},{"name":"amount1Requested","type":"uint128","internalType":"uint128"}],"outputs":[{"name":"amount0","type":"uint128","internalType":"uint128"},{"name":"amount1","type":"uint128","internalType":"uint128"}],"stateMutability":"nonpayable"},{"type":"function","name":"factory","inputs":[],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"fee","inputs":[],"outputs":[{"name":"","type":"uint24","internalType":"uint24"}],"stateMutability":"view"},{"type":"function","name":"feeGrowthGlobal0X128","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"feeGrowthGlobal1X128","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"flash","inputs":[{"name":"recipient","type":"address","internalType":"address"},{"name":"amount0","type":"uint256","internalType":"uint256"},{"name":"amount1","type":"uint256","internalType":"uint256"},{"name":"data","type":"bytes","internalType":"bytes"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"increaseObservationCardinalityNext","inputs":[{"name":"observationCardinalityNext","type":"uint16","internalType":"uint16"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"initialize","inputs":[{"name":"sqrtPriceX96","type":"uint160","internalType":"uint160"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"liquidity","inputs":[],"outputs":[{"name":"","type":"uint128","internalType":"uint128"}],"stateMutability":"view"},{"type":"function","name":"maxLiquidityPerTick","inputs":[],"outputs":[{"name":"","type":"uint128","internalType":"uint128"}],"stateMutability":"view"},{"type":"function","name":"mint","inputs":[{"name":"recipient","type":"address","internalType":"address"},{"name":"tickLower","type":"int24","internalType":"int24"},{"name":"tickUpper","type":"int24","internalType":"int24"},{"name":"amount","type":"uint128","internalType":"uint128"},{"name":"data","type":"bytes","internalType":"bytes"}],"outputs":[{"name":"amount0","type":"uint256","internalType":"uint256"},{"name":"amount1","type":"uint256","internalType":"uint256"}],"stateMutability":"nonpayable"},{"type":"function","name":"observations","inputs":[{"name":"index","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"blockTimestamp","type":"uint32","internalType":"uint32"},{"name":"tickCumulative","type":"int56","internalType":"int56"},{"name":"secondsPerLiquidityCumulativeX128","type":"uint160","internalType":"uint160"},{"name":"initialized","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"observe","inputs":[{"name":"secondsAgos","type":"uint32[]","internalType":"uint32[]"}],"outputs":[{"name":"tickCumulatives","type":"int56[]","internalType":"int56[]"},{"name":"secondsPerLiquidityCumulativeX128s","type":"uint160[]","internalType":"uint160[]"}],"stateMutability":"view"},{"type":"function","name":"positions","inputs":[{"name":"key","type":"bytes32","internalType":"bytes32"}],"outputs":[{"name":"_liquidity","type":"uint128","internalType":"uint128"},{"name":"feeGrowthInside0LastX128","type":"uint256","internalType":"uint256"},{"name":"feeGrowthInside1LastX128","type":"uint256","internalType":"uint256"},{"name":"tokensOwed0","type":"uint128","internalType":"uint128"},{"name":"tokensOwed1","type":"uint128","internalType":"uint128"}],"stateMutability":"view"},{"type":"function","name":"protocolFees","inputs":[],"outputs":[{"name":"token0","type":"uint128","internalType":"uint128"},{"name":"token1","type":"uint128","internalType":"uint128"}],"stateMutability":"view"},{"type":"function","name":"setFeeProtocol","inputs":[{"name":"feeProtocol0","type":"uint8","internalType":"uint8"},{"name":"feeProtocol1","type":"uint8","internalType":"uint8"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"slot0","inputs":[],"outputs":[{"name":"sqrtPriceX96","type":"uint160","internalType":"uint160"},{"name":"tick","type":"int24","internalType":"int24"},{"name":"observationIndex","type":"uint16","internalType":"uint16"},{"name":"observationCardinality","type":"uint16","internalType":"uint16"},{"name":"observationCardinalityNext","type":"uint16","internalType":"uint16"},{"name":"feeProtocol","type":"uint8","internalType":"uint8"},{"name":"unlocked","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"snapshotCumulativesInside","inputs":[{"name":"tickLower","type":"int24","internalType":"int24"},{"name":"tickUpper","type":"int24","internalType":"int24"}],"outputs":[{"name":"tickCumulativeInside","type":"int56","internalType":"int56"},{"name":"secondsPerLiquidityInsideX128","type":"uint160","internalType":"uint160"},{"name":"secondsInside","type":"uint32","internalType":"uint32"}],"stateMutability":"view"},{"type":"function","name":"swap","inputs":[{"name":"recipient","type":"address","internalType":"address"},{"name":"zeroForOne","type":"bool","internalType":"bool"},{"name":"amountSpecified","type":"int256","internalType":"int256"},{"name":"sqrtPriceLimitX96","type":"uint160","internalType":"uint160"},{"name":"data","type":"bytes","internalType":"bytes"}],"outputs":[{"name":"amount0","type":"int256","internalType":"int256"},{"name":"amount1","type":"int256","internalType":"int256"}],"stateMutability":"nonpayable"},{"type":"function","name":"tickBitmap","inputs":[{"name":"wordPosition","type":"int16","internalType":"int16"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"tickSpacing","inputs":[],"outputs":[{"name":"","type":"int24","internalType":"int24"}],"stateMutability":"view"},{"type":"function","name":"ticks","inputs":[{"name":"tick","type":"int24","internalType":"int24"}],"outputs":[{"name":"liquidityGross","type":"uint128","internalType":"uint128"},{"name":"liquidityNet","type":"int128","internalType":"int128"},{"name":"feeGrowthOutside0X128","type":"uint256","internalType":"uint256"},{"name":"feeGrowthOutside1X128","type":"uint256","internalType":"uint256"},{"name":"tickCumulativeOutside","type":"int56","internalType":"int56"},{"name":"secondsPerLiquidityOutsideX128","type":"uint160","internalType":"uint160"},{"name":"secondsOutside","type":"uint32","internalType":"uint32"},{"name":"initialized","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"token0","inputs":[],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"token1","inputs":[],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"event","name":"Burn","inputs":[{"name":"owner","type":"address","indexed":true,"internalType":"address"},{"name":"tickLower","type":"int24","indexed":true,"internalType":"int24"},{"name":"tickUpper","type":"int24","indexed":true,"internalType":"int24"},{"name":"amount","type":"uint128","indexed":false,"internalType":"uint128"},{"name":"amount0","type":"uint256","indexed":false,"internalType":"uint256"},{"name":"amount1","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"Collect","inputs":[{"name":"owner","type":"address","indexed":true,"internalType":"address"},{"name":"recipient","type":"address","indexed":false,"internalType":"address"},{"name":"tickLower","type":"int24","indexed":true,"internalType":"int24"},{"name":"tickUpper","type":"int24","indexed":true,"internalType":"int24"},{"name":"amount0","type":"uint128","indexed":false,"internalType":"uint128"},{"name":"amount1","type":"uint128","indexed":false,"internalType":"uint128"}],"anonymous":false},{"type":"event","name":"CollectProtocol","inputs":[{"name":"sender","type":"address","indexed":true,"internalType":"address"},{"name":"recipient","type":"address","indexed":true,"internalType":"address"},{"name":"amount0","type":"uint128","indexed":false,"internalType":"uint128"},{"name":"amount1","type":"uint128","indexed":false,"internalType":"uint128"}],"anonymous":false},{"type":"event","name":"Flash","inputs":[{"name":"sender","type":"address","indexed":true,"internalType":"address"},{"name":"recipient","type":"address","indexed":true,"internalType":"address"},{"name":"amount0","type":"uint256","indexed":false,"internalType":"uint256"},{"name":"amount1","type":"uint256","indexed":false,"internalType":"uint256"},{"name":"paid0","type":"uint256","indexed":false,"internalType":"uint256"},{"name":"paid1","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"IncreaseObservationCardinalityNext","inputs":[{"name":"observationCardinalityNextOld","type":"uint16","indexed":false,"internalType":"uint16"},{"name":"observationCardinalityNextNew","type":"uint16","indexed":false,"internalType":"uint16"}],"anonymous":false},{"type":"event","name":"Initialize","inputs":[{"name":"sqrtPriceX96","type":"uint160","indexed":false,"internalType":"uint160"},{"name":"tick","type":"int24","indexed":false,"internalType":"int24"}],"anonymous":false},{"type":"event","name":"Mint","inputs":[{"name":"sender","type":"address","indexed":false,"internalType":"address"},{"name":"owner","type":"address","indexed":true,"internalType":"address"},{"name":"tickLower","type":"int24","indexed":true,"internalType":"int24"},{"name":"tickUpper","type":"int24","indexed":true,"internalType":"int24"},{"name":"amount","type":"uint128","indexed":false,"internalType":"uint128"},{"name":"amount0","type":"uint256","indexed":false,"internalType":"uint256"},{"name":"amount1","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"SetFeeProtocol","inputs":[{"name":"feeProtocol0Old","type":"uint8","indexed":false,"internalType":"uint8"},{"name":"feeProtocol1Old","type":"uint8","indexed":false,"internalType":"uint8"},{"name":"feeProtocol0New","type":"uint8","indexed":false,"internalType":"uint8"},{"name":"feeProtocol1New","type":"uint8","indexed":false,"internalType":"uint8"}],"anonymous":false},{"type":"event","name":"Swap","inputs":[{"name":"sender","type":"address","indexed":true,"internalType":"address"},{"name":"recipient","type":"address","indexed":true,"internalType":"address"},{"name":"amount0","type":"int256","indexed":false,"internalType":"int256"},{"name":"amount1","type":"int256","indexed":false,"internalType":"int256"},{"name":"sqrtPriceX96","type":"uint160","indexed":false,"internalType":"uint160"},{"name":"liquidity","type":"uint128","indexed":false,"internalType":"uint128"},{"name":"tick","type":"int24","indexed":false,"internalType":"int24"}],"anonymous":false}]`

var IPoolAbi = &IPool{}

type IPool struct {
	abi *abi.ABI
}

func (i *IPool) GetAbi() (*abi.ABI, error) {
	if i.abi == nil {
		_erc721Abi, err := abi.JSON(strings.NewReader(poolAbiJsonStr))
		if err != nil {
			return nil, err
		}
		i.abi = &_erc721Abi
	}
	return i.abi, nil
}

func (i *IPool) Method(fn string, param ...interface{}) ([]byte, error) {
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

func (i *IPool) MustAbi() *abi.ABI {
	abi, _ := i.GetAbi()
	return abi
}

func (i *IPool) Slot0() ([]byte, error) {
	return i.Method("slot0")
}

func (i *IPool) UnpackSlot0(data []byte) (*Slot0, error) {
	result, err := i.abi.Unpack("slot0", data)
	if err != nil {
		return nil, errors.Wrap(err, "unpack slot0 err")
	}

	return &Slot0{
		SqrtPriceX96:               result[0].(*big.Int),
		Tick:                       result[1].(*big.Int),
		ObservationIndex:           result[2].(uint16),
		ObservationCardinality:     result[3].(uint16),
		ObservationCardinalityNext: result[4].(uint16),
		FeeProtocol:                result[5].(uint8),
		Unlocked:                   result[6].(bool),
	}, nil
}

func (i *IPool) Token0() ([]byte, error) {
	return i.Method("token0")
}

func (i *IPool) UnpackToken0(data []byte) (common.Address, error) {
	result, err := i.abi.Unpack("token0", data)
	if err != nil {
		return common.Address{}, errors.WithStack(err)
	}
	return result[0].(common.Address), err
}

func (i *IPool) Token1() ([]byte, error) {
	return i.Method("token1")
}

func (i *IPool) UnpackToken1(data []byte) (common.Address, error) {
	result, err := i.abi.Unpack("token1", data)
	if err != nil {
		return common.Address{}, errors.WithStack(err)
	}
	return result[0].(common.Address), err
}

func (i *IPool) Mint(recipient common.Address, tickLower, tickUpper, amount *big.Int) ([]byte, error) {
	return i.Method("mint", recipient, tickLower, tickUpper, amount)
}

func (i *IPool) UnpackMint(data []byte) (*big.Int, *big.Int, error) {
	result, err := i.abi.Unpack("mint", data)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	return result[0].(*big.Int), result[1].(*big.Int), nil
}
