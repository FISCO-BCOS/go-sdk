package conf

import (
	"testing"
)

const (
	standardKeyHex        = "b89d42f12290070f235fb8fb61dcf96e3b11516c5d4f6333f26e49bb955f8b62"
	standardCurve         = "secp256k1"
	standardFileContent   = "-----BEGIN PRIVATE KEY-----\nMIGEAgEAMBAGByqGSM49AgEGBSuBBAAKBG0wawIBAQQguJ1C8SKQBw8jX7j7Ydz5\nbjsRUWxdT2Mz8m5Ju5Vfi2KhRANCAAQf+Lvzi2JmjS3R4rkA9+O4aVk/db7Hc7H+\nuzWl00qEH2Esk9fGeiCLuuQKGX3+TspTTvWA99FFi67RXgF+Sj15\n-----END PRIVATE KEY-----\n"
	standardSMKeyHex      = "389bb3e29db735b5dc4f114923f1ac5136891efda282a18dc0768e34305c861b"
	standardSMCurve       = "sm2p256v1"
	standardSMFileContent = "-----BEGIN PRIVATE KEY-----\nMIGHAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBG0wawIBAQQgOJuz4p23NbXcTxFJ\nI/GsUTaJHv2igqGNwHaONDBchhuhRANCAARPdpGofuy2PV77J2ydxWYBWu6p7Rl+\n6ExUKuJsBgI4p5/WkELL1+wtwvwQDI/0x2QqBarnAHp5FLxjFQqbVjVo\n-----END PRIVATE KEY-----\n"
)

func TestParsepem(t *testing.T) {
	// test nosm private
	keyHex, curve, fileContent, err := LoadECPrivateKeyFromPEM("../.ci/0x83309d045a19c44dc3722d15a6abd472f95866ac.pem")
	if err != nil {
		t.Fatalf("parse nosm private key failed, err: %v", err)
	}
	if keyHex != standardKeyHex {
		t.Fatalf("nosm keyHex is inconsistent with the anticipation, keyHex: %s", keyHex)
	}
	if curve != standardCurve {
		t.Fatalf("nosm curve is inconsistent with the anticipation, curve: %s", curve)
	}
	if fileContent != standardFileContent {
		t.Fatalf("nosm fileContent is inconsistent with the anticipation, fileContent:\n %s", fileContent)
	}
	t.Logf("the output of parsing nosm private key, keyHex: %s\n curve: %s\n fileContent:\n %s", keyHex, curve, fileContent)

	// test sm private
	keyHex, curve, fileContent, err = LoadECPrivateKeyFromPEM("../.ci/sm2p256v1_0x791a0073e6dfd9dc5e5061aebc43ab4f7aa4ae8b.pem")
	if err != nil {
		t.Fatalf("parse sm private key failed, err: %v", err)
	}
	if keyHex != standardSMKeyHex {
		t.Fatalf("sm keyHex is inconsistent with the anticipation, keyHex: %s", keyHex)
	}
	if curve != standardSMCurve {
		t.Fatalf("sm curve is inconsistent with the anticipation, curve: %s", curve)
	}
	if fileContent != standardSMFileContent {
		t.Fatalf("sm fileContent is inconsistent with the anticipation, fileContent:\n %s", fileContent)
	}
	t.Logf("the output of parsing sm private key, keyHex: %s\n curve: %s\n fileContent:\n %s", keyHex, curve, fileContent)

}
