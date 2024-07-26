package grpc

import (
	"fmt"
	"net"
	"time"

	"github.com/avGenie/gophkeeper/server/internal/config"
	"github.com/avGenie/gophkeeper/server/internal/storage"
	"github.com/avGenie/gophkeeper/server/internal/storage/postgres"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	pb "github.com/avGenie/gophkeeper/proto"
)

const requestTimeout = 3 * time.Second

type GRPCServer struct {
	pb.GophkeeperServer

	server *grpc.Server

	config  config.Config
	storage storage.Storage
}

func NewServer(config config.Config) (*GRPCServer, error) {
	storage, err := postgres.NewStorage(config.Storage)
	if err != nil {
		return nil, fmt.Errorf("couldn't create storage: %w", err)
	}

	return &GRPCServer{
		server:  grpc.NewServer(),
		config:  config,
		storage: storage,
	}, nil
}

func (s *GRPCServer) Start() {
	listen, err := net.Listen("tcp", s.config.Server.GRPCAddress)
	if err != nil {
		zap.L().Fatal("failed to listen grpc address", zap.Error(err))

		return
	}

	// регистрируем сервис
	pb.RegisterGophkeeperServer(s.server, s)

	zap.L().Info("Server gRPC starts", zap.String("address:", s.config.Server.GRPCAddress))

	if err := s.server.Serve(listen); err != nil {
		zap.L().Fatal("failed to server grpc server", zap.Error(err))

		return
	}
}

func (s *GRPCServer) Stop() {
	s.storage.Close()
	s.server.GracefulStop()
}

// func (s *GRPCServer)SaveLoginPassword(ctx context.Context, loginPasswordData *pb.LoginPasswordData) (*emptypb.Empty, error)
// func (s *GRPCServer)SaveCard(ctx context.Context, cardData *pb.CardData) (*emptypb.Empty, error)
// func (s *GRPCServer)SaveText(ctx context.Context, textData *pb.TextData) (*emptypb.Empty, error)
// func (s *GRPCServer)SaveBinary(ctx context.Context, binaryData *pb.BinaryData) (*emptypb.Empty, error)

// func (s *GRPCServer)GetListData(ctx context.Context, dataGetterRequest *pb.DataGetterRequest) (*pb.DataListResponse, error)
// func (s *GRPCServer)GetLoginPasswordObject(ctx context.Context, request *pb.DataRequest) (*pb.LoginPasswordData, error)
// func (s *GRPCServer)GetCardObject(ctx context.Context, request *pb.DataRequest) (*pb.CardData, error)
// func (s *GRPCServer)GetTextObject(ctx context.Context, request *pb.DataRequest) (*pb.TextData, error)
// func (s *GRPCServer)GetBinaryObject(ctx context.Context, request *pb.DataRequest) (*pb.BinaryData, error)

// func (s *GRPCServer)DeleteObject(ctx context.Context, dataGetterRequest *pb.DataGetterRequest) (*emptypb.Empty, error)
