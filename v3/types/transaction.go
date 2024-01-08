package types

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"

	"github.com/FISCO-BCOS/go-sdk/v3/smcrypto/sm3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
)

var (
	ErrInvalidSig = errors.New("invalid raw transaction v, r, s values")
)

type Transaction struct {
	Data       transactionData
	DataHash   *common.Hash    `json:"dataHash"`
	Signature  []byte          `json:"signature"`
	ImportTime int64           `json:"importTime"`
	Attribute  int             `json:"attribute"`
	Sender     *common.Address `json:"sender"` // nil means contract creation
	ExtraData  string          `json:"extraData"`
	SMCrypto   bool            `json:"-"`

	// Signature values
	V *big.Int `json:"v"`
	R *big.Int `json:"r"`
	S *big.Int `json:"s"`
}

type transactionData struct {
	Version              int32           `json:"version"`
	ChainID              string          `json:"chainID"`
	GroupID              string          `json:"groupID"`
	BlockLimit           int64           `json:"blockLimit"`
	Nonce                string          `json:"nonce"`
	To                   *common.Address `json:"to"` // nil means contract creation
	Input                []byte          `json:"input"`
	Abi                  string          `json:"abi"`
	Value                *big.Int        `json:"value"`
	GasPrice             *big.Int        `json:"gasPrice"`
	GasLimit             int64           `json:"gasLimit"`
	MaxFeePerGas         *big.Int        `json:"maxFeePerGas"`
	MaxPriorityFeePerGas *big.Int        `json:"maxPriorityFeePerGas"`
}

// NewTransaction returns a new transaction
func NewTransaction(to common.Address, amount *big.Int, gasLimit int64, gasPrice *big.Int, blockLimit int64, data []byte, nonce, chainId, groupId, extraData string, smcrypto bool) *Transaction {
	return newTransaction(&to, amount, gasLimit, gasPrice, blockLimit, data, nonce, chainId, groupId, extraData, smcrypto)
}

// NewContractCreation creates a contract transaction
func NewContractCreation(amount *big.Int, gasLimit int64, blockLimit int64, data []byte, nonce, chainId, groupId, extraData string, smcrypto bool) *Transaction {
	return newTransaction(nil, amount, gasLimit, nil, blockLimit, data, nonce, chainId, groupId, extraData, smcrypto)
}

// NewSimpleTx creates a contract transaction, if nonce is empty string, the nonce will be auto generated
func NewSimpleTx(to *common.Address, data []byte, abi, nonce, extraData string, smcrypto bool) *Transaction {
	if len(data) > 0 {
		data = common.CopyBytes(data)
	}
	d := transactionData{
		To:    to,
		Input: data,
		Abi:   abi,
		Nonce: nonce,
	}
	return &Transaction{Data: d, SMCrypto: smcrypto, ExtraData: extraData}
}

func newTransaction(to *common.Address, amount *big.Int, gasLimit int64, gasPrice *big.Int, blockLimit int64, data []byte, nonce, chainId, groupId string, extraData string, smcrypto bool) *Transaction {
	if len(data) > 0 {
		data = common.CopyBytes(data)
	}
	d := transactionData{
		Nonce:      nonce,
		To:         to,
		Input:      data,
		Value:      new(big.Int),
		GasLimit:   gasLimit,
		BlockLimit: blockLimit,
		GasPrice:   new(big.Int),
		ChainID:    chainId,
		GroupID:    groupId,
	}
	if amount != nil {
		d.Value.Set(amount)
	}
	if gasPrice != nil {
		d.GasPrice.Set(gasPrice)
	}

	return &Transaction{Data: d, SMCrypto: smcrypto, ExtraData: extraData}
}

// ChainID returns which chain id this transaction was signed for (if at all)
func (tx *Transaction) ChainID() *big.Int {
	return deriveChainID(tx.V)
}

// Protected returns whether the transaction is protected from replay protection.
func (tx *Transaction) Protected() bool {
	return isProtectedV(tx.V)
}

func isProtectedV(V *big.Int) bool {
	if V.BitLen() <= 8 {
		v := V.Uint64()
		return v != 27 && v != 28
	}
	// anything not 27 or 28 is considered protected
	return true
}

func (tx *Transaction) Input() []byte      { return common.CopyBytes(tx.Data.Input) }
func (tx *Transaction) ABI() string        { return tx.Data.Abi }
func (tx *Transaction) GasPrice() *big.Int { return new(big.Int).Set(tx.Data.GasPrice) }
func (tx *Transaction) Value() *big.Int    { return new(big.Int).Set(tx.Data.Value) }
func (tx *Transaction) Nonce() string      { return tx.Data.Nonce }
func (tx *Transaction) CheckNonce() bool   { return true }

// To returns the recipient address of the transaction.
// It returns nil if the transaction is a contract creation.
func (tx *Transaction) To() *common.Address {
	if tx.Data.To == nil {
		return nil
	}
	to := *tx.Data.To
	return &to
}

// Hash hashes the RLP encoding of tx.
// It uniquely identifies the transaction.
func (tx *Transaction) Hash() common.Hash {
	if hash := tx.DataHash; hash != nil {
		return *hash
	}
	if tx.SMCrypto {
		return tx.sm3HashWithSig()
	}
	v := rlpHash(tx)
	tx.DataHash = &v
	return v
}

// SM3HashNonSig hashes the RLP encoding of tx use sm3.
// It uniquely identifies the transaction.
func (tx *Transaction) SM3HashNonSig() (h common.Hash) {
	var src []byte
	buf := bytes.NewBuffer(src)
	rlp.Encode(buf, []interface{}{
		tx.Data.Nonce,
		tx.Data.GasPrice,
		tx.Data.GasLimit,
		tx.Data.BlockLimit,
		tx.Data.To,
		tx.Data.Value,
		tx.Data.Input,
		tx.Data.ChainID,
		tx.Data.GroupID,
	})
	v := sm3.Hash(buf.Bytes())
	copy(h[:], v)
	return h
}

func (tx *Transaction) sm3HashWithSig() (h common.Hash) {

	var src []byte
	buf := bytes.NewBuffer(src)
	rlp.Encode(buf, tx)
	v := sm3.Hash(buf.Bytes())
	copy(h[:], v)
	tx.DataHash = &h
	return h
}

func rlpHash(x interface{}) (h common.Hash) {
	hw := sha3.NewLegacyKeccak256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
}

// WithSignature returns a new transaction with the given signature.
// This signature needs to be in the [R || S || V] format where V is 0 or 1.
func (tx *Transaction) WithSignature(signer Signer, sig []byte) (*Transaction, error) {
	r, s, v, err := signer.SignatureValues(tx, sig)
	if err != nil {
		return nil, err
	}
	cpy := &Transaction{Data: tx.Data}
	cpy.R, cpy.S, cpy.V = r, s, v
	cpy.SMCrypto = tx.SMCrypto
	return cpy, nil
}

// WithSM2Signature returns a new transaction with the given signature.
// This signature needs to be in the [R || S || V] format where V is 0 or 1.
func (tx *Transaction) WithSM2Signature(signer Signer, sig []byte) (*Transaction, error) {
	if len(sig) != 128 {
		panic(fmt.Sprintf("wrong size for SM2Signature: got %d, want 128", len(sig)))
	}
	r := new(big.Int).SetBytes(sig[:32])
	s := new(big.Int).SetBytes(sig[32:64])
	v := new(big.Int).SetBytes(sig[64:])
	cpy := &Transaction{Data: tx.Data}
	cpy.R, cpy.S, cpy.V = r, s, v
	cpy.SMCrypto = tx.SMCrypto
	return cpy, nil
}

// Cost returns amount + gasprice * gaslimit.
func (tx *Transaction) Cost() *big.Int {
	total := new(big.Int).Mul(tx.Data.GasPrice, new(big.Int).SetInt64(tx.Data.GasLimit))
	total.Add(total, tx.Data.Value)
	return total
}

// SignatureValues returns the V, R, S signature values of the transaction.
// The return values should not be modified by the caller.
func (tx *Transaction) SignatureValues() (v, r, s *big.Int) {
	return tx.V, tx.R, tx.S
}

// Transactions is a Transaction slice type for basic sorting.
type Transactions []*Transaction

// Len returns the length of s.
func (s Transactions) Len() int { return len(s) }

// Swap swaps the i'th and the j'th element in s.
func (s Transactions) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// GetRlp implements Rlpable and returns the i'th element of s in rlp.
func (s Transactions) GetRlp(i int) []byte {
	enc, _ := rlp.EncodeToBytes(s[i])
	return enc
}

// TxDifference returns a new set which is the difference between a and b.
func TxDifference(a, b Transactions) Transactions {
	keep := make(Transactions, 0, len(a))

	remove := make(map[common.Hash]struct{})
	for _, tx := range b {
		remove[tx.Hash()] = struct{}{}
	}

	for _, tx := range a {
		if _, ok := remove[tx.Hash()]; !ok {
			keep = append(keep, tx)
		}
	}

	return keep
}
