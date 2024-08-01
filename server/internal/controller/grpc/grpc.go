package grpc

import (
	"fmt"
	"net"
	"time"

	"github.com/avGenie/gophkeeper/package/tls"
	"github.com/avGenie/gophkeeper/server/internal/config"
	"github.com/avGenie/gophkeeper/server/internal/storage"
	"github.com/avGenie/gophkeeper/server/internal/storage/postgres"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	pb "github.com/avGenie/gophkeeper/proto"
)

const requestTimeout = 3 * time.Second

// GRPCServer GRPC server struct
type GRPCServer struct {
	pb.GophkeeperServer

	server *grpc.Server

	config  config.Config
	storage storage.Storage
}

// NewServer Creates GRPC server
func NewServer(config config.Config) (*GRPCServer, error) {
	storage, err := postgres.NewStorage(config.Storage)
	if err != nil {
		return nil, fmt.Errorf("couldn't create storage: %w", err)
	}

	tlsCreds, err := tls.GenerateServerTLSCreds(config.Server.CertsPath)
	if err != nil {
		return nil, fmt.Errorf("couldn't generate tls creds: %w", err)
	}

	return &GRPCServer{
		server:  grpc.NewServer(grpc.Creds(tlsCreds)),
		config:  config,
		storage: storage,
	}, nil
}

// Start Starts GRPC server
func (s *GRPCServer) Start() {
	listen, err := net.Listen("tcp", s.config.Server.GRPCAddress)
	if err != nil {
		zap.L().Fatal("failed to listen grpc address", zap.Error(err))

		return
	}

	pb.RegisterGophkeeperServer(s.server, s)

	zap.L().Info("Server gRPC starts", zap.String("address:", s.config.Server.GRPCAddress))

	if err := s.server.Serve(listen); err != nil {
		zap.L().Fatal("failed to serve grpc server", zap.Error(err))

		return
	}
}

// Stop Stops GRPC server
func (s *GRPCServer) Stop() {
	s.storage.Close()
	s.server.GracefulStop()
}
