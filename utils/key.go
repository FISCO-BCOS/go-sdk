package utils

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"github.com/liuxinfeng96/bc-crypto/ecdsa"
	bcx509 "github.com/liuxinfeng96/bc-crypto/x509"
)

// GenerateSecp256k1Key
// generate the private key of the Secp256k1 algorithm
// return the key byte array with PEM
func GenerateSecp256k1Key() ([]byte, error) {
	key, err := bcx509.GenerateKey(bcx509.EC_Secp256k1)
	if err != nil {
		return nil, err
	}

	ecKey, ok := key.(*ecdsa.PrivateKey)
	if !ok {
		return nil, errors.New("the wrong type of key")
	}

	skDer, err := bcx509.MarshalECPrivateKey(ecKey)
	if err != nil {
		return nil, err
	}

	skBlock := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: skDer,
	}

	skBuf := new(bytes.Buffer)
	if err = pem.Encode(skBuf, skBlock); err != nil {
		return nil, err
	}

	return skBuf.Bytes(), nil
}

// GenerateSm2Key
// generate the private key of the SM2 algorithm
// return the key byte array with PEM
func GenerateSm2Key() ([]byte, error) {
	key, err := bcx509.GenerateKey(bcx509.EC_SM2)
	if err != nil {
		return nil, err
	}

	ecKey, ok := key.(*ecdsa.PrivateKey)
	if !ok {
		return nil, errors.New("the wrong type of key")
	}

	skDer, err := bcx509.MarshalECPrivateKey(ecKey)
	if err != nil {
		return nil, err
	}

	skBlock := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: skDer,
	}

	skBuf := new(bytes.Buffer)
	if err = pem.Encode(skBuf, skBlock); err != nil {
		return nil, err
	}

	return skBuf.Bytes(), nil
}

// GenerateRSA2048Key
// generate the private key of the RSA2048 algorithm
// return the key byte array with PEM
func GenerateRSA2048Key() ([]byte, error) {
	key, err := bcx509.GenerateKey(bcx509.RSA2048)
	if err != nil {
		return nil, err
	}

	rsaKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("the wrong type of key")
	}

	skDer, err := x509.MarshalPKCS8PrivateKey(rsaKey)
	if err != nil {
		return nil, err
	}

	skBlock := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: skDer,
	}

	skBuf := new(bytes.Buffer)
	if err = pem.Encode(skBuf, skBlock); err != nil {
		return nil, err
	}

	return skBuf.Bytes(), nil
}
