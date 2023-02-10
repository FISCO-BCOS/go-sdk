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

get_csdk_lib()
{
	if [ ! -d "/usr/local/lib/" ];then
    	sudo mkdir -p /usr/local/lib
	fi
	local suffix="so"
	if [ "$(uname)" == "Darwin" ];then # macOS
		suffix="dylib"
	fi
	if [ ! -f "/usr/local/lib/libbcos-c-sdk.${suffix}" ];then
		curl -#LO "https://github.com/FISCO-BCOS/bcos-c-sdk/releases/download/v3.2.0/libbcos-c-sdk.${suffix}"
		sudo cp "libbcos-c-sdk.${suffix}" /usr/local/lib/
	fi
    export GODEBUG=cgocheck=0
}

calculate_coverage() {
    # start blockchain demo
    latest_version=$(curl -sS https://gitee.com/api/v5/repos/FISCO-BCOS/FISCO-BCOS/tags | grep -oe "\"name\":\"v[3-9]*\.[0-9]*\.[0-9]*\"" | cut -d \" -f 4 | sort -V | tail -n 1)
    curl -#LO https://github.com/FISCO-BCOS/FISCO-BCOS/releases/download/"${latest_version}"/build_chain.sh && chmod u+x build_chain.sh
    bash build_chain.sh -l 127.0.0.1:4 -o nodes
    get_csdk_lib
    cp nodes/127.0.0.1/sdk/* ./
    cp nodes/127.0.0.1/sdk/* ./client/
    bash nodes/127.0.0.1/start_all.sh

    # generate code coverage report
    go test -ldflags="-r /usr/local/lib/" ./client -race -coverprofile=coverage.txt -covermode=atomic
}

calculate_coverage