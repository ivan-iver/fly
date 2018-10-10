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

pcolor "INFO" "Creating output folders";
mkdir -p "${OUTPUT_PATH}" && \
  mkdir -p "${OUTPUT_PATH}"/etc && \
  mkdir -p "${OUTPUT_PATH}"/logs;

pcolor "INFO" "Starting to pack ${BIN}";
cp "${ACTUAL}"/app.conf "${OUTPUT_PATH}"/etc/;
cp -R "${ACTUAL}"/templates "${OUTPUT_PATH}";
cp -R "${ACTUAL}"/assets "${OUTPUT_PATH}";
cd "${ACTUAL}";
tar -zcf ${PKG} ${OUTPUT_NAME};

pcolor "INFO" "Done!";

