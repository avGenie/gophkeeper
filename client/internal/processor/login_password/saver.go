package loginpassword

import "github.com/avGenie/gophkeeper/client/internal/entity"

// SaveData Save data to storage
func (p *LoginPasswordProcessor) SaveData(object entity.LoginPassword) error {
	err := p.client.SaveLoginPasswordData(object)
	if err != nil {
		return err
	}

	p.data[object.Name] = object

	return nil
}
