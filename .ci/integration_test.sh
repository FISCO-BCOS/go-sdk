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
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		log.Fatalf("ParseConfigFile failed, err: %v", err)
	}
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
	_, _, err = hello.Set("fisco")
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
	_, _, err = counter.Set(big.NewInt(111))
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
	_, _, err = counter.Add()
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
    latest_version=$(curl -sS https://gitee.com/api/v5/repos/FISCO-BCOS/FISCO-BCOS/tags | grep -oe "\"name\":\"v[2-9]*\.[0-9]*\.[0-9]*\"" | cut -d \" -f 4 | sort -V | tail -n 1)
    curl -#LO https://github.com/FISCO-BCOS/FISCO-BCOS/releases/download/"${latest_version}"/build_chain.sh && chmod u+x build_chain.sh
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
    execute_cmd "bash tools/download_solc.sh -v 0.4.25"

    # abigen std
    execute_cmd "./solc-0.4.25 --bin --abi -o .ci/hello .ci/hello/HelloWorld.sol"
    execute_cmd "./abigen --bin .ci/hello/HelloWorld.bin --abi .ci/hello/HelloWorld.abi  --type Hello --pkg main --out=hello.go"
    generate_hello Hello hello.go
    execute_cmd "go build -o hello hello.go"
    execute_cmd "go build -o bn256 .ci/ethPrecompiled/bn256.go"
    LOG_INFO "generate hello.go and build hello done."

    bash build_chain.sh -v "${latest_version}" -l 127.0.0.1:4 -o nodes
    cp nodes/127.0.0.1/sdk/* ./
    bash nodes/127.0.0.1/start_all.sh
    if [ -z "$(./hello | grep address)" ];then LOG_ERROR "std deploy contract failed." && exit 1;fi
    if [ ! -z "$(./hello | grep failed)" ];then LOG_ERROR "call hello failed." && exit 1;fi
    # if [ ! -z "$(./bn256 | grep failed)" ];then ./bn256 && LOG_ERROR "call bn256 failed." && exit 1;fi
    precompiled_test

    execute_cmd "./solc-0.4.25 --bin --abi -o .ci/counter .ci/counter/Counter.sol"
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
    execute_cmd "bash tools/download_solc.sh -v 0.4.25 -g"

    # abigen gm
    execute_cmd "./solc-0.4.25-gm --bin --abi  --overwrite -o .ci/hello .ci/hello/HelloWorld.sol"
    execute_cmd "./abigen --bin .ci/hello/HelloWorld.bin --abi .ci/hello/HelloWorld.abi --type Hello --pkg main --out=hello_gm.go --smcrypto=true"
    generate_hello Hello hello_gm.go
    execute_cmd "go build -o hello_gm hello_gm.go"
    execute_cmd "go build -o bn256_gm .ci/ethPrecompiled/bn256_gm.go"
    LOG_INFO "generate hello_gm.go and build hello_gm done."

    bash build_chain.sh -v "${latest_version}" -l 127.0.0.1:4 -g -o nodes_gm
    cp -r nodes_gm/127.0.0.1/sdk/* ./
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

integration_amop() {
    LOG_INFO "integration_amop testing..."
    bash nodes/127.0.0.1/start_all.sh

    execute_cmd "go build -o subscriber examples/amop/sub/subscriber.go"
    execute_cmd "go build -o unicast_publisher examples/amop/unicast_pub/publisher.go"
    nohup ./unicast_publisher 127.0.0.1:20200 hello > output.file 2>&1 &
    nohup ./subscriber 127.0.0.1:20201 hello > subscriber0.out 2>&1 &
    sleep 13s
    cat subscriber0.out
    if ! grep "hello, FISCO BCOS" ./subscriber0.out >> /dev/null ;then LOG_ERROR "amop unique broadcast failed." && exit 1;fi
    pid=$(ps -ef | grep -v grep | grep unicast_publisher | awk '{print $2}')
    if [[ ! -z "${pid}" ]];then kill -9 "${pid}";fi
    pid=$(ps -ef | grep -v grep | grep subscriber | awk '{print $2}')
    if [[ ! -z "${pid}" ]];then kill -9 "${pid}";fi
    LOG_INFO "amop unique broadcast test success!"

    execute_cmd "go build -o multicast_publisher examples/amop/broadcast_pub/publisher.go"
    nohup ./multicast_publisher 127.0.0.1:20202 hello1 > output.file 2>&1 &
    nohup ./subscriber 127.0.0.1:20203 hello1 > subscriber1.out 2>&1 &
    sleep 13s
    cat subscriber1.out
    if ! grep "hello, FISCO BCOS" ./subscriber1.out >> /dev/null ;then LOG_ERROR "amop multi broadcast failed." && exit 1;fi
    pid=$(ps -ef | grep -v grep | grep multicast_publisher | awk '{print $2}')
    if [[ ! -z "${pid}" ]];then kill -9 "${pid}";fi
    pid=$(ps -ef | grep -v grep | grep subscriber | awk '{print $2}')
    if [[ ! -z "${pid}" ]];then kill -9 "${pid}";fi
    LOG_INFO "amop multi broadcast test success!"

    bash nodes/127.0.0.1/stop_all.sh
    LOG_INFO "integration_amop testing pass."
}

check_env
compile_and_ut
get_build_chain
integration_std
integration_amop

if [ -z "${macOS}" ];then integration_gm ; fi
