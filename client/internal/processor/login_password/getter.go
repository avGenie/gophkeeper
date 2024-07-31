package loginpassword

import "github.com/avGenie/gophkeeper/client/internal/entity"

func (p *LoginPasswordProcessor) GetData(name entity.ObjectName) (entity.LoginPassword, error) {
	if data, ok := p.data[name]; ok {
		return data, nil
	}

	data, err := p.client.GetLoginPasswordData(name)
	if err != nil {
		return entity.LoginPassword{}, err
	}

	p.data[entity.ObjectName(data.Name)] = data

	return data, nil
}

func (p *LoginPasswordProcessor) GetObjects() (entity.LoginPasswordObjects, error) {
	data, err := p.client.GetLoginPasswordObjects()
	if err != nil {
		return nil, err
	}

	for _, object := range data {
		p.data[entity.ObjectName(object.Name)] = object
	}

	return data, nil
}
