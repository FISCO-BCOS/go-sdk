package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateKey(t *testing.T) {
	secp256k1Key, err := GenerateSecp256k1Key()
	require.Nil(t, err)

	fmt.Print(string(secp256k1Key))
	fmt.Println("---------------------------测试分割线--------------------------------")

	sm2Key, err := GenerateSm2Key()
	require.Nil(t, err)

	fmt.Print(string(sm2Key))
	fmt.Println("---------------------------测试分割线--------------------------------")

	rsa2048Key, err := GenerateRSA2048Key()
	require.Nil(t, err)

	fmt.Print(string(rsa2048Key))
	fmt.Println("---------------------------测试分割线--------------------------------")

}
