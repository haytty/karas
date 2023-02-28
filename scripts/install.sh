#!/bin/bash

PROJECT_DIR=$(cd $(dirname $0)/..;pwd)
APP_NAME=$(basename $PROJECT_DIR)
SOURCE_PKG_PATH="$PROJECT_DIR/bin/$APP_NAME"

__build(){
  "$PROJECT_DIR/scripts/build.sh"
}

__copy(){
  local prefix="$1"
  cp -pr "$SOURCE_PKG_PATH" "$prefix/$APP_NAME"
}

__install(){
  local prefix="$1"
  __build
  __copy "$prefix"
}

__uninstall(){
  local prefix="$1"
  rm -fr "$prefix/$APP_NAME"
}

__reinstall(){
  local prefix="$1"
  __uninstall "$prefix"
  __install "$prefix"
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

    if [ -f "$prefix/$APP_NAME" ]; then
      __reinstall "$prefix"
      exit 0
    fi
    __install "$prefix"
    echo "INSTALL FINISHED!!!"
    exit 0
}

main $@ && exit 1
