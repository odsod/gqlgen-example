package resolver

import (
	"context"
	"fmt"

	todov1beta1 "github.com/odsod/gqlgen-example/internal/gen/proto/go/odsod/todo/v1beta1"
	"go.uber.org/zap"
)

type Query struct {
	TodoServiceClient todov1beta1.TodoServiceClient
	Logger            *zap.Logger
}

func (r *Query) Todos(ctx context.Context) ([]*todov1beta1.Todo, error) {
	r.Logger.Debug("query: todos")
	res, err := r.TodoServiceClient.ListTodos(ctx, &todov1beta1.ListTodosRequest{
		// TODO: Add paging support
		PageSize: 100,
	})
	if err != nil {
		return nil, fmt.Errorf("resolve todos: %w", err)
	}
	return res.Todos, nil
}
