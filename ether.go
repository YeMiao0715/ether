package ether

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"math/big"
	"strings"
)

type EngineType int

const (
	EIP155Signer  EngineType = 1
	EIP2930Signer EngineType = 2
)

type Engine struct {
	logger         *zap.Logger
	rpcClient      *rpc.Client
	ethClient      *ethclient.Client
	chainId        *big.Int
	gasPrice       *big.Int
	gasLimitOffset uint64

	isWs   bool   // 当前是否为ws链接，当为ws链接时可以使用一些订阅事件
	rpcUrl string // 设置rcp链接
	wsUrl  string // 设置ws链接

	txType EngineType
}

// NewEngine 新建一个链接引擎，全局唯一
func NewEngine(logger *zap.Logger, rpcUrl string, wsUrl string) *Engine {
	return &Engine{logger: logger, rpcUrl: rpcUrl, wsUrl: wsUrl, gasPrice: decimal.New(2, 9).BigInt(), txType: 0, gasLimitOffset: 0}
}

func NewEngineWithType(logger *zap.Logger, rpcUrl string, wsUrl string, txType EngineType) *Engine {
	return &Engine{logger: logger, rpcUrl: rpcUrl, wsUrl: wsUrl, gasPrice: decimal.New(2, 9).BigInt(), txType: txType, gasLimitOffset: 0}
}

func (c *Engine) GetRpcClient() (*rpc.Client, bool, error) {

	if c.rpcClient == nil {
		if c.wsUrl != "" {
			_rpcClient, err := rpc.DialWebsocket(context.Background(), c.wsUrl, "")
			if err != nil {
				c.logger.Error("node ws Engine error", zap.Error(err))
				goto rpcClient
			}
			c.rpcClient = _rpcClient
			c.isWs = true
			return c.rpcClient, c.isWs, nil
		}
	}

rpcClient:
	_rpcClient, err := rpc.Dial(c.rpcUrl)
	if err != nil {
		return nil, false, errors.Wrap(err, "node rpc Engine error")
	}
	c.rpcClient = _rpcClient
	c.isWs = false
	goto end

end:
	return c.rpcClient, c.isWs, nil
}

func (c *Engine) MustRpcClient() *rpc.Client {
	rpcClient, _, _ := c.GetRpcClient()
	return rpcClient
}

func (c *Engine) GetEthClient() (*ethclient.Client, bool, error) {
	if c.ethClient == nil {
		_rpcClient, isWs, err := c.GetRpcClient()
		if err != nil {
			return nil, isWs, errors.Wrap(err, "rcp node Engine error")
		}
		c.ethClient = ethclient.NewClient(_rpcClient)
	}

	return c.ethClient, c.isWs, nil
}

func (c *Engine) MustEthClient() *ethclient.Client {
	if c.ethClient == nil {
		_rpcClient, _, _ := c.GetRpcClient()
		c.ethClient = ethclient.NewClient(_rpcClient)
	}
	return c.ethClient
}

func (c *Engine) SetUrl(rpc string, ws string) {
	c.rpcUrl = rpc
	c.wsUrl = ws
}

func (c *Engine) SetGasPrice(gasPrice *big.Int) {
	c.gasPrice = gasPrice
}

func (c *Engine) SetGasLimitOffset(offset uint64) {
	c.gasLimitOffset = offset
}

func (c *Engine) SetLogger(logger *zap.Logger) {
	c.logger = logger
}

func (c *Engine) Logger() *zap.Logger {
	return c.logger
}

func (c *Engine) GetPendingNonce(address common.Address) (uint64, error) {
	client, _, err := c.GetEthClient()
	if err != nil {
		return 0, errors.Wrap(err, "get nonce err")
	}

	return client.PendingNonceAt(context.Background(), address)
}

func (c *Engine) GetNonce(address common.Address) (uint64, error) {
	client, _, err := c.GetEthClient()
	if err != nil {
		return 0, errors.Wrap(err, "get nonce err")
	}
	lastBlockNumber, err := c.GetBlockNumber()
	if err != nil {
		return 0, err
	}
	return client.NonceAt(context.Background(), address, decimal.NewFromInt(int64(lastBlockNumber)).BigInt())
}

func (c *Engine) GetChainId() (*big.Int, error) {
	if c.chainId == nil {
		client, _, err := c.GetEthClient()
		if err != nil {
			return nil, err
		}
		c.chainId, err = client.ChainID(context.Background())
		if err != nil {
			return nil, errors.Wrap(err, "获取chainId失败")
		}
	}

	return c.chainId, nil
}

func (c *Engine) GasPrice() *big.Int {
	return c.gasPrice
}

func (c *Engine) BuildTx(from, to common.Address, gas uint64, gasPrice *big.Int, value *big.Int, data []byte, setNonce *uint64) (*types.Transaction, error) {
	var nonce uint64 = 0
	if setNonce == nil {
		_nonce, err := c.GetNonce(from)
		if err != nil {
			return nil, err
		}
		nonce = _nonce
	} else {
		nonce = *setNonce
	}
	var buildTx *types.Transaction
	switch c.txType {
	case EIP2930Signer:
		chainId, err := c.GetChainId()
		if err != nil {
			return nil, err
		}

		buildTx = types.NewTx(&types.DynamicFeeTx{
			ChainID:    chainId,
			Nonce:      nonce,
			GasTipCap:  gasPrice,
			GasFeeCap:  gasPrice,
			Gas:        gas + c.gasLimitOffset,
			To:         &to,
			Value:      value,
			Data:       data,
			AccessList: nil,
			V:          nil,
			R:          nil,
			S:          nil,
		})

	default:
		buildTx = types.NewTx(&types.LegacyTx{
			Nonce:    nonce,
			GasPrice: gasPrice,
			Gas:      gas,
			To:       &to,
			Value:    value,
			Data:     data,
			V:        nil,
			R:        nil,
			S:        nil,
		})
	}

	return buildTx, nil
}

func (c *Engine) BuildTxByContractWithGas(sender, contract common.Address, gas uint64, data []byte) (*types.Transaction, error) {
	buildTx, err := c.BuildTx(sender, contract, gas, c.GasPrice(), nil, data, nil)
	if err != nil {
		return nil, err
	}
	return buildTx, nil
}

func (c *Engine) BuildTxByContract(sender, contract common.Address, data []byte) (*types.Transaction, error) {
	gas, err := c.EstimateGas(sender, contract, nil, data)
	if err != nil {
		return nil, err
	}
	buildTx, err := c.BuildTx(sender, contract, gas, c.GasPrice(), nil, data, nil)
	if err != nil {
		return nil, err
	}
	return buildTx, nil
}

func (c *Engine) BuildTxByContractWithPrivateKey(contract common.Address, data []byte, privateKey string) (*types.Transaction, error) {
	sender, err := c.PrivateKeyToAddress(privateKey)
	if err != nil {
		return nil, err
	}
	gas, err := c.EstimateGas(*sender, contract, nil, data)
	if err != nil {
		return nil, err
	}
	buildTx, err := c.BuildTx(*sender, contract, gas, c.GasPrice(), nil, data, nil)
	if err != nil {
		return nil, err
	}
	return buildTx, nil
}

func (c *Engine) Singer() (types.Signer, error) {
	chain, err := c.GetChainId()
	if err != nil {
		return nil, err
	}
	var signer types.Signer

	switch c.txType {
	case EIP155Signer:
		signer = types.NewEIP2930Signer(chain)
	case EIP2930Signer:
		signer = types.NewLondonSigner(chain)
	default:
		signer = types.NewEIP155Signer(chain)
	}

	return signer, nil
}

func (c *Engine) EstimateGas(from, to common.Address, value *big.Int, data []byte) (uint64, error) {
	client, _, err := c.GetEthClient()
	if err != nil {
		return 0, err
	}

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:       from,
		To:         &to,
		Gas:        0,
		GasPrice:   c.gasPrice,
		GasFeeCap:  c.gasPrice,
		GasTipCap:  c.gasPrice,
		Value:      value,
		Data:       data,
		AccessList: nil,
	})
	if err != nil {
		return 0, errors.Wrap(err, "估算gas失败")
	}

	return gasLimit, nil
}

func (c *Engine) SendTransaction(transaction *types.Transaction) (string, *types.Transaction, error) {
	client, _, err := c.GetEthClient()
	if err != nil {
		return "", nil, err
	}
	if err := client.SendTransaction(context.Background(), transaction); err != nil {
		return "", nil, errors.Wrap(err, "广播交易失败")
	}

	return transaction.Hash().String(), transaction, err
}

func (c *Engine) SendTransactionWithPrivateKey(transaction *types.Transaction, privateKey string) (string, *types.Transaction, error) {
	singer, err := c.Singer()
	if err != nil {
		return "", nil, err
	}
	privKey, err := c.HexToEcdsaPrivateKey(privateKey)
	if err != nil {
		return "", nil, errors.Wrap(err, "解析私钥失败")
	}
	transaction, err = types.SignTx(transaction, singer, privKey)
	if err != nil {
		return "", nil, errors.Wrap(err, "交易加签失败")
	}

	hash, transaction, err := c.SendTransaction(transaction)
	if err != nil {
		return "", nil, err
	}

	return hash, transaction, err
}

func (c *Engine) HexToEcdsaPrivateKey(privateKey string) (*ecdsa.PrivateKey, error) {
	str := privateKey
	pk := common.HexToHash(privateKey)
	if strings.Index(pk.Hex(), "0x") != -1 {
		str = pk.Hex()[2:]
	}
	privKey, err := crypto.HexToECDSA(str)
	if err != nil {
		return nil, errors.Wrap(err, "解析私钥失败")
	}
	return privKey, nil
}

func (c *Engine) PrivateKeyToAddress(privateKey string) (*common.Address, error) {
	privKey, err := c.HexToEcdsaPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	addr := crypto.PubkeyToAddress(privKey.PublicKey)
	return &addr, nil
}

//func (c *Engine) PrivateSign(hash common.Hash, privateKey string) (common.Hash, error) {
//	privKey, err := c.hexToEcdsaPrivateKey(privateKey)
//	if err != nil {
//		return common.Hash{}, err
//	}
//	privKey.Sign()
//
//	return
//}

func (c *Engine) TransactionByHash(hash common.Hash) (*Transaction, error) {
	client, _, err := c.GetEthClient()
	if err != nil {
		return nil, err
	}

	transaction := &Transaction{}

	tx, isPending, err := client.TransactionByHash(context.Background(), hash)
	if err != nil {
		return nil, errors.Wrap(err, "查询hash失败")
	}

	if isPending {
		transaction.Status = TransactionStatusPending
	}

	transaction.tx = tx

	receipt, err := client.TransactionReceipt(context.Background(), hash)
	if err != nil {
		return nil, errors.Wrap(err, "获取交易回执失败")
	}

	if receipt.Status == 1 {
		transaction.Status = TransactionStatusSuccess
	} else {
		transaction.Status = TransactionStatusFail
	}

	transaction.receipt = receipt

	return transaction, nil
}

func (c *Engine) CallContract(contract common.Address, data []byte) ([]byte, error) {
	client, _, err := c.GetEthClient()
	if err != nil {
		return nil, err
	}

	resb, err := client.PendingCallContract(context.Background(), ethereum.CallMsg{
		To:   &contract,
		Data: data,
	})

	if err != nil {
		return nil, errors.Wrap(err, "访问合约失败")
	}

	return resb, nil
}

func (c *Engine) CallContractWithFrom(from, contract common.Address, data []byte) ([]byte, error) {
	client, _, err := c.GetEthClient()
	if err != nil {
		return nil, err
	}

	resb, err := client.PendingCallContract(context.Background(), ethereum.CallMsg{
		From: from,
		To:   &contract,
		Data: data,
	})

	if err != nil {
		return nil, errors.Wrap(err, "访问合约失败")
	}

	return resb, nil
}

func (c *Engine) GetBalance(account common.Address) (*big.Int, error) {
	client, _, err := c.GetEthClient()
	if err != nil {
		return nil, err
	}
	balance, err := client.PendingBalanceAt(context.Background(), account)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (c *Engine) TransferEth(to common.Address, value *big.Int, privateKey string) (string, *types.Transaction, error) {
	sender, err := c.PrivateKeyToAddress(privateKey)
	if err != nil {
		return "", nil, err
	}

	gas, err := c.EstimateGas(*sender, to, value, nil)
	if err != nil {
		return "", nil, err
	}

	buildTx, err := c.BuildTx(*sender, to, gas, c.gasPrice, value, nil, nil)
	if err != nil {
		return "", nil, err
	}
	return c.SendTransactionWithPrivateKey(buildTx, privateKey)
}

func (c *Engine) TransferEthWithNonce(to common.Address, value *big.Int, privateKey string, setNonce *uint64) (string, *types.Transaction, error) {
	sender, err := c.PrivateKeyToAddress(privateKey)
	if err != nil {
		return "", nil, err
	}

	gas, err := c.EstimateGas(*sender, to, value, nil)
	if err != nil {
		return "", nil, err
	}

	buildTx, err := c.BuildTx(*sender, to, gas, c.gasPrice, value, nil, setNonce)
	if err != nil {
		return "", nil, err
	}
	return c.SendTransactionWithPrivateKey(buildTx, privateKey)
}

func (c *Engine) GetBlockNumber() (uint64, error) {
	client, _, err := c.GetEthClient()
	if err != nil {
		return 0, err
	}
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		return 0, err
	}
	return blockNumber, nil
}

func (c *Engine) IsContract(address common.Address) (bool, error) {
	client, _, err := c.GetEthClient()
	if err != nil {
		return false, err
	}
	b, err := client.PendingCodeAt(context.Background(), address)
	if err != nil {
		return false, err
	}

	if len(b) == 0 {
		return false, nil
	}

	return true, nil
}
