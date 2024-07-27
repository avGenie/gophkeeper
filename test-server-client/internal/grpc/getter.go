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
