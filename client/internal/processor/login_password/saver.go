package loginpassword

import (
	"github.com/avGenie/gophkeeper/client/internal/controller"
	"github.com/avGenie/gophkeeper/client/internal/entity"
)

// SaveData Save data to storage
func (p *LoginPasswordProcessor) SaveData(object entity.LoginPassword) error {
	if _, ok := p.data[object.Name]; ok {
		return controller.ErrDataAlreadyExists
	}

	err := p.client.SaveLoginPasswordData(object)
	if err != nil {
		return err
	}

	p.data[object.Name] = object

	return nil
}
