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

func NewServiceWithRouter(engine *ether.Engine, swapRouterV2 common.Address) (*Service, error) {
	if swapRouterV2 == common.HexToAddress("") {
		return nil, errors.New("router contract address is zero")
	}

	serv := &Service{
		engine:  engine,
		pair:    nil,
		router:  NewRouterContract(engine, swapRouterV2),
		tokenA:  nil,
		tokenB:  nil,
		factory: nil,
		symbol:  nil,
	}

	factory, err := serv.Router().Factory()
	if err != nil {
		return nil, err
	}

	serv.factory = NewFactoryContract(engine, factory)

	return serv, nil
}

func NewServiceWithPairAndRouter(
	engine *ether.Engine,
	pairContractAddress common.Address, // lp合约地址
	routerContractAddress common.Address, // 路由合约地址
) (*Service, error) {

	if pairContractAddress == common.HexToAddress("") {
		return nil, errors.New("pair contract address is zero")
	}

	if routerContractAddress == common.HexToAddress("") {
		return nil, errors.New("router contract address is zero")
	}
	serv := &Service{
		engine:  engine,
		pair:    NewPairContract(engine, pairContractAddress),
		router:  NewRouterContract(engine, routerContractAddress),
		tokenA:  nil,
		tokenB:  nil,
		factory: nil,
	}

	return serv, nil
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
		return nil, errors.New("not fount pair contract address, check tokenA and tokenB contract address")
	}
	if routerContractAddress == common.HexToAddress("") {
		return nil, errors.New("router contract address is zero")
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

func (s *Service) MustFactory() *FactoryContract {
	factory, _ := s.Factory()
	return factory
}

func (s *Service) Pair() (*PairContract, error) {
	if s.pair == nil {
		return nil, errors.New("pair is nil")
	}

	return s.pair, nil
}

func (s *Service) Router() *RouterContract {
	return s.router
}

func (s *Service) NewServiceForTokenAndWETH(token common.Address) (*Service, error) {
	WETHContractAddress, err := s.Router().WETH()
	if err != nil {
		return nil, err
	}

	return s.NewServiceForTokenAAndTokenB(WETHContractAddress, token)
}

func (s *Service) NewServiceForTokenAAndTokenB(tokenA, tokenB common.Address) (*Service, error) {
	service, err := NewServiceWithRouter(s.engine, s.router.contract)
	if err != nil {
		return nil, err
	}

	factoryContractAddress, err := service.Router().Factory()
	if err != nil {
		return nil, err
	}

	service.tokenA = erc20.NewErc20WithContract(s.engine, tokenA)
	service.tokenB = erc20.NewErc20WithContract(s.engine, tokenB)
	service.factory = NewFactoryContract(s.engine, factoryContractAddress)
	pair, err := service.MustFactory().GetPair(service.tokenA.Contract(), service.tokenB.Contract())
	if err != nil {
		return nil, err
	}
	service.pair = NewPairContract(s.engine, pair)

	return service, nil
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
	inAmount := decimal.New(1, int32(decimalsByA)).BigInt()
	return s.PriceOfTokenA(inAmount)
}

// PriceOfTokenA 价格 {amount} tokenA: %d tokenB
func (s *Service) PriceOfTokenA(inAmount *big.Int) (*ServicePrice, error) {
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
	outAmount := big.NewInt(0)
	if inAmount.Int64() != 0 {
		_, outAmount, err = s.GetAmountsOut(inAmount)
		if err != nil {
			return nil, err
		}
	}

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

// GetAmountsOut 获取对应金额 A to B
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

// GetAmountsIn 获取对应金额 B to A
func (s *Service) GetAmountsIn(outAmount *big.Int) (*big.Int, *big.Int, error) {
	_tokenA, err := s.TokenA()
	if err != nil {
		return nil, nil, err
	}
	_tokenB, err := s.TokenB()
	if err != nil {
		return nil, nil, err
	}
	res, err := s.router.GetAmountsIn(outAmount, []common.Address{
		_tokenA.Contract(),
		_tokenB.Contract(),
	})
	if err != nil {
		return nil, nil, err
	}

	return res[0], res[1], nil
}

// AddLiquidity 添加流动性
func (s *Service) AddLiquidity(amountA, amountB *big.Int, privateKey string) (addLiquidityTx, tokenATx, tokenBTx *types.Transaction, err error) {
	_tokenA, err := s.TokenA()
	if err != nil {
		return
	}
	owner, err := s.engine.PrivateKeyToAddress(privateKey)
	if err != nil {
		return
	}
	_tokenB, err := s.TokenB()
	if err != nil {
		return
	}

	_, tokenATx, err = s.WaitApprove(_tokenA, s.router.contract, amountA, privateKey)
	if err != nil {
		return
	}
	_, tokenBTx, err = s.WaitApprove(_tokenB, s.router.contract, amountB, privateKey)
	if err != nil {
		return
	}

	hash, addLiquidityTx, err := s.router.AddLiquidity(
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

	symbolA, _ := _tokenA.Symbol()
	symbolB, _ := _tokenB.Symbol()
	s.engine.Logger().Info("添加流动性",
		zap.String("hash", hash),
		zap.String(symbolA, amountA.String()),
		zap.String(symbolB, amountB.String()),
	)

	return
}

// AddLiquidityWithTokenA 添加对应A数量的流动性
func (s *Service) AddLiquidityWithTokenA(tokenA *big.Int, privateKey string) (addLiquidityTx, tokenATx, tokenBTx *types.Transaction, err error) {
	inAmount, outAmount, err := s.GetAmountsOut(tokenA)
	if err != nil {
		return
	}
	return s.AddLiquidity(inAmount, outAmount, privateKey)
}

// AddLiquidityWithTokenB 添加对应B数量的流动性
func (s *Service) AddLiquidityWithTokenB(tokenB *big.Int, privateKey string) (addLiquidityTx, tokenATx, tokenBTx *types.Transaction, err error) {
	inAmount, outAmount, err := s.GetAmountsIn(tokenB)
	if err != nil {
		return
	}
	return s.AddLiquidity(inAmount, outAmount, privateKey)
}

// RemoveLiquidity 移除流动性
func (s *Service) RemoveLiquidity(lpAmount *big.Int, privateKey string) (hash string, tx *types.Transaction, err error) {
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

	amountMinADec := decimal.NewFromBigInt(reserve0, 0).Mul(lpAmountDec.Div(totalSupplyDec)).Mul(decimal.NewFromFloat(0.999))
	amountMinBDec := decimal.NewFromBigInt(reserve1, 0).Mul(lpAmountDec.Div(totalSupplyDec)).Mul(decimal.NewFromFloat(0.999))

	s.engine.Logger().Info("移除数量",
		zap.String("loAmount", lpAmount.String()),
		zap.String("amountA", amountMinADec.BigInt().String()),
		zap.String("amountB", amountMinBDec.BigInt().String()),
	)

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
		amountMinADec.BigInt(),
		amountMinBDec.BigInt(),
		*owner,
		deadline,
		false,
		uint8(signV.Uint64()),
		to32Array(signR.Bytes()),
		to32Array(signS.Bytes()),
		privateKey,
	)
}

// RemoveLiquidityWithTokenA 移除流动性 - tokenA
func (s *Service) RemoveLiquidityWithTokenA(tokenA *big.Int, privateKey string) (hash string, tx *types.Transaction, err error) {
	_pair, err := s.Pair()
	if err != nil {
		return "", nil, err
	}

	totalSupply, err := _pair.TotalSupply()
	if err != nil {
		return "", nil, err
	}
	totalSupplyDec := decimal.NewFromBigInt(totalSupply, 0)

	reserve0, _, _, err := _pair.GetReserves()
	tokenADec := decimal.NewFromBigInt(tokenA, 0)
	lpAmountDec := totalSupplyDec.Mul(tokenADec.Div(decimal.NewFromBigInt(reserve0, 0)))

	return s.RemoveLiquidity(lpAmountDec.BigInt(), privateKey)
}

// RemoveLiquidityWithTokenB 移除流动性 - tokenB
func (s *Service) RemoveLiquidityWithTokenB(tokenB *big.Int, privateKey string) (hash string, tx *types.Transaction, err error) {
	_pair, err := s.Pair()
	if err != nil {
		return "", nil, err
	}

	totalSupply, err := _pair.TotalSupply()
	if err != nil {
		return "", nil, err
	}
	totalSupplyDec := decimal.NewFromBigInt(totalSupply, 0)

	_, reserve1, _, err := _pair.GetReserves()
	tokenBDec := decimal.NewFromBigInt(tokenB, 0)
	lpAmount := totalSupplyDec.Mul(tokenBDec.Div(decimal.NewFromBigInt(reserve1, 0))).BigInt()
	return s.RemoveLiquidity(lpAmount, privateKey)
}

// SwapExactTokensForTokens 交换代币
func (s *Service) SwapExactTokensForTokens(amountIn, amountOutMax *big.Int, tokenA, tokenB *erc20.Erc20Contract, privateKey string) (hash string, tx *types.Transaction, err error) {

	owner, err := s.engine.PrivateKeyToAddress(privateKey)
	if err != nil {
		return
	}

	//amountOutMin = decimal.NewFromBigInt(amountOutMin, 0).Mul(decimal.NewFromFloat(0.99999)).BigInt()

	symbolA, _ := tokenA.Symbol()
	symbolB, _ := tokenB.Symbol()
	s.engine.Logger().Info("交换代币数量",
		zap.String(symbolA, amountIn.String()),
		zap.String(symbolB, amountOutMax.String()),
	)

	return s.router.SwapExactTokensForTokens(
		amountIn,
		amountOutMax,
		[]common.Address{
			tokenA.Contract(),
			tokenB.Contract(),
		},
		*owner,
		big.NewInt(time.Now().Unix()+600),
		privateKey,
	)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens 交换代币
func (s *Service) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn, amountOutMin *big.Int, tokenA, tokenB *erc20.Erc20Contract, privateKey string) (hash string, tx *types.Transaction, err error) {

	owner, err := s.engine.PrivateKeyToAddress(privateKey)
	if err != nil {
		return
	}

	amountOutMin = decimal.NewFromBigInt(amountOutMin, 0).Mul(decimal.NewFromFloat(0.999)).BigInt()

	symbolA, _ := tokenA.Symbol()
	symbolB, _ := tokenB.Symbol()
	s.engine.Logger().Info("交换代币数量",
		zap.String(symbolA, amountIn.String()),
		zap.String(symbolB, amountOutMin.String()),
	)

	return s.router.SwapExactTokensForTokensSupportingFeeOnTransferTokens(
		amountIn,
		amountOutMin,
		[]common.Address{
			tokenA.Contract(),
			tokenB.Contract(),
		},
		*owner,
		big.NewInt(time.Now().Unix()+600),
		privateKey,
	)
}

func (s *Service) WaitApprove(token *erc20.Erc20Contract, sender common.Address, amount *big.Int, privateKey string) (string, *types.Transaction, error) {
	amountDec := decimal.NewFromBigInt(amount, 0)
	owner, err := s.engine.PrivateKeyToAddress(privateKey)
	balance, err := token.BalanceOf(*owner)
	if err != nil {
		return "", nil, err
	}
	balanceDec := decimal.NewFromBigInt(balance, 0)

	if balanceDec.LessThan(amountDec) {
		err = errors.New("balance is low")
		return "", nil, err
	}

	allowance, err := token.Allowance(*owner, sender)
	if err != nil {
		return "", nil, err
	}
	allowanceDec := decimal.NewFromBigInt(allowance, 0)

	if allowanceDec.LessThan(amountDec) {

		hash, tx, _err := token.Approve(sender, amount, privateKey)
		if _err != nil {
			return "", nil, err
		}
		symbol, err := token.Symbol()

		s.engine.Logger().Info(fmt.Sprintf("授权%s", symbol),
			zap.String("hash", hash),
			zap.String("amount", amount.String()),
		)

		_, err = s.WaitTx(tx)
		if err != nil {
			return hash, tx, err
		}

		return hash, tx, err
	}

	return "", nil, nil
}

func (s *Service) SwapTokenA2TokenB(amountIn *big.Int, fee float64, privateKey string) (swapTx, tokenAtx *types.Transaction, err error) {
	amountA, amountB, err := s.GetAmountsOut(amountIn)
	if err != nil {
		return
	}

	_tokenA, err := s.TokenA()
	if err != nil {
		return
	}

	_tokenB, err := s.TokenB()
	if err != nil {
		return
	}

	amountBDec := decimal.NewFromBigInt(amountB, 0)
	amountBMax := amountBDec.Mul(decimal.NewFromInt(1).Sub(decimal.NewFromFloat(fee)))

	_, tokenAtx, err = s.WaitApprove(_tokenA, s.router.contract, amountIn, privateKey)
	if err != nil {
		return
	}

	_, swapTx, err = s.SwapExactTokensForTokens(amountA, amountBMax.BigInt(), _tokenA, _tokenB, privateKey)
	return
}

func (s *Service) SwapTokenB2TokenA(amountIn *big.Int, fee float64, privateKey string) (swapTx, tokenBtx *types.Transaction, err error) {
	amountA, amountB, err := s.GetAmountsIn(amountIn)
	if err != nil {
		return
	}

	_tokenA, err := s.TokenA()
	if err != nil {
		return
	}

	_tokenB, err := s.TokenB()
	if err != nil {
		return
	}

	amountADec := decimal.NewFromBigInt(amountA, 0)
	amountAMax := amountADec.Mul(decimal.NewFromInt(1).Sub(decimal.NewFromFloat(fee)))

	_, tokenBtx, err = s.WaitApprove(_tokenB, s.router.contract, amountIn, privateKey)
	if err != nil {
		return
	}

	_, swapTx, err = s.SwapExactTokensForTokens(amountB, amountAMax.BigInt(), _tokenB, _tokenA, privateKey)
	return
}

func (s *Service) SwapTokenA2TokenBWithSupportingFee(amountIn *big.Int, fee float64, privateKey string) (swapTx, tokenAtx *types.Transaction, err error) {
	amountA, amountB, err := s.GetAmountsOut(amountIn)
	if err != nil {
		return
	}

	_tokenA, err := s.TokenA()
	if err != nil {
		return
	}

	_tokenB, err := s.TokenB()
	if err != nil {
		return
	}

	amountBDec := decimal.NewFromBigInt(amountB, 0)
	amountBMin := amountBDec.Mul(decimal.NewFromInt(1).Sub(decimal.NewFromFloat(fee)))

	_, tokenAtx, err = s.WaitApprove(_tokenA, s.router.contract, amountIn, privateKey)
	if err != nil {
		return
	}

	_, swapTx, err = s.SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountA, amountBMin.BigInt(), _tokenA, _tokenB, privateKey)
	return
}

func (s *Service) SwapTokenB2TokenAWithSupportingFee(amountIn *big.Int, fee float64, privateKey string) (swapTx, tokenBtx *types.Transaction, err error) {
	amountA, amountB, err := s.GetAmountsIn(amountIn)
	if err != nil {
		return
	}

	_tokenA, err := s.TokenA()
	if err != nil {
		return
	}

	_tokenB, err := s.TokenB()
	if err != nil {
		return
	}

	amountADec := decimal.NewFromBigInt(amountA, 0)
	amountBMin := amountADec.Mul(decimal.NewFromInt(1).Sub(decimal.NewFromFloat(fee)))

	_, tokenBtx, err = s.WaitApprove(_tokenB, s.router.contract, amountIn, privateKey)
	if err != nil {
		return
	}

	_, swapTx, err = s.SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountB, amountBMin.BigInt(), _tokenB, _tokenA, privateKey)
	return
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
