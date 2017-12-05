#!/bin/bash
set -o errexit
set -o nounset

export ACTUAL=`pwd`
export PKG=github.com/dherby/lipu-api
export BIN=lipu-api

echo "READ VERSION FROM TAGS ...";

TAGS=`git fetch --tags`
CURRENT_TAG=`git describe --tags $(git rev-list --tags --max-count=1) | cut -d'-' -f2`
HASH=`git rev-parse --short HEAD`

export VERSION="${CURRENT_TAG}-(${HASH})"

echo "VERSION CREATED: ${VERSION}"

UNAMESTR=`uname -s`;

if [[ "${UNAMESTR}" == 'Darwin' ]]; then
  CGO_ENABLED=0 GOOS=linux
fi

go build -a                                                     \
    -ldflags "-X main.version=${VERSION}"                       \
    -o bin/${BIN}                                               \
    ./cmd/...

