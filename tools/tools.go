// +build tools

package tools

import (
	_ "github.com/axw/gocov/gocov"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/matm/gocov-html"
	_ "github.com/vasi-stripe/gogroup/cmd/gogroup"
	_ "golang.org/x/tools/cmd/cover"
)
