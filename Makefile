all: \
	gqlgen-generate \
	go-mod-tidy

include tools/gqlgen/rules.mk

.PHONY: gqlgen-generate
gqlgen-generate: $(gqlgen)
	$(gqlgen) generate -v

gqlgen.yml: $(gqlgen)
	$(gqlgen) init

.PHONY: go-mod-tidy
go-mod-tidy:
	go mod tidy -v
