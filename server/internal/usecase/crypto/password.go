package crypto

import (
	"errors"
	"fmt"

	"github.com/avGenie/gophkeeper/server/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost = 14
)

var (
	ErrWrongPassword = errors.New("wrong password")
)

func HashPassword(user entity.User) (entity.User, error) {
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return entity.User{}, err
	}

	user.Password = hashedPassword

	return user, nil
}

func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return ErrWrongPassword
		}

		return fmt.Errorf("error while checking password hash: %w", err)
	}

	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return "", fmt.Errorf("failed to generate hash: %w", err)
	}

	return string(bytes), nil
}
