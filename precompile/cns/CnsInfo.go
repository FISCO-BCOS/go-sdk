package cns

import "encoding/json"

// CnsInfo is used for the CNSService
type CnsInfo struct {
	Name      string  `json:"name"`
	Version   string  `json:"version"`
	Address   string  `json:"address"`
	Abi       string  `json:"abi"`
}

// GetName returns the name of the CNS
func (cns *CnsInfo) GetName() string {
	return cns.Name
}

// GetVersion returns the version of the CNS
func (cns *CnsInfo) GetVersion() string {
	return cns.Version
}

// GetAddress returns the address of the CNS
func (cns *CnsInfo) GetAddress() string {
	return cns.Address
}

// GetAbi returns the abi of the CNS
func (cns *CnsInfo) GetAbi() string {
	return cns.Abi
}

// SetName modified the name of CNS
func (cns *CnsInfo) SetName(newName string) {
	cns.Name = newName
}

// SetVersion modified the version of CNS
func (cns *CnsInfo) SetVersion(newVersion string) {
	cns.Version =  newVersion
}

// SetAddress modified the address of CNS
func (cns *CnsInfo) SetAddress(newAddress string)  {
	cns.Address = newAddress
}

// SetAbi modified the abi of CNS
func (cns *CnsInfo) SetAbi(newAbi string)  {
	cns.Abi = newAbi
}

// String method of CnsInfo
func (cns *CnsInfo) String() string {
	out, err := json.MarshalIndent(cns, "", "\t")
	if err != nil {
		return "get string of CnsInfo failed"
	}
	return string(out)
}