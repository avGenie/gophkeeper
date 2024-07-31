package loginpassword

import (
	"github.com/avGenie/gophkeeper/client/internal/controller"
	"github.com/avGenie/gophkeeper/client/internal/entity"
)

// UpdateData Update data in storage
func (p *LoginPasswordProcessor) UpdateData(object entity.LoginPassword) error {
	if _, ok := p.data[object.Name]; ok {
		return controller.ErrDataAlreadyExists
	}

	err := p.client.UpdateLoginPasswordData(object)
	if err != nil {
		return err
	}

	p.data[object.Name] = object

	return nil
}
