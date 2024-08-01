package text

import (
	"github.com/avGenie/gophkeeper/client/internal/controller"
	"github.com/avGenie/gophkeeper/client/internal/entity"
)

type TextProcessor struct {
	client controller.GophkeeperClient

	data map[entity.ObjectName]entity.TextData
}

func NewText(client controller.GophkeeperClient) *TextProcessor {
	return &TextProcessor{
		client: client,
		data:   make(map[entity.ObjectName]entity.TextData),
	}
}
