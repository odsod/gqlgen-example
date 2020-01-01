package resolver

import (
	"context"
	"fmt"

	todov1beta1 "github.com/odsod/gqlgen-example/internal/gen/proto/go/odsod/todo/v1beta1"
	"github.com/odsod/gqlgen-example/internal/storage"
	"go.uber.org/zap"
)

type Query struct {
	Storage *storage.InMemory
	Logger  *zap.Logger
}

func (r *Query) Todos(ctx context.Context) ([]*todov1beta1.Todo, error) {
	r.Logger.Debug("query: todos")
	todos, err := r.Storage.ListTodos(ctx)
	if err != nil {
		return nil, fmt.Errorf("resolve todos: %w", err)
	}
	return todos, nil
}
