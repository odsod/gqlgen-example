package resolver

import (
	"context"
	"fmt"

	todov1beta1 "github.com/odsod/gqlgen-example/internal/gen/proto/go/odsod/todo/v1beta1"
	"github.com/odsod/gqlgen-example/internal/middleware"
	"github.com/odsod/gqlgen-example/internal/model"
	"go.uber.org/zap"
)

type Todo struct {
	Logger *zap.Logger
}

func (t *Todo) User(ctx context.Context, todo *todov1beta1.Todo) (*model.User, error) {
	t.Logger.Debug("todo: user", zap.Any("todo", todo))
	userLoader, ok := middleware.UserLoaderFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("resolve todo %s user %s: no user data loader in context", todo.Id, todo.UserId)
	}
	user, err := userLoader.Load(todo.UserId)
	if err != nil {
		return nil, fmt.Errorf("resolve todo %s user %s: %w", todo.Id, todo.UserId, err)
	}
	return user, nil
}
