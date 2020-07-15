package cns

import "encoding/json"

// Info is used for the CNSService
type Info struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Address string `json:"address"`
	Abi     string `json:"abi"`
}

// String method of Info
func (cns *Info) String() string {
	out, err := json.MarshalIndent(cns, "", "\t")
	if err != nil {
		return "get string of Info failed"
	}
	return string(out)
}
