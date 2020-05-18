package smcrypto

import (
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"

	"github.com/FISCO-BCOS/crypto/ecdsa"
	"github.com/FISCO-BCOS/crypto/elliptic"
	"github.com/FISCO-BCOS/crypto/x509"
	"github.com/FISCO-BCOS/go-sdk/smcrypto/sm3"
	"github.com/ethereum/go-ethereum/common"
)

const publicKeyLength = 64

type sm2Signature struct {
	R *big.Int
	S *big.Int
}

// ECDSAPubBytes return esdsa public key as slice
func ECDSAPubBytes(pub *ecdsa.PublicKey) []byte {
	if pub == nil || pub.X == nil || pub.Y == nil {
		return nil
	}
	pubBytes := make([]byte, publicKeyLength)
	copy(pubBytes[:], pub.X.Bytes())
	copy(pubBytes[publicKeyLength/2:], pub.Y.Bytes())
	return pubBytes
}

// PubkeyToAddress calculate address from sm2p256v1 private key
func PubkeyToAddress(p ecdsa.PublicKey) common.Address {
	pubBytes := ECDSAPubBytes(&p)
	sm3digest := sm3.Hash(pubBytes)
	return common.BytesToAddress(sm3digest[12:])
}

// HexKeyToAddress calculate address from sm2p256v1 private key
func HexKeyToAddress(hexKey string) common.Address {
	key, _ := HexToECDSA(hexKey)
	pubBytes := ECDSAPubBytes(&key.PublicKey)
	sm3digest := sm3.Hash(pubBytes)
	return common.BytesToAddress(sm3digest[12:])
}

// toECDSA creates a private key with the given D value.
func toECDSA(d []byte) (*ecdsa.PrivateKey, error) {
	priv := new(ecdsa.PrivateKey)
	curve := elliptic.Sm2p256v1()
	curveOrder := curve.Params().N
	k := new(big.Int).SetBytes(d)
	if k.Cmp(curveOrder) >= 0 {
		return nil, errors.New("x509: invalid elliptic curve private key value")
	}
	priv.D = k
	priv.Curve = curve

	privateKey := make([]byte, (curveOrder.BitLen()+7)/8)

	// Some private keys have leading zero padding. This is invalid
	// according to [SEC1], but this code will ignore it.
	for len(d) > len(privateKey) {
		if d[0] != 0 {
			return nil, errors.New("x509: invalid private key length")
		}
		d = d[1:]
	}

	// Some private keys remove all leading zeros, this is also invalid
	// according to [SEC1] but since OpenSSL used to do this, we ignore
	// this too.
	copy(privateKey[len(privateKey)-len(d):], d)
	priv.X, priv.Y = curve.ScalarBaseMult(privateKey)
	if priv.PublicKey.X == nil {
		return nil, errors.New("invalid private key")
	}
	return priv, nil
}

// HexToECDSA parses a secp256k1 private key.
func HexToECDSA(hexkey string) (*ecdsa.PrivateKey, error) {
	b, err := hex.DecodeString(hexkey)
	if err != nil {
		return nil, errors.New("invalid hex string")
	}
	return toECDSA(b)
}

// HexToPem parses a secp256k1 private key.
func HexToPem(hexkey string) (string, error) {
	key, err := HexToECDSA(hexkey)
	if err != nil {
		return "", err
	}
	return ECDSAToPem(key)
}

// ECDSAToPem parses a secp256k1 private key.
func ECDSAToPem(key *ecdsa.PrivateKey) (string, error) {
	b, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		return "", fmt.Errorf("encode private key failed %w", err)
	}
	pemdata := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: b,
		},
	)
	return string(pemdata), nil
}

// Sign calculates an sm2 signature.
//
// This function is susceptible to chosen plaintext attacks that can leak
// information about the private key that is used for signing. Callers must
// be aware that the given hash cannot be chosen by an adversery. Common
// solution is to hash any input before calculating the signature.
//
// The produced signature is in the [R || S || V] format where V is public key.
func Sign(hash []byte, hexKey string) (sig []byte, err error) {
	if len(hash) != 32 {
		return nil, fmt.Errorf("hash is required to be exactly 32 bytes (%d)", len(hash))
	}
	if len(hexKey) < 64 {
		return nil, fmt.Errorf("hex private key is required to be exactly 64 bytes (%d)", len(hexKey))
	}
	key, err := HexToECDSA(hexKey)
	if err != nil {
		return nil, err
	}
	pubBytes := ECDSAPubBytes(&key.PublicKey)

	r, s, err := SM2Sign(hash, key)
	sig = r.Bytes()
	sig = append(sig, s.Bytes()...)
	sig = append(sig, pubBytes...)

	return sig, nil
}
