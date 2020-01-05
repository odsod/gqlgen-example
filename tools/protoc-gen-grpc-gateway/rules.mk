protoc_gen_grpc_gateway_cwd := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
protoc_gen_grpc_gateway := $(grpcurl_cwd)/bin/protoc-gen-grpc-gateway

PATH := $(dir $(protoc_gen_grpc_gateway)):$(PATH)

$(protoc_gen_grpc_gateway): $(protoc_gen_grpc_gateway_cwd)/../../go.mod
	cd $(protoc_gen_grpc_gateway_cwd)/../.. && go build -o $@ github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
