#!/bin/bash

GOPATH_BIN=$(go env GOPATH)/bin

LOG_ERROR() {
    content=${1}
    echo -e "\033[31m${content}\033[0m"
}

LOG_INFO() {
    content=${1}
    echo -e "\033[32m${content}\033[0m"
}

execute_cmd() {
    command="${1}"
    eval "${command}"
    ret=$?
    if [ $ret -ne 0 ]; then
        LOG_ERROR "FAILED of command: ${command}"
        exit 1
    else
        LOG_INFO "SUCCESS of command: ${command}"
    fi
}

generate_hello() {
    local struct="${1}"
    local output="${2}"
    go get -u github.com/sqs/goreturns   # to format code style
cat << EOF >> "${output}"

func main() {
	configs := conf.ParseConfig("config.toml")
	client, err := client.Dial(&configs[0])
	if err != nil {
		fmt.Printf("Dial Client failed, err:%v", err)
		return
	}
	address, tx, _, err := Deploy${struct}(client.GetTransactOpts(), client)
	if err != nil {
		fmt.Printf("Deploy failed, err:%v", err)
		return
	}
	fmt.Println(address.Hex()) // the address should be saved
	fmt.Println(tx.Hash().Hex())
}
EOF
    "${GOPATH_BIN}"/goreturns -w  "${output}"
}

calculate_coverage() {
    # start blockchain demo
    curl -LO https://raw.githubusercontent.com/FISCO-BCOS/FISCO-BCOS/master/tools/build_chain.sh && chmod u+x build_chain.sh
    bash build_chain.sh -l 127.0.0.1:4 -o nodes
    cp nodes/127.0.0.1/sdk/* ./
    bash nodes/127.0.0.1/start_all.sh

    # generate code coverage report
    go test ./... -race -coverprofile=coverage.txt -covermode=atomic
}

calculate_coverage