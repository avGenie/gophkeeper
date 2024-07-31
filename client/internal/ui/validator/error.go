package validator

import "fmt"

var (
	ErrInvalidCardExpirationMonth = fmt.Errorf("invalid expiration month")
	ErrInvalidCardExpirationYear  = fmt.Errorf("invalid expiration year")
	ErrInvalidCardCode            = fmt.Errorf("invalid code")
)
