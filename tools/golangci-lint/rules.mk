cwd := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
golangci_lint := $(cwd)/bin/golangci-lint

$(golangci_lint): $(cwd)/go.mod
	cd $(cwd) && go build -o $@ github.com/golangci/golangci-lint/cmd/golangci-lint
