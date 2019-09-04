package types

import (
	"container/heap"
	"errors"
	"io"
	"math/big"
	"sync/atomic"
	
	"github.com/KasperLiu/gobcos/common"
	"github.com/KasperLiu/gobcos/common/hexutil"
	"github.com/KasperLiu/gobcos/crypto"
	"github.com/KasperLiu/gobcos/rlp"
	"golang.org/x/crypto/sha3"
)

//go:generate gencodec -type rawtxdata -field-override txdataMarshaling -out gen_rawtx_json.go

var (
	ErrInvalidRawSig = errors.New("invalid raw transaction v, r, s values")
)

type RawTransaction struct {
	data rawtxdata
	// caches
	hash atomic.Value
	size atomic.Value
	from atomic.Value
}

type rawtxdata struct {
	AccountNonce *big.Int        `json:"nonce"    gencodec:"required"`
	Price        *big.Int        `json:"gasPrice"   gencodec:"required"`
	GasLimit     *big.Int         `json:"gas"        gencodec:"required"`
	BlockLimit   *big.Int        `json:"blocklimit" gencodec:"required"` 
	Recipient    *common.Address `json:"to"         rlp:"nil"` // nil means contract creation
	Amount       *big.Int        `json:"value"      gencodec:"required"`
	Payload      []byte          `json:"input"      gencodec:"required"`
	// for fisco bcos 2.0
	ChainId       *big.Int        `json:"chainId"    gencodec:"required"`
	GroupId       *big.Int        `json:"groupId"    gencodec:"required"`
	ExtraData      []byte         `json:"extraData"    rlp:"nil"` //gencodec:"required"`

	// Signature values
	V *big.Int `json:"v" gencodec:"required"`
	R *big.Int `json:"r" gencodec:"required"`
	S *big.Int `json:"s" gencodec:"required"`

	// This is only used when marshaling to JSON.
	Hash *common.Hash `json:"hash" rlp:"-"`
}

type rawtxdataMarshaling struct {
	AccountNonce *hexutil.Big
	Price        *hexutil.Big
	GasLimit     *hexutil.Big
	BlockLimit   *hexutil.Big
	Amount       *hexutil.Big
	Payload      hexutil.Bytes
	// for fisco bcos 2.0
	ChainId     *hexutil.Big
	GroupId     *hexutil.Big
	ExtraData   hexutil.Bytes

	V            *hexutil.Big
	R            *hexutil.Big
	S            *hexutil.Big
}

// NewRawTransaction returns a new raw transaction
func NewRawTransaction(nonce *big.Int, to common.Address, amount *big.Int, gasLimit *big.Int, gasPrice *big.Int, blockLimit *big.Int, data []byte, chainId *big.Int , groupId *big.Int, extraData []byte) *RawTransaction {
	return newRawTransaction(nonce, &to, amount, gasLimit, gasPrice, blockLimit, data, chainId, groupId, extraData)
}

// NewRawContractCreation creates a contract transaction
func NewRawContractCreation(nonce *big.Int, amount *big.Int, gasLimit *big.Int, gasPrice *big.Int, blockLimit *big.Int, data []byte, chainId *big.Int , groupId *big.Int, extraData []byte) *RawTransaction {
	return newRawTransaction(nonce, nil, amount, gasLimit, gasPrice, blockLimit, data,chainId, groupId, extraData)
}

func newRawTransaction(nonce *big.Int, to *common.Address, amount *big.Int, gasLimit *big.Int, gasPrice *big.Int, blockLimit *big.Int, data []byte, chainId *big.Int , groupId *big.Int, extraData []byte) *RawTransaction {
	if len(data) > 0 {
		data = common.CopyBytes(data)
	}
	d := rawtxdata{
		AccountNonce: nonce,
		Recipient:    to,
		Payload:      data,
		Amount:       new(big.Int),
		GasLimit:     gasLimit,
		BlockLimit:   blockLimit,
		Price:        new(big.Int),
		ChainId:	  new(big.Int),
		GroupId:	  new(big.Int),
		ExtraData:	  extraData,
		V:            new(big.Int),
		R:            new(big.Int),
		S:            new(big.Int),
	}
	if amount != nil {
		d.Amount.Set(amount)
	}
	if gasPrice != nil {
		d.Price.Set(gasPrice)
	}
	if chainId != nil {
		d.ChainId.Set(chainId)
	}
	if groupId != nil {
		d.GroupId.Set(groupId)
	}
	if extraData != nil {
		d.ExtraData = extraData
	}


	return &RawTransaction{data: d}
}

// ChainId returns which chain id this transaction was signed for (if at all)
func (tx *RawTransaction) ChainId() *big.Int {
	return deriveChainId(tx.data.V)
}

// Protected returns whether the transaction is protected from replay protection.
func (tx *RawTransaction) Protected() bool {
	return isProtectedV(tx.data.V)
}

func isProtectedV(V *big.Int) bool {
	if V.BitLen() <= 8 {
		v := V.Uint64()
		return v != 27 && v != 28
	}
	// anything not 27 or 28 is considered protected
	return true
}

// EncodeRLP implements rlp.Encoder
func (tx *RawTransaction) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, &tx.data)
}

// DecodeRLP implements rlp.Decoder
func (tx *RawTransaction) DecodeRLP(s *rlp.Stream) error {
	_, size, _ := s.Kind()
	err := s.Decode(&tx.data)
	if err == nil {
		tx.size.Store(common.StorageSize(rlp.ListSize(size)))
	}

	return err
}

// MarshalJSON encodes the web3 RPC transaction format.
func (tx *RawTransaction) MarshalJSON() ([]byte, error) {
	hash := tx.Hash()
	data := tx.data
	data.Hash = &hash
	return data.MarshalJSON()
}

// UnmarshalJSON decodes the web3 RPC transaction format.
func (tx *RawTransaction) UnmarshalJSON(input []byte) error {
	var dec rawtxdata
	if err := dec.UnmarshalJSON(input); err != nil {
		return err
	}

	withSignature := dec.V.Sign() != 0 || dec.R.Sign() != 0 || dec.S.Sign() != 0
	if withSignature {
		var V byte
		if isProtectedV(dec.V) {
			chainID := deriveChainId(dec.V).Uint64()
			V = byte(dec.V.Uint64() - 35 - 2*chainID)
		} else {
			V = byte(dec.V.Uint64() - 27)
		}
		if !crypto.ValidateSignatureValues(V, dec.R, dec.S, false) {
			return ErrInvalidRawSig
		}
	}

	*tx = RawTransaction{data: dec}
	return nil
}

func (tx *RawTransaction) Data() []byte       { return common.CopyBytes(tx.data.Payload) }
func (tx *RawTransaction) Gas() *big.Int        { return tx.data.GasLimit }
func (tx *RawTransaction) GasPrice() *big.Int { return new(big.Int).Set(tx.data.Price) }
func (tx *RawTransaction) Value() *big.Int    { return new(big.Int).Set(tx.data.Amount) }
func (tx *RawTransaction) Nonce() *big.Int      { return tx.data.AccountNonce }
func (tx *RawTransaction) CheckNonce() bool   { return true }

// To returns the recipient address of the transaction.
// It returns nil if the transaction is a contract creation.
func (tx *RawTransaction) To() *common.Address {
	if tx.data.Recipient == nil {
		return nil
	}
	to := *tx.data.Recipient
	return &to
}

// Hash hashes the RLP encoding of tx.
// It uniquely identifies the transaction.
func (tx *RawTransaction) Hash() common.Hash {
	if hash := tx.hash.Load(); hash != nil {
		return hash.(common.Hash)
	}
	v := rlpHash(tx)
	tx.hash.Store(v)
	return v
}

func rlpHash(x interface{}) (h common.Hash) {
	hw := sha3.NewLegacyKeccak256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
}

// Size returns the true RLP encoded storage size of the transaction, either by
// encoding and returning it, or returning a previsouly cached value.
func (tx *RawTransaction) Size() common.StorageSize {
	if size := tx.size.Load(); size != nil {
		return size.(common.StorageSize)
	}
	c := writeCounter(0)
	rlp.Encode(&c, &tx.data)
	tx.size.Store(common.StorageSize(c))
	return common.StorageSize(c)
}

type writeCounter common.StorageSize

func (c *writeCounter) Write(b []byte) (int, error) {
	*c += writeCounter(len(b))
	return len(b), nil
}


// AsMessage returns the transaction as a core.Message.
//
// AsMessage requires a signer to derive the sender.
//
// XXX Rename message to something less arbitrary?
func (tx *RawTransaction) AsMessage(s RawSigner) (RawMessage, error) {
	msg := RawMessage{
		nonce:      tx.data.AccountNonce,
		gasLimit:   tx.data.GasLimit,
		gasPrice:   new(big.Int).Set(tx.data.Price),
		blockLimit: tx.data.BlockLimit,
		to:         tx.data.Recipient,
		amount:     tx.data.Amount,
		data:       tx.data.Payload,
		checkNonce: true,
	}

	var err error
	msg.from, err = RawSender(s, tx)
	return msg, err
}

// WithSignature returns a new transaction with the given signature.
// This signature needs to be in the [R || S || V] format where V is 0 or 1.
func (tx *RawTransaction) WithSignature(signer RawSigner, sig []byte) (*RawTransaction, error) {
	r, s, v, err := signer.SignatureValues(tx, sig)
	if err != nil {
		return nil, err
	}
	cpy := &RawTransaction{data: tx.data}
	cpy.data.R, cpy.data.S, cpy.data.V = r, s, v
	return cpy, nil
}

// Cost returns amount + gasprice * gaslimit.
func (tx *RawTransaction) Cost() *big.Int {
	total := new(big.Int).Mul(tx.data.Price, new(big.Int).Set(tx.data.GasLimit))
	total.Add(total, tx.data.Amount)
	return total
}

// RawSignatureValues returns the V, R, S signature values of the transaction.
// The return values should not be modified by the caller.
func (tx *RawTransaction) RawSignatureValues() (v, r, s *big.Int) {
	return tx.data.V, tx.data.R, tx.data.S
}

// RawTransactions is a Transaction slice type for basic sorting.
type RawTransactions []*RawTransaction

// Len returns the length of s.
func (s RawTransactions) Len() int { return len(s) }

// Swap swaps the i'th and the j'th element in s.
func (s RawTransactions) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// GetRlp implements Rlpable and returns the i'th element of s in rlp.
func (s RawTransactions) GetRlp(i int) []byte {
	enc, _ := rlp.EncodeToBytes(s[i])
	return enc
}

// RawTxDifference returns a new set which is the difference between a and b.
func RawTxDifference(a, b RawTransactions) RawTransactions {
	keep := make(RawTransactions, 0, len(a))

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

// RawTxByNonce implements the sort interface to allow sorting a list of transactions
// by their nonces. This is usually only useful for sorting transactions from a
// single account, otherwise a nonce comparison doesn't make much sense.
type RawTxByNonce RawTransactions

func (s RawTxByNonce) Len() int           { return len(s) }
func (s RawTxByNonce) Less(i, j int) bool { return s[i].data.AccountNonce.Cmp(s[j].data.AccountNonce) > 0} //{ return s[i].data.AccountNonce < s[j].data.AccountNonce }
func (s RawTxByNonce) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// RawTxByPrice implements both the sort and the heap interface, making it useful
// for all at once sorting as well as individually adding and removing elements.
type RawTxByPrice RawTransactions

func (s RawTxByPrice) Len() int           { return len(s) }
func (s RawTxByPrice) Less(i, j int) bool { return s[i].data.Price.Cmp(s[j].data.Price) > 0 }
func (s RawTxByPrice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (s *RawTxByPrice) Push(x interface{}) {
	*s = append(*s, x.(*RawTransaction))
}

func (s *RawTxByPrice) Pop() interface{} {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[0 : n-1]
	return x
}

// RawTransactionsByPriceAndNonce represents a set of transactions that can return
// transactions in a profit-maximizing sorted order, while supporting removing
// entire batches of transactions for non-executable accounts.
type RawTransactionsByPriceAndNonce struct {
	txs    map[common.Address]RawTransactions // Per account nonce-sorted list of transactions
	heads  RawTxByPrice                       // Next transaction for each unique account (price heap)
	signer RawSigner                          // Signer for the set of transactions
}

// NewRawTransactionsByPriceAndNonce creates a transaction set that can retrieve
// price sorted transactions in a nonce-honouring way.
//
// Note, the input map is reowned so the caller should not interact any more with
// if after providing it to the constructor.
func NewRawTransactionsByPriceAndNonce(signer RawSigner, txs map[common.Address]RawTransactions) *RawTransactionsByPriceAndNonce {
	// Initialize a price based heap with the head transactions
	heads := make(RawTxByPrice, 0, len(txs))
	for from, accTxs := range txs {
		heads = append(heads, accTxs[0])
		// Ensure the sender address is from the signer
		acc, _ := RawSender(signer, accTxs[0])
		txs[acc] = accTxs[1:]
		if from != acc {
			delete(txs, from)
		}
	}
	heap.Init(&heads)

	// Assemble and return the transaction set
	return &RawTransactionsByPriceAndNonce{
		txs:    txs,
		heads:  heads,
		signer: signer,
	}
}

// Peek returns the next transaction by price.
func (t *RawTransactionsByPriceAndNonce) Peek() *RawTransaction {
	if len(t.heads) == 0 {
		return nil
	}
	return t.heads[0]
}

// Shift replaces the current best head with the next one from the same account.
func (t *RawTransactionsByPriceAndNonce) Shift() {
	acc, _ := RawSender(t.signer, t.heads[0])
	if txs, ok := t.txs[acc]; ok && len(txs) > 0 {
		t.heads[0], t.txs[acc] = txs[0], txs[1:]
		heap.Fix(&t.heads, 0)
	} else {
		heap.Pop(&t.heads)
	}
}

// Pop removes the best transaction, *not* replacing it with the next one from
// the same account. This should be used when a transaction cannot be executed
// and hence all subsequent ones should be discarded from the same account.
func (t *RawTransactionsByPriceAndNonce) Pop() {
	heap.Pop(&t.heads)
}

// Message is a fully derived transaction and implements core.Message
//
// NOTE: In a future PR this will be removed.
type RawMessage struct {
	to         *common.Address
	from       common.Address
	nonce      *big.Int
	amount     *big.Int
	gasLimit   *big.Int
	gasPrice   *big.Int
	blockLimit *big.Int
	data       []byte
	checkNonce bool
}

func NewRawMessage(from common.Address, to *common.Address, nonce *big.Int, amount *big.Int, gasLimit *big.Int, gasPrice *big.Int, blockLimit *big.Int, data []byte, checkNonce bool) RawMessage {
	return RawMessage{
		from:       from,
		to:         to,
		nonce:      nonce,
		amount:     amount,
		gasLimit:   gasLimit,
		gasPrice:   gasPrice,
		blockLimit: blockLimit,
		data:       data,
		checkNonce: checkNonce,
	}
}

func (m RawMessage) From() common.Address { return m.from }
func (m RawMessage) To() *common.Address  { return m.to }
func (m RawMessage) GasPrice() *big.Int   { return m.gasPrice }
func (m RawMessage) Value() *big.Int      { return m.amount }
func (m RawMessage) Gas() *big.Int          { return m.gasLimit }
func (m RawMessage) Nonce() *big.Int        { return m.nonce }
func (m RawMessage) Data() []byte         { return m.data }
func (m RawMessage) CheckNonce() bool     { return m.checkNonce }


// =================================== kasperliu ===============================
type newRawTransactionStruct struct {
	data newrawtxdata
	// caches
	hash atomic.Value
	size atomic.Value
	from atomic.Value
}

type newrawtxdata struct {
	AccountNonce *big.Int        `json:"nonce"    gencodec:"required"`
	Price        *big.Int        `json:"gasPrice"   gencodec:"required"`
	GasLimit     *big.Int        `json:"gas"        gencodec:"required"`
	BlockLimit   *big.Int        `json:"blocklimit" gencodec:"required"`
	Recipient    string          `json:"to"         rlp:"nil"` // nil means contract creation
	Amount       *big.Int        `json:"value"      gencodec:"required"`
	Payload      string          `json:"input"      gencodec:"required"`

	// Signature values
	V *big.Int `json:"v" gencodec:"required"`
	R *big.Int `json:"r" gencodec:"required"`
	S *big.Int `json:"s" gencodec:"required"`

	// This is only used when marshaling to JSON.
	Hash *common.Hash `json:"hash" rlp:"-"`
}

func (tx *RawTransaction) ConverToNewRawTx() (*newRawTransactionStruct) {
	return &newRawTransactionStruct{
		data: newrawtxdata{
			AccountNonce: tx.data.AccountNonce,
			Price       : tx.data.Price,
			GasLimit    : tx.data.GasLimit,
			BlockLimit  : tx.data.BlockLimit,
			Recipient   : tx.data.Recipient.String(),
			Amount      : tx.data.Amount,
			Payload     : hexutil.Bytes(tx.data.Payload).String(),
			V           : tx.data.V,
			R           : tx.data.R,
			S           : tx.data.S,
		},
		hash: tx.hash,
		size: tx.size,
		from: tx.from,
	}
}