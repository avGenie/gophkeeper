package validator

import "fmt"

var (
	ErrInvalidCardExpirationMonth = fmt.Errorf("invalid expiration month")
	ErrInvalidCardCode            = fmt.Errorf("invalid code")
)
