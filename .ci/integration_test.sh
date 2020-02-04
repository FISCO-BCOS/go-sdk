# !/bin/bash

set -e

macOS=

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
}

compile_and_ut()
{
    export GO111MODULE="on"
    execute_cmd "go build cmd/console.go"
    execute_cmd "go build -o abigen ./cmd/abigen/main.go"

    execute_cmd "go test -v ./smcrypto"
}

generate_main() {
    local output="${1}"
cat << EOF >> ${output}

func main() {
	configs := conf.ParseConfig("config.toml")
	client, err := client.Dial(&configs[0])
	if err != nil {
		fmt.Printf("Dial Client failed, err:%v", err)
		return
	}
	address, tx, instance, err := DeployMain(client.GetTransactOpts(), client)
	if err != nil {
		fmt.Printf("Deploy failed, err:%v", err)
		return
	}
	fmt.Println("contract address: ", address.Hex()) // the address should be saved
	fmt.Println("transaction hash: ", tx.Hash().Hex())
	hello := &MainSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}
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
}

get_build_chain()
{
    # latest_version=$(curl -s https://api.github.com/repos/FISCO-BCOS/FISCO-BCOS/releases | grep "\"v2\.[0-9]\.[0-9]\"" | sort -u | tail -n 1 | cut -d \" -f 4)
    # curl -LO https://github.com/FISCO-BCOS/FISCO-BCOS/releases/download/${latest_version}/build_chain.sh && chmod u+x build_chain.sh
    curl -LO https://raw.githubusercontent.com/FISCO-BCOS/FISCO-BCOS/master/tools/build_chain.sh && chmod u+x build_chain.sh
}

precompiled_test(){
    if [ -z "$(go test -v ./precompiled/cns| grep FAIL)";then LOG_ERROR "test precompiled/cns failed" && exit 1; fi
    if [ -z "$(go test -v ./precompiled/config| grep FAIL)";then LOG_ERROR "test precompiled/config failed" && exit 1; fi
    if [ -z "$(go test -v ./precompiled/consensus| grep FAIL)";then LOG_ERROR "test precompiled/consensus failed" && exit 1; fi
    if [ -z "$(go test -v ./precompiled/crud| grep FAIL)";then LOG_ERROR "test precompiled/crud failed" && exit 1; fi
    if [ -z "$(go test -v ./precompiled/permission| grep FAIL)";then LOG_ERROR "test precompiled/permission failed" && exit 1; fi
}

integration_std()
{
    LOG_INFO "integration_std testing..."

    # abigen std
    execute_cmd "./abigen --bin=.ci/hello/HelloWorld.bin --abi=.ci/hello/HelloWorld.abi --pkg=main --out=hello.go"
    generate_main hello.go
    execute_cmd "go build -o hello hello.go"
    LOG_INFO "generate hello.go and build hello done."

    bash build_chain.sh -l 127.0.0.1:4 -o nodes
    cp nodes/127.0.0.1/sdk/* ./
    bash nodes/127.0.0.1/start_all.sh
    if [ -z "$(./hello | grep address)" ];then LOG_ERROR "std deploy contract failed." && exit 1;fi
    if [ ! -z "$(./hello | grep failed)" ];then LOG_ERROR "call contract interface failed." && exit 1;fi
    precompiled_test
    bash nodes/127.0.0.1/stop_all.sh
    LOG_INFO "integration_std testing pass."

}

integration_gm()
{
    LOG_INFO "integration_gm testing..."

    # abigen gm 
    execute_cmd "./abigen --bin=.ci/hello/HelloWorld_gm.bin --abi=.ci/hello/HelloWorld.abi --pkg=main --out=hello_gm.go --smcrypto=true"
    generate_main hello_gm.go
    execute_cmd "go build -o hello_gm hello_gm.go"
    LOG_INFO "generate hello_gm.go and build hello_gm done."

    bash build_chain.sh -l 127.0.0.1:4 -g -o nodes_gm
    cp nodes_gm/127.0.0.1/sdk/* ./
    bash nodes_gm/127.0.0.1/start_all.sh
    sed -i "s/SMCrypto=false/SMCrypto=true/g" config.toml
    sed -i "s#KeyFile=\".ci/0x83309d045a19c44dc3722d15a6abd472f95866ac.pem\"#KeyFile=\".ci/sm2p256v1_0x791a0073e6dfd9dc5e5061aebc43ab4f7aa4ae8b.pem\"#g" config.toml
    if [ -z "$(./hello_gm | grep address)" ];then LOG_ERROR "gm deploy contract failed." && exit 1;fi
    if [ ! -z "$(./hello_gm | grep failed)" ];then LOG_ERROR "gm call contract interface failed." && exit 1;fi
    precompiled_test
    bash nodes_gm/127.0.0.1/stop_all.sh
    LOG_INFO "integration_gm testing pass."

}

check_env
compile_and_ut
get_build_chain
integration_std
if [ -z "${macOS}" ];then integration_gm ; fi
