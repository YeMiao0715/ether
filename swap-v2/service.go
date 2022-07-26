package swap_v2

import (
	"context"
	"fmt"
	"github.com/YeMiao0715/ether"
	"github.com/YeMiao0715/ether/erc20"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"math/big"
	"strings"
	"time"
)

type Service struct {
	engine  *ether.Engine
	pair    *PairContract
	router  *RouterContract
	tokenA  *erc20.Erc20Contract
	tokenB  *erc20.Erc20Contract
	factory *FactoryContract
	symbol  *string
}

func NewServiceWithPairAndRouter(
	engine *ether.Engine,
	pairContractAddress common.Address, // lp合约地址
	routerContractAddress common.Address, // 路由合约地址
) *Service {
	serv := &Service{
		engine:  engine,
		pair:    NewPairContract(engine, pairContractAddress),
		router:  NewRouterContract(engine, routerContractAddress),
		tokenA:  nil,
		tokenB:  nil,
		factory: nil,
	}

	return serv
}

func NewServiceWithFactory(
	engine *ether.Engine,
	factory *FactoryContract, // factory合约地址
	routerContractAddress common.Address, // 路由合约地址
	TokenAAddress common.Address,
	TokenBAddress common.Address,
) (*Service, error) {

	pairAddress, err := factory.GetPair(TokenAAddress, TokenBAddress)
	if err != nil {
		return nil, err
	}

	if pairAddress == common.HexToAddress("") {
		return nil, errors.New("not fount pair address")
	}

	serv := &Service{
		engine:  engine,
		pair:    NewPairContract(engine, pairAddress),
		router:  NewRouterContract(engine, routerContractAddress),
		tokenA:  nil,
		tokenB:  nil,
		factory: factory,
	}

	return serv, nil
}

// Symbol 交易对
func (s *Service) Symbol() (string, error) {
	if s.symbol == nil {
		tokenA, err := s.TokenA()
		if err != nil {
			return "", err
		}
		tokenB, err := s.TokenB()
		if err != nil {
			return "", err
		}
		tokenAName, err := tokenA.Symbol()
		if err != nil {
			return "", err
		}
		tokenBName, err := tokenB.Symbol()
		if err != nil {
			return "", err
		}
		str := fmt.Sprintf("%s/%s", tokenAName, tokenBName)
		s.symbol = &str
	}

	return *s.symbol, nil
}

func (s *Service) TokenA() (*erc20.Erc20Contract, error) {
	if s.tokenA == nil {
		_pair, err := s.Pair()
		if err != nil {
			return nil, err
		}
		token0, err := _pair.Token0()
		if err != nil {
			return nil, err
		}

		s.tokenA = erc20.NewErc20WithContract(s.engine, token0)
	}

	return s.tokenA, nil
}

func (s *Service) TokenB() (*erc20.Erc20Contract, error) {
	if s.tokenB == nil {
		_pair, err := s.Pair()
		if err != nil {
			return nil, err
		}
		token1, err := _pair.Token1()
		if err != nil {
			return nil, err
		}

		s.tokenB = erc20.NewErc20WithContract(s.engine, token1)
	}

	return s.tokenB, nil
}

func (s *Service) Factory() (*FactoryContract, error) {
	if s.factory == nil {
		_pair, err := s.Pair()
		factoryAddress, err := _pair.Factory()
		if err != nil {
			return nil, err
		}

		s.factory = NewFactoryContract(s.engine, factoryAddress)
	}

	return s.factory, nil
}

func (s *Service) Pair() (*PairContract, error) {
	if s.pair == nil {
		return nil, errors.New("pair is nil")
	}

	return s.pair, nil
}

type ServicePriceCoin struct {
	Symbol   string
	Decimals uint8
	Amount   *big.Int
}

type ServicePrice struct {
	Coin   ServicePriceCoin
	ToCoin ServicePriceCoin
	Price  *big.Int
}

func (s ServicePrice) ToDecimal() decimal.Decimal {
	return decimal.NewFromBigInt(s.Price, 0).Div(decimal.New(1, int32(s.ToCoin.Decimals)))
}

func (s ServicePrice) ToString() string {
	return fmt.Sprintf("1%s-%s%s", s.Coin.Symbol, s.ToDecimal().String(), s.ToCoin.Symbol)
}

func (s *Service) AmountByTokenAFromFloat(amount float64) (*big.Int, error) {
	_tokenA, err := s.TokenA()
	if err != nil {
		return nil, err
	}

	decimals, err := _tokenA.Decimals()
	if err != nil {
		return nil, err
	}

	return decimal.NewFromFloat(amount).Mul(decimal.New(1, int32(decimals))).BigInt(), nil
}

func (s *Service) AmountByTokenBFromFloat(amount float64) (*big.Int, error) {
	_tokenB, err := s.TokenB()
	if err != nil {
		return nil, err
	}

	decimals, err := _tokenB.Decimals()
	if err != nil {
		return nil, err
	}

	return decimal.NewFromFloat(amount).Mul(decimal.New(1, int32(decimals))).BigInt(), nil
}

func (s *Service) AmountByLpFromFloat(amount float64) (*big.Int, error) {
	_pair, err := s.Pair()
	if err != nil {
		return nil, err
	}

	decimals, err := _pair.Decimals()
	if err != nil {
		return nil, err
	}

	return decimal.NewFromFloat(amount).Mul(decimal.New(1, int32(decimals))).BigInt(), nil
}

// Price 价格 1 tokenA: %d tokenB
func (s *Service) Price() (*ServicePrice, error) {
	_tokenA, err := s.TokenA()
	if err != nil {
		return nil, err
	}
	decimalsByA, err := _tokenA.Decimals()
	if err != nil {
		return nil, err
	}
	symbolByA, err := _tokenA.Symbol()
	if err != nil {
		return nil, err
	}

	inAmount := decimal.New(1, int32(decimalsByA)).BigInt()
	_, outAmount, err := s.GetAmountsOut(inAmount)

	_tokenB, _ := s.TokenB()
	decimalsByB, err := _tokenB.Decimals()
	if err != nil {
		return nil, err
	}
	symbolByB, err := _tokenB.Symbol()
	if err != nil {
		return nil, err
	}

	return &ServicePrice{
		Coin: ServicePriceCoin{
			Symbol:   symbolByA,
			Decimals: decimalsByA,
			Amount:   inAmount,
		},
		ToCoin: ServicePriceCoin{
			Symbol:   symbolByB,
			Decimals: decimalsByB,
			Amount:   outAmount,
		},
		Price: outAmount,
	}, nil
}

// GetAmountsOut 获取对应金额
func (s *Service) GetAmountsOut(inAmount *big.Int) (*big.Int, *big.Int, error) {
	_tokenA, err := s.TokenA()
	if err != nil {
		return nil, nil, err
	}
	_tokenB, err := s.TokenB()
	if err != nil {
		return nil, nil, err
	}
	res, err := s.router.GetAmountsOut(inAmount, []common.Address{
		_tokenA.Contract(),
		_tokenB.Contract(),
	})
	if err != nil {
		return nil, nil, err
	}

	return res[0], res[1], nil
}

// AddLiquidity 添加流动性
func (s *Service) AddLiquidity(amountA, amountB *big.Int, autoApprove bool, privateKey string) (addLiquidityTx, tokenATx, tokenBTx *types.Transaction, err error) {
	_tokenA, err := s.TokenA()
	if err != nil {
		return
	}
	owner, err := s.engine.PrivateKeyToAddress(privateKey)
	if err != nil {
		return
	}

	balance, err := _tokenA.BalanceOf(*owner)
	if decimal.NewFromBigInt(balance, 0).LessThan(decimal.NewFromBigInt(amountA, 0)) {
		err = errors.New("tokenA balance low")
		return
	}

	allowanceAmount, err := _tokenA.Allowance(*owner, s.router.contract)
	if err != nil {
		return
	}
	// 判断路由地址授权金额是否一致
	if decimal.NewFromBigInt(allowanceAmount, 0).LessThan(decimal.NewFromBigInt(amountA, 0)) {
		if autoApprove {
			// 进行授权
			hash, tx, _err := _tokenA.Approve(s.router.contract, amountA, privateKey)
			if _err != nil {
				err = _err
				return
			}
			tokenATx = tx
			s.engine.Logger().Info("授权TokenA",
				zap.String("amount", amountA.String()),
				zap.String("hash", hash),
			)
			_, err = s.WaitTx(tx)
			if err != nil {
				return
			}
		}
	}

	_tokenB, err := s.TokenB()
	if err != nil {
		return
	}

	balance, err = _tokenB.BalanceOf(*owner)
	if decimal.NewFromBigInt(balance, 0).LessThan(decimal.NewFromBigInt(amountB, 0)) {
		err = errors.New("tokenA balance low")
		return
	}

	allowanceAmount, err = _tokenB.Allowance(*owner, s.router.contract)
	if err != nil {
		return
	}
	// 判断路由地址授权金额是否一致
	if decimal.NewFromBigInt(allowanceAmount, 0).LessThan(decimal.NewFromBigInt(amountB, 0)) {
		if autoApprove {
			// 进行授权
			hash, tx, _err := _tokenB.Approve(s.router.contract, amountB, privateKey)
			if _err != nil {
				err = _err
				return
			}
			tokenBTx = tx
			s.engine.Logger().Info("授权TokenB",
				zap.String("amount", amountB.String()),
				zap.String("hash", hash),
			)
			_, err = s.WaitTx(tx)
			if err != nil {
				return
			}
		}
	}

	hash, tx, err := s.router.AddLiquidity(
		_tokenA.Contract(),
		_tokenB.Contract(),
		amountA,
		amountB,
		decimal.Zero.BigInt(),
		decimal.Zero.BigInt(),
		*owner,
		big.NewInt(time.Now().Unix()+600),
		privateKey,
	)
	if err != nil {
		return
	}

	s.engine.Logger().Info("添加流动性",
		zap.String("hash", hash),
		zap.String("amountA", amountA.String()),
		zap.String("amountB", amountB.String()),
	)

	addLiquidityTx = tx
	return
}

// AddLiquidityWithTokenA 添加对应A数量的流动性
func (s *Service) AddLiquidityWithTokenA(tokenA *big.Int, autoApprove bool, privateKey string) (addLiquidityTx, tokenATx, tokenBTx *types.Transaction, err error) {
	inAmount, outAmount, err := s.GetAmountsOut(tokenA)
	if err != nil {
		return
	}
	return s.AddLiquidity(inAmount, outAmount, autoApprove, privateKey)
}

// RemoveLiquidityWithPermit 移除流动性
func (s *Service) RemoveLiquidityWithPermit(lpAmount *big.Int, privateKey string) (hash string, tx *types.Transaction, err error) {
	owner, err := s.engine.PrivateKeyToAddress(privateKey)
	if err != nil {
		return "", nil, err
	}

	_pair, err := s.Pair()
	if err != nil {
		return "", nil, err
	}

	totalSupply, err := _pair.TotalSupply()
	if err != nil {
		return "", nil, err
	}
	totalSupplyDec := decimal.NewFromBigInt(totalSupply, 0)
	balance, err := _pair.BalanceOf(*owner)
	if err != nil {
		return "", nil, err
	}
	balanceDec := decimal.NewFromBigInt(balance, 0)
	lpAmountDec := decimal.NewFromBigInt(lpAmount, 0)
	if balanceDec.LessThan(lpAmountDec) {
		return "", nil, errors.New("pair balance low")
	}
	deadline := big.NewInt(time.Now().Unix() + 600)
	signV, signR, signS, err := _pair.PermitSign(s.router.contract, lpAmount, deadline, privateKey)
	if err != nil {
		return "", nil, err
	}

	_tokenA, err := s.TokenA()
	if err != nil {
		return "", nil, err
	}
	_tokenB, err := s.TokenB()
	if err != nil {
		return "", nil, err
	}

	reserve0, reserve1, _, err := _pair.GetReserves()
	if err != nil {
		return "", nil, err
	}

	amountMinA := decimal.NewFromBigInt(reserve0, 0).Mul(lpAmountDec.Div(totalSupplyDec)).BigInt()
	amountMinB := decimal.NewFromBigInt(reserve1, 0).Mul(lpAmountDec.Div(totalSupplyDec)).BigInt()

	to32Array := func(bytes []byte) [32]byte {
		res := [32]byte{}
		for i, _ := range res {
			res[i] = bytes[i]
		}
		return res
	}

	return s.router.RemoveLiquidityWithPermit(
		_tokenA.Contract(),
		_tokenB.Contract(),
		lpAmount,
		amountMinA,
		amountMinB,
		*owner,
		deadline,
		false,
		uint8(signV.Uint64()),
		to32Array(signR.Bytes()),
		to32Array(signS.Bytes()),
		privateKey,
	)
}

// WaitTx 等待交易成功
func (s *Service) WaitTx(tx *types.Transaction) (bool, error) {
	ethClient, _, err := s.engine.GetEthClient()
	if err != nil {
		return false, err
	}

	goto Loop
SleepLoop:
	time.Sleep(time.Second * 5)
Loop:
	receipt, err := ethClient.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		if strings.Index(err.Error(), "not found") != -1 {
			s.engine.Logger().Info("查询交易",
				zap.String("hash", tx.Hash().String()),
				zap.String("status", "not found"),
			)
			goto SleepLoop
		}

		return false, err
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		s.engine.Logger().Info("查询交易",
			zap.String("hash", tx.Hash().String()),
			zap.String("status", "fail"),
		)
		return false, errors.New(fmt.Sprintf("hash: %s is fail", tx.Hash()))
	}

	s.engine.Logger().Info("查询交易",
		zap.String("hash", tx.Hash().String()),
		zap.String("status", "success"),
	)
	return true, nil
}
