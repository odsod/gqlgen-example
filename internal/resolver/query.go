package resolver

import (
	"context"
	"fmt"

	"github.com/odsod/gqlgen-getting-started/internal/model"
	"github.com/odsod/gqlgen-getting-started/internal/storage"
	"go.uber.org/zap"
)

type Query struct {
	Storage *storage.InMemory
	Logger  *zap.Logger
}

func (r *Query) Todos(ctx context.Context) ([]*model.Todo, error) {
	r.Logger.Debug("query: todos")
	todos, err := r.Storage.ListTodos(ctx)
	if err != nil {
		return nil, fmt.Errorf("resolve todos: %w", err)
	}
	return todos, nil
}
