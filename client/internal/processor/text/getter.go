package text

import "github.com/avGenie/gophkeeper/client/internal/entity"

func (p *TextProcessor) GetData(name entity.ObjectName) (entity.TextData, error) {
	if data, ok := p.data[name]; ok {
		return data, nil
	}

	data, err := p.client.GetTextData(name)
	if err != nil {
		return entity.TextData{}, err
	}

	p.data[entity.ObjectName(data.Name)] = data

	return data, nil
}

func (p *TextProcessor) GetObjects() (entity.TextObjects, error) {
	data, err := p.client.GetTextObjects()
	if err != nil {
		return nil, err
	}

	for _, object := range data {
		p.data[entity.ObjectName(object.Name)] = object
	}

	return data, nil
}