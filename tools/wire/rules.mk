wire_cwd := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
wire := $(wire_cwd)/bin/wire

$(wire): $(wire_cwd)/../../go.mod
	cd $(wire_cwd)/../.. && go build -o $@ github.com/google/wire/cmd/wire
