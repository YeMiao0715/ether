package erc721

import (
	"fmt"
	"github.com/YeMiao0715/ether"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"math/big"
	"testing"
)

var engine *ether.Engine

func init() {
	logger, _ := zap.NewDevelopment()
	engine = ether.NewEngine(logger, "http://120.24.95.205:8545", "")
	engine.SetGasPrice(decimal.New(10, 9).BigInt())
}

func TestErc721_Name(t *testing.T) {
	e := NewErc721(engine)
	t.Log(e.Name(common.HexToAddress("0x10C29E9abcCc3312cAe8563f91ffc24D20D84EDE")))
	t.Log(e.Symbol(common.HexToAddress("0x10C29E9abcCc3312cAe8563f91ffc24D20D84EDE")))
	t.Log(e.BalanceOf(
		common.HexToAddress("0x10C29E9abcCc3312cAe8563f91ffc24D20D84EDE"),
		common.HexToAddress("0x7eBEc3488D6951eDC00FA3fB63657F70e2D20bCf"),
	))
	t.Log(e.OwnerOf(
		common.HexToAddress("0x10C29E9abcCc3312cAe8563f91ffc24D20D84EDE"),
		big.NewInt(51),
	))
	t.Log(e.TokenURI(
		common.HexToAddress("0x10C29E9abcCc3312cAe8563f91ffc24D20D84EDE"),
		big.NewInt(1),
	))
	t.Log(e.GetApproved(
		common.HexToAddress("0x10C29E9abcCc3312cAe8563f91ffc24D20D84EDE"),
		big.NewInt(100),
	))
	t.Log(e.IsApprovedForAll(
		common.HexToAddress("0x10C29E9abcCc3312cAe8563f91ffc24D20D84EDE"),
		common.HexToAddress("0x7eBEc3488D6951eDC00FA3fB63657F70e2D20bCf"),
		common.HexToAddress("0x06eea78c7722d79b5B4B4681cB0E5798146f193d"),
	))
}

func TestErc721_Approve(t *testing.T) {
	e := NewErc721(engine)

	hash, tx, err := e.Approve(
		common.HexToAddress("0x10C29E9abcCc3312cAe8563f91ffc24D20D84EDE"),
		common.HexToAddress("0x06eea78c7722d79b5B4B4681cB0E5798146f193d"),
		big.NewInt(100),
		"",
	)
	if err != nil {
		fmt.Printf("%+v", err)
	} else {
		t.Log(hash)
		t.Log(tx.RawSignatureValues())
	}
}

func TestErc721_SetApprovalForAll(t *testing.T) {
	e := NewErc721(engine)

	hash, tx, err := e.SetApprovalForAll(
		common.HexToAddress("0x10C29E9abcCc3312cAe8563f91ffc24D20D84EDE"),
		common.HexToAddress("0x06eea78c7722d79b5B4B4681cB0E5798146f193d"),
		true,
		"9d3f5dd0dd61eff0c54bea1b32234220d87135f231a5633f6bcb6df0aaa1acf4",
	)
	if err != nil {
		fmt.Printf("%+v", err)
	} else {
		t.Log(hash)
		t.Log(tx.RawSignatureValues())
	}

	t.Log(e.GetApproved(
		common.HexToAddress("0x10C29E9abcCc3312cAe8563f91ffc24D20D84EDE"),
		big.NewInt(102),
	))
}

func TestErc721_TransferFrom(t *testing.T) {
	e := NewErc721(engine)

	hash, tx, err := e.TransferFrom(
		common.HexToAddress("0x10C29E9abcCc3312cAe8563f91ffc24D20D84EDE"),
		common.HexToAddress("0x7eBEc3488D6951eDC00FA3fB63657F70e2D20bCf"),
		common.HexToAddress("0x06eea78c7722d79b5B4B4681cB0E5798146f193d"),
		big.NewInt(50),
		"",
	)
	if err != nil {
		fmt.Printf("%+v", err)
	} else {
		t.Log(hash)
		t.Log(tx.RawSignatureValues())
	}

	t.Log(e.OwnerOf(
		common.HexToAddress("0x10C29E9abcCc3312cAe8563f91ffc24D20D84EDE"),
		big.NewInt(50),
	))
}

func TestErc721_SafeTransferFrom(t *testing.T) {
	e := NewErc721(engine)

	hash, tx, err := e.SafeTransferFrom(
		common.HexToAddress("0x10C29E9abcCc3312cAe8563f91ffc24D20D84EDE"),
		common.HexToAddress("0x7eBEc3488D6951eDC00FA3fB63657F70e2D20bCf"),
		common.HexToAddress("0x06eea78c7722d79b5B4B4681cB0E5798146f193d"),
		big.NewInt(51),
		"",
	)
	if err != nil {
		fmt.Printf("%+v", err)
	} else {
		t.Log(hash)
		t.Log(tx.RawSignatureValues())
	}

	t.Log(e.OwnerOf(
		common.HexToAddress("0x10C29E9abcCc3312cAe8563f91ffc24D20D84EDE"),
		big.NewInt(51),
	))
}

func TestErc721Contract_Approve(t *testing.T) {
	e := NewErc721WithContract(engine, common.HexToAddress("0x10C29E9abcCc3312cAe8563f91ffc24D20D84EDE"))
	t.Log(e.OwnerOf(big.NewInt(51)))
}
