//+build wireinject

package server

import (
	"context"

	"github.com/google/wire"
	"github.com/odsod/gqlgen-example/internal/middleware"
	"github.com/odsod/gqlgen-example/internal/resolver"
)

func Init(ctx context.Context, cfg *Config) (*App, func(), error) {
	panic(
		wire.Build(
			wire.Struct(new(App), "*"),
			InitLogger,
			InitHTTPServer,
			InitHTTPServeMux,
			InitExecutableSchema,
			InitGRPCServer,
			InitGRPCGatewayServeMux,
			InitUserServiceServer,
			InitUserServiceClient,
			InitTodoServiceServer,
			InitTodoServiceClient,
			wire.Struct(new(resolver.Root), "*"),
			wire.Struct(new(resolver.Query), "*"),
			wire.Struct(new(resolver.Mutation), "*"),
			wire.Struct(new(resolver.Todo), "*"),
			wire.Struct(new(middleware.Dataloader), "*"),
		),
	)
}
