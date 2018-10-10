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

# shellcheck source=/dev/null
. "${SCRIPTS_PATH}"/env

go run -a main.go

pcolor "INFO" "Build done!";
pcolor "INFO" "Execute binary file: ${BIN}";
cd "${OUTPUT_PATH}" && ${BIN};
