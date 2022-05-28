package ether

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type TransactionStatus string

const (
	TransactionStatusPending TransactionStatus = "pending"
	TransactionStatusSuccess TransactionStatus = "success"
	TransactionStatusFail    TransactionStatus = "fail"
)

type Transaction struct {
	Status  TransactionStatus `json:"status"`
	tx      *types.Transaction
	receipt *types.Receipt
}

func (t Transaction) Transaction() (*types.Transaction, bool) {
	return t.tx, t.tx != nil
}

func (t Transaction) Receipt() (*types.Receipt, bool) {
	return t.receipt, t.receipt != nil
}

func (t Transaction) IsContract(engine *Engine) (isContract bool, code []byte, err error) {
	if tx, ok := t.Transaction(); ok {
		if tx.To().Hex() != "0x0000000000000000000000000000000000000000" {
			client, _, _err := engine.GetEthClient()
			if _err != nil {
				err = _err
				return
			}
			blockNumber, _err := client.BlockNumber(context.Background())
			if _err != nil {
				err = _err
				return
			}
			b, _err := client.CodeAt(context.Background(), *tx.To(), big.NewInt(int64(blockNumber)))
			if _err != nil {
				err = _err
				return
			}
			if len(b) > 0 {
				isContract = true
			}
			code = b
		}
	}
	return
}

func (t Transaction) ContractType(engine *Engine) {
	//isContract, code, err := t.IsContract(engine)
	//vm.Contract{}
}

func (t Transaction) AsMessage(engine *Engine) (*types.Message, error) {
	transaction, _ := t.Transaction()

	singer, err := engine.Singer()
	if err != nil {
		return nil, err
	}
	message, err := transaction.AsMessage(singer, transaction.GasPrice())
	if err != nil {
		return nil, err
	}

	return &message, nil
}
