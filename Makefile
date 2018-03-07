SHELL := bash

export ACTUAL := $(shell pwd)
export SCRIPTS_PATH := ${ACTUAL}/scripts
export OUTPUT_PATH := ${ACTUAL}/bin/
export DEPLOY_CI_PATH=${ACTUAL}/.ci
export BIN := fly
export PKG := ${BIN}.tar.gz
export REPO_NAME=github.com/iver/fly

# Tarjet options are: MAC | LINUX
export TARJET_OS := MAC

build:
	${SCRIPTS_PATH}/build.sh;

run:
	${SCRIPTS_PATH}/run.sh;

test:
	${SCRIPTS_PATH}/test.sh;

package: TARJET_OS := LINUX
package: clean build
	${SCRIPTS_PATH}/package.sh;

deploy: package
	${SCRIPTS_PATH}/deploy.sh;

install:
	${SCRIPTS_PATH}/install.sh;

clean:
	@echo "Removing files ...";
	@rm -rf ${OUTPUT_PATH} ${PKG}
	@echo "Done!";

uninstall:
	# Start uninstall
	@if [[ -d ${GOPATH}/src/${PKG} ]]; then \
		echo "Removing package directory"; \
		rm -rf ${GOPATH}/src/${PKG}; \
	else \
		echo "Unlink package"; \
		unlink ${GOPATH}/src/${PKG}; \
	fi;
	@rm -f ${GOPATH}/bin/${PKG};
	go clean -i github.com/iver/fly;

