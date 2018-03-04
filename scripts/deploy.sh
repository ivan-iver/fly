#!/bin/bash
# --
# Author
# --
# Iván Jaimes <ivan@iver.mx>
#
# Goal script: Deploy artifacts.

set -o errexit
set -o nounset

export USER=deploy

export SSH_PORT=22
export DEPLOY_KEY=${DEPLOY_CI_PATH}/deploy_rsa

. ${SCRIPTS_PATH}/env;
source ${SCRIPTS_PATH}/branch-detect.sh;

pcolor "DEBUG" "Current branch is: ${branch}";

export TARGET=$(cat ${DEPLOY_CI_PATH}/${branch})

pcolor "DEBUG" "Deploying target: ${TARGET}"

if [ ! -f ${PKG} ]; then
  pcolor "ERROR" "Web files are not yet ready. Use make package."
  exit 0;
fi

pcolor "INFO" "Sending files to server";
scp -vv -P ${SSH_PORT} -i ${DEPLOY_KEY} -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null ${PKG} ${USER}@${TARGET}:/tmp;
ssh -vv -p ${SSH_PORT} -i ${DEPLOY_KEY} -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -l ${USER} ${TARGET} "
    sudo /usr/local/bin/deploy.sh;

";

pcolor "INFO" "Deleting zip: ${PKG}";
rm ${PKG};
pcolor "INFO" "Deployment done!";

