package validation

import "github.com/avGenie/gophkeeper/server/internal/entity"

// ValidateCreateUserRequest Validates create user request. Checks if password and user length more than zero size
func ValidateCreateUserRequest(user entity.User) bool {
	return len(user.Login) > 0 && len(user.Password) > 0
}