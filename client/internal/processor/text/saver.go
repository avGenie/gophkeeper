package text

import (
	"github.com/avGenie/gophkeeper/client/internal/controller"
	"github.com/avGenie/gophkeeper/client/internal/entity"
)

// SaveData Save data to storage
func (p *TextProcessor) SaveData(object entity.TextData) error {
	if _, ok := p.data[object.Name]; ok {
		return controller.ErrDataAlreadyExists
	}

	err := p.client.SaveTextData(object)
	if err != nil {
		return err
	}

	p.data[object.Name] = object

	return nil
}
