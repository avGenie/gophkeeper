package grpc

import (
	"context"
	"time"

	"github.com/avGenie/gophkeeper/test-server-client/internal/converter"
	"github.com/avGenie/gophkeeper/test-server-client/internal/entity"
)

func (c *Client) DeleteDataObject(requestType entity.DataRequestType, name, token string) error {
	request := converter.CreatePbDataGetterRequest(name, requestType)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ctx = saveTokenToContext(ctx, token)

	_, err := c.client.DeleteObject(ctx, request)

	return err
}
