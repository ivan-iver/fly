#!/bin/bash
# --
# Author
# --
# Iván Jaimes <ivan@iver.mx>
#
# Goal script: Build with version from tag names.
set -o errexit
set -o nounset

. ${SCRIPTS_PATH}/env
source ${SCRIPTS_PATH}/version.sh

UNAMESTR=`uname -s`;
pcolor "DEBUG" "Tarjet Operative System is : ${TARJET_OS}";

if [[ "${UNAMESTR}" == 'Darwin' && "${TARJET_OS}" == 'LINUX' ]]; then
  pcolor "INFO" "Build on Darwin ... for linux"
  export CGO_ENABLED=0;
  export GOOS=linux;
else
  pcolor "INFO" "Build on ${UNAMESTR} the bin=>${BIN}"
fi

go build -a                                                     \
    -ldflags "-X main.version=${VERSION}"                       \
    -o ${OUTPUT_PATH}/${BIN}

