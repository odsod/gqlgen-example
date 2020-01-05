protoc_gen_go_cwd := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
protoc_gen_go := $(protoc_gen_go_cwd)/bin/protoc-gen-go

PATH := $(dir $(protoc_gen_go)):$(PATH)

$(protoc_gen_go): $(protoc_gen_go_cwd)/../../go.mod
	cd $(protoc_gen_go_cwd)/../.. && go build -o $@ github.com/golang/protobuf/protoc-gen-go
