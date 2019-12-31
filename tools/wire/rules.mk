cwd := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
wire := $(cwd)/bin/wire

$(wire): $(cwd)/../../go.mod
	cd $(cwd)/../.. && go build -o $@ github.com/google/wire/cmd/wire
