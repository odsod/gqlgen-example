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
	todoServiceClient, cleanup2, err := InitTodoServiceClient(ctx, cfg, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	mutation := &resolver.Mutation{
		TodoServiceClient: todoServiceClient,
		Logger:            logger,
	}
	query := &resolver.Query{
		TodoServiceClient: todoServiceClient,
		Logger:            logger,
	}
	todo := &resolver.Todo{
		Logger: logger,
	}
	root := &resolver.Root{
		Logger:           logger,
		MutationResolver: mutation,
		QueryResolver:    query,
		TodoResolver:     todo,
	}
	executableSchema := InitExecutableSchema(root)
	userServiceClient, cleanup3, err := InitUserServiceClient(ctx, cfg, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	dataloader := &middleware.Dataloader{
		UserServiceClient: userServiceClient,
		Logger:            logger,
	}
	userServiceServer, err := InitUserServiceServer(ctx, logger)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serveMux, err := InitGRPCGatewayServeMux(ctx, userServiceServer)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	httpServeMux := InitHTTPServeMux(cfg, logger, executableSchema, dataloader, serveMux)
	server := InitHTTPServer(httpServeMux)
	todoServiceServer, err := InitTodoServiceServer(ctx, logger)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	grpcServer := InitGRPCServer(userServiceServer, todoServiceServer)
	app := &App{
		Config:     cfg,
		HTTPServer: server,
		GRPCServer: grpcServer,
		Logger:     logger,
	}
	return app, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
