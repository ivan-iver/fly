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

pcolor "INFO" "READ VERSION FROM TAGS ...";

TAGS=`git fetch --tags`
CURRENT_TAG=`git describe --tags $(git rev-list --tags --max-count=1) | cut -d'-' -f2`
HASH=`git rev-parse --short HEAD`

export VERSION="${CURRENT_TAG}-(${HASH})"

pcolor "DEBUG" "VERSION CREATED: ${VERSION} "

