package storage

import (
	"context"

	todov1beta1 "github.com/odsod/gqlgen-example/internal/gen/proto/go/odsod/todo/v1beta1"
	userv1beta1 "github.com/odsod/gqlgen-example/internal/gen/proto/go/odsod/user/v1beta1"
)

type InMemory struct {
	todos map[string]*todov1beta1.Todo
	users map[string]*userv1beta1.User
}

func NewInMemory() *InMemory {
	return &InMemory{
		todos: map[string]*todov1beta1.Todo{},
		users: map[string]*userv1beta1.User{},
	}
}

func (s *InMemory) UpdateUser(_ context.Context, user *userv1beta1.User) (*userv1beta1.User, error) {
	s.users[user.Id] = user
	return user, nil
}

func (s *InMemory) GetUser(_ context.Context, id string) (*userv1beta1.User, bool, error) {
	user, ok := s.users[id]
	return user, ok, nil
}

func (s *InMemory) BatchGetUsers(_ context.Context, ids []string) ([]*userv1beta1.User, []string, error) {
	var foundUsers []*userv1beta1.User
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
