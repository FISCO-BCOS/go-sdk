module github.com/FISCO-BCOS/go-sdk

go 1.16

replace (
	golang.org/x/net => github.com/golang/net v0.0.0-20190415214537-1da14a5a36f2
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190412183630-56d357773e84
)

require (
	github.com/FISCO-BCOS/bcos-c-sdk v0.0.0-20230720102808-34f03aa15ebe
	github.com/FISCO-BCOS/crypto v0.0.0-20200202032121-bd8ab0b5d4f1
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/ethereum/go-ethereum v1.10.21
	github.com/go-kit/kit v0.9.0 // indirect
	github.com/go-ole/go-ole v1.2.5 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/urfave/cli/v2 v2.10.2
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
	golang.org/x/net v0.2.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
)
