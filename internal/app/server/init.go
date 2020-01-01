package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/handler"
	"github.com/golang/protobuf/ptypes"
	"github.com/odsod/gqlgen-example/internal/gen/graph"
	todov1beta1 "github.com/odsod/gqlgen-example/internal/gen/proto/go/odsod/todo/v1beta1"
	userv1beta1 "github.com/odsod/gqlgen-example/internal/gen/proto/go/odsod/user/v1beta1"
	"github.com/odsod/gqlgen-example/internal/middleware"
	"github.com/odsod/gqlgen-example/internal/resolver"
	"github.com/odsod/gqlgen-example/internal/storage"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitInMemoryStorage(
	ctx context.Context,
) (*storage.InMemory, error) {
	inMemoryStorage := storage.NewInMemory()
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
		if _, err := inMemoryStorage.UpdateTodo(ctx, todo); err != nil {
			return nil, fmt.Errorf("init in-memory storage: %w", err)
		}
	}
	for _, user := range []*userv1beta1.User{
		{
			Id:          "user1",
			DisplayName: "User 1",
			CreateTime:  ptypes.TimestampNow(),
			UpdateTime:  ptypes.TimestampNow(),
		},
		{
			Id:          "user2",
			DisplayName: "User 2",
			CreateTime:  ptypes.TimestampNow(),
			UpdateTime:  ptypes.TimestampNow(),
		},
		{
			Id:          "user3",
			DisplayName: "User 3",
			Deleted:     true,
			CreateTime:  ptypes.TimestampNow(),
			UpdateTime:  ptypes.TimestampNow(),
			DeleteTime:  ptypes.TimestampNow(),
		},
	} {
		if _, err := inMemoryStorage.UpdateUser(ctx, user); err != nil {
			return nil, fmt.Errorf("init in-memory storage: %w", err)
		}
	}
	return inMemoryStorage, nil
}

func InitExecutableSchema(
	rootResolver *resolver.Root,
) graphql.ExecutableSchema {
	return graph.NewExecutableSchema(graph.Config{
		Resolvers: rootResolver,
	})
}

func InitHTTPServeMux(
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
			pattern: "/graphql",
			handler: dataLoaderMiddleware.ApplyMiddleware(handler.GraphQL(executableSchema)),
		},
		{
			pattern: "/",
			handler: handler.Playground("GraphQL playground", "/graphql"),
		},
	} {
		logger.Info("route", zap.String("pattern", route.pattern))
		mux.Handle(route.pattern, route.handler)
	}
	return mux
}

func InitHTTPServer(
	c *Config,
	httpServeMux *http.ServeMux,
) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", c.HTTPServer.Port),
		Handler: httpServeMux,
	}
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
