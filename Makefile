all: \
	go-dataloaders-generate \
	go-gqlgen-generate \
	go-wire-generate \
	go-lint \
	go-mod-tidy

include tools/dataloaden/rules.mk
include tools/golangci-lint/rules.mk
include tools/gqlgen/rules.mk
include tools/wire/rules.mk

.PHONY: go-dataloaders-generate
go-dataloaders-generate: \
	internal/dataloader/userloader_gen.go

internal/dataloader/userloader_gen.go: $(dataloaden)
	cd internal/dataloader && $(dataloaden) UserLoader string '*github.com/odsod/gqlgen-getting-started/internal/model.User'

.PHONY: go-gqlgen-generate
go-gqlgen-generate: $(gqlgen)
	$(gqlgen) generate -v

.PHONY: go-wire-generate
go-wire-generate: $(shell find . -type f -name 'wire.go' | sed 's/wire.go/wire_gen.go/')

%wire_gen.go: $(wire) $(shell find . -type f -not -name 'wire_gen.go')
	$(wire) gen ./$(dir $@)

.PHONY: go-mod-tidy
go-mod-tidy:
	go mod tidy -v

.PHONY: go-lint
go-lint: $(golangci_lint)
	$(golangci_lint) run
