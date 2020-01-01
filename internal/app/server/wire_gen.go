// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package server

import (
	"context"
	"github.com/odsod/gqlgen-example/internal/middleware"
	"github.com/odsod/gqlgen-example/internal/resolver"
)

// Injectors from wire.go:

func Init(ctx context.Context, cfg *Config) (*App, func(), error) {
	logger, cleanup, err := InitLogger(cfg)
	if err != nil {
		return nil, nil, err
	}
	inMemory, err := InitInMemoryStorage(ctx)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	mutation := &resolver.Mutation{
		Storage: inMemory,
		Logger:  logger,
	}
	query := &resolver.Query{
		Storage: inMemory,
		Logger:  logger,
	}
	todo := &resolver.Todo{
		Logger: logger,
	}
	root := &resolver.Root{
		Storage:          inMemory,
		Logger:           logger,
		MutationResolver: mutation,
		QueryResolver:    query,
		TodoResolver:     todo,
	}
	executableSchema := InitExecutableSchema(root)
	dataloader := &middleware.Dataloader{
		Storage: inMemory,
		Logger:  logger,
	}
	serveMux := InitHTTPServeMux(logger, executableSchema, dataloader)
	server := InitHTTPServer(cfg, serveMux)
	grpcServer := InitGRPCServer()
	app := &App{
		Config:     cfg,
		HTTPServer: server,
		GRPCServer: grpcServer,
		Logger:     logger,
	}
	return app, func() {
		cleanup()
	}, nil
}
