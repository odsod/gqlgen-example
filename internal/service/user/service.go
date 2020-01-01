package user

import (
	"context"
	"sync"

	userv1beta1 "github.com/odsod/gqlgen-example/internal/gen/proto/go/odsod/user/v1beta1"
	"go.uber.org/zap"
)

type Service struct {
	logger *zap.Logger
	mu     sync.Mutex
	users  map[string]*userv1beta1.User
}

func NewService(logger *zap.Logger) *Service {
	return &Service{
		logger: logger,
		users:  map[string]*userv1beta1.User{},
	}
}

func (s *Service) ListUsers(context.Context, *userv1beta1.ListUsersRequest) (*userv1beta1.ListUsersResponse, error) {
	panic("implement me")
}

func (s *Service) GetUser(context.Context, *userv1beta1.GetUserRequest) (*userv1beta1.GetUserResponse, error) {
	panic("implement me")
}

func (s *Service) BatchGetUsers(_ context.Context, req *userv1beta1.BatchGetUsersRequest) (*userv1beta1.BatchGetUsersResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var res userv1beta1.BatchGetUsersResponse
	for _, id := range req.Ids {
		if user, ok := s.users[id]; ok {
			res.FoundUsers = append(res.FoundUsers, user)
		} else {
			res.MissingIds = append(res.MissingIds, id)
		}
	}
	return &res, nil
}

func (s *Service) CreateUser(_ context.Context, req *userv1beta1.CreateUserRequest) (*userv1beta1.CreateUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	req.User.Id = req.UserId
	s.users[req.UserId] = req.User
	return &userv1beta1.CreateUserResponse{
		User: req.User,
	}, nil
}

func (s *Service) UpdateUser(context.Context, *userv1beta1.UpdateUserRequest) (*userv1beta1.UpdateUserResponse, error) {
	panic("implement me")
}

func (s *Service) DeleteUser(context.Context, *userv1beta1.DeleteUserRequest) (*userv1beta1.DeleteUserResponse, error) {
	panic("implement me")
}
