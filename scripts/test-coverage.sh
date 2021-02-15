#!/usr/bin/env bash

go test -race -coverpkg=./... -v -coverprofile coverage.out.tmp ./...
# shellcheck disable=SC2002
cat coverage.out.tmp | grep -v "cmd/" > coverage.out
rm -rf coverage.out.tmp
gocov convert coverage.out > coverage.out.json
gocov report coverage.out.json
gocov-html coverage.out.json > coverage.out.html
go tool cover -html=coverage.out