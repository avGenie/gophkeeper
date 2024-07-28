package controller

import "fmt"

var (
	ErrUserExists   = fmt.Errorf("user with this credentials exists")
	ErrUserNotFound = fmt.Errorf("user with this credentials not found")
	ErrUserInvalid  = fmt.Errorf("invalid input user credentials")
)
