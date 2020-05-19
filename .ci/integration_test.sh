# !/bin/bash

set -e

macOS=
GOPATH_BIN=$(go env GOPATH)/bin
SHELL_FOLDER=$(
    cd $(dirname $0)
    pwd
)

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
    eval ${command}
    ret=$?
    if [ $ret -ne 0 ]; then
        LOG_ERROR "FAILED of command: ${command}"
        exit 1
    else
        LOG_INFO "SUCCESS of command: ${command}"
    fi
}

check_env(){
    if [ "$(uname)" == "Darwin" ];then
        # export PATH="/usr/local/opt/openssl/bin:$PATH"
        macOS="macOS"
    fi
    go get -u github.com/sqs/goreturns
}

compile_and_ut()
{
    export GO111MODULE="on"
    execute_cmd "go build cmd/console.go"
    execute_cmd "go build -o abigen ./cmd/abigen/main.go"

    execute_cmd "go test -v ./smcrypto"
}

generate_main() {
    local struct="${1}"
    local output="${2}"
cat << EOF >> "${output}"

func main() {
	configs := conf.ParseConfig("config.toml")
	client, err := client.Dial(&configs[0])
	if err != nil {
		fmt.Printf("Dial Client failed, err:%v", err)
		return
	}
	address, tx, instance, err := Deploy${struct}(client.GetTransactOpts(), client)
	if err != nil {
		fmt.Printf("Deploy failed, err:%v", err)
		return
	}
	fmt.Println("contract address: ", address.Hex()) // the address should be saved
	fmt.Println("transaction hash: ", tx.Hash().Hex())
EOF
}

generate_hello() {
    local struct="${1}"
    local output="${2}"
    generate_main "${1}" "${2}"
cat << EOF >> "${output}"

	hello := &${struct}Session{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}
	ret, err := hello.Get()
	if err != nil {
		fmt.Printf("hello.Get() failed: %v", err)
		return
	}
	fmt.Printf("Get: %s\n", ret)
	_, err = hello.Set("fisco")
	if err != nil {
		fmt.Printf("hello.Set failed: %v", err)
		return
	}
	ret, err = hello.Get()
	if err != nil {
		fmt.Printf("hello.Get() failed: %v", err)
		return
	}
	fmt.Printf("Get: %s\n", ret)
}
EOF
    "${GOPATH_BIN}"/goreturns -w  "${output}"
}

generate_counter() {
    local struct="${1}"
    local output="${2}"
    generate_main "${1}" "${2}"
cat << EOF >> "${output}"

	counter := &${struct}Session{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}
	ret, err := counter.Get()
	if err != nil {
		fmt.Printf("counter.Get() failed: %v", err)
		return
	}
	fmt.Printf("Get: %d\n", ret)
	_, err = counter.Set(big.NewInt(111))
	if err != nil {
		fmt.Printf("counter.Set failed: %v", err)
		return
	}
	ret, err = counter.Get()
	if err != nil {
		fmt.Printf("counter.Get() failed: %v", err)
		return
	}
	if big.NewInt(111).Cmp(ret) != 0 {
		fmt.Printf("counter.Set() failed, expected 111 (got %d)", ret)
		return
	}
	fmt.Printf("Get: %s\n", ret)
	ret, err = counter.Version()
	if err != nil {
		fmt.Printf("counter.Version() failed: %v", err)
		return
	}
	if big.NewInt(0).Cmp(ret) != 0 {
		fmt.Printf("counter.Version() failed, expected 0 (got %d)", ret)
		return
	}
	_, err = counter.Add()
	if err != nil {
		fmt.Printf("counter.Add() failed: %v", err)
		return
	}
	ret, err = counter.Get()
	if err != nil {
		fmt.Printf("counter.Get() failed: %v", err)
		return
	}
	if big.NewInt(112).Cmp(ret) != 0 {
		fmt.Printf("counter.Add() failed, expected 111 (got %d)", ret)
		return
	}
}

EOF
    "${GOPATH_BIN}"/goreturns -w  "${output}"
}

get_build_chain()
{
    # latest_version=$(curl -s https://api.github.com/repos/FISCO-BCOS/FISCO-BCOS/releases | grep "\"v2\.[0-9]\.[0-9]\"" | sort -u | tail -n 1 | cut -d \" -f 4)
    # curl -LO https://github.com/FISCO-BCOS/FISCO-BCOS/releases/download/${latest_version}/build_chain.sh && chmod u+x build_chain.sh
    curl -LO https://raw.githubusercontent.com/FISCO-BCOS/FISCO-BCOS/master/tools/build_chain.sh && chmod u+x build_chain.sh
}

precompiled_test(){
# TODO: consensus test use getSealer first
    precompileds=(config cns crud permission)
    for pkg in ${precompileds[*]}; do
        if [ ! -z "$(go test -v ./precompiled/${pkg}| grep FAIL)" ];then
            LOG_ERROR "test precompiled/${pkg} failed"
            exit 1;
        fi
    done
}

integration_std()
{
    LOG_INFO "integration_std testing..."

    # abigen std
    execute_cmd "./abigen --bin .ci/hello/HelloWorld.bin --abi .ci/hello/HelloWorld.abi  --type Hello --pkg main --out=hello.go"
    generate_hello Hello hello.go
    execute_cmd "go build -o hello hello.go"
    execute_cmd "go build -o bn256 .ci/Precompiledbn256/bn256.go"
    LOG_INFO "generate hello.go and build hello done."

    bash build_chain.sh -l 127.0.0.1:4 -o nodes
    cp nodes/127.0.0.1/sdk/* ./
    bash nodes/127.0.0.1/start_all.sh
    if [ -z "$(./hello | grep address)" ];then LOG_ERROR "std deploy contract failed." && exit 1;fi
    if [ ! -z "$(./hello | grep failed)" ];then LOG_ERROR "call hello failed." && exit 1;fi
    # if [ ! -z "$(./bn256 | grep failed)" ];then ./bn256 && LOG_ERROR "call bn256 failed." && exit 1;fi
    precompiled_test

    execute_cmd "./abigen --bin .ci/counter/Counter.bin --abi .ci/counter/Counter.abi  --type Counter --pkg main --out=counter.go"
    generate_counter Counter counter.go
    execute_cmd "go build -o counter counter.go"
    if [ -z "$(./counter | grep address)" ];then LOG_ERROR "std deploy contract failed." && exit 1;fi
    if [ ! -z "$(./counter | grep failed)" ];then LOG_ERROR "call counter failed." && exit 1;fi

    bash nodes/127.0.0.1/stop_all.sh
    LOG_INFO "integration_std testing pass."

}

integration_gm()
{
    LOG_INFO "integration_gm testing..."

    # abigen gm
    execute_cmd "./abigen --bin .ci/hello/HelloWorld_gm.bin --abi .ci/hello/HelloWorld.abi  --type Hello --pkg main --out=hello_gm.go --smcrypto=true"
    generate_hello Hello hello_gm.go
    execute_cmd "go build -o hello_gm hello_gm.go"
    execute_cmd "go build -o bn256_gm .ci/Precompiledbn256/bn256_gm.go"
    LOG_INFO "generate hello_gm.go and build hello_gm done."

    bash build_chain.sh -l 127.0.0.1:4 -g -o nodes_gm
    cp nodes_gm/127.0.0.1/sdk/* ./
    bash nodes_gm/127.0.0.1/start_all.sh
    sed -i "s/SMCrypto=false/SMCrypto=true/g" config.toml
    sed -i "s#KeyFile=\".ci/0x83309d045a19c44dc3722d15a6abd472f95866ac.pem\"#KeyFile=\".ci/sm2p256v1_0x791a0073e6dfd9dc5e5061aebc43ab4f7aa4ae8b.pem\"#g" config.toml
    if [ -z "$(./hello_gm | grep address)" ];then LOG_ERROR "gm deploy contract failed." && exit 1;fi
    if [ ! -z "$(./hello_gm | grep failed)" ];then LOG_ERROR "gm call hello_gm failed." && exit 1;fi
    # if [ ! -z "$(./bn256_gm | grep failed)" ];then ./bn256_gm && LOG_ERROR "gm call bn256_gm failed." && exit 1;fi
    # precompiled_test
    bash nodes_gm/127.0.0.1/stop_all.sh
    LOG_INFO "integration_gm testing pass."

}

check_env
compile_and_ut
get_build_chain
integration_std
if [ -z "${macOS}" ];then integration_gm ; fi
