package conf

import (
	"encoding/json"
	"testing"
)

const (
	standardJSON = "[{\"IsHTTP\":false,\"ChainID\":1,\"CAFile\":\"ca.crt\",\"Key\":\"sdk.key\",\"Cert\":\"sdk.crt\",\"IsSMCrypto\":false,\"PrivateKey\":\"b89d42f12290070f235fb8fb61dcf96e3b11516c5d4f6333f26e49bb955f8b62\",\"GroupID\":1,\"NodeURL\":\"127.0.0.1:20200\"}]"
	fileContent  = "[Network]\n#type rpc or channel\nType=\"channel\"\nCAFile=\"ca.crt\"\nCert=\"sdk.crt\"\nKey=\"sdk.key\"\n[[Network.Connection]]\nNodeURL=\"127.0.0.1:20200\"\nGroupID=1\n# [[Network.Connection]]\n# NodeURL=\"127.0.0.1:20200\"\n# GroupID=2\n\n[Account]\n# only support PEM format for now\nKeyFile=\"../.ci/0x83309d045a19c44dc3722d15a6abd472f95866ac.pem\"\n\n[Chain]\nChainID=1\nSMCrypto=false\n\n[log]\nPath=\"./\""
)

func TestConfig(t *testing.T) {
	// test parseConfig
	configs, err := ParseConfig([]byte(fileContent))
	if err != nil {
		t.Fatalf("TestConfig failed, err: %v", err)
	}
	jsons, err := json.Marshal(configs)
	if err != nil {
		t.Fatalf("failed when struct transfers to json, err: %v", err)
	}
	if standardJSON != string(jsons) {
		t.Fatalf("parsing the output of test.toml is inconsistent with the standardJson\n the output of test.toml: %v\n standardJson: %v", string(jsons), standardJSON)
	}
	t.Logf("the output of test.toml: %v", string(jsons))
}
