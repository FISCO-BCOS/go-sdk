package types

import (
	"fmt"
	"encoding/json"

	"github.com/KasperLiu/gobcos/common"
)

// Receipt represents the results of a transaction.
type Receipt struct {
	TransactionHash      string   `json:"transactionHash"`
	TransactionIndex     string   `json:"transactionIndex"`
	BlockHash            string   `json:"blockHash"`
	BlockNumber          string   `json:"blockNumber"`
	GasUsed              string   `json:"gasUsed"`
	ContractAddress      string   `json:"contractAddress"`
	Root                 string   `json:"root"`
	Status               string   `json:"status"`
	From                 string   `json:"from"`
	To                   string   `json:"to"`
	Input                string   `json:"input"`
	Output               string   `json:"output"`
	Logs                 []*NewLog   `json:"logs"`
	LogsBloom            string   `json:"logsBloom"`
}

// GetTransactionHash returns the transaction hash string
func (r *Receipt) GetTransactionHash() string {
	return r.TransactionHash
}

// GetTransactionIndex returns the transaction index string
func (r *Receipt) GetTransactionIndex() string {
	return r.TransactionIndex
}

// GetBlockHash returns the block hash string
func (r *Receipt) GetBlockHash() string {
	return r.BlockHash
}

// GetBlockNumber returns the block number string
func (r *Receipt) GetBlockNumber() string {
	return r.BlockNumber
}

// GetGasUsed returns the used gas
func (r *Receipt) GetGasUsed() string {
	return r.GasUsed
}

// GetContractAddress returns the contract address
func (r *Receipt) GetContractAddress() common.Address {
	return common.HexToAddress(r.ContractAddress)
}

// GetRoot returns the transactions root
func (r *Receipt) GetRoot() string {
	return r.Root
}

// GetStatus returns the transaction status
func (r *Receipt) GetStatus() string {
	return r.Status
}

// GetFrom returns the transaction sender address
func (r *Receipt) GetFrom() string {
	return r.From
}

// GetTo returns the transaction receiver address
func (r *Receipt) GetTo() string {
	return r.To
}

// GetInput returns the transaction content
func (r *Receipt) GetInput() string {
	return r.Input
}

// GetOutput returns the transaction output
func (r *Receipt) GetOutput() string {
	return r.Output
}

// String returns the string representation of Receipt sturct.
func (r *Receipt) String() string {
    out, err := json.MarshalIndent(r, "", "\t")
    if err != nil {
        return fmt.Sprintf("%v", err)
	}
	return fmt.Sprintf(string(out))
}