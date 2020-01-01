grpcurl_cwd := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
grpcurl := $(grpcurl_cwd)/bin/grpcurl

$(grpcurl): $(grpcurl_cwd)/go.mod
	cd $(grpcurl_cwd) && go build -o $@ github.com/fullstorydev/grpcurl/cmd/grpcurl
