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

func (c *Client) SaveText(name, text, meta, token string) error {
	data := converter.CreatePbTextData(name, text, meta)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ctx = saveTokenToContext(ctx, token)

	_, err := c.client.SaveText(ctx, data)

	return err
}

func (c *Client) SaveCard(name, number, cardholder string, expireMonth, expireYear, code int, meta, token string) error {
	data := converter.CreatePbCardData(name, number, cardholder, expireMonth, expireYear, code, meta)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ctx = saveTokenToContext(ctx, token)

	_, err := c.client.SaveCard(ctx, data)

	return err
}
