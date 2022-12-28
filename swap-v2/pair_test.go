package swap_v2

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"math/big"
	"testing"
)

var pair *Pair

var pairAddress = common.HexToAddress("0xdf5B1f505837D763B6AA98C003Df3C53625239AB")

func init() {
	pair = NewPair(engine)
}

func TestPair_Token0(t *testing.T) {
	t.Log(pair.Factory(pairAddress))
	t.Log(pair.Token0(pairAddress))
	t.Log(pair.Token1(pairAddress))
	t.Log(pair.Name(pairAddress))
	t.Log(pair.Price0CumulativeLast(pairAddress))
	t.Log(pair.Price1CumulativeLast(pairAddress))
	t.Log(pair.KLast(pairAddress))
	t.Log(pair.GetReserves(pairAddress))
	t.Log(pair.DOMAIN_SEPARATOR(pairAddress))
	t.Log(pair.PERMIT_TYPEHASH(pairAddress))
	t.Log(pair.Nonces(pairAddress, common.HexToAddress("0xadD275a8Ee37acC0E82F6EB1F5ccAa3b2B51E8C0")))
}

func TestPair_PermitSign(t *testing.T) {

	//bytes32, _ := abi.NewType("bytes32", "", nil)
	//address, _ := abi.NewType("address", "", nil)
	//uint256, _ := abi.NewType("uint256", "", nil)
	//args := abi.Arguments{
	//	{Name: "PERMIT_TYPEHASH", Type: bytes32},
	//	{Name: "owner", Type: address},
	//	{Name: "spender", Type: address},
	//	{Name: "permit_value", Type: uint256},
	//	{Name: "nonce", Type: uint256},
	//	{Name: "deadline", Type: uint256},
	//}
	//fmt.Println(args)
	//fmt.Println(args.Pack([32]byte{1, 23, 2}, sender))

	//uint8t, _ := abi.NewType("uint8", "", nil)
	//args := abi.Arguments{
	//	{Type: uint8t},
	//}
	//
	//b, err := args.Pack(uint8(1))
	//fmt.Println(hexutil.Encode(b), err)
	//fmt.Println(crypto.Keccak256Hash(b))
	//fmt.Println(crypto.Keccak256Hash([]byte("asd")))
	pairContract := NewPairContract(engine, pairAddress)
	t.Log(pairContract.PermitSign(router2,
		decimal.NewFromInt(197499999999950).BigInt(),
		big.NewInt(1657767358),
		senderPrivateKey))
}
