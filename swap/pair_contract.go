package swap

import (
	"fmt"
	"github.com/YeMiao0715/ether"
	"github.com/YeMiao0715/ether/erc20"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

type PairContract struct {
	contract common.Address
	*Pair
	*erc20.Erc20Contract
}

func NewPairContract(engine *ether.Engine, contract common.Address) *PairContract {
	return &PairContract{
		contract:      contract,
		Pair:          NewPair(engine),
		Erc20Contract: erc20.NewErc20WithContract(engine, contract),
	}
}

func (p *PairContract) Factory() (common.Address, error) {
	return p.Pair.Factory(p.contract)
}

func (p *PairContract) Token0() (common.Address, error) {
	return p.Pair.Token0(p.contract)
}

func (p *PairContract) Token1() (common.Address, error) {
	return p.Pair.Token1(p.contract)
}

func (p *PairContract) Price0CumulativeLast() (*big.Int, error) {
	return p.Pair.Price0CumulativeLast(p.contract)
}

func (p *PairContract) Price1CumulativeLast() (*big.Int, error) {
	return p.Pair.Price1CumulativeLast(p.contract)
}

func (p *PairContract) KLast() (*big.Int, error) {
	return p.Pair.KLast(p.contract)
}

func (p *PairContract) GetReserves() (reserve0, reserve1 *big.Int, blockTimestampLast uint32, err error) {
	return p.Pair.GetReserves(p.contract)
}

func (p *PairContract) DOMAIN_SEPARATOR() (common.Hash, error) {
	return p.Pair.DOMAIN_SEPARATOR(p.contract)
}

func (p *PairContract) PERMIT_TYPEHASH() (common.Hash, error) {
	return p.Pair.PERMIT_TYPEHASH(p.contract)
}

func (p *PairContract) Nonces(addr common.Address) (*big.Int, error) {
	return p.Pair.Nonces(p.contract, addr)
}

func (p *PairContract) PermitSign(sender common.Address, amount *big.Int, deadline *big.Int, privateKey string) (v, r, s *big.Int, err error) {
	PERMIT_TYPEHASH, err := p.Pair.PERMIT_TYPEHASH(p.contract)
	if err != nil {
		return
	}
	DOMAIN_SEPARATOR, err := p.Pair.DOMAIN_SEPARATOR(p.contract)
	if err != nil {
		return
	}
	fmt.Println(PERMIT_TYPEHASH)
	fmt.Println(DOMAIN_SEPARATOR)
	bytes32, _ := abi.NewType("bytes32", "", nil)
	address, _ := abi.NewType("address", "", nil)
	uint256, _ := abi.NewType("uint256", "", nil)
	args := abi.Arguments{
		{Name: "PERMIT_TYPEHASH", Type: bytes32},
		{Name: "owner", Type: address},
		{Name: "spender", Type: address},
		{Name: "permit_value", Type: uint256},
		{Name: "nonce", Type: uint256},
		{Name: "deadline", Type: uint256},
	}

	owner, err := p.engine.PrivateKeyToAddress(privateKey)
	if err != nil {
		return
	}
	nonce, err := p.Nonces(*owner)
	if err != nil {
		return
	}
	fmt.Println(PERMIT_TYPEHASH, *owner, sender, amount, nonce, deadline)
	b, err := args.Pack(PERMIT_TYPEHASH, *owner, sender, amount, nonce, deadline)
	if err != nil {
		return
	}

	DETAIL_HASH := crypto.Keccak256(b)
	fmt.Println(hexutil.Encode(DETAIL_HASH))
	//bytes, _ := abi.NewType("bytes", "", nil)
	//args = abi.Arguments{
	//	{Name: "", Type: bytes},
	//	{Name: "DOMAIN_SEPARATOR", Type: bytes},
	//	{Name: "DETAIL_HASH", Type: bytes},
	//}
	//fmt.Println(args)
	//b, err = args.Pack([]byte("\\x19\\x01"), DOMAIN_SEPARATOR[:], DETAIL_HASH)
	//if err != nil {
	//	return
	//}
	fmt.Println(hexutil.Encode([]byte("\x19\x01")), hexutil.Encode(DOMAIN_SEPARATOR[:]), hexutil.Encode(DETAIL_HASH))
	hash := crypto.Keccak256Hash([]byte("\x19\x01"), DOMAIN_SEPARATOR[:], DETAIL_HASH)
	fmt.Println(len(hash), hash)

	priv, err := p.engine.HexToEcdsaPrivateKey(privateKey)
	if err != nil {
		return
	}

	sig, err := crypto.Sign(hash.Bytes(), priv)
	if err != nil {
		return
	}

	r = new(big.Int).SetBytes(sig[:32])
	s = new(big.Int).SetBytes(sig[32:64])
	v = new(big.Int).SetBytes([]byte{sig[64] + 27})

	return
}
