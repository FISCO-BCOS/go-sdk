package types

// NewLog is used for the receipt
type NewLog struct {
	// The Removed field is true if this log was reverted due to a chain reorganisation.
	// You must pay attention to this field if you receive logs through a filter query.
	Removed            bool `json:"removed"`
	// index of the log in the block
	LogIndex           string `json:"logIndex" `
	// index of the transaction in the block
	TransactionIndex   string `json:"transactionIndex" `
	// hash of the transaction
	TransactionHash    string `json:"transactionHash"`
	// hash of the block in which the transaction was included
	BlockHash          string `json:"blockHash"`
	// Derived fields. These fields are filled in by the node
	// but not secured by consensus.
	// block in which the transaction was included
	BlockNumber        string `json:"blockNumber"`
	// Consensus fields:
	// address of the contract that generated the event
	Address            string `json:"address"`
	// supplied by the contract, usually ABI-encoded
	Data               string `json:"data"`
	// Type for FISCO BCOS
	Type               string `json:"type"`
	// list of topics provided by the contract.
	Topics             []interface{} `json:"topics" `
}