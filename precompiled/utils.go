package precompiled

import (
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/core/types"

	"github.com/FISCO-BCOS/go-sdk/abi"
)

func GetPreServiceOutput(abiStr string, name string, receipt *types.Receipt) (int, error) {
	parsed, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		fmt.Printf("parse ABI failed, err: %v", err)
	}
	var ret *big.Int
	if len(receipt.Output) < 2 {
		return types.GoErrorCode, fmt.Errorf("the leghth of receipt.Output %v is less than 2", receipt.Output)
	}
	b, err := hex.DecodeString(receipt.Output[2:])
	if err != nil {
		return types.GoErrorCode, fmt.Errorf("decode receipt.Output[2:] failed, err: %v", err)
	}
	err = parsed.Unpack(&ret, name, b)
	if err != nil {
		return types.GoErrorCode, fmt.Errorf("unpack %v failed, err: %v", name, err)
	}
	return types.GetServiceOutputCode(ret)
}

func BigIntToUint64(bigNum *big.Int) (uint64, error) {
	boolean := bigNum.IsUint64()
	if !boolean {
		return 0, fmt.Errorf("bigNum %v can't transfer to Uint64", bigNum)
	}
	return bigNum.Uint64(), nil
}

func Uint64ToInt64(num uint64) (int64, error) {
	if num > math.MaxInt64 {
		return 0, fmt.Errorf("uint64 %v can't tranfer to int64", num)
	}
	return int64(num), nil
}
