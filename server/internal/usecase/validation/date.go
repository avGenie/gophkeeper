package validation

import "strconv"

// ValidateCardExpireMonth Checks if month value positive and less than 13
func ValidateCardExpireMonth(expireMonth int) bool {
	return expireMonth > 0 && expireMonth < 13
}

// ValidateCardExpireYear Checks if year value positive
func ValidateCardExpireYear(expireYear int) bool {
	return expireYear > 0
}

// ValidateCardCode Checks if code value positive and digits count equal 3
func ValidateCardCode(code int) bool {
	digitsCount := len(strconv.Itoa(code))

	return code > 0 && digitsCount == 3
}
