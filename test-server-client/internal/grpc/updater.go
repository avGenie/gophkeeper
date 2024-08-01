package grpc

import (
	"context"
	"time"

	"github.com/avGenie/gophkeeper/test-server-client/internal/converter"
)

func (c *Client) UpdateLoginPasswordUser(name, login, password, meta, token string) error {
	loginPasswordData := converter.CreatePbLoginPasswordData(name, login, password, meta)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ctx = saveTokenToContext(ctx, token)

	_, err := c.client.UpdateLoginPassword(ctx, loginPasswordData)

	return err
}

func (c *Client) UpdateText(name, text, meta, token string) error {
	data := converter.CreatePbTextData(name, text, meta)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ctx = saveTokenToContext(ctx, token)

	_, err := c.client.UpdateText(ctx, data)

	return err
}

func (c *Client) UpdateCard(name, number, cardholder string, expireMonth, expireYear, code int, meta, token string) error {
	data := converter.CreatePbCardData(name, number, cardholder, expireMonth, expireYear, code, meta)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ctx = saveTokenToContext(ctx, token)

	_, err := c.client.UpdateCard(ctx, data)

	return err
}