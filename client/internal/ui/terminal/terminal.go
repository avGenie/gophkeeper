package terminal

import (
	"github.com/avGenie/gophkeeper/client/internal/controller"
)

type Terminal struct {
	client controller.GophkeeperClient
}

func NewTerminal(client controller.GophkeeperClient) *Terminal {
	return &Terminal{
		client: client,
	}
}

func (t *Terminal) Login() error {
	return nil
}

func (t *Terminal) Menu() {

}
