all: \
	gqlgen-generate \
	go-lint \
	go-mod-tidy

include tools/gqlgen/rules.mk
include tools/golangci-lint/rules.mk

.PHONY: gqlgen-generate
gqlgen-generate: $(gqlgen)
	$(gqlgen) generate -v

gqlgen.yml: $(gqlgen)
	$(gqlgen) init

.PHONY: go-mod-tidy
go-mod-tidy:
	go mod tidy -v

.PHONY: go-lint
go-lint: $(golangci_lint)
	$(golangci_lint) run
