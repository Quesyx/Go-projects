#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

if [ -z "${OS:-OS}" ]; then
    echo "OS must be set"
    exit 1
fi
if [ -z "${ARCH:-ARCH}" ]; then
    echo "ARCH must be set"
    exit 1
fi
if [ -z "${VERSION:-VERSION}" ]; then
    echo "VERSION must be set"
    exit 1
fi

export CGO_ENABLED=0
export GOARCH="amd64"
export GOOS="linux"
export GO111MODULE=on
export GOFLAGS="-mod=mod"

go install                                                      \
    -installsuffix "static"                                     \
    -ldflags "-X $(go list -m)/pkg/version.Version=go 1.14.4"  \
    ./...
