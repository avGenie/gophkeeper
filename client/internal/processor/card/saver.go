package card

import (
	"github.com/avGenie/gophkeeper/client/internal/controller"
	"github.com/avGenie/gophkeeper/client/internal/entity"
)

// SaveData Save data to storage
func (p *CardProcessor) SaveData(object entity.CardData) error {
	if _, ok := p.data[object.Name]; ok {
		return controller.ErrDataAlreadyExists
	}

	err := p.client.SaveCardData(object)
	if err != nil {
		return err
	}

	p.data[object.Name] = object

	return nil
}
