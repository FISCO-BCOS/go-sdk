#!/bin/bash
source="https://github.com/FISCO-BCOS/solidity/releases/download"
install_path="${HOME}/.fisco"
version="0.4.25"
OS="linux"
crypto=

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
    -v <solc version>           Default 0.4.25, 0.5.2 is supported
    -g <gm version>             if set download solc gm version
    -h Help
e.g 
    $0 -v 0.4.25 -g
EOF

exit 0
}

check_env() {
    if [ "$(uname)" == "Darwin" ];then
        OS="mac"
    fi
    if [ "$(uname -m)" != "x86_64" ];then
        LOG_WARN "We only offer x86_64 precompiled solc binary, your OS architecture is not x86_64. Please compile from source."
        exit 1
    fi
}

parse_params()
{
    while getopts "v:o:gh" option;do
        case $option in
        v) [ -n "$OPTARG" ] && version="$OPTARG";;
        o) [ -n "$OPTARG" ] && install_path="$OPTARG";;
        g) crypto="-gm";;
        h) help;;
        *) LOG_WARN "invalid option $option";;
        esac
    done
}

main()
{
    package_name="solc-${OS}${crypto}.tar.gz"
    download_link="${source}/v${version}/${package_name}"
    echo "Downloading solc ${version} ${package_name} from ${download_link}"

    if [ ! -f "${install_path}/solc-${version}${crypto}" ];then
        if curl -LO "${download_link}" ;then
            tar -zxf "${package_name}" 
            rm -rf "${package_name}"
            mkdir -p "${install_path}"
            mv solc "${install_path}/solc-${version}${crypto}"
        else
            LOG_WARN "Download from ${download_link} failed, please retry."
            exit 1
        fi
    fi
    ln -s "${install_path}/solc-${version}${crypto}" "./solc-${version}${crypto}"

}

print_result()
{
    echo "=============================================================="
    LOG_INFO "os            : ${OS}"
    LOG_INFO "solc version  : ${version}"
    LOG_INFO "solc location : ./solc-${version}${crypto}"
    echo "=============================================================="
    LOG_INFO "./solc-${version}${crypto} --version"
    "./solc-${version}${crypto}" --version
}

parse_params "$@"
check_env
main
print_result
