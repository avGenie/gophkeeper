package grpc

import (
	"context"
	"time"

	"github.com/avGenie/gophkeeper/test-server-client/internal/converter"
	"github.com/avGenie/gophkeeper/test-server-client/internal/entity"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (c *Client) GetLoginPasswordObjects(token string) (entity.LoginPasswordObjects, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ctx = saveTokenToContext(ctx, token)

	obj, err := c.client.GetLoginPasswordObjects(ctx, &emptypb.Empty{})
	if err != nil {
		return entity.LoginPasswordObjects{}, err
	}

	return converter.ConvertPbLoginPasswordObjectsToLoginPasswordObjects(obj), nil
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

func (c *Client) GetTextObjects(token string) (entity.TextObjects, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ctx = saveTokenToContext(ctx, token)

	obj, err := c.client.GetTextObjects(ctx, &emptypb.Empty{})
	if err != nil {
		return entity.TextObjects{}, err
	}

	return converter.ConvertPbTextObjectsToTextObjects(obj), nil
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

func (c *Client) GetCardObjects(token string) (entity.CardObjects, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ctx = saveTokenToContext(ctx, token)

	obj, err := c.client.GetCardObjects(ctx, &emptypb.Empty{})
	if err != nil {
		return entity.CardObjects{}, err
	}

	return converter.ConvertPbCardObjectsToCardObjects(obj), nil
}
