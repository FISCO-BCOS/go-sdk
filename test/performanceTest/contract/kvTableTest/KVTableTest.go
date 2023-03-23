// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kvTableTest

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

// KVTableTestABI is the input ABI used to generate the binding from.
const KVTableTestABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"count\",\"type\":\"int256\"}],\"name\":\"InsertResult\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"age\",\"type\":\"string\"}],\"name\":\"insert\",\"outputs\":[{\"internalType\":\"int32\",\"name\":\"\",\"type\":\"int32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"select\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// KVTableTestBin is the compiled bytecode used for deploying new contracts.
var KVTableTestBin = "0x60806040523480156200001157600080fd5b506060600267ffffffffffffffff811180156200002d57600080fd5b506040519080825280602002602001820160405280156200006357816020015b60608152602001906001900390816200004d5790505b5090506040518060400160405280600481526020017f6e616d650000000000000000000000000000000000000000000000000000000081525081600081518110620000aa57fe5b60200260200101819052506040518060400160405280600381526020017f616765000000000000000000000000000000000000000000000000000000000081525081600181518110620000f957fe5b60200260200101819052506200010e620003ab565b60405180604001604052806040518060400160405280600281526020017f6964000000000000000000000000000000000000000000000000000000000000815250815260200183815250905061100273ffffffffffffffffffffffffffffffffffffffff166331a5a51e6040518060400160405280600681526020017f745f746573740000000000000000000000000000000000000000000000000000815250836040518363ffffffff1660e01b8152600401620001ce929190620005ec565b602060405180830381600087803b158015620001e957600080fd5b505af1158015620001fe573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200022491906200041f565b50600061100273ffffffffffffffffffffffffffffffffffffffff1663f23f63c96040518060400160405280600681526020017f745f7465737400000000000000000000000000000000000000000000000000008152506040518263ffffffff1660e01b8152600401620002999190620005c8565b60206040518083038186803b158015620002b257600080fd5b505afa158015620002c7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620002ed9190620003f3565b9050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141562000362576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620003599062000627565b60405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050506200076b565b604051806040016040528060608152602001606081525090565b600081519050620003d68162000737565b92915050565b600081519050620003ed8162000751565b92915050565b6000602082840312156200040657600080fd5b60006200041684828501620003c5565b91505092915050565b6000602082840312156200043257600080fd5b60006200044284828501620003dc565b91505092915050565b6000620004598383620004e2565b905092915050565b60006200046e8262000659565b6200047a81856200067c565b9350836020820285016200048e8562000649565b8060005b85811015620004d05784840389528151620004ae85826200044b565b9450620004bb836200066f565b925060208a0199505060018101905062000492565b50829750879550505050505092915050565b6000620004ef8262000664565b620004fb81856200068d565b93506200050d818560208601620006f0565b620005188162000726565b840191505092915050565b6000620005308262000664565b6200053c81856200069e565b93506200054e818560208601620006f0565b620005598162000726565b840191505092915050565b6000620005736000836200069e565b9150600082019050919050565b600060408301600083015184820360008601526200059f8282620004e2565b91505060208301518482036020860152620005bb828262000461565b9150508091505092915050565b60006020820190508181036000830152620005e4818462000523565b905092915050565b6000604082019050818103600083015262000608818562000523565b905081810360208301526200061e818462000580565b90509392505050565b60006020820190508181036000830152620006428162000564565b9050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b6000620006bc82620006d0565b9050919050565b60008160030b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60005b8381101562000710578082015181840152602081019050620006f3565b8381111562000720576000848401525b50505050565b6000601f19601f8301169050919050565b6200074281620006af565b81146200074e57600080fd5b50565b6200075c81620006c3565b81146200076857600080fd5b50565b610a03806200077b6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632fe99bdc1461003b578063fcd7e3c11461006b575b600080fd5b61005560048036038101906100509190610577565b61009c565b60405161006291906107c7565b60405180910390f35b61008560048036038101906100809190610536565b610234565b604051610093929190610804565b60405180910390f35b60006060600267ffffffffffffffff811180156100b857600080fd5b506040519080825280602002602001820160405280156100ec57816020015b60608152602001906001900390816100d75790505b50905083816000815181106100fd57fe5b6020026020010181905250828160018151811061011657fe5b6020026020010181905250610129610346565b604051806040016040528087815260200183815250905060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635c6e105f836040518263ffffffff1660e01b815260040161019c919061083b565b602060405180830381600087803b1580156101b657600080fd5b505af11580156101ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101ee919061050d565b90507fc57b01fa77f41df77eaab79a0e2623fab2e7ae3e9530d9b1cab225ad65f2b7ce8160405161021f91906107ac565b60405180910390a18093505050509392505050565b60608061023f610346565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fcd7e3c1856040518263ffffffff1660e01b815260040161029991906107e2565b60006040518083038186803b1580156102b157600080fd5b505afa1580156102c5573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906102ee919061060e565b90506060806002836020015151141561033857826020015160008151811061031257fe5b60200260200101519150826020015160018151811061032d57fe5b602002602001015190505b818194509450505050915091565b604051806040016040528060608152602001606081525090565b600082601f83011261037157600080fd5b815161038461037f8261088a565b61085d565b9150818183526020840193506020810190508360005b838110156103ca57815186016103b0888261043d565b84526020840193506020830192505060018101905061039a565b5050505092915050565b6000815190506103e3816109b6565b92915050565b600082601f8301126103fa57600080fd5b813561040d610408826108b2565b61085d565b9150808252602083016020830185838301111561042957600080fd5b610434838284610963565b50505092915050565b600082601f83011261044e57600080fd5b815161046161045c826108b2565b61085d565b9150808252602083016020830185838301111561047d57600080fd5b610488838284610972565b50505092915050565b6000604082840312156104a357600080fd5b6104ad604061085d565b9050600082015167ffffffffffffffff8111156104c957600080fd5b6104d58482850161043d565b600083015250602082015167ffffffffffffffff8111156104f557600080fd5b61050184828501610360565b60208301525092915050565b60006020828403121561051f57600080fd5b600061052d848285016103d4565b91505092915050565b60006020828403121561054857600080fd5b600082013567ffffffffffffffff81111561056257600080fd5b61056e848285016103e9565b91505092915050565b60008060006060848603121561058c57600080fd5b600084013567ffffffffffffffff8111156105a657600080fd5b6105b2868287016103e9565b935050602084013567ffffffffffffffff8111156105cf57600080fd5b6105db868287016103e9565b925050604084013567ffffffffffffffff8111156105f857600080fd5b610604868287016103e9565b9150509250925092565b60006020828403121561062057600080fd5b600082015167ffffffffffffffff81111561063a57600080fd5b61064684828501610491565b91505092915050565b600061065b83836106f6565b905092915050565b600061066e826108ee565b6106788185610911565b93508360208202850161068a856108de565b8060005b858110156106c657848403895281516106a7858261064f565b94506106b283610904565b925060208a0199505060018101905061068e565b50829750879550505050505092915050565b6106e181610951565b82525050565b6106f081610944565b82525050565b6000610701826108f9565b61070b8185610922565b935061071b818560208601610972565b610724816109a5565b840191505092915050565b600061073a826108f9565b6107448185610933565b9350610754818560208601610972565b61075d816109a5565b840191505092915050565b6000604083016000830151848203600086015261078582826106f6565b9150506020830151848203602086015261079f8282610663565b9150508091505092915050565b60006020820190506107c160008301846106d8565b92915050565b60006020820190506107dc60008301846106e7565b92915050565b600060208201905081810360008301526107fc818461072f565b905092915050565b6000604082019050818103600083015261081e818561072f565b90508181036020830152610832818461072f565b90509392505050565b600060208201905081810360008301526108558184610768565b905092915050565b6000604051905081810181811067ffffffffffffffff8211171561088057600080fd5b8060405250919050565b600067ffffffffffffffff8211156108a157600080fd5b602082029050602081019050919050565b600067ffffffffffffffff8211156108c957600080fd5b601f19601f8301169050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b60008160030b9050919050565b600061095c82610944565b9050919050565b82818337600083830152505050565b60005b83811015610990578082015181840152602081019050610975565b8381111561099f576000848401525b50505050565b6000601f19601f8301169050919050565b6109bf81610944565b81146109ca57600080fd5b5056fea264697066735822122096b6c562707a1d7d796e386e2ecf842b857037477455dbd940ba564db2eabd9864736f6c634300060a0033"

// DeployKVTableTest deploys a new contract, binding an instance of KVTableTest to it.
func DeployKVTableTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Receipt, *KVTableTest, error) {
	parsed, err := abi.JSON(strings.NewReader(KVTableTestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, receipt, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KVTableTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &KVTableTest{KVTableTestCaller: KVTableTestCaller{contract: contract}, KVTableTestTransactor: KVTableTestTransactor{contract: contract}, KVTableTestFilterer: KVTableTestFilterer{contract: contract}}, nil
}

func AsyncDeployKVTableTest(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(KVTableTestABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(KVTableTestBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// KVTableTest is an auto generated Go binding around a Solidity contract.
type KVTableTest struct {
	KVTableTestCaller     // Read-only binding to the contract
	KVTableTestTransactor // Write-only binding to the contract
	KVTableTestFilterer   // Log filterer for contract events
}

// KVTableTestCaller is an auto generated read-only Go binding around a Solidity contract.
type KVTableTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KVTableTestTransactor is an auto generated write-only Go binding around a Solidity contract.
type KVTableTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KVTableTestFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type KVTableTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KVTableTestSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type KVTableTestSession struct {
	Contract     *KVTableTest      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KVTableTestCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type KVTableTestCallerSession struct {
	Contract *KVTableTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// KVTableTestTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type KVTableTestTransactorSession struct {
	Contract     *KVTableTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// KVTableTestRaw is an auto generated low-level Go binding around a Solidity contract.
type KVTableTestRaw struct {
	Contract *KVTableTest // Generic contract binding to access the raw methods on
}

// KVTableTestCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type KVTableTestCallerRaw struct {
	Contract *KVTableTestCaller // Generic read-only contract binding to access the raw methods on
}

// KVTableTestTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type KVTableTestTransactorRaw struct {
	Contract *KVTableTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKVTableTest creates a new instance of KVTableTest, bound to a specific deployed contract.
func NewKVTableTest(address common.Address, backend bind.ContractBackend) (*KVTableTest, error) {
	contract, err := bindKVTableTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KVTableTest{KVTableTestCaller: KVTableTestCaller{contract: contract}, KVTableTestTransactor: KVTableTestTransactor{contract: contract}, KVTableTestFilterer: KVTableTestFilterer{contract: contract}}, nil
}

// NewKVTableTestCaller creates a new read-only instance of KVTableTest, bound to a specific deployed contract.
func NewKVTableTestCaller(address common.Address, caller bind.ContractCaller) (*KVTableTestCaller, error) {
	contract, err := bindKVTableTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KVTableTestCaller{contract: contract}, nil
}

// NewKVTableTestTransactor creates a new write-only instance of KVTableTest, bound to a specific deployed contract.
func NewKVTableTestTransactor(address common.Address, transactor bind.ContractTransactor) (*KVTableTestTransactor, error) {
	contract, err := bindKVTableTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KVTableTestTransactor{contract: contract}, nil
}

// NewKVTableTestFilterer creates a new log filterer instance of KVTableTest, bound to a specific deployed contract.
func NewKVTableTestFilterer(address common.Address, filterer bind.ContractFilterer) (*KVTableTestFilterer, error) {
	contract, err := bindKVTableTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KVTableTestFilterer{contract: contract}, nil
}

// bindKVTableTest binds a generic wrapper to an already deployed contract.
func bindKVTableTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KVTableTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KVTableTest *KVTableTestRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KVTableTest.Contract.KVTableTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KVTableTest *KVTableTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _KVTableTest.Contract.KVTableTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KVTableTest *KVTableTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _KVTableTest.Contract.KVTableTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KVTableTest *KVTableTestCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KVTableTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KVTableTest *KVTableTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _KVTableTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KVTableTest *KVTableTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _KVTableTest.Contract.contract.Transact(opts, method, params...)
}

// Select is a free data retrieval call binding the contract method 0xfcd7e3c1.
//
// Solidity: function select(string id) constant returns(string, string)
func (_KVTableTest *KVTableTestCaller) Select(opts *bind.CallOpts, id string) (string, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _KVTableTest.contract.Call(opts, out, "select", id)
	return *ret0, *ret1, err
}

// Select is a free data retrieval call binding the contract method 0xfcd7e3c1.
//
// Solidity: function select(string id) constant returns(string, string)
func (_KVTableTest *KVTableTestSession) Select(id string) (string, string, error) {
	return _KVTableTest.Contract.Select(&_KVTableTest.CallOpts, id)
}

// Select is a free data retrieval call binding the contract method 0xfcd7e3c1.
//
// Solidity: function select(string id) constant returns(string, string)
func (_KVTableTest *KVTableTestCallerSession) Select(id string) (string, string, error) {
	return _KVTableTest.Contract.Select(&_KVTableTest.CallOpts, id)
}

// Insert is a paid mutator transaction binding the contract method 0x2fe99bdc.
//
// Solidity: function insert(string id, string name, string age) returns(int32)
func (_KVTableTest *KVTableTestTransactor) Insert(opts *bind.TransactOpts, id string, name string, age string) (*types.Transaction, *types.Receipt, error) {
	return _KVTableTest.contract.Transact(opts, "insert", id, name, age)
}

func (_KVTableTest *KVTableTestTransactor) AsyncInsert(handler func(*types.Receipt, error), opts *bind.TransactOpts, id string, name string, age string) (*types.Transaction, error) {
	return _KVTableTest.contract.AsyncTransact(opts, handler, "insert", id, name, age)
}

// Insert is a paid mutator transaction binding the contract method 0x2fe99bdc.
//
// Solidity: function insert(string id, string name, string age) returns(int32)
func (_KVTableTest *KVTableTestSession) Insert(id string, name string, age string) (*types.Transaction, *types.Receipt, error) {
	return _KVTableTest.Contract.Insert(&_KVTableTest.TransactOpts, id, name, age)
}

func (_KVTableTest *KVTableTestSession) AsyncInsert(handler func(*types.Receipt, error), id string, name string, age string) (*types.Transaction, error) {
	return _KVTableTest.Contract.AsyncInsert(handler, &_KVTableTest.TransactOpts, id, name, age)
}

// Insert is a paid mutator transaction binding the contract method 0x2fe99bdc.
//
// Solidity: function insert(string id, string name, string age) returns(int32)
func (_KVTableTest *KVTableTestTransactorSession) Insert(id string, name string, age string) (*types.Transaction, *types.Receipt, error) {
	return _KVTableTest.Contract.Insert(&_KVTableTest.TransactOpts, id, name, age)
}

func (_KVTableTest *KVTableTestTransactorSession) AsyncInsert(handler func(*types.Receipt, error), id string, name string, age string) (*types.Transaction, error) {
	return _KVTableTest.Contract.AsyncInsert(handler, &_KVTableTest.TransactOpts, id, name, age)
}

// KVTableTestInsertResult represents a InsertResult event raised by the KVTableTest contract.
type KVTableTestInsertResult struct {
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// WatchInsertResult is a free log subscription operation binding the contract event 0xc57b01fa77f41df77eaab79a0e2623fab2e7ae3e9530d9b1cab225ad65f2b7ce.
//
// Solidity: event InsertResult(int256 count)
func (_KVTableTest *KVTableTestFilterer) WatchInsertResult(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _KVTableTest.contract.WatchLogs(fromBlock, handler, "InsertResult")
}

func (_KVTableTest *KVTableTestFilterer) WatchAllInsertResult(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _KVTableTest.contract.WatchLogs(fromBlock, handler, "InsertResult")
}

// ParseInsertResult is a log parse operation binding the contract event 0xc57b01fa77f41df77eaab79a0e2623fab2e7ae3e9530d9b1cab225ad65f2b7ce.
//
// Solidity: event InsertResult(int256 count)
func (_KVTableTest *KVTableTestFilterer) ParseInsertResult(log types.Log) (*KVTableTestInsertResult, error) {
	event := new(KVTableTestInsertResult)
	if err := _KVTableTest.contract.UnpackLog(event, "InsertResult", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchInsertResult is a free log subscription operation binding the contract event 0xc57b01fa77f41df77eaab79a0e2623fab2e7ae3e9530d9b1cab225ad65f2b7ce.
//
// Solidity: event InsertResult(int256 count)
func (_KVTableTest *KVTableTestSession) WatchInsertResult(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _KVTableTest.Contract.WatchInsertResult(fromBlock, handler)
}

func (_KVTableTest *KVTableTestSession) WatchAllInsertResult(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _KVTableTest.Contract.WatchAllInsertResult(fromBlock, handler)
}

// ParseInsertResult is a log parse operation binding the contract event 0xc57b01fa77f41df77eaab79a0e2623fab2e7ae3e9530d9b1cab225ad65f2b7ce.
//
// Solidity: event InsertResult(int256 count)
func (_KVTableTest *KVTableTestSession) ParseInsertResult(log types.Log) (*KVTableTestInsertResult, error) {
	return _KVTableTest.Contract.ParseInsertResult(log)
}
