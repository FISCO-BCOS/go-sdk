// Package conf parse config to configuration
package conf

import (
	"fmt"
)

// Config contains configuration items for sdk
type Config struct {
	ConfigFile     string
	CAFile         string
	TLSCAContext   []byte
	TLSKeyFile     string
	TLSKeyContext  []byte
	TLSCertFile    string
	TLSCertContext []byte
	IsSMCrypto     bool
	PrivateKey     []byte
	GroupID        string
	NodeURL        string
	Host           string
	Port           int
}

// ParseConfigOptions parses from arguments
func ParseConfigOptions(caFile string, key string, cert, keyFile string, groupId string, ipPort string, isSMCrypto bool) (*Config, error) {
	config := Config{
		CAFile:      caFile,
		TLSKeyFile:  key,
		TLSCertFile: cert,
		IsSMCrypto:  isSMCrypto,
		GroupID:     groupId,
		NodeURL:     ipPort,
	}
	keyBytes, curve, err := LoadECPrivateKeyFromPEM(keyFile)
	if err != nil {
		return nil, fmt.Errorf("parse private key failed, err: %v", err)
	}
	if config.IsSMCrypto && curve != sm2p256v1 {
		return nil, fmt.Errorf("smcrypto must use sm2p256v1 private key, but found %s", curve)
	}
	if !config.IsSMCrypto && curve != secp256k1 {
		return nil, fmt.Errorf("must use secp256k1 private key, but found %s", curve)
	}
	config.PrivateKey = keyBytes
	return &config, nil
}
