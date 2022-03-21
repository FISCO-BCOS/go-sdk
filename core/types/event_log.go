package types

type EventLogParams struct {
	FromBlock string   `json:"fromBlock"`
	ToBlock   string   `json:"toBlock"`
	Addresses []string `json:"addresses"`
	Topics    []string `json:"topics"`
	GroupID   string   `json:"groupID"`
	FilterID  string   `json:"filterID"`
}

type EventLog struct {
	Removed          bool     `json:"removed"`
	LogIndex         string   `json:"logIndex"`
	TransactionIndex string   `json:"transactionIndex"`
	TransactionHash  string   `json:"transactionHash"`
	BlockHash        string   `json:"blockHash"`
	BlockNumber      string   `json:"blockNumber"`
	Address          string   `json:"address"`
	Data             string   `json:"data"`
	Type             string   `json:"type"`
	Topics           []string `json:"topics"`
}
