package storage

import (
	"context"

	todov1beta1 "github.com/odsod/gqlgen-example/internal/gen/proto/go/odsod/todo/v1beta1"
)

type InMemory struct {
	todos map[string]*todov1beta1.Todo
}

func NewInMemory() *InMemory {
	return &InMemory{
		todos: map[string]*todov1beta1.Todo{},
	}
}

func (s *InMemory) UpdateTodo(_ context.Context, todo *todov1beta1.Todo) (*todov1beta1.Todo, error) {
	s.todos[todo.Id] = todo
	return todo, nil
}

func (s *InMemory) GetTodo(_ context.Context, id string) (*todov1beta1.Todo, bool, error) {
	user, ok := s.todos[id]
	return user, ok, nil
}

func (s *InMemory) BatchGetTodos(_ context.Context, ids []string) ([]*todov1beta1.Todo, []string, error) {
	var foundTodos []*todov1beta1.Todo
	var notFoundIDs []string
	for _, id := range ids {
		if user, ok := s.todos[id]; ok {
			foundTodos = append(foundTodos, user)
		} else {
			notFoundIDs = append(notFoundIDs, id)
		}
	}
	return foundTodos, notFoundIDs, nil
}

func (s *InMemory) ListTodos(_ context.Context) ([]*todov1beta1.Todo, error) {
	result := make([]*todov1beta1.Todo, 0, len(s.todos))
	for _, todo := range s.todos {
		result = append(result, todo)
	}
	return result, nil
}
