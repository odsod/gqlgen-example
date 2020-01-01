package resolver

import (
	"context"
	"fmt"

	"github.com/odsod/gqlgen-example/internal/middleware"
	"github.com/odsod/gqlgen-example/internal/model"
	"go.uber.org/zap"
)

type Todo struct {
	Logger *zap.Logger
}

func (t *Todo) User(ctx context.Context, todo *model.Todo) (*model.User, error) {
	t.Logger.Debug("todo: user", zap.Any("todo", todo))
	userLoader, ok := middleware.UserLoaderFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("resolve todo %s user %s: no user data loader in context", todo.ID, todo.UserID)
	}
	user, err := userLoader.Load(todo.UserID)
	if err != nil {
		return nil, fmt.Errorf("resolve todo %s user %s: %w", todo.ID, todo.UserID, err)
	}
	return user, nil
}
