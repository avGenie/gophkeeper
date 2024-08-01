package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/avGenie/gophkeeper/server/internal/entity"
	"github.com/avGenie/gophkeeper/server/internal/usecase/crypto"
	"google.golang.org/grpc/metadata"
)

const (
	tokenKey = "token"
)

var (
	ErrEmptyMetadata   = fmt.Errorf("empty data in metadata context")
	ErrEmptyTokenValue = fmt.Errorf("empty context value by key")
	ErrExpiredToken    = fmt.Errorf("token is expired")
	ErrInvalidToken    = fmt.Errorf("invalid token")
)

// GetUserIDFromContext Gets user id from context
func GetUserIDFromContext(ctx context.Context) (entity.UserID, error) {
	token, err := getTokenFromContext(ctx)
	if err != nil {
		return entity.UserID(""), err
	}

	userID, err := crypto.GetUserID(token)
	if err != nil {
		if errors.Is(err, crypto.ErrTokenExpired) {
			return entity.UserID(""), ErrExpiredToken
		}

		if errors.Is(err, crypto.ErrTokenNotValid) {
			return entity.UserID(""), ErrInvalidToken
		}

		return entity.UserID(""), fmt.Errorf("unknown error: %w", err)
	}

	return userID, nil
}

func getTokenFromContext(ctx context.Context) (entity.Token, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", ErrEmptyMetadata
	}

	values := md.Get(tokenKey)
	if len(values) == 0 {
		return "", ErrEmptyTokenValue
	}

	return entity.Token(values[0]), nil
}
