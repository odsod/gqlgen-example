package resolver

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/odsod/gqlgen-getting-started/internal/model"
)

type mutationResolver struct {
	*Root
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	log.Printf("mutation: createTodo: %+v", input)
	todo := &model.Todo{
		ID:     fmt.Sprintf("T%d", rand.Int()),
		Text:   input.Text,
		UserID: input.UserID,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}
