package grpc

import (
	"log"

	pb "github.com/avGenie/gophkeeper/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn *grpc.ClientConn
	client pb.GophkeeperClient
}

func CreateClient() Client {
	conn, err := grpc.NewClient(":9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	
	client := pb.NewGophkeeperClient(conn)

	return Client{
		conn: conn,
		client: client,
	}
}

func (c *Client) Stop() {
	c.conn.Close()
}
