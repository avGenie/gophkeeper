package storage

import "fmt"

var (
	ErrUserExists = fmt.Errorf("user exists")
	ErrUserNotFound = fmt.Errorf("user not found")
)