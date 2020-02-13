# go-sdk

Golang SDK For FISCO BCOS 2.0.0

[![CodeFactor](https://www.codefactor.io/repository/github/fisco-bcos/go-sdk/badge)](https://www.codefactor.io/repository/github/fisco-bcos/go-sdk) [![Codacy Badge](https://api.codacy.com/project/badge/Grade/afbb696df3a8436a9e446d39251b2158)](https://www.codacy.com/gh/FISCO-BCOS/go-sdk?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=FISCO-BCOS/go-sdk&amp;utm_campaign=Badge_Grade)


![FISCO-BCOS Go-SDK GitHub Actions](https://github.com/FISCO-BCOS/go-sdk/workflows/FISCO-BCOS%20Go-SDK%20GitHub%20Actions/badge.svg)  [![Build Status](https://travis-ci.org/FISCO-BCOS/go-sdk.svg?branch=master)](https://travis-ci.org/FISCO-BCOS/go-sdk)  [![codecov](https://codecov.io/gh/FISCO-BCOS/go-sdk/branch/master/graph/badge.svg)](https://codecov.io/gh/FISCO-BCOS/go-sdk)  ![Code Lines](https://tokei.rs/b1/github/FISCO-BCOS/go-sdk?category=code)
____

FISCO BCOS Go语言版本的SDK，借助以太坊代码进行改进，主要实现的功能有：

- [FISCO BCOS 2.0 JSON-RPC服务](https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/api.html)
- `Solidity`合约编译为Go文件
- 部署、查询、写入智能合约
- 控制台

`go-sdk`的使用可以当做是一个`package`进行使用，亦可对项目代码进行编译，直接使用**控制台**通过配置文件来进行访问FISCO BCOS。

# 环境准备

- [Golang](https://golang.org/), 版本需不低于`1.13.6`，本项目采用`go module`进行包管理。具体可查阅[Using Go Modules](https://blog.golang.org/using-go-modules)
- [FISCO BCOS 2.2.0+](https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/), **需要提前运行** FISCO BCOS 区块链平台，可参考[安装搭建](https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/installation.html#fisco-bcos)
- Solidity编译器，默认[0.4.25版本](https://github.com/ethereum/solidity/releases/tag/v0.4.25)

# 控制台使用
在使用控制台需要先拉取代码或下载代码，然后对配置文件`config.toml`进行更改:

```bash
git clone https://github.com/FISCO-BCOS/go-sdk.git
cd go-sdk
```

本项目使用了`go module`的[特性](https://blog.golang.org/using-go-modules)，可以在旧版本的`$GOPATH`路径之外直接运行`go`命令，如果项目仍然在`$GOPATH`路径之下，则需要显示开启`GO111MODULE`以支持该特性:

```
export GO111MODULE=on
```

编译代码后会在`$GOBIN`下生成控制台程序，请确保添加`$GOBIN`到系统路径`$PATH`中，关于`$GOBIN`等的设置可以[参考这里](https://www.cnblogs.com/zhaof/p/7906722.html)，以便能正常执行`go`生成的程序:

```go
go build cmd/console.go
```

如果不能访问外网，则可以设置开源代理进行依赖下载(需使用`go module`的特性)：
```bash
export GOPROXY=https://goproxy.io
```

如若仍然无法解决依赖问题，则可以[参考文章](https://shockerli.net/post/go-get-golang-org-x-solution/)，使用手动下载的方式，但无法支持具体版本的依赖库 :(

最后，运行控制台查看可用指令:

```bash
./console help
```

# Package功能使用

以下的示例是通过`import`的方式来使用`go-sdk`，如引入RPC控制台库:

```go
import "github.com/FISCO-BCOS/go-sdk/client"
```

## Solidity合约编译为Go文件

在利用SDK进行项目开发时，对智能合约进行操作时需要将Solidity智能合约利用go-sdk的`abigen`工具转换为`Go`文件代码。整体上主要包含了五个流程：

- 准备需要编译的智能合约
- 配置好相应版本的`solc`编译器
- 构建go-sdk的合约编译工具`abigen`
- 编译生成go文件
- 使用生成的go文件进行合约调用

下面的内容作为一个示例进行使用介绍。

1.提供一份简单的样例智能合约`Store.sol`如下:

```solidity
pragma solidity ^0.4.25;

contract Store {
  event ItemSet(bytes32 key, bytes32 value);

  string public version;
  mapping (bytes32 => bytes32) public items;

  constructor(string _version) public {
    version = _version;
  }

  function setItem(bytes32 key, bytes32 value) external {
    items[key] = value;
    emit ItemSet(key, value);
  }
}
```

2.安装对应版本的[`solc`编译器](https://github.com/ethereum/solidity/releases/tag/v0.4.25)，目前FISCO BCOS默认的`solc`编译器版本为[0.4.25](https://github.com/ethereum/solidity/releases/tag/v0.4.25)。

```bash
solc --version
# solc, the solidity compiler commandline interface
# Version: 0.4.25+commit.59dbf8f1.Linux.g++
```

3.构建`go-sdk`的代码生成工具`abigen`

```bash
git clone https://github.com/FISCO-BCOS/go-sdk.git # 下载go-sdk代码，如已下载请跳过
cd go-sdk # 进入代码目录
go build ./cmd/abigen # 编译生成abigen工具
```

执行命令后，检查根目录下是否存在`abigen`，并将生成的`abigen`以及所准备的智能合约`Store.sol`放置在一个新的目录下：

```
mkdir ../newdir && cp ./abigen ../newdir && cd ../newdir
# cp your/Store.sol path/to/newdir
```

4.编译生成go文件，先利用`solc`将合约文件生成`abi`和`bin`文件，以前面所提供的`Store.sol`为例：

```bash
solc --bin --abi -o ./ Store.sol
```

`Store.sol`目录下会生成`Store.bin`和`Store.abi`。此时利用`abigen`工具将`Store.bin`和`Store.abi`转换成`Store.go`：

```bash
./abigen --bin Store.bin --abi Store.abi --pkg store --out Store.go
```

最后目录下面存在以下文件：

```bash
>>ls
abigen  Store.abi  Store.bin  Store.go  Store.sol
```

5.调用生成的`Store.go`文件进行合约调用

至此，合约已成功转换为go文件，用户可根据此文件在项目中利用SDK进行合约操作。具体的使用可参阅下一节。

## 部署、查询、写入智能合约

此部分承接上一节的内容，同时也简单涵盖了SDK的合约使用部分以及账户私钥的加载。

### 创建外部账户

SDK发送交易需要一个外部账户，导入go-sdk的`crypto`包，该包提供用于生成随机私钥的`GenerateKey`方法：

```go
privateKey, err := crypto.GenerateKey()
if err != nil {
    log.Fatal(err)
}
```

然后我们可以通过导入golang`crypto/ecdsa`包并使用`FromECDSA`方法将其转换为字节：

```go
privateKeyBytes := crypto.FromECDSA(privateKey)
```

我们现在可以使用go-sdk的`common/hexutil`包将它转换为十六进制字符串，该包提供了一个带有字节切片的`Encode`方法。 然后我们在十六进制编码之后删除“0x”。

```go
fmt.Println(hexutil.Encode(privateKeyBytes)[2:])
```

这就是`用于签署交易的私钥，将被视为密码，永远不应该被共享给别人`。

由于公钥是从私钥派生的，加密私钥具有一个返回公钥的`Public`方法：

```go
publicKey := privateKey.Public()
```

将其转换为十六进制的过程与我们使用转化私钥的过程类似。 我们剥离了`0x`和前2个字符`04`，它始终是EC前缀，不是必需的。

```go
publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
if !ok {
    log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
}

publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
fmt.Println(hexutil.Encode(publicKeyBytes)[4:])
```

现在我们拥有公钥，就可以轻松生成你经常看到的公共地址。 加密包里有一个`PubkeyToAddress`方法，它接受一个ECDSA公钥，并返回公共地址。

```go
address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
fmt.Println(address) // 0x96216849c49358B10257cb55b28eA603c874b05E
```

公共地址可以查询合约信息。

整体的代码示例为：

```go
package main

import (
    "crypto/ecdsa"
    "fmt"
    "log"
    "os"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/common/hexutil"
)

func main() {
    privateKey, err := crypto.GenerateKey()
    if err != nil {
        log.Fatal(err)
    }

    privateKeyBytes := crypto.FromECDSA(privateKey)
    fmt.Println("private key: ", hexutil.Encode(privateKeyBytes)[2:]) // privateKey in hex without "0x"

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }

    publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
    fmt.Println("publick key: ", hexutil.Encode(publicKeyBytes)[4:])  // publicKey in hex without "0x"

    address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
    fmt.Println("address: ", address)  // account address
}
```

### 部署智能合约

首先在利用`abigen`生成的`Store.go`文件下，创建一个新的`contract_run.go`文件用来调用`Store.go`文件，并创建一个新的文件夹来放置`Store.go`以方便调用，同时利用`go mod`进行包管理，初始化为一个`contract`包：

```bash
mkdir testfile
mv ./Store.go testfile
touch contract_run.go
go mod init contract
```

此时目录下会生成`go.mod`包管理文件。而在`contract_deploy.go`部署合约之前，需要先从`go-sdk`中导入`accounts/abi/bind`包，然后调用传入`client.GetTransactOpts()`：

```go 
package main

import (
    "fmt"
    "log"
    "github.com/FISCO-BCOS/go-sdk/client"
    "github.com/FISCO-BCOS/go-sdk/abi/bind"
    "github.com/ethereum/go-ethereum/crypto"
    store "contract/testfile" // import Store.go
)

func main(){
    config := &conf.Config{IsHTTP: true, ChainID: 1, IsSMCrypto: false, GroupID: 1, PrivateKey:"145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58",NodeURL: "http://localhost:8545"}

    client, err := client.Dial(config)
    if err != nil {
        log.Fatal(err)
    }
    input := "Store deployment 1.0"
    address, tx, instance, err := store.DeployStore(client.GetTransactOpts(), client, input)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("contract address: ", address.Hex())  // the address should be saved
    fmt.Println("transaction hash: ", tx.Hash().Hex())
    _ = instance
}
```

### 加载智能合约
在部署完智能合约后，可以获取到合约的地址，但在进行合约查询以及写入时，需要先加载智能合约，此时需要导入`common`包以获取正确的合约地址，新建`contract_load.go`以加载智能合约：

```go
package main 

import (
    "fmt"
    "log"
    "github.com/ethereum/go-ethereum/common"
    "github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
    store "contract/testfile" // for demo
)

func main() {
    config := &conf.ParseConfig("config.toml")[0]
    client, err := client.Dial(config)
    if err != nil {
        log.Fatal(err)
    }
    contractAddress := common.HexToAddress("contract addree in hex") // 0x0626918C51A1F36c7ad4354BB1197460A533a2B9
    instance, err := store.NewStore(contractAddress, client)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("contract is loaded")
    _ = instance
}
```

### 查询智能合约

在部署过程中设置的`Store.sol`合约中有一个名为`version`的全局变量。 因为它是公开的，这意味着它们将成为我们自动创建的`getter`函数。 常量和`view`函数也接受`bind.CallOpts`作为第一个参数，新建`contract_read.go`文件以查询合约：

```go
package main

import (
    "fmt"
    "log"
    "github.com/ethereum/go-ethereum/common"
    "github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
    "github.com/FISCO-BCOS/go-sdk/abi/bind"
    store "contract/testfile" // for demo
)

func main() {
    config := &conf.ParseConfig("config.toml")[0]
    client, err := client.Dial(config)
    if err != nil {
        log.Fatal(err)
    }

    // load the contract
    contractAddress := common.HexToAddress("contract addree in hex") // 0x0626918C51A1F36c7ad4354BB1197460A533a2B9
    instance, err := store.NewStore(contractAddress, client)
    if err != nil {
        log.Fatal(err)
    }

	storeSession := &StoreSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}

    version, err := storeSession.Version()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("version :", version) // "Store deployment 1.0"
}
```

### 写入智能合约

写入智能合约需要我们用私钥来对交易事务进行签名，我们创建的智能合约有一个名为`SetItem`的外部方法，它接受solidity`bytes32`类型的两个参数（key，value）。 这意味着在Go文件中需要传递一个长度为32个字节的字节数组。 调用`SetItem`方法需要我们传递我们之前创建的`auth`对象（keyed transactor）。 在幕后，此方法将使用它的参数对此函数调用进行编码，将其设置为事务的data属性，并使用私钥对其进行签名。 结果将是一个已签名的事务对象。新建`contract_write.go`来测试写入智能合约：

```go
package main

import (
    "fmt"
    "log"
    "context"
    "github.com/ethereum/go-ethereum/common"
    "github.com/FISCO-BCOS/go-sdk/client"
    "github.com/FISCO-BCOS/go-sdk/abi/bind"
    "github.com/ethereum/go-ethereum/crypto"
    store "contract/testfile" // for demo
)

func main() {
    config := &conf.ParseConfig("config.toml")[0]
    client, err := client.Dial(config)
    if err != nil {
        log.Fatal(err)
    }

    // load the contract
    contractAddress := common.HexToAddress("contract addree in hex") // 0x0626918C51A1F36c7ad4354BB1197460A533a2B9
    instance, err := store.NewStore(contractAddress, client)
    if err != nil {
        log.Fatal(err)
    }

	storeSession := &StoreSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}

    key := [32]byte{}
    value := [32]byte{}
    copy(key[:], []byte("foo"))
    copy(value[:], []byte("bar"))

    tx, err := storeSession.SetItem(key, value)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("tx sent: %s\n", tx.Hash().Hex())

    // wait for the mining
    receipt, err := client.WaitMined(tx)
    if err != nil {
        log.Fatalf("tx mining error:%v\n", err)
    }
    fmt.Printf("transaction hash of receipt: %s\n", receipt.GetTransactionHash())

    // read the result
    result, err := storeSession.Items(key)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(result[:])) // "bar"
}
```
