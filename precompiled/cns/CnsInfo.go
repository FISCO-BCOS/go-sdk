package cns

import "encoding/json"

// Info is used for the CNSService
type Info struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Address string `json:"address"`
	Abi     string `json:"abi"`
}

// GetName returns the name of the CNS
func (cns *Info) GetName() string {
	return cns.Name
}

// GetVersion returns the version of the CNS
func (cns *Info) GetVersion() string {
	return cns.Version
}

// GetAddress returns the address of the CNS
func (cns *Info) GetAddress() string {
	return cns.Address
}

// GetAbi returns the abi of the CNS
func (cns *Info) GetAbi() string {
	return cns.Abi
}

// SetName modified the name of CNS
func (cns *Info) SetName(newName string) {
	cns.Name = newName
}

// SetVersion modified the version of CNS
func (cns *Info) SetVersion(newVersion string) {
	cns.Version = newVersion
}

// SetAddress modified the address of CNS
func (cns *Info) SetAddress(newAddress string) {
	cns.Address = newAddress
}

// SetAbi modified the abi of CNS
func (cns *Info) SetAbi(newAbi string) {
	cns.Abi = newAbi
}

// String method of Info
func (cns *Info) String() string {
	out, err := json.MarshalIndent(cns, "", "\t")
	if err != nil {
		return "get string of Info failed"
	}
	return string(out)
}
