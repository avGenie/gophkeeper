package grpc

import (
	"context"
	"time"

	pb "github.com/avGenie/gophkeeper/proto"
)

func (c *Client) RegisterUser(login, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userCreds := &pb.UserCredentials{
		Login:    login,
		Password: password,
	}

	_, err := c.client.RegisterUser(ctx, userCreds)

	return err
}

func (c *Client) AuthenticateUser(login, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userCreds := &pb.UserCredentials{
		Login:    login,
		Password: password,
	}

	token, err := c.client.AuthenticateUser(ctx, userCreds)
	if err != nil {
		return "", err
	}

	return token.GetToken(), err
}
