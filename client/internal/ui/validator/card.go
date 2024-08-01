package validator

import (
	"fmt"
	"strconv"
	"time"
)

const (
	minExpirationMonth = 1
	maxExpirationMonth = 12
	minMaxYearDuration = 100
	codeDigitCount = 3
)

func ValidateExpirationMonth(month int) error {
	if month > maxExpirationMonth || month < minExpirationMonth {
		return ErrInvalidCardExpirationMonth
	}

	return nil
}

func ValidateExpirationYear(year int) error {
	minExpirationYear := time.Now().Year()
	maxExpirationYear := minExpirationYear + minMaxYearDuration

	if year > maxExpirationYear || year < minExpirationYear {
		return fmt.Errorf("Wrong data. Please, enter the number between %d and %d", minExpirationYear, maxExpirationYear)
	}

	return nil
}

func ValidateCode(code int) error {
	digitsCount := len(strconv.Itoa(code))

	if code < 0 || digitsCount != codeDigitCount {
		return ErrInvalidCardCode
	}

	return nil
}
