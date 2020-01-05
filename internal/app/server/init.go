package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/handler"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
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
	for _, todo := range []struct {
		id       string
		userName string
		text     string
	}{
		{
			id:       "todo1",
			text:     "Todo 1",
			userName: "users/user1",
		},
		{
			id:       "todo2",
			text:     "Todo 2",
			userName: "users/user2",
		},
		{
			id:       "todo3",
			text:     "Todo 3",
			userName: "users/user2",
		},
		{
			id:       "todo4",
			text:     "Todo 4",
			userName: "users/user3",
		},
	} {
		todo := todo
		if _, err := s.CreateTodo(ctx, &todov1beta1.CreateTodoRequest{
			TodoId: todo.id,
			Todo: &todov1beta1.Todo{
				Text:     todo.text,
				UserName: todo.userName,
			},
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
	s, err := userservice.NewMemory(logger)
	if err != nil {
		return nil, fmt.Errorf("init user service server: %w", err)
	}
	for _, user := range []struct {
		id          string
		displayName string
	}{
		{
			id:          "user1",
			displayName: "User 1",
		},
		{
			id:          "user2",
			displayName: "User 2",
		},
		{
			id:          "user3",
			displayName: "User 3",
		},
	} {
		user := user
		if _, err := s.CreateUser(ctx, &userv1beta1.CreateUserRequest{
			UserId: user.id,
			User: &userv1beta1.User{
				DisplayName: user.displayName,
			},
		}); err != nil {
			return nil, fmt.Errorf("init user service server: %w", err)
		}
	}
	if _, err := s.DeleteUser(ctx, &userv1beta1.DeleteUserRequest{
		Name: "users/user3",
	}); err != nil {
		return nil, fmt.Errorf("init user service server: %w", err)
	}
	return s, nil
}

func InitGRPCGatewayServeMux(
	ctx context.Context,
	userServiceServer userv1beta1.UserServiceServer,
) (*runtime.ServeMux, error) {
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{}),
	)
	if err := userv1beta1.RegisterUserServiceHandlerServer(ctx, mux, userServiceServer); err != nil {
		return nil, fmt.Errorf("init gRPC gateway serve mux: %w", err)
	}
	return mux, nil
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
	grpcGatewayServeMux *runtime.ServeMux,
) *http.ServeMux {
	mux := http.NewServeMux()
	for _, route := range []struct {
		pattern string
		handler http.Handler
	}{
		{
			pattern: c.HTTPServeMux.Patterns.GRPCGateway,
			handler: grpcGatewayServeMux,
		},
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
