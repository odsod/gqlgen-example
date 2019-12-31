package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/odsod/gqlgen-getting-started/internal/app/server"
)

func main() {
	configFile := flag.String("config", "config.toml", "config file")
	flag.Parse()
	var cfg server.Config
	if err := cfg.LoadFile(*configFile); err != nil {
		log.Panic(err)
	}
	shutdownSignals := []os.Signal{syscall.SIGINT, syscall.SIGTERM}
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sigChan := make(chan os.Signal, len(shutdownSignals))
		signal.Notify(sigChan, shutdownSignals...)
		<-sigChan
		cancel()
		signal.Stop(sigChan)
		close(sigChan)
	}()
	app, cleanup, err := server.Init(ctx, &cfg)
	if err != nil {
		log.Panic(err)
	}
	defer cleanup()
	if err := app.Run(ctx); err != nil {
		log.Panic(err)
	}
}
