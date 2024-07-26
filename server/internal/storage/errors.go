package storage

import "fmt"

// Storage output errors
//
// ErrUserExists - returned if user exists in storage
// ErrUserNotFound - returned if user not found in storage
var (
	ErrUserExists = fmt.Errorf("user exists")
	ErrUserNotFound = fmt.Errorf("user not found")
)