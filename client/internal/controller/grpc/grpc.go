package grpc

import (
	"fmt"

	"github.com/avGenie/gophkeeper/client/internal/config"
	"github.com/avGenie/gophkeeper/client/internal/ui"
	"github.com/avGenie/gophkeeper/client/internal/ui/terminal"
	pb "github.com/avGenie/gophkeeper/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	conn   *grpc.ClientConn
	client pb.GophkeeperClient
	ui     ui.UI
}

func NewClient(config config.Config) (*GRPCClient, error) {
	conn, err := grpc.NewClient(config.GRPCConnection, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("couldn't create client: %w", err)
	}

	client := &GRPCClient{
		conn:   conn,
		client: pb.NewGophkeeperClient(conn),
	} 

	client.ui = terminal.NewTerminal(client)

	return client, nil
}

func (c *GRPCClient) Start() {
	
	err := c.ui.Login()
	if err != nil {
		zap.S().Fatal("incorrect login credentials")
	}

	c.ui.Menu()
}

func (c *GRPCClient) Stop() {
	c.conn.Close()
}
