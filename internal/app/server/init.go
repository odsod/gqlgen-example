package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/handler"
	"github.com/odsod/gqlgen-getting-started/internal/dataloader"
	"github.com/odsod/gqlgen-getting-started/internal/graph"
	"github.com/odsod/gqlgen-getting-started/internal/model"
	"github.com/odsod/gqlgen-getting-started/internal/resolver"
	"github.com/odsod/gqlgen-getting-started/internal/storage"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitInMemoryStorage(
	ctx context.Context,
) (*storage.InMemory, error) {
	inMemoryStorage := storage.NewInMemory()
	for _, todo := range []*model.Todo{
		{
			ID:     "todo1",
			UserID: "user1",
			Text:   "Todo 1",
		},
		{
			ID:     "todo2",
			UserID: "user2",
			Text:   "Todo 2",
		},
		{
			ID:     "todo3",
			UserID: "user2",
			Text:   "Todo 3",
		},
		{
			ID:     "todo4",
			UserID: "user3",
			Text:   "Todo 4",
		},
	} {
		if _, err := inMemoryStorage.UpdateTodo(ctx, todo); err != nil {
			return nil, fmt.Errorf("init in-memory storage: %w", err)
		}
	}
	for _, user := range []*model.User{
		{ID: "user1", Name: "User 1"},
		{ID: "user2", Name: "User 2"},
		{ID: "user3", Name: "User 3"},
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
	dataLoaderMiddleware *dataloader.Middleware,
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
