package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type App struct {
	Config     *Config
	HTTPServer *http.Server
	GRPCServer *grpc.Server
	Logger     *zap.Logger
}

func (a *App) Run(ctx context.Context) error {
	a.Logger.Info("running")
	defer a.Logger.Info("stopped")
	var listenConfig net.ListenConfig
	grpcServerListener, err := listenConfig.Listen(ctx, "tcp", fmt.Sprintf(":%d", a.Config.GRPCServer.Port))
	if err != nil {
		return fmt.Errorf("bind gRPC server listener: %w", err)
	}
	httpServerListener, err := listenConfig.Listen(ctx, "tcp", fmt.Sprintf(":%d", a.Config.HTTPServer.Port))
	if err != nil {
		return fmt.Errorf("bind HTTP server listener: %w", err)
	}
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		a.Logger.Info("serving HTTP", zap.Int("port", a.Config.HTTPServer.Port))
		if err := a.HTTPServer.Serve(httpServerListener); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})
	g.Go(func() error {
		<-ctx.Done()
		a.Logger.Info("closing HTTP server")
		return a.HTTPServer.Close()
	})
	g.Go(func() error {
		a.Logger.Info("serving gRPC", zap.Int("port", a.Config.GRPCServer.Port))
		return a.GRPCServer.Serve(grpcServerListener)
	})
	g.Go(func() error {
		<-ctx.Done()
		a.Logger.Info("stopping gRPC server")
		a.GRPCServer.GracefulStop()
		return nil
	})
	return g.Wait()
}
