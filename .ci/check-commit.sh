#!/bin/bash

set -e

SHELL_FOLDER=$(
    cd $(dirname $0)
    pwd
)

# check_script=gofmt -e -s -w
check_script="goimports -d"
commit_limit=2
file_limit=35
insert_limit=1000
delete_limit=500
new_file_header_length=35
skip_check_words="sync code"

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

function check_codeFormat() {
    go version
    go install golang.org/x/tools/cmd/goimports@latest || true
    # go get golang.org/x/tools/cmd/goimports || true
    sum=0
    for file in $(git diff-index --name-status HEAD^ | grep -v D | grep -E '\.go' | awk '{print $2}'); do
        checkResult=$(eval "${check_script} ${file}")
        if [ -n "${checkResult}" ]; then
            LOG_ERROR "file ${file} does not meet the format requirements"
            echo "${checkResult}"
            sum=1
        fi
    done
    if [ ${sum} -ne 0 ]; then
        LOG_ERROR "######### ERROR: Format check failed, please adjust them before commit"
        exit 1
    fi
}

function check_PR_limit() {
    # if [ "${TRAVIS_PULL_REQUEST}" != "false" ]; then
    #     local skip=$(curl -s https://api.github.com/repos/FISCO-BCOS/FISCO-BCOS/pulls/${TRAVIS_PULL_REQUEST} | grep "title\"" | grep "${skip_check_words}")
    #     if [ ! -z "${skip}" ]; then
    #         LOG_INFO "sync code PR, skip PR limit check!"
    #         exit 0
    #     else
    #         LOG_INFO "PR-${TRAVIS_PULL_REQUEST}, checking PR limit..."
    #     fi
    # fi
    local files=$(git diff --shortstat HEAD^ | awk -F ' ' '{print $1}')
    # if [ ${file_limit} -lt ${files} ]; then
    #     LOG_ERROR "modify ${files} files, limit is ${file_limit}"
    #     exit 1
    # fi
    local new_files=$(git diff HEAD^ | grep "new file" | wc -l)
    local ci_insertions=$(git diff --numstat HEAD^ | grep ".ci/" | awk -F ' ' '{sum+=$1}END{print sum}')
    local insertions=$(git diff --shortstat HEAD^ | awk -F ' ' '{print $4}')
    local valid_insertions=$((insertions - new_files * new_file_header_length - ci_insertions))
    if [ ${insert_limit} -lt ${valid_insertions} ]; then
        LOG_ERROR "insert ${insertions} lines, valid is ${valid_insertions}, limit is ${insert_limit}"
        exit 1
    fi
    local deletions=$(git diff --shortstat HEAD^ | awk -F ' ' '{print $6}')
    #if [ ${delete_limit} -lt ${deletions} ];then
    #    LOG_ERROR "delete ${deletions} lines, limit is ${delete_limit}"
    #    exit 1
    #fi
    local commits=$(git rev-list --count HEAD^..HEAD)
    if [ ${commit_limit} -lt ${commits} ]; then
        LOG_ERROR "${commits} commits, limit is ${commit_limit}"
        git --no-pager log -"${commits}"
        exit 1
    fi
    local unique_commit=$(git log --format=%s HEAD^..HEAD | sort -u | wc -l)
    if [ ${unique_commit} -ne ${commits} ]; then
        LOG_ERROR "${commits} != ${unique_commit}, please make commit message unique!"
        exit 1
    fi
    local merges=$(git log --format=%s HEAD^..HEAD | grep -i merge | wc -l)
    if [ ${merges} -gt 2 ]; then
        LOG_ERROR "PR contain merge : ${merges}, Please rebase!"
        exit 1
    fi
    LOG_INFO "modify ${files} files, insert ${insertions} lines, valid insertion ${valid_insertions}, delete ${deletions} lines. Total ${commits} commits."
}

check_codeFormat
# check_PR_limit
