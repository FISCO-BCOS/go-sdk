package precompiled

import (
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/core/types"
)

const (
	// common precompiled contract error code
	permissionDenied  int64 = -50000
	tableAlreadyExist int64 = -50001

	// non-precompiled contract error code, only used when logic errors occur
	DefaultErrorCode int64 = -1

	bigIntHexStrLen int = 66
)

// GetCommonErrorCodeMessage returns the message of common error code
func GetCommonErrorCodeMessage(errorCode int64) string {
	var message string
	switch errorCode {
	case permissionDenied:
		message = "permission denied"
	case tableAlreadyExist:
		message = "table already exist"
	default:
		message = ""
	}
	return message
}

func ParseBigIntFromOutput(abiStr, name string, receipt *types.Receipt) (*big.Int, error) {
	parsed, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		fmt.Printf("parse ABI failed, err: %v", err)
	}
	var ret *big.Int
	if len(receipt.Output) != bigIntHexStrLen {
		return nil, fmt.Errorf("the leghth of receipt.Output %v is inconsistent with %v", len(receipt.Output), bigIntHexStrLen)
	}
	b, err := hex.DecodeString(receipt.Output[2:])
	if err != nil {
		return nil, fmt.Errorf("decode receipt.Output[2:] failed, err: %v", err)
	}
	err = parsed.Unpack(&ret, name, b)
	if err != nil {
		return nil, fmt.Errorf("unpack %v failed, err: %v", name, err)
	}
	return ret, nil
}

func BigIntToUint64(bigNum *big.Int) (uint64, error) {
	boolean := bigNum.IsUint64()
	if !boolean {
		return 0, fmt.Errorf("bigNum %v can't transfer to Uint64", bigNum)
	}
	return bigNum.Uint64(), nil
}

func BigIntToInt64(bigNum *big.Int) (int64, error) {
	boolean := bigNum.IsInt64()
	if !boolean {
		return 0, fmt.Errorf("BigIntToint64 failed, bigNum %v can't transfer to Int64", bigNum)
	}
	return bigNum.Int64(), nil
}

func Uint64ToInt64(num uint64) (int64, error) {
	if num > math.MaxInt64 {
		return 0, fmt.Errorf("uint64 %v can't tranfer to int64", num)
	}
	return int64(num), nil
}
