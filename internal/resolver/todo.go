package resolver

import (
	"context"
	"fmt"

	todov1beta1 "github.com/odsod/gqlgen-example/internal/gen/proto/go/odsod/todo/v1beta1"
	userv1beta1 "github.com/odsod/gqlgen-example/internal/gen/proto/go/odsod/user/v1beta1"
	"github.com/odsod/gqlgen-example/internal/middleware"
	"go.uber.org/zap"
)

type Todo struct {
	Logger *zap.Logger
}

func (t *Todo) User(ctx context.Context, todo *todov1beta1.Todo) (*userv1beta1.User, error) {
	t.Logger.Debug("todo: user", zap.Any("todo", todo))
	userLoader, ok := middleware.UserLoaderFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("resolve todo %s user %s: no user data loader in context", todo.Name, todo.UserName)
	}
	user, err := userLoader.Load(todo.UserName)
	if err != nil {
		return nil, fmt.Errorf("resolve todo %s user %s: %w", todo.UserName, todo.UserName, err)
	}
	return user, nil
}
