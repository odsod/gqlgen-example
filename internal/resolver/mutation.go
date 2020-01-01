package resolver

import (
	"context"
	"fmt"
	"math/rand"

	todov1beta1 "github.com/odsod/gqlgen-example/internal/gen/proto/go/odsod/todo/v1beta1"
	"github.com/odsod/gqlgen-example/internal/model"
	"github.com/odsod/gqlgen-example/internal/storage"
	"go.uber.org/zap"
)

type Mutation struct {
	Storage *storage.InMemory
	Logger  *zap.Logger
}

func (r *Mutation) CreateTodo(ctx context.Context, newTodo model.NewTodo) (*todov1beta1.Todo, error) {
	r.Logger.Debug("create todo", zap.Any("newTodo", newTodo))
	todo, err := r.Storage.UpdateTodo(ctx, &todov1beta1.Todo{
		Id:     fmt.Sprintf("T%d", rand.Int()),
		Text:   newTodo.Text,
		UserId: newTodo.UserID,
	})
	if err != nil {
		return nil, fmt.Errorf("create todo: %w", err)
	}
	return todo, nil
}
