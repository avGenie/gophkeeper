package crypto

import (
	"errors"
	"fmt"
	"time"

	"github.com/avGenie/gophkeeper/server/internal/entity"
	"github.com/golang-jwt/jwt/v4"
)

const (
	tokenTimeout = time.Hour * 3
	secretKey    = "5269889d400bbf2dc66216f37b2839bb"
)

// JWT errors
//
// ErrTokenNotValid - returned if token is not valid
// ErrTokenExpired - returned if token is expired
var (
	ErrTokenNotValid = errors.New("token is not valid")
	ErrTokenExpired  = errors.New("token is expired")
)

// Claims Contains token payload (user id)
type Claims struct {
	jwt.RegisteredClaims
	UserID entity.UserID
}

// BuildJWTString Creates JSON Web Token
func BuildJWTString(userID entity.UserID) (entity.Token, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTimeout)),
		},
		UserID: userID,
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return entity.Token(tokenString), nil
}

// GetUserID Obtaines user id from JSON Web Token
func GetUserID(tokenString string) (entity.UserID, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method while parsing jwt: %v", t.Header["alg"])
			}
			return []byte(secretKey), nil
		})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return entity.UserID(""), ErrTokenExpired
		}

		return entity.UserID(""), fmt.Errorf("error while getting user id from token: %w", err)
	}

	if !token.Valid {
		return entity.UserID(""), ErrTokenNotValid
	}

	return claims.UserID, nil
}
