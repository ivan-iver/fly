SHELL		:= bash
ACTUAL := $(shell pwd)
NAME := fly
PKGNAME := fly.tar.gz
VERSION := "build\:\($(shell git rev-parse --short HEAD)\)"

export ACTUAL
export NAME
export PKGNAME
export VERSION

build: install
	@sed -i '' -E "s/build\:\([a-zA-Z0-9]*\)/${VERSION}/g" ${NAME}/lib/app.go
	@go build -o bin/${NAME} fly;
	@cp ${ACTUAL}/fly/app.conf bin/;
	@cp ${ACTUAL}/fly/README.md bin/;
	@cp -r ${ACTUAL}/fly/assets bin/;
	@cp -r ${ACTUAL}/fly/templates bin/;
	@mkdir -p bin/log/;

install: base get

get:
	go get -v bitbucket.org/ivan-iver/config;
	go get -v gopkg.in/alecthomas/kingpin.v2;
	go get -v bitbucket.org/ivan-iver/config;
	go get -v github.com/sirupsen/logrus;
	go get -v gopkg.in/unrolled/render.v1;
	go get -v github.com/theplant/blackfriday;

base:
	#  ---- Variables ----
	# | GOPATH: ${GOPATH}
	# | ACTUAL: ${ACTUAL}
	#  -------------------
	# Creating working directory ${GOPATH}/src/${NAME}
	# Checking link ${GOPATH}/src/${NAME} to ${ACTUAL}/${NAME}
	@if [[ -L ${GOPATH}/src/${NAME} && -d ${GOPATH}/src/${NAME} ]]; then \
		echo "Skip Linked"; \
	else \
		echo "Linking package ..."; \
		ln -sf ${ACTUAL}/${NAME} ${GOPATH}/src/${NAME}; \
	fi;

package: build
	@cp -r ${ACTUAL}/${NAME}/app.conf bin/;
	@tar -zcf ${PKGNAME} bin;
	@rm -rf ${ACTUAL}/bin;
	@echo "Package done! ... you can run deploy.sh script from yout host machine.";

uninstall:
	# Start uninstall
	@if [[ -d ${GOPATH}/src/${NAME} ]]; then \
		echo "Removing directory"; \
		rm -rf ${GOPATH}/src/${NAME}; \
	else \
		echo "Unlink package"; \
		unlink ${GOPATH}/src/${NAME}; \
	fi;
	@rm -f ${GOPATH}/bin/${NAME};

clean:
	@rm -rf ${ACTUAL}/${PKGNAME}
	@rm -rf ${ACTUAL}/bin;

