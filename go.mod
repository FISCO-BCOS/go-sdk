module github.com/FISCO-BCOS/go-sdk

go 1.21.5

replace (
	// github.com/FISCO-BCOS/bcos-c-sdk v0.0.0-20231221132830-c0cfb1d98eec => ../bcos-c-sdk
	golang.org/x/net => github.com/golang/net v0.0.0-20190415214537-1da14a5a36f2
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190412183630-56d357773e84
)

require (
	github.com/FISCO-BCOS/crypto v0.0.0-20200202032121-bd8ab0b5d4f1
	github.com/ethereum/go-ethereum v1.10.21
	github.com/sirupsen/logrus v1.4.2
)

require (
	github.com/konsorten/go-windows-terminal-sequences v1.0.1 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
)
