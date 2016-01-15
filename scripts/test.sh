#!/usr/bin/env bash
# This script syncs all the dependency packages required for this go web app
export GOPATH="$(dirname "$(dirname "$(dirname "$(dirname "$(pwd)")")")")"
export PATH=$PATH:$GOPATH/bin
echo "Current go path : $GOPATH"

go test github.com/betacraft/circleci -v

