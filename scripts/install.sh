#!/bin/bash
# --
# Author
# --
# Iván Jaimes <ivan@iver.mx>
#
# Goal script:  set color environment variables.
# COMMENT: This script depends on setting global variables from Makefile on root project
set -o errexit
set -o nounset

go env
go get github.com/Masterminds/glide

if [[ -z "${GOPATH}/src/${REPO_NAME}" ]]; then
  mkdir -p $GOPATH/src/$(dirname $REPO_NAME)
  ln -svf ${ACTUAL} ${GOPATH}/src/${REPO_NAME}
fi

cd ${GOPATH}/src/${REPO_NAME}
glide install

