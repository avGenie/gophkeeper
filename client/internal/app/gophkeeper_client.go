package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/avGenie/gophkeeper/client/internal/config"
	"github.com/avGenie/gophkeeper/client/internal/controller"
	"github.com/avGenie/gophkeeper/client/internal/controller/grpc"
	"go.uber.org/zap"
)

// App Main gophkeeper client application
type App struct {
	client controller.GophkeeperClient
}

// NewApp Creates main application
func NewApp(config config.Config) (*App, error) {
	client, err := grpc.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("couldn't create client: %w", err)
	}

	return &App{
		client: client,
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

	go a.client.Start()

	<-ctx.Done()

	zap.L().Info("Got interruption signal. Shutting down HTTP client gracefully...")

	a.client.Stop()
}
