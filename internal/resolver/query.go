package resolver

import (
	"context"
	"log"

	"github.com/odsod/gqlgen-getting-started/internal/model"
)

type query struct {
	root *Root
}

func (r *query) Todos(ctx context.Context) ([]*model.Todo, error) {
	log.Printf("query: todos")
	return r.root.Storage.Todos, nil
}
