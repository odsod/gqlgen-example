package resolver

import (
	"context"
	"fmt"

	todov1beta1 "github.com/odsod/gqlgen-example/internal/gen/proto/go/odsod/todo/v1beta1"
	"github.com/odsod/gqlgen-example/internal/model"
	"go.uber.org/zap"
)

type Mutation struct {
	TodoServiceClient todov1beta1.TodoServiceClient
	Logger            *zap.Logger
}

func (r *Mutation) CreateTodo(ctx context.Context, newTodo model.NewTodo) (*todov1beta1.Todo, error) {
	r.Logger.Debug("create todo", zap.Any("newTodo", newTodo))
	res, err := r.TodoServiceClient.CreateTodo(ctx, &todov1beta1.CreateTodoRequest{
		Todo: &todov1beta1.Todo{
			Text:     newTodo.Text,
			UserName: newTodo.UserName,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("create todo: %w", err)
	}
	return res.Todo, nil
}
