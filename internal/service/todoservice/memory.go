package todoservice

import (
	"context"
	"encoding/base64"
	"fmt"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	todov1beta1 "github.com/odsod/gqlgen-example/internal/gen/proto/go/odsod/todo/v1beta1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Memory struct {
	logger *zap.Logger
	mu     sync.Mutex
	todos  []*todov1beta1.Todo
}

var _ todov1beta1.TodoServiceServer = &Memory{}

func NewMemory(logger *zap.Logger) *Memory {
	return &Memory{
		logger: logger,
	}
}

func (s *Memory) ListTodos(
	_ context.Context,
	r *todov1beta1.ListTodosRequest,
) (*todov1beta1.ListTodosResponse, error) {
	var pageToken listTodosPageToken
	if err := pageToken.UnmarshalString(r.PageToken); err != nil {
		s.logger.Error("parse page token", zap.Error(err), zap.String("pageToken", r.PageToken))
		return nil, status.Error(codes.InvalidArgument, "malformed page token")
	}
	todos := s.todos
	if !r.ShowDeleted {
		todos = make([]*todov1beta1.Todo, 0, len(s.todos))
		for _, u := range s.todos {
			if !u.Deleted {
				todos = append(todos, u)
			}
		}
	}
	res := &todov1beta1.ListTodosResponse{
		Todos:     make([]*todov1beta1.Todo, 0, r.PageSize),
		TotalSize: int32(len(todos)),
	}
	for int(pageToken.Offset) < len(todos) && len(res.Todos) < int(r.PageSize) {
		u := s.todos[pageToken.Offset]
		pageToken.Offset++
		if u.Deleted && !r.ShowDeleted {
			continue
		}
		res.Todos = append(res.Todos, u)
	}
	if int(pageToken.Offset) < len(todos) {
		str, err := pageToken.MarshalString()
		if err != nil {
			s.logger.Error("generate page token", zap.Error(err))
			return nil, status.Error(codes.Internal, "failed to generate page token")
		}
		res.NextPageToken = str
	}
	return res, nil
}

func (s *Memory) GetTodo(
	_ context.Context,
	r *todov1beta1.GetTodoRequest,
) (*todov1beta1.GetTodoResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	todo, _, ok := s.getTodoByName(r.Name)
	if !ok {
		return nil, status.Error(codes.NotFound, "todo not found")
	}
	res := &todov1beta1.GetTodoResponse{
		Todo: proto.Clone(todo).(*todov1beta1.Todo),
	}
	return res, nil
}

func (s *Memory) BatchGetTodos(
	_ context.Context,
	r *todov1beta1.BatchGetTodosRequest,
) (*todov1beta1.BatchGetTodosResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	res := &todov1beta1.BatchGetTodosResponse{
		FoundTodos: make([]*todov1beta1.Todo, 0, len(r.Names)),
	}
	for _, name := range r.Names {
		if todo, _, ok := s.getTodoByName(name); ok {
			res.FoundTodos = append(res.FoundTodos, proto.Clone(todo).(*todov1beta1.Todo))
		} else {
			res.MissingNames = append(res.MissingNames, name)
		}
	}
	return res, nil
}

func (s *Memory) CreateTodo(
	_ context.Context,
	r *todov1beta1.CreateTodoRequest,
) (*todov1beta1.CreateTodoResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	newTodo := proto.Clone(r.Todo).(*todov1beta1.Todo)
	var name TodoResourceName
	if r.TodoId != "" {
		name.ID = r.TodoId
	} else {
		name.ID = uuid.New().String()
	}
	newTodo.Name = name.String()
	if _, _, ok := s.getTodoByName(newTodo.Name); ok {
		s.logger.Error("new todo already exists", zap.String("name", newTodo.Name))
		return nil, status.Error(codes.AlreadyExists, "todo already exists")
	}
	newTodo.CreateTime = ptypes.TimestampNow()
	newTodo.UpdateTime = newTodo.CreateTime
	s.todos = append(s.todos, newTodo)
	res := &todov1beta1.CreateTodoResponse{
		Todo: proto.Clone(newTodo).(*todov1beta1.Todo),
	}
	return res, nil
}

func (s *Memory) UpdateTodo(
	_ context.Context,
	r *todov1beta1.UpdateTodoRequest,
) (*todov1beta1.UpdateTodoResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(r.UpdateMask.Paths) > 0 {
		return nil, status.Error(codes.Unimplemented, "field mask support not yet implemented")
	}
	_, i, ok := s.getTodoByName(r.Todo.Name)
	if !ok {
		return nil, status.Error(codes.NotFound, "todo not found")
	}
	updatedTodo := proto.Clone(r.Todo).(*todov1beta1.Todo)
	updatedTodo.UpdateTime = ptypes.TimestampNow()
	s.todos[i] = updatedTodo
	res := &todov1beta1.UpdateTodoResponse{
		Todo: proto.Clone(updatedTodo).(*todov1beta1.Todo),
	}
	return res, nil
}

func (s *Memory) DeleteTodo(
	_ context.Context,
	r *todov1beta1.DeleteTodoRequest,
) (*todov1beta1.DeleteTodoResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	todo, _, ok := s.getTodoByName(r.Name)
	if !ok {
		return nil, status.Error(codes.NotFound, "todo not found")
	}
	if todo.Deleted {
		return nil, status.Error(codes.FailedPrecondition, "todo already deleted")
	}
	todo.Deleted = true
	todo.DeleteTime = ptypes.TimestampNow()
	todo.UpdateTime = ptypes.TimestampNow()
	res := &todov1beta1.DeleteTodoResponse{
		Todo: proto.Clone(todo).(*todov1beta1.Todo),
	}
	return res, nil
}

func (s *Memory) UndeleteTodo(
	_ context.Context,
	r *todov1beta1.UndeleteTodoRequest,
) (*todov1beta1.UndeleteTodoResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	todo, _, ok := s.getTodoByName(r.Name)
	if !ok {
		return nil, status.Error(codes.NotFound, "todo not found")
	}
	if !todo.Deleted {
		return nil, status.Error(codes.FailedPrecondition, "todo not already deleted")
	}
	todo.Deleted = false
	todo.DeleteTime = nil
	todo.UpdateTime = ptypes.TimestampNow()
	res := &todov1beta1.UndeleteTodoResponse{
		Todo: proto.Clone(todo).(*todov1beta1.Todo),
	}
	return res, nil
}

func (s *Memory) getTodoByName(name string) (*todov1beta1.Todo, int, bool) {
	for i, todo := range s.todos {
		if todo.Name == name {
			return todo, i, true
		}
	}
	return nil, 0, false
}

type listTodosPageToken struct {
	todov1beta1.ListTodosPageToken
}

func (p *listTodosPageToken) UnmarshalString(s string) error {
	data, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return fmt.Errorf("unmarshal page token string '%s': %w", s, err)
	}
	if err := proto.Unmarshal(data, p); err != nil {
		return fmt.Errorf("unmarshal page token string '%s': %w", s, err)
	}
	return nil
}

func (p *listTodosPageToken) MarshalString() (string, error) {
	data, err := proto.Marshal(p)
	if err != nil {
		return "", fmt.Errorf("marshal page token string: %w", err)
	}
	s := base64.URLEncoding.EncodeToString(data)
	return s, nil
}
