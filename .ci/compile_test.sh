# !/bin/bash

set -e

SHELL_FOLDER=$(
    cd $(dirname $0)
    pwd
)

LOG_ERROR() {
    content=${1}
    echo -e "\033[31m"${content}"\033[0m"
}

LOG_INFO() {
    content=${1}
    echo -e "\033[32m"${content}"\033[0m"
}

execute_cmd() {
    command="${1}"
    eval ${command}
    ret=$?
    if [ $ret -ne 0 ]; then
        LOG_ERROR "FAILED of command: ${command}"
        exit 1
    else
        LOG_INFO "SUCCESS of command: ${command}"
    fi
}

execute_cmd "go build console.go"
execute_cmd "go build -o abigen cmd/abigen/main.go"
execute_cmd "go test -v ./crypto ./crypto/smcrypto"

