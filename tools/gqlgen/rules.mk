gqlgen_cwd := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
gqlgen := $(gqlgen_cwd)/bin/gqlgen

$(gqlgen): $(gqlgen_cwd)/go.mod
	cd $(gqlgen_cwd) && go build -o $@ github.com/99designs/gqlgen
