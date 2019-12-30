package storage

import "github.com/odsod/gqlgen-getting-started/internal/model"

type InMemory struct {
	Todos []*model.Todo
}
