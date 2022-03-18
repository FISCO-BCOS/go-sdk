package types

type EventLogParams struct {
	FromBlock string   `json:"fromBlock"`
	ToBlock   string   `json:"toBlock"`
	Addresses []string `json:"addresses"`
	Topics    []string `json:"topics"`
	GroupID   string   `json:"groupID"`
	FilterID  string   `json:"filterID"`
}
