#!/bin/bash

PROJECT_DIR=$(cd $(dirname $0)/../..;pwd)

main() {
  $PROJECT_DIR/template/gitpod/ide/goimports.sh
  $PROJECT_DIR/template/gitpod/aws/awssam.sh
  $PROJECT_DIR/template/gitpod/aws/awscli.sh
}

main
