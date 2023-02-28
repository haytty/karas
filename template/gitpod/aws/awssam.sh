#!/bin/bash

# update AWS CLI
OLD_DIR="$PWD"
TMP_DIR="$(mktemp -d)"
echo "Updating AWS SAM"
cd "${TMP_DIR}" || exit 1

wget "https://github.com/aws/aws-sam-cli/releases/latest/download/aws-sam-cli-linux-x86_64.zip"
unzip aws-sam-cli-linux-x86_64.zip -d sam-installation
sudo ./sam-installation/install
rm aws-sam-cli-linux-x86_64.zip

cd "${OLD_DIR}" || exit 1
rm -rf "${TMP_DIR}"

