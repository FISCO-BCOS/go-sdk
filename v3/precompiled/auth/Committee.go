// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package auth

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/v3/abi"
	"github.com/FISCO-BCOS/go-sdk/v3/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/v3/types"
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

// CommitteeABI is the input ABI used to generate the binding from.
const CommitteeABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"governorList\",\"type\":\"address[]\"},{\"internalType\":\"uint32[]\",\"name\":\"weightList\",\"type\":\"uint32[]\"},{\"internalType\":\"uint8\",\"name\":\"participatesRate\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"winRate\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_participatesRate\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_winRate\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"}],\"name\":\"auth\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCommitteeInfo\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"participatesRate\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"winRate\",\"type\":\"uint8\"},{\"internalType\":\"address[]\",\"name\":\"governors\",\"type\":\"address[]\"},{\"internalType\":\"uint32[]\",\"name\":\"weights\",\"type\":\"uint32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"}],\"name\":\"getWeight\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWeights\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"votes\",\"type\":\"address[]\"}],\"name\":\"getWeights\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"}],\"name\":\"isGovernor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"participatesRate\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"winRate\",\"type\":\"uint8\"}],\"name\":\"setRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"weight\",\"type\":\"uint32\"}],\"name\":\"setWeight\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CommitteeBin is the compiled bytecode used for deploying new contracts.
var CommitteeBin = "0x60806040523480156200001157600080fd5b506040516200299438038062002994833981810160405281019062000037919062000c3e565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060005b84518163ffffffff161015620000f857620000e2858263ffffffff1681518110620000aa57620000a962000cee565b5b6020026020010151858363ffffffff1681518110620000ce57620000cd62000cee565b5b60200260200101516200014a60201b60201c565b8080620000ef9062000d4c565b9150506200007a565b5080600460016101000a81548160ff021916908360ff16021790555081600460006101000a81548160ff021916908360ff16021790555062000140336200039660201b60201c565b505050506200110c565b6200015b336200042c60201b60201c565b6200019d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001949062000ddf565b60405180910390fd5b60008163ffffffff16141562000292573273ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156200021f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620002169062000e51565b60405180910390fd5b600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549063ffffffff02191690556200028c826001620004d660201b620009091790919060201c565b62000392565b620002ad8260016200072e60201b62000b471790919060201c565b15620003175780600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908363ffffffff16021790555062000391565b80600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908363ffffffff160217905550620003908260016200077d60201b62000b961790919060201c565b5b5b5050565b620003a7336200042c60201b60201c565b620003e9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620003e09062000ddf565b60405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60003073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156200046d5760019050620004d1565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415620004cc5760019050620004d1565b600090505b919050565b620004e882826200072e60201b60201c565b6200052a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620005219062000ee9565b60405180910390fd5b600060018360000160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546200057c919062000f15565b905060006001846001018054905062000596919062000f15565b90506000846001018281548110620005b357620005b262000cee565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905080856001018481548110620005fa57620005f962000cee565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018362000651919062000f50565b8560000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508460000160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000905584600101805480620006f257620006f162000fad565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905590555050505050565b6000808360000160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541415905092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415620007f0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620007e79062001052565b60405180910390fd5b6200080282826200072e60201b60201c565b1562000845576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200083c90620010ea565b60405180910390fd5b81600101819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600101805490508260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6200095f8262000914565b810181811067ffffffffffffffff8211171562000981576200098062000925565b5b80604052505050565b600062000996620008fb565b9050620009a4828262000954565b919050565b600067ffffffffffffffff821115620009c757620009c662000925565b5b602082029050602081019050919050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600062000a0a82620009dd565b9050919050565b62000a1c81620009fd565b811462000a2857600080fd5b50565b60008151905062000a3c8162000a11565b92915050565b600062000a5962000a5384620009a9565b6200098a565b9050808382526020820190506020840283018581111562000a7f5762000a7e620009d8565b5b835b8181101562000aac578062000a97888262000a2b565b84526020840193505060208101905062000a81565b5050509392505050565b600082601f83011262000ace5762000acd6200090f565b5b815162000ae084826020860162000a42565b91505092915050565b600067ffffffffffffffff82111562000b075762000b0662000925565b5b602082029050602081019050919050565b600063ffffffff82169050919050565b62000b338162000b18565b811462000b3f57600080fd5b50565b60008151905062000b538162000b28565b92915050565b600062000b7062000b6a8462000ae9565b6200098a565b9050808382526020820190506020840283018581111562000b965762000b95620009d8565b5b835b8181101562000bc3578062000bae888262000b42565b84526020840193505060208101905062000b98565b5050509392505050565b600082601f83011262000be55762000be46200090f565b5b815162000bf784826020860162000b59565b91505092915050565b600060ff82169050919050565b62000c188162000c00565b811462000c2457600080fd5b50565b60008151905062000c388162000c0d565b92915050565b6000806000806080858703121562000c5b5762000c5a62000905565b5b600085015167ffffffffffffffff81111562000c7c5762000c7b6200090a565b5b62000c8a8782880162000ab6565b945050602085015167ffffffffffffffff81111562000cae5762000cad6200090a565b5b62000cbc8782880162000bcd565b935050604062000ccf8782880162000c27565b925050606062000ce28782880162000c27565b91505092959194509250565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600062000d598262000b18565b915063ffffffff82141562000d735762000d7262000d1d565b5b600182019050919050565b600082825260208201905092915050565b7f4f6e6c79206f776e657221000000000000000000000000000000000000000000600082015250565b600062000dc7600b8362000d7e565b915062000dd48262000d8f565b602082019050919050565b6000602082019050818103600083015262000dfa8162000db8565b9050919050565b7f596f752063616e206e6f742072656d6f766520796f757273656c662100000000600082015250565b600062000e39601c8362000d7e565b915062000e468262000e01565b602082019050919050565b6000602082019050818103600083015262000e6c8162000e2a565b9050919050565b7f4c6962416464726573735365743a2076616c756520646f65736e27742065786960008201527f73742e0000000000000000000000000000000000000000000000000000000000602082015250565b600062000ed160238362000d7e565b915062000ede8262000e73565b604082019050919050565b6000602082019050818103600083015262000f048162000ec2565b9050919050565b6000819050919050565b600062000f228262000f0b565b915062000f2f8362000f0b565b92508282101562000f455762000f4462000d1d565b5b828203905092915050565b600062000f5d8262000f0b565b915062000f6a8362000f0b565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111562000fa25762000fa162000d1d565b5b828201905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b7f4c6962416464726573735365743a2076616c75652063616e277420626520307860008201527f3000000000000000000000000000000000000000000000000000000000000000602082015250565b60006200103a60218362000d7e565b9150620010478262000fdc565b604082019050919050565b600060208201905081810360008301526200106d816200102b565b9050919050565b7f4c6962416464726573735365743a2076616c756520616c72656164792065786960008201527f73747320696e20746865207365742e0000000000000000000000000000000000602082015250565b6000620010d2602f8362000d7e565b9150620010df8262001074565b604082019050919050565b600060208201905081810360008301526200110581620010c3565b9050919050565b611878806200111c6000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063ac6c525111610071578063ac6c52511461017e578063b2bdfa7b146101ae578063b6fd9067146101cc578063cd5d2118146101ea578063e43581b81461021a578063f437695a1461024a576100b4565b806313af4035146100b957806322acb867146100d55780635615696f146100f35780635e77fe2014610111578063965b9ff11461013257806399bc9c1b14610162575b600080fd5b6100d360048036038101906100ce9190610e8b565b610266565b005b6100dd6102f1565b6040516100ea9190610ed7565b60405180910390f35b6100fb61030a565b6040516101089190610f0e565b60405180910390f35b61011961031d565b60405161012994939291906110a5565b60405180910390f35b61014c60048036038101906101479190611251565b61045c565b6040516101599190610ed7565b60405180910390f35b61017c600480360381019061017791906112c6565b61050f565b005b61019860048036038101906101939190610e8b565b610591565b6040516101a59190610ed7565b60405180910390f35b6101b66105ea565b6040516101c39190611315565b60405180910390f35b6101d461060e565b6040516101e19190610f0e565b60405180910390f35b61020460048036038101906101ff9190610e8b565b610621565b604051610211919061134b565b60405180910390f35b610234600480360381019061022f9190610e8b565b6106c7565b604051610241919061134b565b60405180910390f35b610264600480360381019061025f9190611392565b6106e4565b005b61026f33610621565b6102ae576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102a59061142f565b60405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60006103056103006001610d06565b61045c565b905090565b600460009054906101000a900460ff1681565b60008060608061032d6001610d06565b9150815167ffffffffffffffff81111561034a5761034961110e565b5b6040519080825280602002602001820160405280156103785781602001602082028036833780820191505090505b50905060005b8251811015610431576003600084838151811061039e5761039d61144f565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900463ffffffff168282815181106104045761040361144f565b5b602002602001019063ffffffff16908163ffffffff16815250508080610429906114b7565b91505061037e565b50600460019054906101000a900460ff169250600460009054906101000a900460ff16935090919293565b6000806000905060005b83518163ffffffff1610156105055760036000858363ffffffff16815181106104925761049161144f565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900463ffffffff16826104f09190611500565b915080806104fd9061153a565b915050610466565b5080915050919050565b61051833610621565b610557576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054e9061142f565b60405180910390fd5b80600460016101000a81548160ff021916908360ff16021790555081600460006101000a81548160ff021916908360ff1602179055505050565b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900463ffffffff169050919050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600460019054906101000a900460ff1681565b60003073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561066057600190506106c2565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156106bd57600190506106c2565b600090505b919050565b60006106dd826001610b4790919063ffffffff16565b9050919050565b6106ed33610621565b61072c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107239061142f565b60405180910390fd5b60008163ffffffff161415610815573273ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156107aa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107a1906115b3565b60405180910390fd5b600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549063ffffffff021916905561081082600161090990919063ffffffff16565b610905565b610829826001610b4790919063ffffffff16565b156108915780600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908363ffffffff160217905550610904565b80600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908363ffffffff160217905550610903826001610b9690919063ffffffff16565b5b5b5050565b6109138282610b47565b610952576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161094990611645565b60405180910390fd5b600060018360000160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546109a29190611665565b90506000600184600101805490506109ba9190611665565b905060008460010182815481106109d4576109d361144f565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905080856001018481548110610a1857610a1761144f565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600183610a6d9190611699565b8560000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508460000160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000905584600101805480610b0b57610b0a6116ef565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905590555050505050565b6000808360000160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541415905092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610c06576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bfd90611790565b60405180910390fd5b610c108282610b47565b15610c50576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c4790611822565b60405180910390fd5b81600101819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600101805490508260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050565b60606000826001018054905067ffffffffffffffff811115610d2b57610d2a61110e565b5b604051908082528060200260200182016040528015610d595781602001602082028036833780820191505090505b50905060005b8360010180549050811015610e0f57836001018181548110610d8457610d8361144f565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16828281518110610dc257610dc161144f565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508080610e07906114b7565b915050610d5f565b5080915050919050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610e5882610e2d565b9050919050565b610e6881610e4d565b8114610e7357600080fd5b50565b600081359050610e8581610e5f565b92915050565b600060208284031215610ea157610ea0610e23565b5b6000610eaf84828501610e76565b91505092915050565b600063ffffffff82169050919050565b610ed181610eb8565b82525050565b6000602082019050610eec6000830184610ec8565b92915050565b600060ff82169050919050565b610f0881610ef2565b82525050565b6000602082019050610f236000830184610eff565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610f5e81610e4d565b82525050565b6000610f708383610f55565b60208301905092915050565b6000602082019050919050565b6000610f9482610f29565b610f9e8185610f34565b9350610fa983610f45565b8060005b83811015610fda578151610fc18882610f64565b9750610fcc83610f7c565b925050600181019050610fad565b5085935050505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61101c81610eb8565b82525050565b600061102e8383611013565b60208301905092915050565b6000602082019050919050565b600061105282610fe7565b61105c8185610ff2565b935061106783611003565b8060005b8381101561109857815161107f8882611022565b975061108a8361103a565b92505060018101905061106b565b5085935050505092915050565b60006080820190506110ba6000830187610eff565b6110c76020830186610eff565b81810360408301526110d98185610f89565b905081810360608301526110ed8184611047565b905095945050505050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b611146826110fd565b810181811067ffffffffffffffff821117156111655761116461110e565b5b80604052505050565b6000611178610e19565b9050611184828261113d565b919050565b600067ffffffffffffffff8211156111a4576111a361110e565b5b602082029050602081019050919050565b600080fd5b60006111cd6111c884611189565b61116e565b905080838252602082019050602084028301858111156111f0576111ef6111b5565b5b835b8181101561121957806112058882610e76565b8452602084019350506020810190506111f2565b5050509392505050565b600082601f830112611238576112376110f8565b5b81356112488482602086016111ba565b91505092915050565b60006020828403121561126757611266610e23565b5b600082013567ffffffffffffffff81111561128557611284610e28565b5b61129184828501611223565b91505092915050565b6112a381610ef2565b81146112ae57600080fd5b50565b6000813590506112c08161129a565b92915050565b600080604083850312156112dd576112dc610e23565b5b60006112eb858286016112b1565b92505060206112fc858286016112b1565b9150509250929050565b61130f81610e4d565b82525050565b600060208201905061132a6000830184611306565b92915050565b60008115159050919050565b61134581611330565b82525050565b6000602082019050611360600083018461133c565b92915050565b61136f81610eb8565b811461137a57600080fd5b50565b60008135905061138c81611366565b92915050565b600080604083850312156113a9576113a8610e23565b5b60006113b785828601610e76565b92505060206113c88582860161137d565b9150509250929050565b600082825260208201905092915050565b7f4f6e6c79206f776e657221000000000000000000000000000000000000000000600082015250565b6000611419600b836113d2565b9150611424826113e3565b602082019050919050565b600060208201905081810360008301526114488161140c565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000819050919050565b60006114c2826114ad565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156114f5576114f461147e565b5b600182019050919050565b600061150b82610eb8565b915061151683610eb8565b92508263ffffffff0382111561152f5761152e61147e565b5b828201905092915050565b600061154582610eb8565b915063ffffffff82141561155c5761155b61147e565b5b600182019050919050565b7f596f752063616e206e6f742072656d6f766520796f757273656c662100000000600082015250565b600061159d601c836113d2565b91506115a882611567565b602082019050919050565b600060208201905081810360008301526115cc81611590565b9050919050565b7f4c6962416464726573735365743a2076616c756520646f65736e27742065786960008201527f73742e0000000000000000000000000000000000000000000000000000000000602082015250565b600061162f6023836113d2565b915061163a826115d3565b604082019050919050565b6000602082019050818103600083015261165e81611622565b9050919050565b6000611670826114ad565b915061167b836114ad565b92508282101561168e5761168d61147e565b5b828203905092915050565b60006116a4826114ad565b91506116af836114ad565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156116e4576116e361147e565b5b828201905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b7f4c6962416464726573735365743a2076616c75652063616e277420626520307860008201527f3000000000000000000000000000000000000000000000000000000000000000602082015250565b600061177a6021836113d2565b91506117858261171e565b604082019050919050565b600060208201905081810360008301526117a98161176d565b9050919050565b7f4c6962416464726573735365743a2076616c756520616c72656164792065786960008201527f73747320696e20746865207365742e0000000000000000000000000000000000602082015250565b600061180c602f836113d2565b9150611817826117b0565b604082019050919050565b6000602082019050818103600083015261183b816117ff565b905091905056fea26469706673582212201b41b43858de334a1cf957ae95eb654f54c1e0be48abe6096cacef9c74fdcb8664736f6c634300080b0033"

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
func (_Committee *CommitteeRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Committee.Contract.CommitteeTransactor.contract.TransactWithResult(opts, result, method, params...)
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
func (_Committee *CommitteeTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Committee.Contract.contract.TransactWithResult(opts, result, method, params...)
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
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _Committee.contract.TransactWithResult(opts, out, "setOwner", owner)
	return transaction, receipt, err
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
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _Committee.contract.TransactWithResult(opts, out, "setRate", participatesRate, winRate)
	return transaction, receipt, err
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
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _Committee.contract.TransactWithResult(opts, out, "setWeight", governor, weight)
	return transaction, receipt, err
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
