package storage

import (
	"context"

	"github.com/odsod/gqlgen-example/internal/model"
)

type InMemory struct {
	todos map[string]*model.Todo
	users map[string]*model.User
}

func NewInMemory() *InMemory {
	return &InMemory{
		todos: map[string]*model.Todo{},
		users: map[string]*model.User{},
	}
}

func (s *InMemory) UpdateUser(_ context.Context, user *model.User) (*model.User, error) {
	s.users[user.ID] = user
	return user, nil
}

func (s *InMemory) GetUser(_ context.Context, id string) (*model.User, bool, error) {
	user, ok := s.users[id]
	return user, ok, nil
}

func (s *InMemory) BatchGetUsers(_ context.Context, ids []string) ([]*model.User, []string, error) {
	var foundUsers []*model.User
	var notFoundIDs []string
	for _, id := range ids {
		if user, ok := s.users[id]; ok {
			foundUsers = append(foundUsers, user)
		} else {
			notFoundIDs = append(notFoundIDs, id)
		}
	}
	return foundUsers, notFoundIDs, nil
}

func (s *InMemory) UpdateTodo(_ context.Context, user *model.Todo) (*model.Todo, error) {
	s.todos[user.ID] = user
	return user, nil
}

func (s *InMemory) GetTodo(_ context.Context, id string) (*model.Todo, bool, error) {
	user, ok := s.todos[id]
	return user, ok, nil
}

func (s *InMemory) BatchGetTodos(_ context.Context, ids []string) ([]*model.Todo, []string, error) {
	var foundTodos []*model.Todo
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

func (s *InMemory) ListTodos(_ context.Context) ([]*model.Todo, error) {
	result := make([]*model.Todo, 0, len(s.todos))
	for _, todo := range s.todos {
		result = append(result, todo)
	}
	return result, nil
}
