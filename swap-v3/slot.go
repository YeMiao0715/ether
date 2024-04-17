package swap_v3

import (
	"github.com/shopspring/decimal"
	"math/big"
)

type Slot0 struct {
	SqrtPriceX96               *big.Int `json:"sqrtPriceX96"`
	Tick                       *big.Int `json:"tick"`
	ObservationIndex           uint16   `json:"observationIndex"`
	ObservationCardinality     uint16   `json:"observationCardinality"`
	ObservationCardinalityNext uint16   `json:"observationCardinalityNext"`
	FeeProtocol                uint8    `json:"feeProtocol"`
	Unlocked                   bool     `json:"unlocked"`
}

func (s Slot0) P() decimal.Decimal {
	sqrtPriceX96 := decimal.NewFromBigInt(s.SqrtPriceX96, 0)
	q := decimal.NewFromInt(2).Pow(decimal.NewFromInt(96))
	p := sqrtPriceX96.Div(q).Pow(decimal.NewFromInt(2))
	return p
}

func (s Slot0) Price(amount0 decimal.Decimal, amount1 decimal.Decimal) decimal.Decimal {
	return s.P().Mul(amount0).Div(amount1)
}
