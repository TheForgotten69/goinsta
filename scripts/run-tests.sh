#!/usr/bin/env sh
set -e

export GO111MODULE=on

go test -v -race  ./...
