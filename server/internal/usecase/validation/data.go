package validation

import "github.com/avGenie/gophkeeper/server/internal/entity"

// ValidateLoginPasswordData Checks if name, login and password is not empty
func ValidateLoginPasswordData(data entity.LoginPassword) bool {
	return data.Name != "" && data.Login != "" && data.Password != ""
}

// ValidateTextData Checks if name, text is not empty
func ValidateTextData(data entity.TextData) bool {
	return data.Name != "" && data.Text != ""
}
