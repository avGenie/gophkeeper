package loginpassword

import "github.com/avGenie/gophkeeper/client/internal/entity"

func (p *LoginPasswordProcessor) DeleteData(name entity.ObjectName) error {
	err := p.client.DeleteLoginPasswordData(name)
	if err != nil {
		return err
	}

	delete(p.data, name)

	return nil
}
