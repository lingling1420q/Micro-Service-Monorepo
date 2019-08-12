#!/bin/bash

ROOT_DIR=${ROOT_DIR:-$(git rev-parse --show-toplevel)}

ARGS=("$@")

RED='\033[0;31m'
NC='\033[0m' # No Color
GREEN='\033[0;32m'

function red {
    echo -e "${RED}$1${NC}"
}

function green {
    echo -e "${GREEN}$1${NC}"
}

function check_process {
    if (($?)); then
        red "Start $1 ERROR"
        if [[ $2 ]]; then
            `$2`
        fi
        exit 1
    else
        green "Start $1 SUCCESS"
    fi
}

stop_process () {
    echo "stop process of "$1
    _PID=`ps ufx | grep "$1" | awk '{print $2}'`
    echo "$_PID"|while read pid
    do
        if [ "$pid" != "" ]; then
            echo "kill process $pid"
            kill $pid
        else
            echo "No "$1" processes detected."
        fi
    done

    echo "Done."
}

remove_file () {
    if [ -f "$1" ]; then
    echo "$1 exist"
    rm -rf $1
    fi
}

build_server() {
    source_file="$ROOT_DIR/build/server"
    remove_file $source_file
    go build -x -o $source_file
}

print_help () {
    echo "#*-----------------------------------------------------------------------------"
    echo "# Usage: dev [start/stop] <target> [--native port]"
    echo "#"
    echo "# Mandatory arguments:"
    echo "#   start:               start component"
    echo "#   stop:                kill component"
    green "#   run:                run dev server"
    green "#   build:              build a native file to repo/build/server"
    echo "#"
    echo "# All avaiable targets are list as follow:"
    echo "#   all         All services at once"
    echo "#"
    echo "# use 'dev help' to list all available comamnds"
    echo "#*-----------------------------------------------------------------------------"
}


case ${ARGS[0]} in
    start)
        case ${ARGS[1]} in
            *)          print_help ;;
        esac ;;
    stop)
        case ${ARGS[1]} in
            *)          print_help ;;
        esac ;;
    run)
        go run main.go ;;
    build)
        build_server ;;
    *)
        print_help ;;
esac
