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
	_ = types.BloomLookup
)

// CommitteeABI is the input ABI used to generate the binding from.
const CommitteeABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"governorList\",\"type\":\"address[]\"},{\"internalType\":\"uint32[]\",\"name\":\"weightList\",\"type\":\"uint32[]\"},{\"internalType\":\"uint8\",\"name\":\"participatesRate\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"winRate\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_participatesRate\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_winRate\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"}],\"name\":\"auth\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommitteeInfo\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"participatesRate\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"winRate\",\"type\":\"uint8\"},{\"internalType\":\"address[]\",\"name\":\"governors\",\"type\":\"address[]\"},{\"internalType\":\"uint32[]\",\"name\":\"weights\",\"type\":\"uint32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"}],\"name\":\"getWeight\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWeights\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"votes\",\"type\":\"address[]\"}],\"name\":\"getWeights\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"}],\"name\":\"isGovernor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"participatesRate\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"winRate\",\"type\":\"uint8\"}],\"name\":\"setRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"weight\",\"type\":\"uint32\"}],\"name\":\"setWeight\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CommitteeBin is the compiled bytecode used for deploying new contracts.
var CommitteeBin = "0x60806040523480156200001157600080fd5b5060405162001d4838038062001d48833981810160405260808110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b838201915060208201858111156200006f57600080fd5b82518660208202830111640100000000821117156200008d57600080fd5b8083526020830192505050908051906020019060200280838360005b83811015620000c6578082015181840152602081019050620000a9565b5050505090500160405260200180516040519392919084640100000000821115620000f057600080fd5b838201915060208201858111156200010757600080fd5b82518660208202830111640100000000821117156200012557600080fd5b8083526020830192505050908051906020019060200280838360005b838110156200015e57808201518184015260208101905062000141565b505050509050016040526020018051906020019092919080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008090505b84518163ffffffff1610156200022c576200021e858263ffffffff1681518110620001ef57fe5b6020026020010151858363ffffffff16815181106200020a57fe5b60200260200101516200027e60201b60201c565b8080600101915050620001c8565b5080600460016101000a81548160ff021916908360ff16021790555081600460006101000a81548160ff021916908360ff16021790555062000274336200052c60201b60201c565b5050505062000ac6565b6200028f33620005f360201b60201c565b62000302576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f4f6e6c79206f776e65722100000000000000000000000000000000000000000081525060200191505060405180910390fd5b60008163ffffffff16141562000428573273ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415620003b5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f596f752063616e206e6f742072656d6f766520796f757273656c66210000000081525060200191505060405180910390fd5b600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549063ffffffff0219169055620004228260016200069e60201b62000c4b1790919060201c565b62000528565b62000443826001620008cf60201b62000e6f1790919060201c565b15620004ad5780600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908363ffffffff16021790555062000527565b80600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908363ffffffff160217905550620005268260016200091e60201b62000ebe1790919060201c565b5b5b5050565b6200053d33620005f360201b60201c565b620005b0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f4f6e6c79206f776e65722100000000000000000000000000000000000000000081525060200191505060405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60003073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141562000634576001905062000699565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141562000694576001905062000699565b600090505b919050565b620006b08282620008cf60201b60201c565b62000707576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602381526020018062001d046023913960400191505060405180910390fd5b600060018360000160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205403905060006001846001018054905003905060008460010182815481106200077157fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905080856001018481548110620007af57fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600183018560000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508460000160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009055846001018054806200089357fe5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905590555050505050565b6000808360000160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541415905092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415620009a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602181526020018062001d276021913960400191505060405180910390fd5b620009b88282620008cf60201b60201c565b1562000a10576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f81526020018062001cd5602f913960400191505060405180910390fd5b81600101819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600101805490508260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050565b6111ff8062000ad66000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063ac6c525111610071578063ac6c525114610322578063b2bdfa7b14610386578063b6fd9067146103d0578063cd5d2118146103f4578063e43581b814610450578063f437695a146104ac576100b4565b806313af4035146100b957806322acb867146100fd5780635615696f146101275780635e77fe201461014b578063965b9ff11461020c57806399bc9c1b146102e4575b600080fd5b6100fb600480360360208110156100cf57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610500565b005b6101056105be565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b61012f6105d7565b604051808260ff1660ff16815260200191505060405180910390f35b6101536105ea565b604051808560ff1660ff1681526020018460ff1660ff1681526020018060200180602001838103835285818151815260200191508051906020019060200280838360005b838110156101b2578082015181840152602081019050610197565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156101f45780820151818401526020810190506101d9565b50505050905001965050505050505060405180910390f35b6102c26004803603602081101561022257600080fd5b810190808035906020019064010000000081111561023f57600080fd5b82018360208201111561025157600080fd5b8035906020019184602083028401116401000000008311171561027357600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050509192919290505050610716565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b610320600480360360408110156102fa57600080fd5b81019080803560ff169060200190929190803560ff1690602001909291905050506107b6565b005b6103646004803603602081101561033857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061086b565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b61038e6108c4565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6103d86108e9565b604051808260ff1660ff16815260200191505060405180910390f35b6104366004803603602081101561040a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506108fc565b604051808215151515815260200191505060405180910390f35b6104926004803603602081101561046657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506109a3565b604051808215151515815260200191505060405180910390f35b6104fe600480360360408110156104c257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803563ffffffff1690602001909291905050506109c0565b005b610509336108fc565b61057b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f4f6e6c79206f776e65722100000000000000000000000000000000000000000081525060200191505060405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60006105d26105cd600161105a565b610716565b905090565b600460009054906101000a900460ff1681565b6000806060806105fa600161105a565b9150815167ffffffffffffffff8111801561061457600080fd5b506040519080825280602002602001820160405280156106435781602001602082028036833780820191505090505b50905060008090505b82518110156106eb576003600084838151811061066557fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900463ffffffff168282815181106106c457fe5b602002602001019063ffffffff16908163ffffffff1681525050808060010191505061064c565b50600460019054906101000a900460ff169250600460009054906101000a900460ff16935090919293565b6000806000905060008090505b83518163ffffffff1610156107ac5760036000858363ffffffff168151811061074857fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900463ffffffff16820191508080600101915050610723565b5080915050919050565b6107bf336108fc565b610831576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f4f6e6c79206f776e65722100000000000000000000000000000000000000000081525060200191505060405180910390fd5b80600460016101000a81548160ff021916908360ff16021790555081600460006101000a81548160ff021916908360ff1602179055505050565b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900463ffffffff169050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600460019054906101000a900460ff1681565b60003073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561093b576001905061099e565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610999576001905061099e565b600090505b919050565b60006109b9826001610e6f90919063ffffffff16565b9050919050565b6109c9336108fc565b610a3b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600b8152602001807f4f6e6c79206f776e65722100000000000000000000000000000000000000000081525060200191505060405180910390fd5b60008163ffffffff161415610b57573273ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610aec576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f596f752063616e206e6f742072656d6f766520796f757273656c66210000000081525060200191505060405180910390fd5b600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549063ffffffff0219169055610b52826001610c4b90919063ffffffff16565b610c47565b610b6b826001610e6f90919063ffffffff16565b15610bd35780600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908363ffffffff160217905550610c46565b80600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908363ffffffff160217905550610c45826001610ebe90919063ffffffff16565b5b5b5050565b610c558282610e6f565b610caa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260238152602001806111866023913960400191505060405180910390fd5b600060018360000160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020540390506000600184600101805490500390506000846001018281548110610d1357fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905080856001018481548110610d5057fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600183018560000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508460000160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000905584600101805480610e3357fe5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905590555050505050565b6000808360000160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541415905092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610f44576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001806111a96021913960400191505060405180910390fd5b610f4e8282610e6f565b15610fa4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f815260200180611157602f913960400191505060405180910390fd5b81600101819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600101805490508260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050565b606080826001018054905067ffffffffffffffff8111801561107b57600080fd5b506040519080825280602002602001820160405280156110aa5781602001602082028036833780820191505090505b50905060005b836001018054905081101561114c578360010181815481106110ce57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682828151811061110557fe5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505080806001019150506110b0565b508091505091905056fe4c6962416464726573735365743a2076616c756520616c72656164792065786973747320696e20746865207365742e4c6962416464726573735365743a2076616c756520646f65736e27742065786973742e4c6962416464726573735365743a2076616c75652063616e277420626520307830a2646970667358221220abe2d0afc892bbdc3006638b96e21fd616f4ab51cf938c3abbc90323ddbb757164736f6c634300060a00334c6962416464726573735365743a2076616c756520616c72656164792065786973747320696e20746865207365742e4c6962416464726573735365743a2076616c756520646f65736e27742065786973742e4c6962416464726573735365743a2076616c75652063616e277420626520307830"

// DeployCommittee deploys a new contract, binding an instance of Committee to it.
func DeployCommittee(auth *bind.TransactOpts, backend bind.ContractBackend, governorList []common.Address, weightList []uint32, participatesRate uint8, winRate uint8) (common.Address, *types.Receipt, *Committee, error) {
	parsed, err := abi.JSON(strings.NewReader(CommitteeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, receipt, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CommitteeBin), backend, governorList, weightList, participatesRate, winRate)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &Committee{CommitteeCaller: CommitteeCaller{contract: contract}, CommitteeTransactor: CommitteeTransactor{contract: contract}, CommitteeFilterer: CommitteeFilterer{contract: contract}}, nil
}

func AsyncDeployCommittee(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend, governorList []common.Address, weightList []uint32, participatesRate uint8, winRate uint8) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(CommitteeABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(CommitteeBin), backend, governorList, weightList, participatesRate, winRate)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Committee is an auto generated Go binding around a Solidity contract.
type Committee struct {
	CommitteeCaller     // Read-only binding to the contract
	CommitteeTransactor // Write-only binding to the contract
	CommitteeFilterer   // Log filterer for contract events
}

// CommitteeCaller is an auto generated read-only Go binding around a Solidity contract.
type CommitteeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitteeTransactor is an auto generated write-only Go binding around a Solidity contract.
type CommitteeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitteeFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type CommitteeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitteeSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type CommitteeSession struct {
	Contract     *Committee        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CommitteeCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type CommitteeCallerSession struct {
	Contract *CommitteeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CommitteeTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type CommitteeTransactorSession struct {
	Contract     *CommitteeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CommitteeRaw is an auto generated low-level Go binding around a Solidity contract.
type CommitteeRaw struct {
	Contract *Committee // Generic contract binding to access the raw methods on
}

// CommitteeCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type CommitteeCallerRaw struct {
	Contract *CommitteeCaller // Generic read-only contract binding to access the raw methods on
}

// CommitteeTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type CommitteeTransactorRaw struct {
	Contract *CommitteeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCommittee creates a new instance of Committee, bound to a specific deployed contract.
func NewCommittee(address common.Address, backend bind.ContractBackend) (*Committee, error) {
	contract, err := bindCommittee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Committee{CommitteeCaller: CommitteeCaller{contract: contract}, CommitteeTransactor: CommitteeTransactor{contract: contract}, CommitteeFilterer: CommitteeFilterer{contract: contract}}, nil
}

// NewCommitteeCaller creates a new read-only instance of Committee, bound to a specific deployed contract.
func NewCommitteeCaller(address common.Address, caller bind.ContractCaller) (*CommitteeCaller, error) {
	contract, err := bindCommittee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CommitteeCaller{contract: contract}, nil
}

// NewCommitteeTransactor creates a new write-only instance of Committee, bound to a specific deployed contract.
func NewCommitteeTransactor(address common.Address, transactor bind.ContractTransactor) (*CommitteeTransactor, error) {
	contract, err := bindCommittee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CommitteeTransactor{contract: contract}, nil
}

// NewCommitteeFilterer creates a new log filterer instance of Committee, bound to a specific deployed contract.
func NewCommitteeFilterer(address common.Address, filterer bind.ContractFilterer) (*CommitteeFilterer, error) {
	contract, err := bindCommittee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CommitteeFilterer{contract: contract}, nil
}

// bindCommittee binds a generic wrapper to an already deployed contract.
func bindCommittee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CommitteeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Committee *CommitteeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Committee.Contract.CommitteeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Committee *CommitteeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Committee.Contract.CommitteeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Committee *CommitteeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Committee.Contract.CommitteeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Committee *CommitteeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Committee.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Committee *CommitteeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Committee.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Committee *CommitteeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Committee.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() constant returns(address)
func (_Committee *CommitteeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Committee.contract.Call(opts, out, "_owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() constant returns(address)
func (_Committee *CommitteeSession) Owner() (common.Address, error) {
	return _Committee.Contract.Owner(&_Committee.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() constant returns(address)
func (_Committee *CommitteeCallerSession) Owner() (common.Address, error) {
	return _Committee.Contract.Owner(&_Committee.CallOpts)
}

// ParticipatesRate is a free data retrieval call binding the contract method 0x5615696f.
//
// Solidity: function _participatesRate() constant returns(uint8)
func (_Committee *CommitteeCaller) ParticipatesRate(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Committee.contract.Call(opts, out, "_participatesRate")
	return *ret0, err
}

// ParticipatesRate is a free data retrieval call binding the contract method 0x5615696f.
//
// Solidity: function _participatesRate() constant returns(uint8)
func (_Committee *CommitteeSession) ParticipatesRate() (uint8, error) {
	return _Committee.Contract.ParticipatesRate(&_Committee.CallOpts)
}

// ParticipatesRate is a free data retrieval call binding the contract method 0x5615696f.
//
// Solidity: function _participatesRate() constant returns(uint8)
func (_Committee *CommitteeCallerSession) ParticipatesRate() (uint8, error) {
	return _Committee.Contract.ParticipatesRate(&_Committee.CallOpts)
}

// WinRate is a free data retrieval call binding the contract method 0xb6fd9067.
//
// Solidity: function _winRate() constant returns(uint8)
func (_Committee *CommitteeCaller) WinRate(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Committee.contract.Call(opts, out, "_winRate")
	return *ret0, err
}

// WinRate is a free data retrieval call binding the contract method 0xb6fd9067.
//
// Solidity: function _winRate() constant returns(uint8)
func (_Committee *CommitteeSession) WinRate() (uint8, error) {
	return _Committee.Contract.WinRate(&_Committee.CallOpts)
}

// WinRate is a free data retrieval call binding the contract method 0xb6fd9067.
//
// Solidity: function _winRate() constant returns(uint8)
func (_Committee *CommitteeCallerSession) WinRate() (uint8, error) {
	return _Committee.Contract.WinRate(&_Committee.CallOpts)
}

// Auth is a free data retrieval call binding the contract method 0xcd5d2118.
//
// Solidity: function auth(address src) constant returns(bool)
func (_Committee *CommitteeCaller) Auth(opts *bind.CallOpts, src common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Committee.contract.Call(opts, out, "auth", src)
	return *ret0, err
}

// Auth is a free data retrieval call binding the contract method 0xcd5d2118.
//
// Solidity: function auth(address src) constant returns(bool)
func (_Committee *CommitteeSession) Auth(src common.Address) (bool, error) {
	return _Committee.Contract.Auth(&_Committee.CallOpts, src)
}

// Auth is a free data retrieval call binding the contract method 0xcd5d2118.
//
// Solidity: function auth(address src) constant returns(bool)
func (_Committee *CommitteeCallerSession) Auth(src common.Address) (bool, error) {
	return _Committee.Contract.Auth(&_Committee.CallOpts, src)
}

// GetCommitteeInfo is a free data retrieval call binding the contract method 0x5e77fe20.
//
// Solidity: function getCommitteeInfo() constant returns(uint8 participatesRate, uint8 winRate, address[] governors, uint32[] weights)
func (_Committee *CommitteeCaller) GetCommitteeInfo(opts *bind.CallOpts) (struct {
	ParticipatesRate uint8
	WinRate          uint8
	Governors        []common.Address
	Weights          []uint32
}, error) {
	ret := new(struct {
		ParticipatesRate uint8
		WinRate          uint8
		Governors        []common.Address
		Weights          []uint32
	})
	out := ret
	err := _Committee.contract.Call(opts, out, "getCommitteeInfo")
	return *ret, err
}

// GetCommitteeInfo is a free data retrieval call binding the contract method 0x5e77fe20.
//
// Solidity: function getCommitteeInfo() constant returns(uint8 participatesRate, uint8 winRate, address[] governors, uint32[] weights)
func (_Committee *CommitteeSession) GetCommitteeInfo() (struct {
	ParticipatesRate uint8
	WinRate          uint8
	Governors        []common.Address
	Weights          []uint32
}, error) {
	return _Committee.Contract.GetCommitteeInfo(&_Committee.CallOpts)
}

// GetCommitteeInfo is a free data retrieval call binding the contract method 0x5e77fe20.
//
// Solidity: function getCommitteeInfo() constant returns(uint8 participatesRate, uint8 winRate, address[] governors, uint32[] weights)
func (_Committee *CommitteeCallerSession) GetCommitteeInfo() (struct {
	ParticipatesRate uint8
	WinRate          uint8
	Governors        []common.Address
	Weights          []uint32
}, error) {
	return _Committee.Contract.GetCommitteeInfo(&_Committee.CallOpts)
}

// GetWeight is a free data retrieval call binding the contract method 0xac6c5251.
//
// Solidity: function getWeight(address governor) constant returns(uint32)
func (_Committee *CommitteeCaller) GetWeight(opts *bind.CallOpts, governor common.Address) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Committee.contract.Call(opts, out, "getWeight", governor)
	return *ret0, err
}

// GetWeight is a free data retrieval call binding the contract method 0xac6c5251.
//
// Solidity: function getWeight(address governor) constant returns(uint32)
func (_Committee *CommitteeSession) GetWeight(governor common.Address) (uint32, error) {
	return _Committee.Contract.GetWeight(&_Committee.CallOpts, governor)
}

// GetWeight is a free data retrieval call binding the contract method 0xac6c5251.
//
// Solidity: function getWeight(address governor) constant returns(uint32)
func (_Committee *CommitteeCallerSession) GetWeight(governor common.Address) (uint32, error) {
	return _Committee.Contract.GetWeight(&_Committee.CallOpts, governor)
}

// GetWeights is a free data retrieval call binding the contract method 0x22acb867.
//
// Solidity: function getWeights() constant returns(uint32)
func (_Committee *CommitteeCaller) GetWeights(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Committee.contract.Call(opts, out, "getWeights")
	return *ret0, err
}

// GetWeights is a free data retrieval call binding the contract method 0x22acb867.
//
// Solidity: function getWeights() constant returns(uint32)
func (_Committee *CommitteeSession) GetWeights() (uint32, error) {
	return _Committee.Contract.GetWeights(&_Committee.CallOpts)
}

// GetWeights is a free data retrieval call binding the contract method 0x22acb867.
//
// Solidity: function getWeights() constant returns(uint32)
func (_Committee *CommitteeCallerSession) GetWeights() (uint32, error) {
	return _Committee.Contract.GetWeights(&_Committee.CallOpts)
}

// GetWeights0 is a free data retrieval call binding the contract method 0x965b9ff1.
//
// Solidity: function getWeights(address[] votes) constant returns(uint32)
func (_Committee *CommitteeCaller) GetWeights0(opts *bind.CallOpts, votes []common.Address) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Committee.contract.Call(opts, out, "getWeights0", votes)
	return *ret0, err
}

// GetWeights0 is a free data retrieval call binding the contract method 0x965b9ff1.
//
// Solidity: function getWeights(address[] votes) constant returns(uint32)
func (_Committee *CommitteeSession) GetWeights0(votes []common.Address) (uint32, error) {
	return _Committee.Contract.GetWeights0(&_Committee.CallOpts, votes)
}

// GetWeights0 is a free data retrieval call binding the contract method 0x965b9ff1.
//
// Solidity: function getWeights(address[] votes) constant returns(uint32)
func (_Committee *CommitteeCallerSession) GetWeights0(votes []common.Address) (uint32, error) {
	return _Committee.Contract.GetWeights0(&_Committee.CallOpts, votes)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address governor) constant returns(bool)
func (_Committee *CommitteeCaller) IsGovernor(opts *bind.CallOpts, governor common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Committee.contract.Call(opts, out, "isGovernor", governor)
	return *ret0, err
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address governor) constant returns(bool)
func (_Committee *CommitteeSession) IsGovernor(governor common.Address) (bool, error) {
	return _Committee.Contract.IsGovernor(&_Committee.CallOpts, governor)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address governor) constant returns(bool)
func (_Committee *CommitteeCallerSession) IsGovernor(governor common.Address) (bool, error) {
	return _Committee.Contract.IsGovernor(&_Committee.CallOpts, governor)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner) returns()
func (_Committee *CommitteeTransactor) SetOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Committee.contract.Transact(opts, "setOwner", owner)
}

func (_Committee *CommitteeTransactor) AsyncSetOwner(handler func(*types.Receipt, error), opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _Committee.contract.AsyncTransact(opts, handler, "setOwner", owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner) returns()
func (_Committee *CommitteeSession) SetOwner(owner common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Committee.Contract.SetOwner(&_Committee.TransactOpts, owner)
}

func (_Committee *CommitteeSession) AsyncSetOwner(handler func(*types.Receipt, error), owner common.Address) (*types.Transaction, error) {
	return _Committee.Contract.AsyncSetOwner(handler, &_Committee.TransactOpts, owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner) returns()
func (_Committee *CommitteeTransactorSession) SetOwner(owner common.Address) (*types.Transaction, *types.Receipt, error) {
	return _Committee.Contract.SetOwner(&_Committee.TransactOpts, owner)
}

func (_Committee *CommitteeTransactorSession) AsyncSetOwner(handler func(*types.Receipt, error), owner common.Address) (*types.Transaction, error) {
	return _Committee.Contract.AsyncSetOwner(handler, &_Committee.TransactOpts, owner)
}

// SetRate is a paid mutator transaction binding the contract method 0x99bc9c1b.
//
// Solidity: function setRate(uint8 participatesRate, uint8 winRate) returns()
func (_Committee *CommitteeTransactor) SetRate(opts *bind.TransactOpts, participatesRate uint8, winRate uint8) (*types.Transaction, *types.Receipt, error) {
	return _Committee.contract.Transact(opts, "setRate", participatesRate, winRate)
}

func (_Committee *CommitteeTransactor) AsyncSetRate(handler func(*types.Receipt, error), opts *bind.TransactOpts, participatesRate uint8, winRate uint8) (*types.Transaction, error) {
	return _Committee.contract.AsyncTransact(opts, handler, "setRate", participatesRate, winRate)
}

// SetRate is a paid mutator transaction binding the contract method 0x99bc9c1b.
//
// Solidity: function setRate(uint8 participatesRate, uint8 winRate) returns()
func (_Committee *CommitteeSession) SetRate(participatesRate uint8, winRate uint8) (*types.Transaction, *types.Receipt, error) {
	return _Committee.Contract.SetRate(&_Committee.TransactOpts, participatesRate, winRate)
}

func (_Committee *CommitteeSession) AsyncSetRate(handler func(*types.Receipt, error), participatesRate uint8, winRate uint8) (*types.Transaction, error) {
	return _Committee.Contract.AsyncSetRate(handler, &_Committee.TransactOpts, participatesRate, winRate)
}

// SetRate is a paid mutator transaction binding the contract method 0x99bc9c1b.
//
// Solidity: function setRate(uint8 participatesRate, uint8 winRate) returns()
func (_Committee *CommitteeTransactorSession) SetRate(participatesRate uint8, winRate uint8) (*types.Transaction, *types.Receipt, error) {
	return _Committee.Contract.SetRate(&_Committee.TransactOpts, participatesRate, winRate)
}

func (_Committee *CommitteeTransactorSession) AsyncSetRate(handler func(*types.Receipt, error), participatesRate uint8, winRate uint8) (*types.Transaction, error) {
	return _Committee.Contract.AsyncSetRate(handler, &_Committee.TransactOpts, participatesRate, winRate)
}

// SetWeight is a paid mutator transaction binding the contract method 0xf437695a.
//
// Solidity: function setWeight(address governor, uint32 weight) returns()
func (_Committee *CommitteeTransactor) SetWeight(opts *bind.TransactOpts, governor common.Address, weight uint32) (*types.Transaction, *types.Receipt, error) {
	return _Committee.contract.Transact(opts, "setWeight", governor, weight)
}

func (_Committee *CommitteeTransactor) AsyncSetWeight(handler func(*types.Receipt, error), opts *bind.TransactOpts, governor common.Address, weight uint32) (*types.Transaction, error) {
	return _Committee.contract.AsyncTransact(opts, handler, "setWeight", governor, weight)
}

// SetWeight is a paid mutator transaction binding the contract method 0xf437695a.
//
// Solidity: function setWeight(address governor, uint32 weight) returns()
func (_Committee *CommitteeSession) SetWeight(governor common.Address, weight uint32) (*types.Transaction, *types.Receipt, error) {
	return _Committee.Contract.SetWeight(&_Committee.TransactOpts, governor, weight)
}

func (_Committee *CommitteeSession) AsyncSetWeight(handler func(*types.Receipt, error), governor common.Address, weight uint32) (*types.Transaction, error) {
	return _Committee.Contract.AsyncSetWeight(handler, &_Committee.TransactOpts, governor, weight)
}

// SetWeight is a paid mutator transaction binding the contract method 0xf437695a.
//
// Solidity: function setWeight(address governor, uint32 weight) returns()
func (_Committee *CommitteeTransactorSession) SetWeight(governor common.Address, weight uint32) (*types.Transaction, *types.Receipt, error) {
	return _Committee.Contract.SetWeight(&_Committee.TransactOpts, governor, weight)
}

func (_Committee *CommitteeTransactorSession) AsyncSetWeight(handler func(*types.Receipt, error), governor common.Address, weight uint32) (*types.Transaction, error) {
	return _Committee.Contract.AsyncSetWeight(handler, &_Committee.TransactOpts, governor, weight)
}
