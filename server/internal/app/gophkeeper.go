package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/avGenie/gophkeeper/server/internal/config"
	"github.com/avGenie/gophkeeper/server/internal/controller"
	"github.com/avGenie/gophkeeper/server/internal/controller/grpc"
	"go.uber.org/zap"
)

// App Main gophkeeper server application
type App struct {
	config config.Config
	server controller.GophkeeperServer
}

// NewApp Creates main application
func NewApp(config config.Config) (*App, error) {
	server, err := grpc.NewServer(config)
	if err != nil {
		return nil, fmt.Errorf("couldn't create server: %w", err)
	}

	return &App{
		config: config,
		server: server,
	}, nil
}

// Run Starts main application
func (a *App) Run() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGQUIT,
		os.Interrupt,
	)
	defer cancel()

	go a.server.Start()

	<-ctx.Done()

	zap.L().Info("Got interruption signal. Shutting down HTTP server gracefully...")

	a.server.Stop()
}
