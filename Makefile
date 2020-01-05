curr_file := $(lastword $(MAKEFILE_LIST))

.PHONY: all
all: \
	protoc-generate \
	dataloaders-generate \
	gqlgen-generate \
	wire-generate \
	go-lint \
	go-mod-tidy

.PHONY: clean
clean:
	rm -rf build internal/gen
	find . -name '*_gen.go' -exec rm {} \+
	find tools -mindepth 2 -maxdepth 2 -type d -exec rm -rf {} \+

include tools/buf/rules.mk
include tools/dataloaden/rules.mk
include tools/golangci-lint/rules.mk
include tools/grpcurl/rules.mk
include tools/gqlgen/rules.mk
include tools/protoc/rules.mk
include tools/protoc-gen-go/rules.mk
include tools/protoc-gen-grpc-gateway/rules.mk
include tools/wire/rules.mk

.PHONY: buf-check-lint
buf-check-lint: $(buf)
	$(buf) check lint

proto_files := $(shell find api/proto -type f -name '*.proto')

build/proto.bin: $(buf) $(proto_files)
	mkdir -p $(dir $@)
	$(buf) image build -o $@

.PHONY: protoc-generate
protoc-generate: build/protoc-generate

build/protoc-generate: go_out := internal/gen/proto/go
build/protoc-generate: build/proto.bin $(curr_file) $(protoc) $(protoc_gen_go) $(protoc_gen_grpc_gateway)
	rm -rf $(go_out)
	mkdir -p $(go_out)
	$(protoc) --descriptor_set_in=$< \
		--go_out=plugins=grpc:$(go_out) \
		--grpc-gateway_out=logtostderr=true:$(go_out) \
		$(shell cd api/proto/src && find odsod/todo -type f)
	$(protoc) --descriptor_set_in=$< \
		--go_out=plugins=grpc:$(go_out) \
		--grpc-gateway_out=logtostderr=true:$(go_out) \
		$(shell cd api/proto/src && find odsod/user -type f)
	touch $@

.PHONY: dataloaders-generate
dataloaders-generate: \
	internal/gen/dataloader/userloader_gen.go

internal/gen/dataloader/package.go:
	mkdir -p $(dir $@)
	echo 'package dataloader' > $@

internal/gen/dataloader/userloader_gen.go: $(dataloaden) $(curr_file) internal/gen/dataloader/package.go
	cd $(dir $@) && $(dataloaden) UserLoader string '*github.com/odsod/gqlgen-example/internal/gen/proto/go/odsod/user/v1beta1.User'

.PHONY: gqlgen-generate
gqlgen-generate: build/gqlgen-generate

graphql_files := $(shell find api/graphql -type f)
model_files := $(shell find internal/model -type f -and -not -name '*_gen.go')

build/gqlgen-generate: $(gqlgen) gqlgen.yml $(graphql_files) $(model_files)
	$(gqlgen) generate -v
	touch $@

.PHONY: wire-generate
wire-generate: $(shell find . -type f -name 'wire.go' | sed 's/wire.go/wire_gen.go/')

%wire_gen.go: $(wire) $(shell find . -type f -name '*.go' -and -not -name 'wire_gen.go')
	$(wire) gen ./$(dir $@)

.PHONY: go-mod-tidy
go-mod-tidy:
	go mod tidy -v

.PHONY: go-lint
go-lint: $(golangci_lint)
	$(golangci_lint) run

.PHONY: grpcurl-list
grpcurl-list: $(grpcurl)
	$(grpcurl) -plaintext localhost:8081 list

.PHONY: curl-list-users
curl-list-users:
	curl -s 'localhost:8080/v1beta1/users?page_size=10' | jq .
