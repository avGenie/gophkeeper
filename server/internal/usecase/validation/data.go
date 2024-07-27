package validation

import "github.com/avGenie/gophkeeper/server/internal/entity"

func ValidateLoginPasswordData(data entity.LoginPassword) bool {
	return data.Name != "" && data.Login != "" && data.Password != ""
}
