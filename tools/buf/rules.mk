buf_cwd := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
buf := $(buf_cwd)/bin/buf

$(buf): $(buf_cwd)/go.mod
	cd $(buf_cwd) && go build -o $@ github.com/bufbuild/buf/cmd/buf
