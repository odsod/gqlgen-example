package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/handler"
	"github.com/odsod/gqlgen-example/internal/gen/graph"
	todov1beta1 "github.com/odsod/gqlgen-example/internal/gen/proto/go/odsod/todo/v1beta1"
	userv1beta1 "github.com/odsod/gqlgen-example/internal/gen/proto/go/odsod/user/v1beta1"
	"github.com/odsod/gqlgen-example/internal/middleware"
	"github.com/odsod/gqlgen-example/internal/resolver"
	"github.com/odsod/gqlgen-example/internal/service/todoservice"
	"github.com/odsod/gqlgen-example/internal/service/userservice"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func InitTodoServiceServer(
	ctx context.Context,
	logger *zap.Logger,
) (todov1beta1.TodoServiceServer, error) {
	s := todoservice.NewMemory(logger)
	for _, todo := range []*todov1beta1.Todo{
		{
			Id:     "todo1",
			UserId: "user1",
			Text:   "Todo 1",
		},
		{
			Id:     "todo2",
			UserId: "user2",
			Text:   "Todo 2",
		},
		{
			Id:     "todo3",
			UserId: "user2",
			Text:   "Todo 3",
		},
		{
			Id:     "todo4",
			UserId: "user3",
			Text:   "Todo 4",
		},
	} {
		if _, err := s.CreateTodo(ctx, &todov1beta1.CreateTodoRequest{
			TodoId: todo.Id,
			Todo:   todo,
		}); err != nil {
			return nil, fmt.Errorf("init todo service server: %w", err)
		}
	}
	return s, nil
}

func InitUserServiceServer(
	ctx context.Context,
	logger *zap.Logger,
) (userv1beta1.UserServiceServer, error) {
	s := userservice.NewMemory(logger)
	for _, u := range []*userv1beta1.User{
		{
			Id:          "user1",
			DisplayName: "User 1",
		},
		{
			Id:          "user2",
			DisplayName: "User 2",
		},
		{
			Id:          "user3",
			DisplayName: "User 3",
		},
	} {
		if _, err := s.CreateUser(ctx, &userv1beta1.CreateUserRequest{
			UserId: u.Id,
			User:   u,
		}); err != nil {
			return nil, fmt.Errorf("init user service server: %w", err)
		}
	}
	if _, err := s.DeleteUser(ctx, &userv1beta1.DeleteUserRequest{
		Id: "user3",
	}); err != nil {
		return nil, fmt.Errorf("init user service server: %w", err)
	}
	return s, nil
}

func InitGRPCServer(
	userServiceServer userv1beta1.UserServiceServer,
	todoServiceServer todov1beta1.TodoServiceServer,
) *grpc.Server {
	s := grpc.NewServer()
	userv1beta1.RegisterUserServiceServer(s, userServiceServer)
	todov1beta1.RegisterTodoServiceServer(s, todoServiceServer)
	reflection.Register(s)
	return s
}

func InitExecutableSchema(
	rootResolver *resolver.Root,
) graphql.ExecutableSchema {
	return graph.NewExecutableSchema(graph.Config{
		Resolvers: rootResolver,
	})
}

func InitHTTPServeMux(
	c *Config,
	logger *zap.Logger,
	executableSchema graphql.ExecutableSchema,
	dataLoaderMiddleware *middleware.Dataloader,
) *http.ServeMux {
	mux := http.NewServeMux()
	for _, route := range []struct {
		pattern string
		handler http.Handler
	}{
		{
			pattern: c.HTTPServeMux.Patterns.GraphQL,
			handler: dataLoaderMiddleware.ApplyMiddleware(handler.GraphQL(executableSchema)),
		},
		{
			pattern: c.HTTPServeMux.Patterns.GraphQLPlayground,
			handler: handler.Playground("GraphQL playground", c.HTTPServeMux.Patterns.GraphQL),
		},
	} {
		logger.Info("HTTP route", zap.String("pattern", route.pattern))
		mux.Handle(route.pattern, route.handler)
	}
	return mux
}

func InitHTTPServer(
	httpServeMux *http.ServeMux,
) *http.Server {
	return &http.Server{
		Handler: httpServeMux,
	}
}

func InitUserServiceClient(
	ctx context.Context,
	c *Config,
	logger *zap.Logger,
) (userv1beta1.UserServiceClient, func(), error) {
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("localhost:%d", c.GRPCServer.Port), grpc.WithInsecure())
	if err != nil {
		return nil, nil, fmt.Errorf("init user service client: %w", err)
	}
	cleanup := func() {
		logger.Info("closing user service connection")
		if err := conn.Close(); err != nil {
			logger.Error("close user service connection", zap.Error(err))
		}
	}
	return userv1beta1.NewUserServiceClient(conn), cleanup, nil
}

func InitTodoServiceClient(
	ctx context.Context,
	c *Config,
	logger *zap.Logger,
) (todov1beta1.TodoServiceClient, func(), error) {
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("localhost:%d", c.GRPCServer.Port), grpc.WithInsecure())
	if err != nil {
		return nil, nil, fmt.Errorf("init todo service client: %w", err)
	}
	cleanup := func() {
		logger.Info("closing todo service connection")
		if err := conn.Close(); err != nil {
			logger.Error("close todo service connection", zap.Error(err))
		}
	}
	return todov1beta1.NewTodoServiceClient(conn), cleanup, nil
}

func InitLogger(
	c *Config,
) (*zap.Logger, func(), error) {
	var zapConfig zap.Config
	if c.Logger.Development {
		zapConfig = zap.NewDevelopmentConfig()
		zapConfig.EncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	} else {
		zapConfig = zap.NewProductionConfig()
	}
	if err := zapConfig.Level.UnmarshalText([]byte(c.Logger.Level)); err != nil {
		return nil, nil, fmt.Errorf("init logger: %w", err)
	}
	logger, err := zapConfig.Build()
	if err != nil {
		return nil, nil, fmt.Errorf("init logger: %w", err)
	}
	logger.Info("init", zap.Any("config", c))
	cleanup := func() {
		logger.Info("closing logger, goodbye")
		_ = logger.Sync()
	}
	return logger, cleanup, nil
}
