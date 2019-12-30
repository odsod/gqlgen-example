package resolver

import (
	"context"
	"log"

	"github.com/odsod/gqlgen-getting-started/internal/model"
)

type todo struct {
	root *Root
}

func (t *todo) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	log.Printf("todo %s: user %s", obj.ID, obj.UserID)
	return &model.User{
		ID:   "1234",
		Name: "John Smith",
	}, nil
}
