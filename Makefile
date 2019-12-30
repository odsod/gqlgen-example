include tools/gqlgen/rules.mk

gqlgen.yml: $(gqlgen)
	$(gqlgen) init

.PHONY: go-mod-tidy
go-mod-tidy:
	go mod tidy
