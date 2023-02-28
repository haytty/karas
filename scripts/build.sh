#!/bin/bash

PROJECT_DIR=$(cd $(dirname $0)/..;pwd)
APP_NAME=$(basename $PROJECT_DIR)
DISTINATION_PATH="$PROJECT_DIR/bin/$APP_NAME"
BUILD_PATH="$PROJECT_DIR/cmd/$APP_NAME"

__build(){
  go build -ldflags="-s -w" -trimpath -o $DISTINATION_PATH $BUILD_PATH
}

__version(){
  echo 'v1.0.0'
}

__help(){
      cat <<EOF
$(basename ${0}) is a tool for ...

Usage:
    $(basename ${0}) [command] [<options>]

Options:
    --version, -v     print $(basename ${0}) version
    --help, -h        print this
EOF
}

__clean() {
  rm -fr "$DISTINATION_PATH"
}

main() {
    local opt optarg clean_flag
    while getopts hmsv-: opt; do
        optarg="${!OPTIND}"
        [[ "$opt" = - ]] && opt="-$OPTARG"

        case "-$opt" in
            --clean)
                clean_flag=1
                shift
                ;;
            -v|--version)
                __version
                exit
                ;;
            -h|--help)
                __help
                exit
                ;;
            --)
                break
                ;;
            -\?)
                exit 1
                ;;
            --*)
                echo "$0: illegal option -- ${opt##-}" >&2
                exit 1
                ;;
        esac
    done

    if [ ! -z "$clean_flag" ]; then
      __clean
    fi

    __build
    echo "BUILD FINISHED!!!"
    exit 0
}

main $@ && exit 1
