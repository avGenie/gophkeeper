package storage

import "fmt"

// Storage output errors
//
// ErrUserExists - returned if user exists in storage
// ErrUserNotFound - returned if user not found in storage
// ErrDataExists - returned if data with provided content exists
// ErrLoginPasswordDataNotFound - returned if login password data not found
var (
	ErrUserExists                = fmt.Errorf("user exists")
	ErrUserNotFound              = fmt.Errorf("user not found")
	ErrDataExists                = fmt.Errorf("data with provided content exists")
	ErrLoginPasswordDataNotFound = fmt.Errorf("login password data not found")
)
