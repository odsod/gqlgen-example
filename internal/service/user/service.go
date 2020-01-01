package user

import (
	"context"

	userv1beta1 "github.com/odsod/gqlgen-example/internal/gen/proto/go/odsod/user/v1beta1"
)

type Service struct{}

func (s *Service) ListUsers(context.Context, *userv1beta1.ListUsersRequest) (*userv1beta1.ListUsersResponse, error) {
	panic("implement me")
}

func (s *Service) GetUser(context.Context, *userv1beta1.GetUserRequest) (*userv1beta1.GetUserResponse, error) {
	panic("implement me")
}

func (s *Service) BatchGetUsers(context.Context, *userv1beta1.BatchGetUsersRequest) (*userv1beta1.BatchGetUsersResponse, error) {
	panic("implement me")
}

func (s *Service) CreateUser(context.Context, *userv1beta1.CreateUserRequest) (*userv1beta1.CreateUserResponse, error) {
	panic("implement me")
}

func (s *Service) UpdateUser(context.Context, *userv1beta1.UpdateUserRequest) (*userv1beta1.UpdateUserResponse, error) {
	panic("implement me")
}

func (s *Service) DeleteUser(context.Context, *userv1beta1.DeleteUserRequest) (*userv1beta1.DeleteUserResponse, error) {
	panic("implement me")
}
