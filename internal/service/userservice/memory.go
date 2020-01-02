package userservice

import (
	"context"
	"encoding/base64"
	"fmt"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	userv1beta1 "github.com/odsod/gqlgen-example/internal/gen/proto/go/odsod/user/v1beta1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Memory struct {
	logger *zap.Logger
	mu     sync.Mutex
	users  []*userv1beta1.User
}

var _ userv1beta1.UserServiceServer = &Memory{}

func NewMemory(logger *zap.Logger) *Memory {
	return &Memory{
		logger: logger,
	}
}

func (s *Memory) ListUsers(
	_ context.Context,
	r *userv1beta1.ListUsersRequest,
) (*userv1beta1.ListUsersResponse, error) {
	var pageToken listUsersPageToken
	if err := pageToken.UnmarshalString(r.PageToken); err != nil {
		s.logger.Error("parse page token", zap.Error(err), zap.String("pageToken", r.PageToken))
		return nil, status.Error(codes.InvalidArgument, "malformed page token")
	}
	users := s.users
	if !r.ShowDeleted {
		users = make([]*userv1beta1.User, 0, len(s.users))
		for _, u := range s.users {
			if !u.Deleted {
				users = append(users, u)
			}
		}
	}
	res := &userv1beta1.ListUsersResponse{
		Users:     make([]*userv1beta1.User, 0, r.PageSize),
		TotalSize: int32(len(users)),
	}
	for int(pageToken.Offset) < len(users) && len(res.Users) < int(r.PageSize) {
		u := s.users[pageToken.Offset]
		pageToken.Offset++
		if u.Deleted && !r.ShowDeleted {
			continue
		}
		res.Users = append(res.Users, u)
	}
	if int(pageToken.Offset) < len(users) {
		str, err := pageToken.MarshalString()
		if err != nil {
			s.logger.Error("generate page token", zap.Error(err))
			return nil, status.Error(codes.Internal, "failed to generate page token")
		}
		res.NextPageToken = str
	}
	return res, nil
}

func (s *Memory) GetUser(
	_ context.Context,
	r *userv1beta1.GetUserRequest,
) (*userv1beta1.GetUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	u, _, ok := s.getUser(r.Id)
	if !ok {
		return nil, status.Error(codes.NotFound, "user not found")
	}
	res := &userv1beta1.GetUserResponse{
		User: proto.Clone(u).(*userv1beta1.User),
	}
	return res, nil
}

func (s *Memory) BatchGetUsers(
	_ context.Context,
	r *userv1beta1.BatchGetUsersRequest,
) (*userv1beta1.BatchGetUsersResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	res := &userv1beta1.BatchGetUsersResponse{
		FoundUsers: make([]*userv1beta1.User, 0, len(r.Ids)),
	}
	for _, id := range r.Ids {
		if user, _, ok := s.getUser(id); ok {
			res.FoundUsers = append(res.FoundUsers, proto.Clone(user).(*userv1beta1.User))
		} else {
			res.MissingIds = append(res.MissingIds, id)
		}
	}
	return res, nil
}

func (s *Memory) CreateUser(
	_ context.Context,
	r *userv1beta1.CreateUserRequest,
) (*userv1beta1.CreateUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	newUser := proto.Clone(r.User).(*userv1beta1.User)
	if r.UserId != "" {
		newUser.Id = r.UserId
	} else {
		newUser.Id = uuid.New().String()
	}
	newUser.CreateTime = ptypes.TimestampNow()
	newUser.UpdateTime = newUser.CreateTime
	if _, _, ok := s.getUser(newUser.Id); ok {
		return nil, status.Error(codes.AlreadyExists, "user already exists")
	}
	s.users = append(s.users, newUser)
	res := &userv1beta1.CreateUserResponse{
		User: proto.Clone(newUser).(*userv1beta1.User),
	}
	return res, nil
}

func (s *Memory) UpdateUser(
	_ context.Context,
	r *userv1beta1.UpdateUserRequest,
) (*userv1beta1.UpdateUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(r.UpdateMask.Paths) > 0 {
		return nil, status.Error(codes.Unimplemented, "field mask support not yet implemented")
	}
	_, i, ok := s.getUser(r.User.Id)
	if !ok {
		return nil, status.Error(codes.NotFound, "user not found")
	}
	updatedUser := proto.Clone(r.User).(*userv1beta1.User)
	updatedUser.UpdateTime = ptypes.TimestampNow()
	s.users[i] = updatedUser
	res := &userv1beta1.UpdateUserResponse{
		User: proto.Clone(updatedUser).(*userv1beta1.User),
	}
	return res, nil
}

func (s *Memory) DeleteUser(
	_ context.Context,
	r *userv1beta1.DeleteUserRequest,
) (*userv1beta1.DeleteUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	u, _, ok := s.getUser(r.Id)
	if !ok {
		return nil, status.Error(codes.NotFound, "user not found")
	}
	if u.Deleted {
		return nil, status.Error(codes.FailedPrecondition, "user already deleted")
	}
	u.Deleted = true
	u.DeleteTime = ptypes.TimestampNow()
	u.UpdateTime = ptypes.TimestampNow()
	res := &userv1beta1.DeleteUserResponse{
		User: proto.Clone(u).(*userv1beta1.User),
	}
	return res, nil
}

func (s *Memory) UndeleteUser(
	_ context.Context,
	r *userv1beta1.UndeleteUserRequest,
) (*userv1beta1.UndeleteUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	u, _, ok := s.getUser(r.Id)
	if !ok {
		return nil, status.Error(codes.NotFound, "user not found")
	}
	if !u.Deleted {
		return nil, status.Error(codes.FailedPrecondition, "user not already deleted")
	}
	u.Deleted = false
	u.DeleteTime = nil
	u.UpdateTime = ptypes.TimestampNow()
	res := &userv1beta1.UndeleteUserResponse{
		User: proto.Clone(u).(*userv1beta1.User),
	}
	return res, nil
}

func (s *Memory) getUser(id string) (*userv1beta1.User, int, bool) {
	for i, u := range s.users {
		if u.Id == id {
			return u, i, true
		}
	}
	return nil, 0, false
}

type listUsersPageToken struct {
	userv1beta1.ListUsersPageToken
}

func (p *listUsersPageToken) UnmarshalString(s string) error {
	data, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return fmt.Errorf("unmarshal page token string '%s': %w", s, err)
	}
	if err := proto.Unmarshal(data, p); err != nil {
		return fmt.Errorf("unmarshal page token string '%s': %w", s, err)
	}
	return nil
}

func (p *listUsersPageToken) MarshalString() (string, error) {
	data, err := proto.Marshal(p)
	if err != nil {
		return "", fmt.Errorf("marshal page token string: %w", err)
	}
	s := base64.URLEncoding.EncodeToString(data)
	return s, nil
}
