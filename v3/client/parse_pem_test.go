package client

import (
	"fmt"
	"os"
	"testing"
)

const (
	standardKeyHex   = "b89d42f12290070f235fb8fb61dcf96e3b11516c5d4f6333f26e49bb955f8b62"
	standardCurve    = "secp256k1"
	standardSMKeyHex = "389bb3e29db735b5dc4f114923f1ac5136891efda282a18dc0768e34305c861b"
	standardSMCurve  = "sm2p256v1"
	pemFileContent   = `-----BEGIN PRIVATE KEY-----
MIGEAgEAMBAGByqGSM49AgEGBSuBBAAKBG0wawIBAQQguJ1C8SKQBw8jX7j7Ydz5
bjsRUWxdT2Mz8m5Ju5Vfi2KhRANCAAQf+Lvzi2JmjS3R4rkA9+O4aVk/db7Hc7H+
uzWl00qEH2Esk9fGeiCLuuQKGX3+TspTTvWA99FFi67RXgF+Sj15
-----END PRIVATE KEY-----`
	secp256k1PemPath     = "./0x83309d045a19c44dc3722d15a6abd472f95866ac.pem"
	sm2p256v1PemPath     = "./sm2p256v1_0x791a0073e6dfd9dc5e5061aebc43ab4f7aa4ae8b.pem"
	sm2p256v1FileContent = `-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBG0wawIBAQQgOJuz4p23NbXcTxFJ
I/GsUTaJHv2igqGNwHaONDBchhuhRANCAARPdpGofuy2PV77J2ydxWYBWu6p7Rl+
6ExUKuJsBgI4p5/WkELL1+wtwvwQDI/0x2QqBarnAHp5FLxjFQqbVjVo
-----END PRIVATE KEY-----`
)

func preparePem(t *testing.T) {
	fmt.Println("init")
	os.WriteFile(secp256k1PemPath, []byte(pemFileContent), 0666)
	os.WriteFile(sm2p256v1PemPath, []byte(sm2p256v1FileContent), 0666)
	t.Cleanup(func() {
		os.Remove(secp256k1PemPath)
		os.Remove(sm2p256v1PemPath)
	})
}

func TestParsePem(t *testing.T) {
	preparePem(t)
	// test nosm private
	keyBytes, curve, err := LoadECPrivateKeyFromPEM(secp256k1PemPath)
	if err != nil {
		t.Fatalf("parse nosm private key failed, err: %v", err)
	}
	if fmt.Sprintf("%064x", keyBytes) != standardKeyHex {
		t.Fatalf("nosm keyHex is inconsistent with the anticipation, keyHex: %s", fmt.Sprintf("%064x", keyBytes))
	}
	if curve != standardCurve {
		t.Fatalf("nosm curve is inconsistent with the anticipation, curve: %s", curve)
	}
	t.Logf("the output of parsing nosm private key, keyHex: %s\n curve: %s\n fileContent:\n %s", fmt.Sprintf("%064x", keyBytes), curve, fileContent)

	// test sm private
	keyBytes, curve, err = LoadECPrivateKeyFromPEM(sm2p256v1PemPath)
	if err != nil {
		t.Fatalf("parse sm private key failed, err: %v", err)
	}
	if fmt.Sprintf("%064x", keyBytes) != standardSMKeyHex {
		t.Fatalf("sm keyHex is inconsistent with the anticipation, keyHex: %s", fmt.Sprintf("%064x", keyBytes))
	}
	if curve != standardSMCurve {
		t.Fatalf("sm curve is inconsistent with the anticipation, curve: %s", curve)
	}
	t.Logf("the output of parsing sm private key, keyHex: %s\n curve: %s\n fileContent:\n %s", fmt.Sprintf("%064x", keyBytes), curve, fileContent)

}
