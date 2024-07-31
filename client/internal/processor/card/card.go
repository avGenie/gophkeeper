package card

import (
	"github.com/avGenie/gophkeeper/client/internal/controller"
	"github.com/avGenie/gophkeeper/client/internal/entity"
)

type CardProcessor struct {
	client controller.GophkeeperClient

	data map[entity.ObjectName]entity.CardData
}

func NewCard(client controller.GophkeeperClient) *CardProcessor {
	return &CardProcessor{
		client: client,
		data:   make(map[entity.ObjectName]entity.CardData),
	}
}
