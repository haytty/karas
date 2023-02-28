#!/bin/bash

PROJECT_DIR=$(cd $(dirname $0)/..;pwd)
APP_NAME=$(basename $PROJECT_DIR)

__uninstall(){
  local prefix="$1"
  rm -fr "$prefix/$APP_NAME"
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
    --prefix          change prefix. default:/usr/local/bin
    --version, -v     print $(basename ${0}) version
    --help, -h        print this
EOF
}

main() {
    local opt optarg prefix
    while getopts hmsv-: opt; do
        optarg="${!OPTIND}"
        [[ "$opt" = - ]] && opt="-$OPTARG"

        case "-$opt" in
            --prefix)
                prefix="$optarg"
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

    if [ -z $prefix ]; then
      prefix="/usr/local/bin"
    fi

    if [ ! -f "$prefix/$APP_NAME" ]; then
      echo "$APP_NAME is not found"
      exit 1
    fi
    __uninstall "$prefix"
    echo "UNINSTALL FINISHED!!!"
    exit 0
}

main $@ && exit 1
