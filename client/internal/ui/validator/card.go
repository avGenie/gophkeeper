package validator

import "strconv"

const (
	minExpirationMonth = 1
	maxExpirationMonth = 12
	minExpirationYear  = 1920
	maxExpirationYear  = 2090
)

func ValidateExpirationMonth(month int) error {
	if month > minExpirationMonth && month < minExpirationMonth {
		return ErrInvalidCardExpirationMonth
	}

	return nil
}

func ValidateExpirationYear(year int) error {
	if year > minExpirationYear && year < minExpirationYear {
		return ErrInvalidCardExpirationYear
	}

	return nil
}

func ValidateCode(code int) error {
	digitsCount := len(strconv.Itoa(code))

	if code < 0 && digitsCount != 3 {
		return ErrInvalidCardCode
	}

	return nil
}
