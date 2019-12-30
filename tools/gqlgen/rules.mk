cwd := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
gqlgen := $(cwd)/bin/gqlgen

$(gqlgen): $(cwd)/go.mod
	cd $(cwd) && go build -o $@ github.com/99designs/gqlgen
