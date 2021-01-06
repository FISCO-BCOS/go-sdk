package fiscobcosios

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type BuildSDKResult struct {
	ErrorId   int    `json:"errorId"`
	ErrorInfo string `json:"errorInfo"`
}

type ContractParams struct {
	ValueType string      `json:"type"`
	Value     interface{} `json:"value"`
}

type DeployContractResult struct {
	ErrorInfo   string       `json:"errorInfo"`
	Address     string       `json:"address"`
	Transaction *Transaction `json:"transaction"`
}

type SendTransactionReceipt struct {
	Receipt   string `json:"receipt"`
	ErrorInfo string `json:"errorInfo"`
}

type TxReceipt struct {
	TransactionHash  string     `json:"transactionHash"`
	TransactionIndex string     `json:"transactionIndex"`
	BlockHash        string     `json:"blockHash"`
	BlockNumber      string     `json:"blockNumber"`
	GasUsed          string     `json:"gasUsed"`
	ContractAddress  string     `json:"contractAddress"`
	Root             string     `json:"root"`
	Status           int        `json:"status"`
	From             string     `json:"from"`
	To               string     `json:"to"`
	Input            string     `json:"input"`
	Output           string     `json:"output"`
	Logs             []EventLog `json:"logs"`
	LogsBloom        string     `json:"logsBloom"`
}

type EventLog struct {
	Address string   `json:"address"`
	Data    string   `json:"data"`
	Topics  []string `json:"topics" `
}

type RPCResult struct {
	QueryResult string `json:"queryResult"`
	ErrorInfo   string `json:"errorInfo"`
}

type TransactResult struct {
	Transaction *Transaction
	Receipt     *TxReceipt
	ErrorInfo   string
}

type Transaction struct {
	Hash string
	Size float64
	Data string
}

type FullTransaction struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	From             string `json:"from"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Hash             string `json:"hash"`
	Input            string `json:"input"`
	Nonce            string `json:"nonce"`
	To               string `json:"to"`
	TransactionIndex string `json:"transactionIndex"`
	Value            string `json:"value"`
}

type RPCTransactionResult struct {
	Transaction FullTransaction `json:"transaction"`
	ErrorInfo   string          `json:"errorInfo"`
}

type CallResult struct {
	Result    string `json:"result"`
	ErrorInfo string `json:"errorInfo"`
}

var clientSdk *client.Client

// Build sdk.
// Connect to the proxy or FISCO BCOS node.
func BuildSDK(basePath string) string {
	// connect node
	configs, err := conf.ParseConfigFile(basePath + "config.toml")
	if err != nil {
		buildResult := BuildSDKResult{ErrorId: -1, ErrorInfo: err.Error()}
		ret, _ := json.Marshal(buildResult)
		return string(ret)
	}
	clientSdk, err = client.Dial(&configs[0])
	if err != nil {
		buildResult := BuildSDKResult{ErrorId: -1, ErrorInfo: err.Error()}
		ret, _ := json.Marshal(buildResult)
		return string(ret)
	}
	buildResult := BuildSDKResult{ErrorId: 0, ErrorInfo: ""}
	ret, _ := json.Marshal(buildResult)
	return string(ret)
}

// Build sdk.
// Connect to the proxy or FISCO BCOS node.
func BuildSDKWithParam(basePath string, groupId string, ipPort string, keyfile string) string {
	var config = "[Network]\n" +
		"Type=\"channel\"\n" +
		"CAFile=\"" + basePath + "/ca.crt\"\n" +
		"Cert=\"" + basePath + "/sdk.crt\"\n" +
		"Key=\"" + basePath + "/sdk.key\"\n" +
		"[[Network.Connection]]\n" +
		"NodeURL=\"" + ipPort + "\"\n" +
		"GroupID=" + groupId + "\n\n" +
		"[Account]\n" +
		"KeyFile= \"" + basePath + "/" + keyfile + "\"\n\n" +
		"[Chain]\n" +
		"ChainID=1\n" +
		"SMCrypto=false"
	// connect node
	configs, err := conf.ParseConfig([]byte(config))
	if err != nil {
		return err.Error()
	}
	clientSdk, err = client.Dial(&configs[0])
	if err != nil {
		return err.Error()
	}
	return "Connect Success to node " + ipPort
}

// Get client version
func GetClientVersion() *RPCResult {
	cv, err := clientSdk.GetClientVersion(context.Background())
	var rpcResult *RPCResult
	if err != nil {
		rpcResult = &RPCResult{QueryResult: "", ErrorInfo: err.Error()}
	} else {
		rpcResult = &RPCResult{QueryResult: string(cv), ErrorInfo: ""}
	}
	return rpcResult
}

// Deploy contract
func DeployContract(abiContract string, binContract string, params string) *DeployContractResult {
	ops := clientSdk.GetTransactOpts()
	parsed, err := abi.JSON(strings.NewReader(abiContract))
	if err != nil {
		return toDeployResult(common.Address{}, nil, err)
	}
	goParam, err := toGoParams(params)
	if err != nil {
		return toDeployResult(common.Address{}, nil, err)
	}
	address, transaction, _, err := bind.DeployContract(ops, parsed, common.FromHex(binContract), clientSdk, goParam...)
	return toDeployResult(address, transaction, err)
}

// Send transaction to call function of contract
func SendTransaction(abiContract string, address string, method string, params string) *TransactResult {
	parsed, err := abi.JSON(strings.NewReader(abiContract))
	if err != nil {
		return toTxResult(nil, nil, err)
	}
	goParams, err := toGoParams(params)
	if err != nil {
		fmt.Println("toGoParams error")
		return toTxResult(nil, nil, err)
	}
	addr := common.HexToAddress(address)
	c := bind.NewBoundContract(addr, parsed, clientSdk, clientSdk, clientSdk)
	fmt.Println("send transaction, address :", addr.Hex(), " method :", method, " params: ", goParams)
	transaction, receipt, err := c.Transact(clientSdk.GetTransactOpts(), method, goParams...)
	return toTxResult(transaction, receipt, err)
}

// Call contract
func Call(abiContract string, address string, method string, params string) *CallResult {
	parsed, err := abi.JSON(strings.NewReader(abiContract))
	if err != nil {
		return toCallResult(nil, err)
	}
	goParams, err := toGoParams(params)
	if err != nil {
		return toCallResult(nil, err)
	}
	addr := common.HexToAddress(address)
	c := bind.NewBoundContract(addr, parsed, clientSdk, clientSdk, clientSdk)

	fmt.Println("call, address :", addr.Hex(), " method :", method, " params: ", goParams)
	// override problem
	mtd := parsed.Methods[method]
	if len(mtd.Outputs) == 1 {
		result := getGoType(mtd.Outputs[0].Type)
		err = c.Call(clientSdk.GetCallOpts(), result, method, goParams...)
		return toCallResult(result, err)
	} else {
		var outputs []interface{}
		for _, one := range mtd.Outputs {
			outputs = append(outputs, getGoType(one.Type))
		}
		err = c.Call(clientSdk.GetCallOpts(), &outputs, method, goParams...)
		return toCallResult(&outputs, err)
	}
}

// Get transaction by tx hash
func GetTransactionByHash(txHash string) *RPCTransactionResult {
	raw, err := clientSdk.GetTransactionByHash(context.Background(), txHash)
	var tx FullTransaction
	if err != nil {
		return &RPCTransactionResult{Transaction: tx, ErrorInfo: err.Error()}
	} else {
		err = json.Unmarshal(raw, &tx)
		if err != nil {
			return &RPCTransactionResult{Transaction: tx, ErrorInfo: err.Error()}
		}
		return &RPCTransactionResult{Transaction: tx, ErrorInfo: ""}
	}
}

// Get transaction receipt by tx hash
func GetTransactionReceipt(txHash string) *RPCResult {
	receipt, err := clientSdk.GetTransactionReceipt(context.Background(), txHash)
	if err != nil {
		return &RPCResult{QueryResult: "", ErrorInfo: err.Error()}
	}
	rec, err := toReceipt(receipt)
	if err != nil {
		return &RPCResult{QueryResult: "", ErrorInfo: err.Error()}
	}
	str, err := json.Marshal(rec)
	if err != nil {
		return &RPCResult{QueryResult: "", ErrorInfo: err.Error()}
	}
	return &RPCResult{QueryResult: string(str), ErrorInfo: ""}
}

func toGoParams(param string) ([]interface{}, error) {
	var objs []ContractParams
	if err := json.Unmarshal([]byte(param), &objs); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	var par []interface{}

	for _, t := range objs {
		fmt.Println("--------------")
		fmt.Println("params type :", t.ValueType, "value :", t.Value)
		value, err := StringToInterface(t.ValueType, t.Value)
		if err != nil {
			return nil, err
		}
		par = append(par, value)
	}
	return par, nil
}

func toDeployResult(addr common.Address, transaction *types.Transaction, err error) *DeployContractResult {
	var txResult DeployContractResult
	if err != nil {
		txResult.ErrorInfo = err.Error()
	}
	if transaction != nil {
		txResult.Transaction = toTransact(transaction)
	}
	var emptyAddress = common.Address{}
	if addr != emptyAddress {
		txResult.Address = addr.Hex()
	}
	return &txResult
}

func toTxResult(transaction *types.Transaction, receipt *types.Receipt, err error) *TransactResult {
	var tr TransactResult
	if err != nil {
		tr.ErrorInfo = err.Error()
		return &tr
	}
	if receipt != nil {
		rec, err := toReceipt(receipt)
		tr.Receipt = rec
		if err != nil {
			tr.ErrorInfo = err.Error()
			return &tr
		}
	}
	if transaction != nil {
		tr.Transaction = toTransact(transaction)
	}
	return &tr
}

func toTransact(transaction *types.Transaction) *Transaction {
	var tx Transaction
	var emptyHash = common.Hash{}
	if transaction.Hash() != emptyHash {
		tx.Hash = transaction.Hash().Hex()
	}
	tx.Size = float64(transaction.Size())
	tx.Data = string(transaction.Data())
	return &tx
}

func toReceipt(_r *types.Receipt) (*TxReceipt, error) {
	if _r == nil {
		return nil, errors.New("receipt is null")
	}
	var rec TxReceipt
	rec.TransactionHash = _r.TransactionHash
	rec.TransactionIndex = _r.TransactionIndex
	rec.BlockHash = _r.BlockHash
	rec.BlockNumber = _r.BlockNumber
	rec.GasUsed = _r.GasUsed
	rec.ContractAddress = _r.ContractAddress.Hex()
	rec.Root = _r.Root
	rec.Status = _r.Status
	rec.From = _r.From
	rec.To = _r.To
	rec.Input = _r.Input
	rec.Output = _r.Output
	var logs []EventLog
	for _, one := range _r.Logs {
		topics, err := interfaceToString(one.Topics)
		var log EventLog
		if err != nil {
			return nil, err
		} else {
			log = EventLog{one.Address, one.Data, topics}
		}
		logs = append(logs, log)
	}
	rec.Logs = logs
	rec.LogsBloom = _r.LogsBloom
	return &rec, nil
}

func toCallResult(result interface{}, err error) *CallResult {
	if err != nil {
		return &CallResult{"", err.Error()}
	}
	resultBytes, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return &CallResult{"", err.Error()}
	}
	return &CallResult{string(resultBytes), ""}
}

// interface to string
func interfaceToString(param []interface{}) ([]string, error) {
	var str []string
	for _, p := range param {
		switch p.(type) {
		case string:
			str = append(str, p.(string))
		case int:
			str = append(str, strconv.FormatInt(int64(p.(int)), 10))
		case int8:
			str = append(str, strconv.FormatInt(int64(p.(int8)), 10))
		case int16:
			str = append(str, strconv.FormatInt(int64(p.(int16)), 10))
		case int32:
			str = append(str, strconv.FormatInt(int64(p.(int32)), 10))
		case int64:
			str = append(str, strconv.FormatInt(p.(int64), 10))
		case uint:
			str = append(str, strconv.FormatUint(p.(uint64), 10))
		case uint8:
			str = append(str, strconv.FormatUint(uint64(p.(uint8)), 10))
		case uint16:
			str = append(str, strconv.FormatUint(uint64(p.(uint16)), 10))
		case uint32:
			str = append(str, strconv.FormatUint(uint64(p.(uint32)), 10))
		case uint64:
			str = append(str, strconv.FormatUint(p.(uint64), 10))
		case bool:
			str = append(str, strconv.FormatBool(p.(bool)))
		case []byte:
			str = append(str, string(p.([]byte)))
		case common.Address:
			str = append(str, p.(common.Address).Hex())
		default:
			return nil, errors.New("unsupport interface type (" + reflect.TypeOf(p).String() + ")")
		}
	}
	return str, nil
}

// string to interface
func StringToInterface(paramType string, value interface{}) (interface{}, error) {
	if strings.Count(paramType, "[") != 0 {
		// split elements
		i := strings.LastIndex(paramType, "[")
		preType := paramType[:i]
		//string[] strs

		//strs := strings.Split(value[1:len(value)-1], ",")
		strs, ok := value.([]interface{})
		if !ok {
			return nil, errors.New("parse data to interface error")
		}

		// get type and construct an array
		obj, err := stringToInterfaceBasic(preType, strs[0])
		if err != nil {
			return nil, err
		}

		// construct array
		arrayType := reflect.ArrayOf(len(strs), reflect.TypeOf(obj))
		array := reflect.New(arrayType).Elem()
		for i, one := range strs {
			obj, err := stringToInterfaceBasic(preType, one)
			if err != nil {
				return nil, err
			}
			if array.Index(i).Kind() == reflect.Ptr {
				newObj := reflect.New(array.Index(i).Type().Elem())
				array.Index(i).Set(newObj)
				set(newObj, reflect.ValueOf(obj))
			} else {
				set(array.Index(i), reflect.ValueOf(obj))
			}
		}
		return array.Interface(), nil
	} else {
		return stringToInterfaceBasic(paramType, value)
	}
}

func TestTransfer(paramType string, value string) reflect.Type {
	fmt.Println("before ", paramType, value)
	v, e := StringToInterface(paramType, value)
	fmt.Println(e)
	//fmt.Println("v:")
	fmt.Println("$$$", reflect.TypeOf(v), v)
	return reflect.TypeOf(v)
}
func getType(paramType string, value string) (result reflect.Type) {
	a, _ := stringToInterfaceBasic(paramType, value)
	return reflect.TypeOf(a)
}

// Parse string params to go interface
func stringToInterfaceBasic(paramType string, value interface{}) (interface{}, error) {
	switch paramType {
	case "string":
		return value, nil
	case "int":
		return int(value.(float64)), nil
	case "int8":
		return int8(value.(float64)), nil
	case "int16":
		return int16(value.(float64)), nil
	case "int32":
		return int32(value.(float64)), nil
	case "int64":
		return int64(value.(float64)), nil
	case "int256":
		in, err := strconv.ParseInt(value.(string), 10, 64)
		if err != nil {
			return nil, err
		}
		return big.NewInt(in), nil
	case "uint":
		return uint(value.(float64)), nil
	case "uint8":
		return uint8(value.(float64)), nil
	case "uint16":
		return uint16(value.(float64)), nil
	case "uint32":
		return uint32(value.(float64)), nil
	case "uint64":
		return uint64(value.(float64)), nil
	case "uint256":
		in, err := strconv.ParseUint(value.(string), 10, 64)
		if err != nil {
			return nil, err
		}
		return big.NewInt(int64(in)), nil
	case "bool":
		return value.(bool), nil
	case "[]byte", "bytes":
		byteValue := common.FromHex(value.(string))
		result := make([]byte, len(byteValue))
		copy(result[:], byteValue)
		return result, nil
	case "bytes1", "bytes2", "bytes3", "bytes4", "bytes5", "bytes6", "bytes7", "bytes8", "bytes9", "bytes10", "bytes11", "bytes12", "bytes13", "bytes14", "bytes15", "bytes16", "bytes17", "bytes18", "bytes19", "bytes20", "bytes21", "bytes22", "bytes23", "bytes24", "bytes25", "bytes26", "bytes27", "bytes28", "bytes29", "bytes30", "bytes31", "bytes32":
		length, err := strconv.ParseInt(paramType[5:], 10, 8)
		if err != nil {
			return nil, err
		}
		byteValue := common.FromHex(value.(string))
		result := make([]byte, length)
		copy(result[:], byteValue)
		return mustByteSliceToArray(reflect.ValueOf(result)).Interface(), nil
	case "address":
		result := common.HexToAddress(value.(string))
		return result, nil
	default:
		err := fmt.Errorf("unsupport interface type (" + paramType + ")")
		return value, err
	}
}

// abi.typ to interface
func getGoType(kind abi.Type) interface{} {
	switch kind.T {
	case abi.AddressTy:
		var result *common.Address
		return result
	case abi.IntTy, abi.UintTy:
		parts := regexp.MustCompile(`(u)?int([0-9]*)`).FindStringSubmatch(kind.String())
		if parts[1] == "u" {
			switch parts[2] {
			case "8":
				return new(uint8)
			case "16":
				return new(uint16)
			case "32":
				return new(uint32)
			case "64":
				return new(uint64)
			case "256":
				return new(*big.Int)
			}
		} else {
			switch parts[2] {
			case "8":
				return new(int8)
			case "16":
				return new(int16)
			case "32":
				return new(int32)
			case "64":
				return new(int64)
			case "256":
				return new(*big.Int)
			}
		}
	case abi.FixedBytesTy:
		return new([]byte)
	case abi.BytesTy:
		return new([]byte)
	case abi.FunctionTy:
		return new([24]byte)
	case abi.BoolTy:
		return new(bool)
	case abi.StringTy:
		return new(string)
	case abi.HashTy:
		return new(common.Hash)
	default:
		return new(interface{})
	}
	return nil
}

func set(dst, src reflect.Value) error {
	dstType, srcType := dst.Type(), src.Type()
	fmt.Println("dstVal:  Kind = ", dstType.Kind(), " dst.CanSet() = ", dst.CanSet(), "     src: value =", src, " srcType.Kind() =", srcType.Kind(), " srcType.AssignableTo(dstType) = ", srcType.AssignableTo(dstType))
	switch {
	case dstType.Kind() == reflect.Interface && dst.Elem().IsValid():
		return set(dst.Elem(), src)
	case dstType.Kind() == reflect.Ptr && src.Kind() == reflect.Ptr:
		return set(dst.Elem(), src.Elem())
	case dstType.Kind() == reflect.Ptr:
		return set(dst.Elem(), src)
	case srcType.AssignableTo(dstType) && dst.CanSet():
		dst.Set(src)
	case dstType.Kind() == reflect.Slice && srcType.Kind() == reflect.Slice:
		return setSlice(dst, src)
	default:
		return fmt.Errorf("abi: cannot unmarshal %v in to %v", src.Type(), dst.Type())
	}
	return nil
}

func setSlice(dst, src reflect.Value) error {
	fmt.Println("????")
	slice := reflect.MakeSlice(dst.Type(), src.Len(), src.Len())
	for i := 0; i < src.Len(); i++ {
		v := src.Index(i)
		reflect.Copy(slice.Index(i), v)
	}
	dst.Set(slice)
	return nil
}
func mustByteSliceToArray(value reflect.Value) reflect.Value {
	arrayType := reflect.ArrayOf(value.Len(), reflect.TypeOf(uint8(0)))
	array := reflect.New(arrayType).Elem()
	for i := 0; i < value.Len(); i++ {
		array.Index(i).Set(value.Index(i))
	}
	return array
}
