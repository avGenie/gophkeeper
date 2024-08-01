package grpc

import (
	"context"

	"google.golang.org/grpc/metadata"
)

const (
	tokenKey = "token"
)

func saveTokenToContext(ctx context.Context, token string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, tokenKey, token)
}
