package card

import "github.com/avGenie/gophkeeper/client/internal/entity"

func (p *CardProcessor) GetData(name entity.ObjectName) (entity.CardData, error) {
	if data, ok := p.data[name]; ok {
		return data, nil
	}

	data, err := p.client.GetCardData(name)
	if err != nil {
		return entity.CardData{}, err
	}

	p.data[data.Name] = data

	return data, nil
}

func (p *CardProcessor) GetObjects() (entity.CardObjects, error) {
	data, err := p.client.GetCardObjects()
	if err != nil {
		return nil, err
	}

	for _, object := range data {
		p.data[object.Name] = object
	}

	return data, nil
}
