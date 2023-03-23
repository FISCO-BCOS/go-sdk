// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package auth

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
)

// ContractAuthPrecompiledABI is the input ABI used to generate the binding from.
const ContractAuthPrecompiledABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"funcSelector\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"checkMethodAuth\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"closeDeployAuth\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"funcSelector\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"closeMethodAuth\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"contractAvailable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"}],\"name\":\"getAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"path\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"funcSelector\",\"type\":\"bytes4\"}],\"name\":\"getMethodAuth\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasDeployAuth\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"openDeployAuth\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"funcSelector\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"openMethodAuth\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"resetAdmin\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isFreeze\",\"type\":\"bool\"}],\"name\":\"setContractStatus\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"setDeployAuthType\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"funcSelector\",\"type\":\"bytes4\"},{\"internalType\":\"uint8\",\"name\":\"authType\",\"type\":\"uint8\"}],\"name\":\"setMethodAuthType\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ContractAuthPrecompiledBin is the compiled bytecode used for deploying new contracts.
var ContractAuthPrecompiledBin = "0x608060405234801561001057600080fd5b5061095a806100206000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c806364efb22b1161008c578063bb0aa40c11610066578063bb0aa40c146102bf578063c53057b4146102ef578063cb7c5c111461031f578063d8662aa41461034f576100ea565b806364efb22b1461022f57806381c81cdc1461025f5780639cc3ca0f1461028f576100ea565b80632c8c4a4f116100c85780632c8c4a4f1461016f57806356bd70841461019f57806361548099146101cf578063630577e5146101ff576100ea565b80630578519a146100ef5780630c82b73d146101215780631749bea914610151575b600080fd5b610109600480360381019061010491906104e3565b61037f565b6040516101189392919061075f565b60405180910390f35b61013b6004803603810190610136919061051f565b61038b565b6040516101489190610729565b60405180910390f35b610159610394565b6040516101669190610744565b60405180910390f35b61018960048036038101906101849190610442565b610399565b604051610196919061070e565b60405180910390f35b6101b960048036038101906101b49190610442565b6103a0565b6040516101c69190610729565b60405180910390f35b6101e960048036038101906101e49190610442565b6103a7565b6040516101f69190610729565b60405180910390f35b61021960048036038101906102149190610442565b6103ae565b604051610226919061070e565b60405180910390f35b61024960048036038101906102449190610442565b6103b5565b60405161025691906106f3565b60405180910390f35b610279600480360381019061027491906104a7565b6103bc565b6040516102869190610729565b60405180910390f35b6102a960048036038101906102a4919061056e565b6103c4565b6040516102b69190610729565b60405180910390f35b6102d960048036038101906102d491906105bd565b6103cd565b6040516102e69190610729565b60405180910390f35b6103096004803603810190610304919061046b565b6103d4565b6040516103169190610729565b60405180910390f35b6103396004803603810190610334919061051f565b6103dc565b6040516103469190610729565b60405180910390f35b6103696004803603810190610364919061051f565b6103e5565b604051610376919061070e565b60405180910390f35b60006060809250925092565b60009392505050565b600090565b6000919050565b6000919050565b6000919050565b6000919050565b6000919050565b600092915050565b60009392505050565b6000919050565b600092915050565b60009392505050565b60009392505050565b6000813590506103fd816108c8565b92915050565b600081359050610412816108df565b92915050565b600081359050610427816108f6565b92915050565b60008135905061043c8161090d565b92915050565b60006020828403121561045457600080fd5b6000610462848285016103ee565b91505092915050565b6000806040838503121561047e57600080fd5b600061048c858286016103ee565b925050602061049d858286016103ee565b9150509250929050565b600080604083850312156104ba57600080fd5b60006104c8858286016103ee565b92505060206104d985828601610403565b9150509250929050565b600080604083850312156104f657600080fd5b6000610504858286016103ee565b925050602061051585828601610418565b9150509250929050565b60008060006060848603121561053457600080fd5b6000610542868287016103ee565b935050602061055386828701610418565b9250506040610564868287016103ee565b9150509250925092565b60008060006060848603121561058357600080fd5b6000610591868287016103ee565b93505060206105a286828701610418565b92505060406105b38682870161042d565b9150509250925092565b6000602082840312156105cf57600080fd5b60006105dd8482850161042d565b91505092915050565b60006105f2838361069c565b905092915050565b610603816107f9565b82525050565b6000610614826107b4565b61061e81856107d7565b935083602082028501610630856107a4565b8060005b8581101561066c578484038952815161064d85826105e6565b9450610658836107ca565b925060208a01995050600181019050610634565b50829750879550505050505092915050565b6106878161080b565b82525050565b61069681610843565b82525050565b60006106a7826107bf565b6106b181856107e8565b93506106c1818560208601610884565b6106ca816108b7565b840191505092915050565b6106de8161086d565b82525050565b6106ed81610877565b82525050565b600060208201905061070860008301846105fa565b92915050565b6000602082019050610723600083018461067e565b92915050565b600060208201905061073e600083018461068d565b92915050565b600060208201905061075960008301846106d5565b92915050565b600060608201905061077460008301866106e4565b81810360208301526107868185610609565b9050818103604083015261079a8184610609565b9050949350505050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b60006108048261084d565b9050919050565b60008115159050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b838110156108a2578082015181840152602081019050610887565b838111156108b1576000848401525b50505050565b6000601f19601f8301169050919050565b6108d1816107f9565b81146108dc57600080fd5b50565b6108e88161080b565b81146108f357600080fd5b50565b6108ff81610817565b811461090a57600080fd5b50565b61091681610877565b811461092157600080fd5b5056fea2646970667358221220da854674d347f7edee1db2fbb9b99d40e5dddcf6c82aadab9036da30eb47987664736f6c634300060a0033"

// DeployContractAuthPrecompiled deploys a new contract, binding an instance of ContractAuthPrecompiled to it.
func DeployContractAuthPrecompiled(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Receipt, *ContractAuthPrecompiled, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractAuthPrecompiledABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, receipt, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractAuthPrecompiledBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &ContractAuthPrecompiled{ContractAuthPrecompiledCaller: ContractAuthPrecompiledCaller{contract: contract}, ContractAuthPrecompiledTransactor: ContractAuthPrecompiledTransactor{contract: contract}, ContractAuthPrecompiledFilterer: ContractAuthPrecompiledFilterer{contract: contract}}, nil
}

func AsyncDeployContractAuthPrecompiled(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractAuthPrecompiledABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(ContractAuthPrecompiledBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// ContractAuthPrecompiled is an auto generated Go binding around a Solidity contract.
type ContractAuthPrecompiled struct {
	ContractAuthPrecompiledCaller     // Read-only binding to the contract
	ContractAuthPrecompiledTransactor // Write-only binding to the contract
	ContractAuthPrecompiledFilterer   // Log filterer for contract events
}

// ContractAuthPrecompiledCaller is an auto generated read-only Go binding around a Solidity contract.
type ContractAuthPrecompiledCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAuthPrecompiledTransactor is an auto generated write-only Go binding around a Solidity contract.
type ContractAuthPrecompiledTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAuthPrecompiledFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type ContractAuthPrecompiledFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAuthPrecompiledSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type ContractAuthPrecompiledSession struct {
	Contract     *ContractAuthPrecompiled // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ContractAuthPrecompiledCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type ContractAuthPrecompiledCallerSession struct {
	Contract *ContractAuthPrecompiledCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// ContractAuthPrecompiledTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type ContractAuthPrecompiledTransactorSession struct {
	Contract     *ContractAuthPrecompiledTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ContractAuthPrecompiledRaw is an auto generated low-level Go binding around a Solidity contract.
type ContractAuthPrecompiledRaw struct {
	Contract *ContractAuthPrecompiled // Generic contract binding to access the raw methods on
}

// ContractAuthPrecompiledCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type ContractAuthPrecompiledCallerRaw struct {
	Contract *ContractAuthPrecompiledCaller // Generic read-only contract binding to access the raw methods on
}

// ContractAuthPrecompiledTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type ContractAuthPrecompiledTransactorRaw struct {
	Contract *ContractAuthPrecompiledTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractAuthPrecompiled creates a new instance of ContractAuthPrecompiled, bound to a specific deployed contract.
func NewContractAuthPrecompiled(address common.Address, backend bind.ContractBackend) (*ContractAuthPrecompiled, error) {
	contract, err := bindContractAuthPrecompiled(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractAuthPrecompiled{ContractAuthPrecompiledCaller: ContractAuthPrecompiledCaller{contract: contract}, ContractAuthPrecompiledTransactor: ContractAuthPrecompiledTransactor{contract: contract}, ContractAuthPrecompiledFilterer: ContractAuthPrecompiledFilterer{contract: contract}}, nil
}

// NewContractAuthPrecompiledCaller creates a new read-only instance of ContractAuthPrecompiled, bound to a specific deployed contract.
func NewContractAuthPrecompiledCaller(address common.Address, caller bind.ContractCaller) (*ContractAuthPrecompiledCaller, error) {
	contract, err := bindContractAuthPrecompiled(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAuthPrecompiledCaller{contract: contract}, nil
}

// NewContractAuthPrecompiledTransactor creates a new write-only instance of ContractAuthPrecompiled, bound to a specific deployed contract.
func NewContractAuthPrecompiledTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractAuthPrecompiledTransactor, error) {
	contract, err := bindContractAuthPrecompiled(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAuthPrecompiledTransactor{contract: contract}, nil
}

// NewContractAuthPrecompiledFilterer creates a new log filterer instance of ContractAuthPrecompiled, bound to a specific deployed contract.
func NewContractAuthPrecompiledFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractAuthPrecompiledFilterer, error) {
	contract, err := bindContractAuthPrecompiled(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractAuthPrecompiledFilterer{contract: contract}, nil
}

// bindContractAuthPrecompiled binds a generic wrapper to an already deployed contract.
func bindContractAuthPrecompiled(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractAuthPrecompiledABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAuthPrecompiled *ContractAuthPrecompiledRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractAuthPrecompiled.Contract.ContractAuthPrecompiledCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAuthPrecompiled *ContractAuthPrecompiledRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _ContractAuthPrecompiled.Contract.ContractAuthPrecompiledTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAuthPrecompiled *ContractAuthPrecompiledRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _ContractAuthPrecompiled.Contract.ContractAuthPrecompiledTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAuthPrecompiled *ContractAuthPrecompiledCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractAuthPrecompiled.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _ContractAuthPrecompiled.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _ContractAuthPrecompiled.Contract.contract.Transact(opts, method, params...)
}

// CheckMethodAuth is a free data retrieval call binding the contract method 0xd8662aa4.
//
// Solidity: function checkMethodAuth(address contractAddr, bytes4 funcSelector, address account) constant returns(bool)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledCaller) CheckMethodAuth(opts *bind.CallOpts, contractAddr common.Address, funcSelector [4]byte, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ContractAuthPrecompiled.contract.Call(opts, out, "checkMethodAuth", contractAddr, funcSelector, account)
	return *ret0, err
}

// CheckMethodAuth is a free data retrieval call binding the contract method 0xd8662aa4.
//
// Solidity: function checkMethodAuth(address contractAddr, bytes4 funcSelector, address account) constant returns(bool)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledSession) CheckMethodAuth(contractAddr common.Address, funcSelector [4]byte, account common.Address) (bool, error) {
	return _ContractAuthPrecompiled.Contract.CheckMethodAuth(&_ContractAuthPrecompiled.CallOpts, contractAddr, funcSelector, account)
}

// CheckMethodAuth is a free data retrieval call binding the contract method 0xd8662aa4.
//
// Solidity: function checkMethodAuth(address contractAddr, bytes4 funcSelector, address account) constant returns(bool)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledCallerSession) CheckMethodAuth(contractAddr common.Address, funcSelector [4]byte, account common.Address) (bool, error) {
	return _ContractAuthPrecompiled.Contract.CheckMethodAuth(&_ContractAuthPrecompiled.CallOpts, contractAddr, funcSelector, account)
}

// ContractAvailable is a free data retrieval call binding the contract method 0x2c8c4a4f.
//
// Solidity: function contractAvailable(address _address) constant returns(bool)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledCaller) ContractAvailable(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ContractAuthPrecompiled.contract.Call(opts, out, "contractAvailable", _address)
	return *ret0, err
}

// ContractAvailable is a free data retrieval call binding the contract method 0x2c8c4a4f.
//
// Solidity: function contractAvailable(address _address) constant returns(bool)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledSession) ContractAvailable(_address common.Address) (bool, error) {
	return _ContractAuthPrecompiled.Contract.ContractAvailable(&_ContractAuthPrecompiled.CallOpts, _address)
}

// ContractAvailable is a free data retrieval call binding the contract method 0x2c8c4a4f.
//
// Solidity: function contractAvailable(address _address) constant returns(bool)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledCallerSession) ContractAvailable(_address common.Address) (bool, error) {
	return _ContractAuthPrecompiled.Contract.ContractAvailable(&_ContractAuthPrecompiled.CallOpts, _address)
}

// DeployType is a free data retrieval call binding the contract method 0x1749bea9.
//
// Solidity: function deployType() constant returns(uint256)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledCaller) DeployType(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ContractAuthPrecompiled.contract.Call(opts, out, "deployType")
	return *ret0, err
}

// DeployType is a free data retrieval call binding the contract method 0x1749bea9.
//
// Solidity: function deployType() constant returns(uint256)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledSession) DeployType() (*big.Int, error) {
	return _ContractAuthPrecompiled.Contract.DeployType(&_ContractAuthPrecompiled.CallOpts)
}

// DeployType is a free data retrieval call binding the contract method 0x1749bea9.
//
// Solidity: function deployType() constant returns(uint256)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledCallerSession) DeployType() (*big.Int, error) {
	return _ContractAuthPrecompiled.Contract.DeployType(&_ContractAuthPrecompiled.CallOpts)
}

// GetAdmin is a free data retrieval call binding the contract method 0x64efb22b.
//
// Solidity: function getAdmin(address contractAddr) constant returns(address)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledCaller) GetAdmin(opts *bind.CallOpts, contractAddr common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractAuthPrecompiled.contract.Call(opts, out, "getAdmin", contractAddr)
	return *ret0, err
}

// GetAdmin is a free data retrieval call binding the contract method 0x64efb22b.
//
// Solidity: function getAdmin(address contractAddr) constant returns(address)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledSession) GetAdmin(contractAddr common.Address) (common.Address, error) {
	return _ContractAuthPrecompiled.Contract.GetAdmin(&_ContractAuthPrecompiled.CallOpts, contractAddr)
}

// GetAdmin is a free data retrieval call binding the contract method 0x64efb22b.
//
// Solidity: function getAdmin(address contractAddr) constant returns(address)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledCallerSession) GetAdmin(contractAddr common.Address) (common.Address, error) {
	return _ContractAuthPrecompiled.Contract.GetAdmin(&_ContractAuthPrecompiled.CallOpts, contractAddr)
}

// GetMethodAuth is a free data retrieval call binding the contract method 0x0578519a.
//
// Solidity: function getMethodAuth(address path, bytes4 funcSelector) constant returns(uint8, string[], string[])
func (_ContractAuthPrecompiled *ContractAuthPrecompiledCaller) GetMethodAuth(opts *bind.CallOpts, path common.Address, funcSelector [4]byte) (uint8, []string, []string, error) {
	var (
		ret0 = new(uint8)
		ret1 = new([]string)
		ret2 = new([]string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _ContractAuthPrecompiled.contract.Call(opts, out, "getMethodAuth", path, funcSelector)
	return *ret0, *ret1, *ret2, err
}

// GetMethodAuth is a free data retrieval call binding the contract method 0x0578519a.
//
// Solidity: function getMethodAuth(address path, bytes4 funcSelector) constant returns(uint8, string[], string[])
func (_ContractAuthPrecompiled *ContractAuthPrecompiledSession) GetMethodAuth(path common.Address, funcSelector [4]byte) (uint8, []string, []string, error) {
	return _ContractAuthPrecompiled.Contract.GetMethodAuth(&_ContractAuthPrecompiled.CallOpts, path, funcSelector)
}

// GetMethodAuth is a free data retrieval call binding the contract method 0x0578519a.
//
// Solidity: function getMethodAuth(address path, bytes4 funcSelector) constant returns(uint8, string[], string[])
func (_ContractAuthPrecompiled *ContractAuthPrecompiledCallerSession) GetMethodAuth(path common.Address, funcSelector [4]byte) (uint8, []string, []string, error) {
	return _ContractAuthPrecompiled.Contract.GetMethodAuth(&_ContractAuthPrecompiled.CallOpts, path, funcSelector)
}

// HasDeployAuth is a free data retrieval call binding the contract method 0x630577e5.
//
// Solidity: function hasDeployAuth(address account) constant returns(bool)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledCaller) HasDeployAuth(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ContractAuthPrecompiled.contract.Call(opts, out, "hasDeployAuth", account)
	return *ret0, err
}

// HasDeployAuth is a free data retrieval call binding the contract method 0x630577e5.
//
// Solidity: function hasDeployAuth(address account) constant returns(bool)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledSession) HasDeployAuth(account common.Address) (bool, error) {
	return _ContractAuthPrecompiled.Contract.HasDeployAuth(&_ContractAuthPrecompiled.CallOpts, account)
}

// HasDeployAuth is a free data retrieval call binding the contract method 0x630577e5.
//
// Solidity: function hasDeployAuth(address account) constant returns(bool)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledCallerSession) HasDeployAuth(account common.Address) (bool, error) {
	return _ContractAuthPrecompiled.Contract.HasDeployAuth(&_ContractAuthPrecompiled.CallOpts, account)
}

// CloseDeployAuth is a paid mutator transaction binding the contract method 0x56bd7084.
//
// Solidity: function closeDeployAuth(address account) returns(int256)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactor) CloseDeployAuth(opts *bind.TransactOpts, account common.Address) (*types.Transaction, *types.Receipt, error) {
	return _ContractAuthPrecompiled.contract.Transact(opts, "closeDeployAuth", account)
}

func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactor) AsyncCloseDeployAuth(handler func(*types.Receipt, error), opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ContractAuthPrecompiled.contract.AsyncTransact(opts, handler, "closeDeployAuth", account)
}

// CloseDeployAuth is a paid mutator transaction binding the contract method 0x56bd7084.
//
// Solidity: function closeDeployAuth(address account) returns(int256)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledSession) CloseDeployAuth(account common.Address) (*types.Transaction, *types.Receipt, error) {
	return _ContractAuthPrecompiled.Contract.CloseDeployAuth(&_ContractAuthPrecompiled.TransactOpts, account)
}

func (_ContractAuthPrecompiled *ContractAuthPrecompiledSession) AsyncCloseDeployAuth(handler func(*types.Receipt, error), account common.Address) (*types.Transaction, error) {
	return _ContractAuthPrecompiled.Contract.AsyncCloseDeployAuth(handler, &_ContractAuthPrecompiled.TransactOpts, account)
}

// CloseDeployAuth is a paid mutator transaction binding the contract method 0x56bd7084.
//
// Solidity: function closeDeployAuth(address account) returns(int256)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactorSession) CloseDeployAuth(account common.Address) (*types.Transaction, *types.Receipt, error) {
	return _ContractAuthPrecompiled.Contract.CloseDeployAuth(&_ContractAuthPrecompiled.TransactOpts, account)
}

func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactorSession) AsyncCloseDeployAuth(handler func(*types.Receipt, error), account common.Address) (*types.Transaction, error) {
	return _ContractAuthPrecompiled.Contract.AsyncCloseDeployAuth(handler, &_ContractAuthPrecompiled.TransactOpts, account)
}

// CloseMethodAuth is a paid mutator transaction binding the contract method 0xcb7c5c11.
//
// Solidity: function closeMethodAuth(address contractAddr, bytes4 funcSelector, address account) returns(int256)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactor) CloseMethodAuth(opts *bind.TransactOpts, contractAddr common.Address, funcSelector [4]byte, account common.Address) (*types.Transaction, *types.Receipt, error) {
	return _ContractAuthPrecompiled.contract.Transact(opts, "closeMethodAuth", contractAddr, funcSelector, account)
}

func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactor) AsyncCloseMethodAuth(handler func(*types.Receipt, error), opts *bind.TransactOpts, contractAddr common.Address, funcSelector [4]byte, account common.Address) (*types.Transaction, error) {
	return _ContractAuthPrecompiled.contract.AsyncTransact(opts, handler, "closeMethodAuth", contractAddr, funcSelector, account)
}

// CloseMethodAuth is a paid mutator transaction binding the contract method 0xcb7c5c11.
//
// Solidity: function closeMethodAuth(address contractAddr, bytes4 funcSelector, address account) returns(int256)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledSession) CloseMethodAuth(contractAddr common.Address, funcSelector [4]byte, account common.Address) (*types.Transaction, *types.Receipt, error) {
	return _ContractAuthPrecompiled.Contract.CloseMethodAuth(&_ContractAuthPrecompiled.TransactOpts, contractAddr, funcSelector, account)
}

func (_ContractAuthPrecompiled *ContractAuthPrecompiledSession) AsyncCloseMethodAuth(handler func(*types.Receipt, error), contractAddr common.Address, funcSelector [4]byte, account common.Address) (*types.Transaction, error) {
	return _ContractAuthPrecompiled.Contract.AsyncCloseMethodAuth(handler, &_ContractAuthPrecompiled.TransactOpts, contractAddr, funcSelector, account)
}

// CloseMethodAuth is a paid mutator transaction binding the contract method 0xcb7c5c11.
//
// Solidity: function closeMethodAuth(address contractAddr, bytes4 funcSelector, address account) returns(int256)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactorSession) CloseMethodAuth(contractAddr common.Address, funcSelector [4]byte, account common.Address) (*types.Transaction, *types.Receipt, error) {
	return _ContractAuthPrecompiled.Contract.CloseMethodAuth(&_ContractAuthPrecompiled.TransactOpts, contractAddr, funcSelector, account)
}

func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactorSession) AsyncCloseMethodAuth(handler func(*types.Receipt, error), contractAddr common.Address, funcSelector [4]byte, account common.Address) (*types.Transaction, error) {
	return _ContractAuthPrecompiled.Contract.AsyncCloseMethodAuth(handler, &_ContractAuthPrecompiled.TransactOpts, contractAddr, funcSelector, account)
}

// OpenDeployAuth is a paid mutator transaction binding the contract method 0x61548099.
//
// Solidity: function openDeployAuth(address account) returns(int256)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactor) OpenDeployAuth(opts *bind.TransactOpts, account common.Address) (*types.Transaction, *types.Receipt, error) {
	return _ContractAuthPrecompiled.contract.Transact(opts, "openDeployAuth", account)
}

func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactor) AsyncOpenDeployAuth(handler func(*types.Receipt, error), opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ContractAuthPrecompiled.contract.AsyncTransact(opts, handler, "openDeployAuth", account)
}

// OpenDeployAuth is a paid mutator transaction binding the contract method 0x61548099.
//
// Solidity: function openDeployAuth(address account) returns(int256)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledSession) OpenDeployAuth(account common.Address) (*types.Transaction, *types.Receipt, error) {
	return _ContractAuthPrecompiled.Contract.OpenDeployAuth(&_ContractAuthPrecompiled.TransactOpts, account)
}

func (_ContractAuthPrecompiled *ContractAuthPrecompiledSession) AsyncOpenDeployAuth(handler func(*types.Receipt, error), account common.Address) (*types.Transaction, error) {
	return _ContractAuthPrecompiled.Contract.AsyncOpenDeployAuth(handler, &_ContractAuthPrecompiled.TransactOpts, account)
}

// OpenDeployAuth is a paid mutator transaction binding the contract method 0x61548099.
//
// Solidity: function openDeployAuth(address account) returns(int256)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactorSession) OpenDeployAuth(account common.Address) (*types.Transaction, *types.Receipt, error) {
	return _ContractAuthPrecompiled.Contract.OpenDeployAuth(&_ContractAuthPrecompiled.TransactOpts, account)
}

func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactorSession) AsyncOpenDeployAuth(handler func(*types.Receipt, error), account common.Address) (*types.Transaction, error) {
	return _ContractAuthPrecompiled.Contract.AsyncOpenDeployAuth(handler, &_ContractAuthPrecompiled.TransactOpts, account)
}

// OpenMethodAuth is a paid mutator transaction binding the contract method 0x0c82b73d.
//
// Solidity: function openMethodAuth(address contractAddr, bytes4 funcSelector, address account) returns(int256)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactor) OpenMethodAuth(opts *bind.TransactOpts, contractAddr common.Address, funcSelector [4]byte, account common.Address) (*types.Transaction, *types.Receipt, error) {
	return _ContractAuthPrecompiled.contract.Transact(opts, "openMethodAuth", contractAddr, funcSelector, account)
}

func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactor) AsyncOpenMethodAuth(handler func(*types.Receipt, error), opts *bind.TransactOpts, contractAddr common.Address, funcSelector [4]byte, account common.Address) (*types.Transaction, error) {
	return _ContractAuthPrecompiled.contract.AsyncTransact(opts, handler, "openMethodAuth", contractAddr, funcSelector, account)
}

// OpenMethodAuth is a paid mutator transaction binding the contract method 0x0c82b73d.
//
// Solidity: function openMethodAuth(address contractAddr, bytes4 funcSelector, address account) returns(int256)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledSession) OpenMethodAuth(contractAddr common.Address, funcSelector [4]byte, account common.Address) (*types.Transaction, *types.Receipt, error) {
	return _ContractAuthPrecompiled.Contract.OpenMethodAuth(&_ContractAuthPrecompiled.TransactOpts, contractAddr, funcSelector, account)
}

func (_ContractAuthPrecompiled *ContractAuthPrecompiledSession) AsyncOpenMethodAuth(handler func(*types.Receipt, error), contractAddr common.Address, funcSelector [4]byte, account common.Address) (*types.Transaction, error) {
	return _ContractAuthPrecompiled.Contract.AsyncOpenMethodAuth(handler, &_ContractAuthPrecompiled.TransactOpts, contractAddr, funcSelector, account)
}

// OpenMethodAuth is a paid mutator transaction binding the contract method 0x0c82b73d.
//
// Solidity: function openMethodAuth(address contractAddr, bytes4 funcSelector, address account) returns(int256)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactorSession) OpenMethodAuth(contractAddr common.Address, funcSelector [4]byte, account common.Address) (*types.Transaction, *types.Receipt, error) {
	return _ContractAuthPrecompiled.Contract.OpenMethodAuth(&_ContractAuthPrecompiled.TransactOpts, contractAddr, funcSelector, account)
}

func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactorSession) AsyncOpenMethodAuth(handler func(*types.Receipt, error), contractAddr common.Address, funcSelector [4]byte, account common.Address) (*types.Transaction, error) {
	return _ContractAuthPrecompiled.Contract.AsyncOpenMethodAuth(handler, &_ContractAuthPrecompiled.TransactOpts, contractAddr, funcSelector, account)
}

// ResetAdmin is a paid mutator transaction binding the contract method 0xc53057b4.
//
// Solidity: function resetAdmin(address contractAddr, address admin) returns(int256)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactor) ResetAdmin(opts *bind.TransactOpts, contractAddr common.Address, admin common.Address) (*types.Transaction, *types.Receipt, error) {
	return _ContractAuthPrecompiled.contract.Transact(opts, "resetAdmin", contractAddr, admin)
}

func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactor) AsyncResetAdmin(handler func(*types.Receipt, error), opts *bind.TransactOpts, contractAddr common.Address, admin common.Address) (*types.Transaction, error) {
	return _ContractAuthPrecompiled.contract.AsyncTransact(opts, handler, "resetAdmin", contractAddr, admin)
}

// ResetAdmin is a paid mutator transaction binding the contract method 0xc53057b4.
//
// Solidity: function resetAdmin(address contractAddr, address admin) returns(int256)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledSession) ResetAdmin(contractAddr common.Address, admin common.Address) (*types.Transaction, *types.Receipt, error) {
	return _ContractAuthPrecompiled.Contract.ResetAdmin(&_ContractAuthPrecompiled.TransactOpts, contractAddr, admin)
}

func (_ContractAuthPrecompiled *ContractAuthPrecompiledSession) AsyncResetAdmin(handler func(*types.Receipt, error), contractAddr common.Address, admin common.Address) (*types.Transaction, error) {
	return _ContractAuthPrecompiled.Contract.AsyncResetAdmin(handler, &_ContractAuthPrecompiled.TransactOpts, contractAddr, admin)
}

// ResetAdmin is a paid mutator transaction binding the contract method 0xc53057b4.
//
// Solidity: function resetAdmin(address contractAddr, address admin) returns(int256)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactorSession) ResetAdmin(contractAddr common.Address, admin common.Address) (*types.Transaction, *types.Receipt, error) {
	return _ContractAuthPrecompiled.Contract.ResetAdmin(&_ContractAuthPrecompiled.TransactOpts, contractAddr, admin)
}

func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactorSession) AsyncResetAdmin(handler func(*types.Receipt, error), contractAddr common.Address, admin common.Address) (*types.Transaction, error) {
	return _ContractAuthPrecompiled.Contract.AsyncResetAdmin(handler, &_ContractAuthPrecompiled.TransactOpts, contractAddr, admin)
}

// SetContractStatus is a paid mutator transaction binding the contract method 0x81c81cdc.
//
// Solidity: function setContractStatus(address _address, bool isFreeze) returns(int256)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactor) SetContractStatus(opts *bind.TransactOpts, _address common.Address, isFreeze bool) (*types.Transaction, *types.Receipt, error) {
	return _ContractAuthPrecompiled.contract.Transact(opts, "setContractStatus", _address, isFreeze)
}

func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactor) AsyncSetContractStatus(handler func(*types.Receipt, error), opts *bind.TransactOpts, _address common.Address, isFreeze bool) (*types.Transaction, error) {
	return _ContractAuthPrecompiled.contract.AsyncTransact(opts, handler, "setContractStatus", _address, isFreeze)
}

// SetContractStatus is a paid mutator transaction binding the contract method 0x81c81cdc.
//
// Solidity: function setContractStatus(address _address, bool isFreeze) returns(int256)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledSession) SetContractStatus(_address common.Address, isFreeze bool) (*types.Transaction, *types.Receipt, error) {
	return _ContractAuthPrecompiled.Contract.SetContractStatus(&_ContractAuthPrecompiled.TransactOpts, _address, isFreeze)
}

func (_ContractAuthPrecompiled *ContractAuthPrecompiledSession) AsyncSetContractStatus(handler func(*types.Receipt, error), _address common.Address, isFreeze bool) (*types.Transaction, error) {
	return _ContractAuthPrecompiled.Contract.AsyncSetContractStatus(handler, &_ContractAuthPrecompiled.TransactOpts, _address, isFreeze)
}

// SetContractStatus is a paid mutator transaction binding the contract method 0x81c81cdc.
//
// Solidity: function setContractStatus(address _address, bool isFreeze) returns(int256)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactorSession) SetContractStatus(_address common.Address, isFreeze bool) (*types.Transaction, *types.Receipt, error) {
	return _ContractAuthPrecompiled.Contract.SetContractStatus(&_ContractAuthPrecompiled.TransactOpts, _address, isFreeze)
}

func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactorSession) AsyncSetContractStatus(handler func(*types.Receipt, error), _address common.Address, isFreeze bool) (*types.Transaction, error) {
	return _ContractAuthPrecompiled.Contract.AsyncSetContractStatus(handler, &_ContractAuthPrecompiled.TransactOpts, _address, isFreeze)
}

// SetDeployAuthType is a paid mutator transaction binding the contract method 0xbb0aa40c.
//
// Solidity: function setDeployAuthType(uint8 _type) returns(int256)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactor) SetDeployAuthType(opts *bind.TransactOpts, _type uint8) (*types.Transaction, *types.Receipt, error) {
	return _ContractAuthPrecompiled.contract.Transact(opts, "setDeployAuthType", _type)
}

func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactor) AsyncSetDeployAuthType(handler func(*types.Receipt, error), opts *bind.TransactOpts, _type uint8) (*types.Transaction, error) {
	return _ContractAuthPrecompiled.contract.AsyncTransact(opts, handler, "setDeployAuthType", _type)
}

// SetDeployAuthType is a paid mutator transaction binding the contract method 0xbb0aa40c.
//
// Solidity: function setDeployAuthType(uint8 _type) returns(int256)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledSession) SetDeployAuthType(_type uint8) (*types.Transaction, *types.Receipt, error) {
	return _ContractAuthPrecompiled.Contract.SetDeployAuthType(&_ContractAuthPrecompiled.TransactOpts, _type)
}

func (_ContractAuthPrecompiled *ContractAuthPrecompiledSession) AsyncSetDeployAuthType(handler func(*types.Receipt, error), _type uint8) (*types.Transaction, error) {
	return _ContractAuthPrecompiled.Contract.AsyncSetDeployAuthType(handler, &_ContractAuthPrecompiled.TransactOpts, _type)
}

// SetDeployAuthType is a paid mutator transaction binding the contract method 0xbb0aa40c.
//
// Solidity: function setDeployAuthType(uint8 _type) returns(int256)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactorSession) SetDeployAuthType(_type uint8) (*types.Transaction, *types.Receipt, error) {
	return _ContractAuthPrecompiled.Contract.SetDeployAuthType(&_ContractAuthPrecompiled.TransactOpts, _type)
}

func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactorSession) AsyncSetDeployAuthType(handler func(*types.Receipt, error), _type uint8) (*types.Transaction, error) {
	return _ContractAuthPrecompiled.Contract.AsyncSetDeployAuthType(handler, &_ContractAuthPrecompiled.TransactOpts, _type)
}

// SetMethodAuthType is a paid mutator transaction binding the contract method 0x9cc3ca0f.
//
// Solidity: function setMethodAuthType(address contractAddr, bytes4 funcSelector, uint8 authType) returns(int256)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactor) SetMethodAuthType(opts *bind.TransactOpts, contractAddr common.Address, funcSelector [4]byte, authType uint8) (*types.Transaction, *types.Receipt, error) {
	return _ContractAuthPrecompiled.contract.Transact(opts, "setMethodAuthType", contractAddr, funcSelector, authType)
}

func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactor) AsyncSetMethodAuthType(handler func(*types.Receipt, error), opts *bind.TransactOpts, contractAddr common.Address, funcSelector [4]byte, authType uint8) (*types.Transaction, error) {
	return _ContractAuthPrecompiled.contract.AsyncTransact(opts, handler, "setMethodAuthType", contractAddr, funcSelector, authType)
}

// SetMethodAuthType is a paid mutator transaction binding the contract method 0x9cc3ca0f.
//
// Solidity: function setMethodAuthType(address contractAddr, bytes4 funcSelector, uint8 authType) returns(int256)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledSession) SetMethodAuthType(contractAddr common.Address, funcSelector [4]byte, authType uint8) (*types.Transaction, *types.Receipt, error) {
	return _ContractAuthPrecompiled.Contract.SetMethodAuthType(&_ContractAuthPrecompiled.TransactOpts, contractAddr, funcSelector, authType)
}

func (_ContractAuthPrecompiled *ContractAuthPrecompiledSession) AsyncSetMethodAuthType(handler func(*types.Receipt, error), contractAddr common.Address, funcSelector [4]byte, authType uint8) (*types.Transaction, error) {
	return _ContractAuthPrecompiled.Contract.AsyncSetMethodAuthType(handler, &_ContractAuthPrecompiled.TransactOpts, contractAddr, funcSelector, authType)
}

// SetMethodAuthType is a paid mutator transaction binding the contract method 0x9cc3ca0f.
//
// Solidity: function setMethodAuthType(address contractAddr, bytes4 funcSelector, uint8 authType) returns(int256)
func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactorSession) SetMethodAuthType(contractAddr common.Address, funcSelector [4]byte, authType uint8) (*types.Transaction, *types.Receipt, error) {
	return _ContractAuthPrecompiled.Contract.SetMethodAuthType(&_ContractAuthPrecompiled.TransactOpts, contractAddr, funcSelector, authType)
}

func (_ContractAuthPrecompiled *ContractAuthPrecompiledTransactorSession) AsyncSetMethodAuthType(handler func(*types.Receipt, error), contractAddr common.Address, funcSelector [4]byte, authType uint8) (*types.Transaction, error) {
	return _ContractAuthPrecompiled.Contract.AsyncSetMethodAuthType(handler, &_ContractAuthPrecompiled.TransactOpts, contractAddr, funcSelector, authType)
}
