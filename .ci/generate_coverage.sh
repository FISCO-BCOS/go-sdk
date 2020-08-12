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

calculate_coverage() {
    # start blockchain demo
    curl -LO https://raw.githubusercontent.com/FISCO-BCOS/FISCO-BCOS/master/tools/build_chain.sh && chmod u+x build_chain.sh
    bash build_chain.sh -v 2.5.0 -l 127.0.0.1:4 -o nodes
    cp nodes/127.0.0.1/sdk/* ./
    bash nodes/127.0.0.1/start_all.sh

    # generate code coverage report
    go test ./... -race -coverprofile=coverage.txt -covermode=atomic
}

calculate_coverage