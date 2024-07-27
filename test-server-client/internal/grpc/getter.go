package grpc

import (
	"context"
	"time"

	"github.com/avGenie/gophkeeper/test-server-client/internal/converter"
	"github.com/avGenie/gophkeeper/test-server-client/internal/entity"
)

func (c *Client) GetLoginPasswordUser(name, token string) (entity.LoginPassword, error) {
	dataRequest := converter.CreatePbDataRequest(name)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ctx = saveTokenToContext(ctx, token)

	obj, err := c.client.GetLoginPasswordObject(ctx, dataRequest)
	if err != nil {
		return entity.LoginPassword{}, err
	}

	return converter.ConvertPbLoginPasswordDataToLoginPassword(obj), nil
}

func (c *Client) GetTextData(name, token string) (entity.TextData, error) {
	dataRequest := converter.CreatePbDataRequest(name)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ctx = saveTokenToContext(ctx, token)

	obj, err := c.client.GetTextObject(ctx, dataRequest)
	if err != nil {
		return entity.TextData{}, err
	}

	return converter.ConvertPbTextDataToText(obj), nil
}

func (c *Client) GetCardData(name, token string) (entity.CardData, error) {
	dataRequest := converter.CreatePbDataRequest(name)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ctx = saveTokenToContext(ctx, token)

	obj, err := c.client.GetCardObject(ctx, dataRequest)
	if err != nil {
		return entity.CardData{}, err
	}

	return converter.ConvertPbCardDataToCard(obj), nil
}
