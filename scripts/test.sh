#!/bin/sh
#set -o errexit
#set -o nounset
#set -o pipefail
#set -o xtrace

go test -v $(go list ./... |grep -v vendor/)

