package grpc

import (
	"fmt"
	"time"

	"github.com/avGenie/gophkeeper/client/internal/config"
	"github.com/avGenie/gophkeeper/client/internal/entity"
	"github.com/avGenie/gophkeeper/client/internal/ui"
	"github.com/avGenie/gophkeeper/client/internal/ui/terminal"
	"github.com/avGenie/gophkeeper/package/tls"
	pb "github.com/avGenie/gophkeeper/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	requestTimeout = 5 * time.Second
)

type GRPCClient struct {
	conn   *grpc.ClientConn
	client pb.GophkeeperClient
	ui     ui.UI

	userToken entity.Token
}

func NewClient(config config.Config) (*GRPCClient, error) {
	tlsCreds, err := tls.GenerateClientTLSCreds(config.Client.CertsPath)
	if err != nil {
		return nil, fmt.Errorf("couldn't generate tls creds: %w", err)
	}

	conn, err := grpc.NewClient(config.Client.GRPCAddress, grpc.WithTransportCredentials(tlsCreds))
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
		zap.S().Fatal("failed to login", zap.Error(err))
	}

	c.ui.Menu()
}

func (c *GRPCClient) Stop() {
	c.conn.Close()
}
