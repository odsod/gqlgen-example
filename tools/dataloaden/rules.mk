dataloaden_cwd := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
dataloaden := $(dataloaden_cwd)/bin/dataloaden

$(dataloaden): $(dataloaden_cwd)/go.mod
	cd $(dataloaden_cwd) && go build -o $@ github.com/vektah/dataloaden
