#!/bin/bash
set -e

install_path="/usr/local/lib/"
version="3.4.0"
OS="linux"
versions=(3.2.0 3.4.0)

LOG_WARN()
{
    local content="${1}"
    echo -e "\033[31m[WARN] ${content}\033[0m"
}

LOG_INFO()
{
    local content="${1}"
    echo -e "\033[32m[INFO] ${content}\033[0m"
}

help() {
    cat << EOF
Usage:
    -v <bcos-c-sdk version>           Default v3.4.0
    -o <lib install path>             Default /usr/loacl/lib
    -h Help
e.g
    $0 -0 ./lib
EOF

    exit 0
}

parse_params()
{
    while getopts "v:o:h" option;do
        case $option in
            v) [ -n "$OPTARG" ] && version="$OPTARG"
                if ! echo "${versions[*]}" | grep -i "${version}" &>/dev/null; then
                    LOG_WARN "${version} is not supported. Please set one of ${versions[*]}"
                    exit 1;
                fi
            ;;
            o) [ -n "$OPTARG" ] && install_path="$OPTARG";;
            h) help ;;
            *) LOG_WARN "invalid option $option";;
        esac
    done
}

get_csdk_lib()
{
    if [ ! -d "${install_path}" ];then
        mkdir -p "${install_path}"
    fi
    local OS_ARCH=
    local suffix="so"
    if [ "$(uname)" == "Darwin" ];then
        if [[ "$(uname -m)" == "arm64" ]];then
            OS_ARCH="-aarch64"
        fi
        # ldflags="-ldflags=\"-r /usr/local/lib/\""
        OS="macOS"
        suffix="dylib"
    elif [ "$(uname -s)" == "Linux" ];then
        OS="linux"
        if [[ "$(uname -p)" == "aarch64" ]];then
            OS_ARCH="-aarch64"
        fi
    else
        OS="win"
        suffix="dll"
    fi
    local package_name="libbcos-c-sdk${OS_ARCH}.${suffix}"
    LOG_INFO "downloading ${package_name} to ${install_path} on ${OS} ..."
    if [ ! -f "${install_path}/${package_name}" ];then
        curl -o "${install_path}/${package_name}" -#L "https://github.com/FISCO-BCOS/bcos-c-sdk/releases/download/v${version}/${package_name}"
    fi
}

parse_params "$@"
get_csdk_lib
