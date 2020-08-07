#!/bin/bash

LOG_ERROR()
{
    local content=${1}
    echo -e "\033[31m[ERROR] ${content}\033[0m"
    exit 1
}

execute_cmd()
{
    command="${1}"
    #LOG_INFO "RUN: ${command}"
    eval ${command}
    ret=$?
    if [ $ret -ne 0 ];then
        LOG_ERROR "FAILED execution of command: ${command}"
	if [ -d "nodes" ];then
	   bash nodes/127.0.0.1/stop_all.sh && rm -rf nodes
        fi
    fi
}

LOG_INFO()
{
    local content=${1}
    echo -e "\033[32m[INFO] ${content}\033[0m"
}

cur_path=$(execute_cmd "pwd")

# build blockchain
function build_blockchain()
{
  execute_cmd "rm -rf nodes"
  # download build_chain.sh
  execute_cmd "curl -LO https://raw.githubusercontent.com/FISCO-BCOS/FISCO-BCOS/master/tools/build_chain.sh && chmod u+x build_chain.sh"
  # get_buildchain.sh may fail due to access github api failed
  #bash <(curl -s https://raw.githubusercontent.com/FISCO-BCOS/FISCO-BCOS/dev/tools/get_buildchain.sh)
  
  if [ ! -f "build_chain.sh" ];then
     LOG_ERROR "get build_chain.sh failed!"
  fi
  execute_cmd "chmod a+x build_chain.sh"
  # build the blockchain
  ./build_chain.sh -v 2.5.0 -l "127.0.0.1:4"
  # copy certificate
  # execute_cmd "cp nodes/127.0.0.1/sdk/* bin/"
}

# start the nodes
function start_nodes()
{
   execute_cmd "./nodes/127.0.0.1/start_all.sh"
}

# stop the nodes
function stop_nodes()
{
   execute_cmd "./nodes/127.0.0.1/stop_all.sh"
}

function main()
{
   build_blockchain
   start_nodes
}

main