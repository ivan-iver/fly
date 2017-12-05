#!/bin/bash
set -o errexit
set -o nounset

go get -v golang.org/x/tools/cmd/cover
go get -v github.com/mattn/goveralls

go get -v bitbucket.org/ivan-iver/config;
go get -v gopkg.in/alecthomas/kingpin.v2;
go get -v bitbucket.org/ivan-iver/config;
go get -v github.com/op/go-logging;
go get -v gopkg.in/unrolled/render.v1;
go get -v github.com/theplant/blackfriday;
