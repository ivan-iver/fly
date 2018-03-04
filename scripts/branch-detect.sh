#!/bin/bash
# --
# Author
# --
# Iván Jaimes <ivan@iver.mx>
#
# Goal script:  Detect current branch on CI application (travis in this case).
set -o errexit
set -o nounset

RAMA=`git branch -a| grep \* | cut -d ' ' -f2`;

if [[ "${RAMA}" == *"HEAD"* ]] || [[ "${RAMA}" == *"detached"* ]]; then
  [ -n ${TRAVIS_BRANCH-} ] && RAMA=${TRAVIS_BRANCH}
  pcolor "INFO" "SET RAMA: ${RAMA}"
fi

if [[ "${RAMA}" == "master" ]]; then
  export branch=prod
elif [[ "${RAMA}" == *"release"* ]]; then
  export branch=uat
else
  export branch=dev
fi

pcolor "INFO" "Evironment ${PURPLE}${branch}${NC} and branch ${PURPLE}${RAMA}${NC}";

