package types

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
)

// Receipt represents the results of a transaction.
type Receipt struct {
	BlockNumber     int       `json:"blockNumber"`
	ContractAddress string    `json:"contractAddress"`
	From            string    `json:"from"`
	GasUsed         string    `json:"gasUsed"`
	Hash            string    `json:"hash"`
	Input           string    `json:"input"`
	Logs            []*NewLog `json:"logEntries"`
	Message         string    `json:"message"`
	Output          string    `json:"output"`
	Status          int       `json:"status"`
	To              string    `json:"to"`
	TransactionHash string    `json:"transactionHash"`
	ReceiptProof    []string  `json:"txReceiptProof"`
	Version         uint64    `json:"version"`
}

const (
	Success = iota
	Unknown
	BadRLP
	InvalidFormat
	OutOfGasIntrinsic
	InvalidSignature
	InvalidNonce
	NotEnoughCash
	OutOfGasBase
	BlockGasLimitReached
	BadInstruction
	BadJumpDestination
	OutOfGas
	OutOfStack
	StackUnderflow
	NonceCheckFail
	BlockLimitCheckFail
	FilterCheckFail
	NoDeployPermission
	NoCallPermission
	NoTxPermission
	PrecompiledError
	RevertInstruction
	InvalidZeroSignatureFormat
	AddressAlreadyUsed
	PermissionDenied
	CallAddressError
)

// GetStatusMessage returns the status message
func getStatusMessage(status int) string {
	var message string
	switch status {
	case Success:
		message = "success"
	case Unknown:
		message = "unknown"
	case BadRLP:
		message = "bad RLP"
	case InvalidFormat:
		message = "invalid format"
	case OutOfGasIntrinsic:
		message = "out of gas intrinsic"
	case InvalidSignature:
		message = "invalid signature"
	case InvalidNonce:
		message = "invalid nonce"
	case NotEnoughCash:
		message = "not enough cash"
	case OutOfGasBase:
		message = "out of gas base"
	case BlockGasLimitReached:
		message = "block gas limit reached"
	case BadInstruction:
		message = "bad instruction"
	case BadJumpDestination:
		message = "bad jump destination"
	case OutOfGas:
		message = "out of gas"
	case OutOfStack:
		message = "out of stack"
	case StackUnderflow:
		message = "stack underflow"
	case NonceCheckFail:
		message = "nonce check fail"
	case BlockLimitCheckFail:
		message = "block limit check fail"
	case FilterCheckFail:
		message = "filter check fail"
	case NoDeployPermission:
		message = "no deploy permission"
	case NoCallPermission:
		message = "no call permission"
	case NoTxPermission:
		message = "no tx permission"
	case PrecompiledError:
		message = "precompiled error"
	case RevertInstruction:
		message = "revert instruction"
	case InvalidZeroSignatureFormat:
		message = "invalid zero signature format"
	case AddressAlreadyUsed:
		message = "address already used"
	case PermissionDenied:
		message = "permission denied"
	case CallAddressError:
		message = "call address error"
	default:
		message = strconv.Itoa(status)
	}

	return message
}

// GetBlockNumber returns the block number string
func (r *Receipt) GetBlockNumber() int {
	return r.BlockNumber
}

// GetContractAddress returns the contract address
func (r *Receipt) GetContractAddress() string {
	return r.ContractAddress
}

// GetFrom returns the transaction sender address
func (r *Receipt) GetFrom() string {
	return r.From
}

// GetGasUsed returns the used gas
func (r *Receipt) GetGasUsed() string {
	return r.GasUsed
}

func (r *Receipt) GetHash() string {
	return r.Hash
}

// GetInput returns the transaction content
func (r *Receipt) GetInput() string {
	return r.Input
}

// GetTransactionHash returns the transaction hash string
func (r *Receipt) GetTransactionHash() string {
	return r.TransactionHash
}

// GetOutput returns the transaction output
func (r *Receipt) GetOutput() string {
	return r.Output
}

// GetTransactionIndex returns the transaction index string
func (r *Receipt) GetReceiptProof() []string {
	return r.ReceiptProof
}

// GetStatus returns the transaction status
func (r *Receipt) GetStatus() int {
	return r.Status
}

// GetTo returns the transaction receiver address
func (r *Receipt) GetTo() string {
	return r.To
}

// GetTo returns the transaction receiver address
func (r *Receipt) GetVersion() uint64 {
	return r.Version
}

// ParseErrorMessage gets unusual output value from Receipt
func (r *Receipt) GetErrorMessage() string {
	var errorMessage string
	if r.GetStatus() == Success {
		return ""
	}
	errorMessage = getStatusMessage(r.Status)
	if len(r.Output) >= 138 { // 0x + 4 bytes funcName + 32 bytes offset + 32 bytes string length
		outputBytes, err := hex.DecodeString(r.Output[2:])
		if err != nil {
			panic(fmt.Sprintf("GetErrorMessage failed, hex.DecodeString err: %v", err))
		}
		errorMessage += ", " + string(outputBytes[68:])
		return fmt.Sprintf("receipt error code: %v, receipt error message: %v", r.Status, errorMessage)
	}
	return fmt.Sprintf("receipt error code: %v, receipt error message: %v", r.Status, errorMessage)
}

// String returns the string representation of Receipt sturct.
func (r *Receipt) String() string {
	out, err := json.MarshalIndent(r, "", "\t")
	if err != nil {
		return fmt.Sprintf("%v", err)
	}
	return string(out)
}
