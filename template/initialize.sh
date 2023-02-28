#!/bin/bash

PROJECT_DIR=$(cd $(dirname $0)/..;pwd)

APP_NAME=$(basename $PROJECT_DIR)
BASE_PROJECT_NAME="golang_cli_template"

camelize() {
  local str="$1"
  echo "$str" | perl -pe 's/(?:^|_|-)(.)/\U$1/g'
}

CAMELIZED_BASE_PROJECT_NAME=$(camelize $BASE_PROJECT_NAME)
CAMELIZED_APP_NAME=$(camelize $APP_NAME)

pushd $PROJECT_DIR

# Replace Camelized BASE_PROJECT_NAME
grep -lr $CAMELIZED_BASE_PROJECT_NAME ./ | grep -v '.git' | grep -v template/initialize.sh | xargs perl -pi -e "s/${CAMELIZED_BASE_PROJECT_NAME}/${CAMELIZED_APP_NAME}/g"

# Replace BASE_PROJECT_NAME
grep -lr $BASE_PROJECT_NAME ./ | grep -v '.git' | grep -v template/initialize.sh | xargs perl -pi -e "s/${BASE_PROJECT_NAME}/${APP_NAME}/g"

## Change FileName 
mv $PROJECT_DIR/internal/handler/$BASE_PROJECT_NAME/$BASE_PROJECT_NAME.go $PROJECT_DIR/internal/handler/$BASE_PROJECT_NAME/$APP_NAME.go
mv $PROJECT_DIR/internal/model/$BASE_PROJECT_NAME.go $PROJECT_DIR/internal/model/$APP_NAME.go
mv $PROJECT_DIR/cmd/$BASE_PROJECT_NAME/$BASE_PROJECT_NAME.go $PROJECT_DIR/cmd/$BASE_PROJECT_NAME/$APP_NAME.go

## Change DirName
mv $PROJECT_DIR/cmd/$BASE_PROJECT_NAME $PROJECT_DIR/cmd/$APP_NAME
mv $PROJECT_DIR/internal/handler/$BASE_PROJECT_NAME $PROJECT_DIR/internal/handler/$APP_NAME


