package grpc

import (
	"context"
	"time"

	"github.com/avGenie/gophkeeper/test-server-client/internal/converter"
)

func (c *Client) SaveLoginPasswordUser(name, login, password, meta, token string) error {
	loginPasswordData := converter.CreatePbLoginPasswordData(name, login, password, meta)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ctx = saveTokenToContext(ctx, token)

	_, err := c.client.SaveLoginPassword(ctx, loginPasswordData)

	return err
}
