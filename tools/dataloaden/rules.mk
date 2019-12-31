cwd := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
dataloaden := $(cwd)/bin/dataloaden

$(dataloaden): $(cwd)/go.mod
	cd $(cwd) && go build -o $@ github.com/vektah/dataloaden
