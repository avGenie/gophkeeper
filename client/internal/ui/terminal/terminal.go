package terminal

import (
	"os"

	"github.com/avGenie/gophkeeper/client/internal/controller"
	"github.com/avGenie/gophkeeper/client/internal/entity"
)

type Terminal struct {
	client controller.GophkeeperClient
	in     *os.File
	out    *os.File

	userToken entity.Token
}

func NewTerminal(client controller.GophkeeperClient) *Terminal {
	return &Terminal{
		client: client,
		in:     os.Stdin,
		out:    os.Stdout,
	}
}
