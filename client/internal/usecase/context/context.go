package grpc

import (
	"context"

	"github.com/avGenie/gophkeeper/client/internal/entity"
	"google.golang.org/grpc/metadata"
)

const (
	tokenKey = "token"
)

func SaveTokenToContext(ctx context.Context, token entity.Token) context.Context {
	return metadata.AppendToOutgoingContext(ctx, tokenKey, string(token))
}
