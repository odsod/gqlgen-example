package server

import (
	"context"
	"net/http"

	"go.uber.org/zap"
)

type App struct {
	HTTPServer *http.Server
	Logger     *zap.Logger
}

func (a *App) Run(ctx context.Context) error {
	a.Logger.Info("running")
	defer a.Logger.Info("stopped")
	go func() {
		<-ctx.Done()
		a.Logger.Info("closing HTTP server")
		if err := a.HTTPServer.Close(); err != nil {
			a.Logger.Error("close HTTP server", zap.Error(err))
		}
	}()
	a.Logger.Info("serving HTTP", zap.String("address", a.HTTPServer.Addr))
	if err := a.HTTPServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}
