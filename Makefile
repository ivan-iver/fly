SHELL	:= bash
ACTUAL := $(shell pwd)
NAME := fly
PKGNAME := fly.tar.gz
VERSION := "build\:\($(shell git rev-parse --short HEAD)\)"

export ACTUAL
export NAME
export PKGNAME
export VERSION

build:
#	@sed -i '' -e "s/build\:\([a-zA-Z0-9]*\)/${VERSION}/g" lib/app.go
	go build -ldflags "-X github.com/iver/fly/lib.hash=${VERSION}" -o bin/${NAME} github.com/iver/fly;
	@cp ${ACTUAL}/app.conf bin/;
	@cp ${ACTUAL}/README.md bin/;
	@cp -r ${ACTUAL}/assets bin/;
	@cp -r ${ACTUAL}/templates bin/;
	@mkdir -p bin/log/;

test:
	@go test -v ./lib/... -covermode=count -coverprofile=coverage.out
	@goveralls -coverprofile=coverage.out -service=travis-ci -repotoken ${COVERALLS_TOKEN}

init:
	@mkdir -p bin;
	@./scripts/install.sh;
#	@curl https://glide.sh/get | sh;

package: build
	@cp -r ${ACTUAL}/app.conf bin/;
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

