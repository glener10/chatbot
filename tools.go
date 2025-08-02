//go:build tools

package main

import (
	_ "github.com/go-delve/delve/cmd/dlv"
	_ "github.com/golangci/golangci-lint/v2/cmd/golangci-lint"
)
