package types

import (
	"bytes"
	"testing"

	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
	"github.com/ethereum/go-ethereum/common"
)

const (
	DEPLOY_HELLOWORLD_TX_DATA    = `1c2606636861696e30360667726f75703041020a5628313638383630333033393337383835393933333531353931323138383539373634313130333138307d0001060e60806040526040805190810160405280600181526020017f31000000000000000000000000000000000000000000000000000000000000008152506001908051906020019061004f9291906100ae565b5034801561005c57600080fd5b506040805190810160405280600d81526020017f48656c6c6f2c20576f726c642100000000000000000000000000000000000000815250600090805190602001906100a89291906100ae565b50610153565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100ef57805160ff191683800117855561011d565b8280016001018555821561011d579182015b8281111561011c578251825591602001919060010190610101565b5b50905061012a919061012e565b5090565b61015091905b8082111561014c576000816000905550600101610134565b5090565b90565b6104ac806101626000396000f300608060405260043610610057576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680634ed3885e1461005c57806354fd4d50146100c55780636d4ce63c14610155575b600080fd5b34801561006857600080fd5b506100c3600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506101e5565b005b3480156100d157600080fd5b506100da61029b565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561011a5780820151818401526020810190506100ff565b50505050905090810190601f1680156101475780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561016157600080fd5b5061016a610339565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101aa57808201518184015260208101905061018f565b50505050905090810190601f1680156101d75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b80600090805190602001906101fb9291906103db565b507f93a093529f9c8a0c300db4c55fcd27c068c4f5e0e8410bc288c7e76f3d71083e816040518080602001828103825283818151815260200191508051906020019080838360005b8381101561025e578082015181840152602081019050610243565b50505050905090810190601f16801561028b5780820380516001836020036101000a031916815260200191505b509250505060405180910390a150565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103315780601f1061030657610100808354040283529160200191610331565b820191906000526020600020905b81548152906001019060200180831161031457829003601f168201915b505050505081565b606060008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103d15780601f106103a6576101008083540402835291602001916103d1565b820191906000526020600020905b8154815290600101906020018083116103b457829003601f168201915b5050505050905090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061041c57805160ff191683800117855561044a565b8280016001018555821561044a579182015b8281111561044957825182559160200191906001019061042e565b5b509050610457919061045b565b5090565b61047d91905b80821115610479576000816000905550600101610461565b5090565b905600a165627a7a72305820fd433a091cb8e1aba3f49e5efb35f937e4b22a85a46f35574834d120699d7ae5002987000002755b7b22636f6e7374616e74223a66616c73652c22696e70757473223a5b7b226e616d65223a2276222c2274797065223a22737472696e67227d5d2c226e616d65223a22736574222c226f757470757473223a5b5d2c2270617961626c65223a66616c73652c2273746174654d75746162696c697479223a226e6f6e70617961626c65222c2274797065223a2266756e6374696f6e227d2c7b22636f6e7374616e74223a747275652c22696e70757473223a5b5d2c226e616d65223a2276657273696f6e222c226f757470757473223a5b7b226e616d65223a22222c2274797065223a22737472696e67227d5d2c2270617961626c65223a66616c73652c2273746174654d75746162696c697479223a2276696577222c2274797065223a2266756e6374696f6e227d2c7b22636f6e7374616e74223a747275652c22696e70757473223a5b5d2c226e616d65223a22676574222c226f757470757473223a5b7b226e616d65223a22222c2274797065223a22737472696e67227d5d2c2270617961626c65223a66616c73652c2273746174654d75746162696c697479223a2276696577222c2274797065223a2266756e6374696f6e227d2c7b22696e70757473223a5b5d2c2270617961626c65223a66616c73652c2273746174654d75746162696c697479223a226e6f6e70617961626c65222c2274797065223a22636f6e7374727563746f72227d2c7b22616e6f6e796d6f7573223a66616c73652c22696e70757473223a5b7b22696e6465786564223a66616c73652c226e616d65223a22222c2274797065223a22737472696e67227d5d2c226e616d65223a2273657456616c7565222c2274797065223a226576656e74227d5d`
	DEPLOY_HELLOWORLD_TX         = `1a1c2606636861696e30360667726f75703041020a5628313638383630333033393337383835393933333531353931323138383539373634313130333138307d0001060e60806040526040805190810160405280600181526020017f31000000000000000000000000000000000000000000000000000000000000008152506001908051906020019061004f9291906100ae565b5034801561005c57600080fd5b506040805190810160405280600d81526020017f48656c6c6f2c20576f726c642100000000000000000000000000000000000000815250600090805190602001906100a89291906100ae565b50610153565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100ef57805160ff191683800117855561011d565b8280016001018555821561011d579182015b8281111561011c578251825591602001919060010190610101565b5b50905061012a919061012e565b5090565b61015091905b8082111561014c576000816000905550600101610134565b5090565b90565b6104ac806101626000396000f300608060405260043610610057576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680634ed3885e1461005c57806354fd4d50146100c55780636d4ce63c14610155575b600080fd5b34801561006857600080fd5b506100c3600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506101e5565b005b3480156100d157600080fd5b506100da61029b565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561011a5780820151818401526020810190506100ff565b50505050905090810190601f1680156101475780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561016157600080fd5b5061016a610339565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101aa57808201518184015260208101905061018f565b50505050905090810190601f1680156101d75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b80600090805190602001906101fb9291906103db565b507f93a093529f9c8a0c300db4c55fcd27c068c4f5e0e8410bc288c7e76f3d71083e816040518080602001828103825283818151815260200191508051906020019080838360005b8381101561025e578082015181840152602081019050610243565b50505050905090810190601f16801561028b5780820380516001836020036101000a031916815260200191505b509250505060405180910390a150565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103315780601f1061030657610100808354040283529160200191610331565b820191906000526020600020905b81548152906001019060200180831161031457829003601f168201915b505050505081565b606060008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103d15780601f106103a6576101008083540402835291602001916103d1565b820191906000526020600020905b8154815290600101906020018083116103b457829003601f168201915b5050505050905090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061041c57805160ff191683800117855561044a565b8280016001018555821561044a579182015b8281111561044957825182559160200191906001019061042e565b5b509050610457919061045b565b5090565b61047d91905b80821115610479576000816000905550600101610461565b5090565b905600a165627a7a72305820fd433a091cb8e1aba3f49e5efb35f937e4b22a85a46f35574834d120699d7ae5002987000002755b7b22636f6e7374616e74223a66616c73652c22696e70757473223a5b7b226e616d65223a2276222c2274797065223a22737472696e67227d5d2c226e616d65223a22736574222c226f757470757473223a5b5d2c2270617961626c65223a66616c73652c2273746174654d75746162696c697479223a226e6f6e70617961626c65222c2274797065223a2266756e6374696f6e227d2c7b22636f6e7374616e74223a747275652c22696e70757473223a5b5d2c226e616d65223a2276657273696f6e222c226f757470757473223a5b7b226e616d65223a22222c2274797065223a22737472696e67227d5d2c2270617961626c65223a66616c73652c2273746174654d75746162696c697479223a2276696577222c2274797065223a2266756e6374696f6e227d2c7b22636f6e7374616e74223a747275652c22696e70757473223a5b5d2c226e616d65223a22676574222c226f757470757473223a5b7b226e616d65223a22222c2274797065223a22737472696e67227d5d2c2270617961626c65223a66616c73652c2273746174654d75746162696c697479223a2276696577222c2274797065223a2266756e6374696f6e227d2c7b22696e70757473223a5b5d2c2270617961626c65223a66616c73652c2273746174654d75746162696c697479223a226e6f6e70617961626c65222c2274797065223a22636f6e7374727563746f72227d2c7b22616e6f6e796d6f7573223a66616c73652c22696e70757473223a5b7b22696e6465786564223a66616c73652c226e616d65223a22222c2274797065223a22737472696e67227d5d2c226e616d65223a2273657456616c7565222c2274797065223a226576656e74227d5d0b2d0000207e9f0440919e2bd719e29cc512b8ed364a7ce81091d460ea31272106424e93e03d000041b8227e0be95ea1195f8fd605b3d5c288ac4481c366749c9ab12d8e5eef7e57885bd474a427fdba3e6fd3d28171b3fa2e47c373e77c45514d8985abf5715cb5be00`
	HELLOWORLD_SET_TX_DATA       = `1c2606636861696e30360667726f75703041020a56273136313235353631323839383636383633323732353836393034303732333238373635373533336628613663646334666564393662626638633963303133636331393230306233633862613935633933657d0000644ed3885e0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000c68656c6c6f2c20776f726c640000000000000000000000000000000000000000`
	HELLOWORLD_SET_TX            = `1a1c2606636861696e30360667726f75703041020a56273136313235353631323839383636383633323732353836393034303732333238373635373533336628613663646334666564393662626638633963303133636331393230306233633862613935633933657d0000644ed3885e0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000c68656c6c6f2c20776f726c6400000000000000000000000000000000000000000b2d000020b1e82e9052f066250532968bf9ad3a05f19dba10291089a27cea931ddfc8e9b53d000041aa923c250f2e8b64f401766b9348b27234858ee7cceefedda8c72cefc94d4142723d0b2584e753629a69404b8bfb33748747ec6b5a286488f9bebce79b2ef27101`
	DEPLOY_HELLOWORLD_TX_DATA_GM = `1c2606636861696e30360667726f7570304101f456273136353333333933373031333131333936353533313836343438303132323233393434303735327d0001060e60806040526040805190810160405280600181526020017f31000000000000000000000000000000000000000000000000000000000000008152506001908051906020019061004f9291906100ae565b5034801561005c57600080fd5b506040805190810160405280600d81526020017f48656c6c6f2c20576f726c642100000000000000000000000000000000000000815250600090805190602001906100a89291906100ae565b50610153565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100ef57805160ff191683800117855561011d565b8280016001018555821561011d579182015b8281111561011c578251825591602001919060010190610101565b5b50905061012a919061012e565b5090565b61015091905b8082111561014c576000816000905550600101610134565b5090565b90565b6104ac806101626000396000f300608060405260043610610057576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680634ed3885e1461005c57806354fd4d50146100c55780636d4ce63c14610155575b600080fd5b34801561006857600080fd5b506100c3600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506101e5565b005b3480156100d157600080fd5b506100da61029b565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561011a5780820151818401526020810190506100ff565b50505050905090810190601f1680156101475780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561016157600080fd5b5061016a610339565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101aa57808201518184015260208101905061018f565b50505050905090810190601f1680156101d75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b80600090805190602001906101fb9291906103db565b507f93a093529f9c8a0c300db4c55fcd27c068c4f5e0e8410bc288c7e76f3d71083e816040518080602001828103825283818151815260200191508051906020019080838360005b8381101561025e578082015181840152602081019050610243565b50505050905090810190601f16801561028b5780820380516001836020036101000a031916815260200191505b509250505060405180910390a150565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103315780601f1061030657610100808354040283529160200191610331565b820191906000526020600020905b81548152906001019060200180831161031457829003601f168201915b505050505081565b606060008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103d15780601f106103a6576101008083540402835291602001916103d1565b820191906000526020600020905b8154815290600101906020018083116103b457829003601f168201915b5050505050905090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061041c57805160ff191683800117855561044a565b8280016001018555821561044a579182015b8281111561044957825182559160200191906001019061042e565b5b509050610457919061045b565b5090565b61047d91905b80821115610479576000816000905550600101610461565b5090565b905600a165627a7a72305820fd433a091cb8e1aba3f49e5efb35f937e4b22a85a46f35574834d120699d7ae5002987000002755b7b22636f6e7374616e74223a66616c73652c22696e70757473223a5b7b226e616d65223a2276222c2274797065223a22737472696e67227d5d2c226e616d65223a22736574222c226f757470757473223a5b5d2c2270617961626c65223a66616c73652c2273746174654d75746162696c697479223a226e6f6e70617961626c65222c2274797065223a2266756e6374696f6e227d2c7b22636f6e7374616e74223a747275652c22696e70757473223a5b5d2c226e616d65223a2276657273696f6e222c226f757470757473223a5b7b226e616d65223a22222c2274797065223a22737472696e67227d5d2c2270617961626c65223a66616c73652c2273746174654d75746162696c697479223a2276696577222c2274797065223a2266756e6374696f6e227d2c7b22636f6e7374616e74223a747275652c22696e70757473223a5b5d2c226e616d65223a22676574222c226f757470757473223a5b7b226e616d65223a22222c2274797065223a22737472696e67227d5d2c2270617961626c65223a66616c73652c2273746174654d75746162696c697479223a2276696577222c2274797065223a2266756e6374696f6e227d2c7b22696e70757473223a5b5d2c2270617961626c65223a66616c73652c2273746174654d75746162696c697479223a226e6f6e70617961626c65222c2274797065223a22636f6e7374727563746f72227d2c7b22616e6f6e796d6f7573223a66616c73652c22696e70757473223a5b7b22696e6465786564223a66616c73652c226e616d65223a22222c2274797065223a22737472696e67227d5d2c226e616d65223a2273657456616c7565222c2274797065223a226576656e74227d5d`
	DEPLOY_HELLOWORLD_TX_GM      = `1a1c2606636861696e30360667726f7570304101f456273136353333333933373031333131333936353533313836343438303132323233393434303735327d0001060e60806040526040805190810160405280600181526020017f31000000000000000000000000000000000000000000000000000000000000008152506001908051906020019061004f9291906100ae565b5034801561005c57600080fd5b506040805190810160405280600d81526020017f48656c6c6f2c20576f726c642100000000000000000000000000000000000000815250600090805190602001906100a89291906100ae565b50610153565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100ef57805160ff191683800117855561011d565b8280016001018555821561011d579182015b8281111561011c578251825591602001919060010190610101565b5b50905061012a919061012e565b5090565b61015091905b8082111561014c576000816000905550600101610134565b5090565b90565b6104ac806101626000396000f300608060405260043610610057576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680634ed3885e1461005c57806354fd4d50146100c55780636d4ce63c14610155575b600080fd5b34801561006857600080fd5b506100c3600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506101e5565b005b3480156100d157600080fd5b506100da61029b565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561011a5780820151818401526020810190506100ff565b50505050905090810190601f1680156101475780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561016157600080fd5b5061016a610339565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101aa57808201518184015260208101905061018f565b50505050905090810190601f1680156101d75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b80600090805190602001906101fb9291906103db565b507f93a093529f9c8a0c300db4c55fcd27c068c4f5e0e8410bc288c7e76f3d71083e816040518080602001828103825283818151815260200191508051906020019080838360005b8381101561025e578082015181840152602081019050610243565b50505050905090810190601f16801561028b5780820380516001836020036101000a031916815260200191505b509250505060405180910390a150565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103315780601f1061030657610100808354040283529160200191610331565b820191906000526020600020905b81548152906001019060200180831161031457829003601f168201915b505050505081565b606060008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103d15780601f106103a6576101008083540402835291602001916103d1565b820191906000526020600020905b8154815290600101906020018083116103b457829003601f168201915b5050505050905090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061041c57805160ff191683800117855561044a565b8280016001018555821561044a579182015b8281111561044957825182559160200191906001019061042e565b5b509050610457919061045b565b5090565b61047d91905b80821115610479576000816000905550600101610461565b5090565b905600a165627a7a72305820fd433a091cb8e1aba3f49e5efb35f937e4b22a85a46f35574834d120699d7ae5002987000002755b7b22636f6e7374616e74223a66616c73652c22696e70757473223a5b7b226e616d65223a2276222c2274797065223a22737472696e67227d5d2c226e616d65223a22736574222c226f757470757473223a5b5d2c2270617961626c65223a66616c73652c2273746174654d75746162696c697479223a226e6f6e70617961626c65222c2274797065223a2266756e6374696f6e227d2c7b22636f6e7374616e74223a747275652c22696e70757473223a5b5d2c226e616d65223a2276657273696f6e222c226f757470757473223a5b7b226e616d65223a22222c2274797065223a22737472696e67227d5d2c2270617961626c65223a66616c73652c2273746174654d75746162696c697479223a2276696577222c2274797065223a2266756e6374696f6e227d2c7b22636f6e7374616e74223a747275652c22696e70757473223a5b5d2c226e616d65223a22676574222c226f757470757473223a5b7b226e616d65223a22222c2274797065223a22737472696e67227d5d2c2270617961626c65223a66616c73652c2273746174654d75746162696c697479223a2276696577222c2274797065223a2266756e6374696f6e227d2c7b22696e70757473223a5b5d2c2270617961626c65223a66616c73652c2273746174654d75746162696c697479223a226e6f6e70617961626c65222c2274797065223a22636f6e7374727563746f72227d2c7b22616e6f6e796d6f7573223a66616c73652c22696e70757473223a5b7b22696e6465786564223a66616c73652c226e616d65223a22222c2274797065223a22737472696e67227d5d2c226e616d65223a2273657456616c7565222c2274797065223a226576656e74227d5d0b2d00002099f8ac271b22847c54a93e1cecf82c9b017430aafb0364c7d3fe76093ba051f83d000100805d7ea32ba1a14674da4ebeb87c88310b624b3ff53174c07d401cc125eacc58b70f66ee126e73d450958ecd483468c815e1f739094516a5d436a95af0a217d26e4c5a43dcbecdb28436cf138f5b74f8f4809415be473fa31e923e27a0e0eaebabe002537bbf5f787e87df7fa9acf7b13391ccf85618009ccb7413376641b19133`
	HELLOWORLD_SET_TX_DATA_GM    = `1c2606636861696e30360667726f7570304101f456273134323433393934373830323139353931303230323835343436343433333339333430333931356628633035323364626464393462613237653134623033333664373939343839333430636132346364667d0000644ed3885e0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000c68656c6c6f2c20776f726c640000000000000000000000000000000000000000`
	HELLOWORLD_SET_TX_GM         = `1a1c2606636861696e30360667726f7570304101f456273134323433393934373830323139353931303230323835343436343433333339333430333931356628633035323364626464393462613237653134623033333664373939343839333430636132346364667d0000644ed3885e0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000c68656c6c6f2c20776f726c6400000000000000000000000000000000000000000b2d000020db9d06a9258d4e8dbbe490a4f4c3f0bfa4fe8aa783ad77c2720c749410ee9bca3d0001008045e0d42a74157f03aaf9e7d2631df0c21e74ce72dbee2c3c047ac12e0435d24f8563b5b3b7b0a62bd0f95b297530a4e8d789bfb29762f5c47e789e6ffe366fa74c5a43dcbecdb28436cf138f5b74f8f4809415be473fa31e923e27a0e0eaebabe002537bbf5f787e87df7fa9acf7b13391ccf85618009ccb7413376641b19133`
)

func checkTransaction(dataHex, txHex string, isSmCrypto bool, t *testing.T) {
	deployTxData := common.FromHex(dataHex)
	readBuf := codec.NewReader(deployTxData)
	txData := TransactionData{}
	txData.ReadFrom(readBuf)

	deployTx := common.FromHex(txHex)
	txReadBuf := codec.NewReader(deployTx)
	tx := &Transaction{SMCrypto: isSmCrypto}
	tx.ReadFrom(txReadBuf)
	if tx.Data.Version != txData.Version {
		t.Errorf("tx.Data.Version should be %d, but got %d", tx.Data.Version, txData.Version)
	}
	if tx.Data.ChainID != txData.ChainID {
		t.Errorf("tx.Data.ChainID should be %s, but got %s", tx.Data.ChainID, txData.ChainID)
	}
	if tx.Data.GroupID != txData.GroupID {
		t.Errorf("tx.Data.GroupID should be %s, but got %s", tx.Data.GroupID, txData.GroupID)
	}
	if tx.Data.BlockLimit != txData.BlockLimit {
		t.Errorf("tx.Data.BlockLimit should be %d, but got %d", tx.Data.BlockLimit, txData.BlockLimit)
	}
	if tx.Data.Nonce != txData.Nonce {
		t.Errorf("tx.Data.Nonce should be %s, but got %s", tx.Data.Nonce, txData.Nonce)
	}
	if tx.Data.To != nil && txData.To != nil {
		if *tx.Data.To != *txData.To {
			t.Errorf("tx.Data.To should be %s, but got %s", tx.Data.To.String(), txData.To.String())
		}
	} else if tx.Data.To == nil && txData.To == nil {

	} else {
		t.Errorf("tx.Data.To should be %s, but got %s", tx.Data.To.String(), txData.To.String())
	}
	if !bytes.Equal(tx.Data.Input, txData.Input) {
		t.Errorf("tx.Data.Input should be %s, but got %s", common.Bytes2Hex(tx.Data.Input), common.Bytes2Hex(txData.Input))
	}
	if tx.Data.Abi != txData.Abi {
		t.Errorf("tx.Data.Abi should be %s, but got %s", tx.Data.Abi, txData.Abi)
	}
	if tx.Data.Value != nil {
		t.Errorf("tx.Data.Value should be nil, but got %s", tx.Data.Value)
	}
	// if tx.Data.Value != txData.Value {
	// 	t.Errorf("tx.Data.Value should be %s, but got %s", tx.Data.Value, txData.Value)
	// }
	// if tx.Data.GasPrice != txData.GasPrice {
	// 	t.Errorf("tx.Data.GasPrice should be %s, but got %s", tx.Data.GasPrice, txData.GasPrice)
	// }
	if tx.Data.GasLimit != txData.GasLimit {
		t.Errorf("tx.Data.Gas should be %v, but got %v", tx.Data.GasLimit, txData.GasLimit)
	}
	// if tx.Data.MaxFeePerGas != txData.MaxFeePerGas {
	// 	t.Errorf("tx.Data.MaxFeePerGas should be %s, but got %s", tx.Data.MaxFeePerGas, txData.MaxFeePerGas)
	// }
	// if tx.Data.MaxPriorityFeePerGas != txData.MaxPriorityFeePerGas {
	// 	t.Errorf("tx.Data.MaxPriorityFeePerGas should be %s, but got %s", tx.Data.MaxPriorityFeePerGas, txData.MaxPriorityFeePerGas)
	// }
	rightHash := tx.DataHash
	newHash := tx.Hash()
	if !bytes.Equal(rightHash.Bytes(), newHash.Bytes()) {
		t.Errorf("tx.DataHash should be %s, but got %s", rightHash.String(), newHash.String())
	}
	if !bytes.Equal(tx.Data.Bytes(), deployTxData) {
		t.Errorf("tx.Data should be %#x, but got %#x", deployTxData, tx.Data.Bytes())
	}
}

func TestUnmarshalTransactionDataDeploy(t *testing.T) {
	checkTransaction(DEPLOY_HELLOWORLD_TX_DATA, DEPLOY_HELLOWORLD_TX, false, t)
	checkTransaction(HELLOWORLD_SET_TX_DATA, HELLOWORLD_SET_TX, false, t)
	checkTransaction(DEPLOY_HELLOWORLD_TX_DATA_GM, DEPLOY_HELLOWORLD_TX_GM, true, t)
	checkTransaction(HELLOWORLD_SET_TX_DATA_GM, HELLOWORLD_SET_TX_GM, true, t)
}
