package terminal

import (
	"os"

	"github.com/avGenie/gophkeeper/client/internal/controller"
	"github.com/avGenie/gophkeeper/client/internal/entity"
	"github.com/avGenie/gophkeeper/client/internal/processor/card"
	loginpassword "github.com/avGenie/gophkeeper/client/internal/processor/login_password"
	"github.com/avGenie/gophkeeper/client/internal/processor/text"
)

type Terminal struct {
	client controller.GophkeeperClient
	in     *os.File
	out    *os.File

	userToken entity.Token

	loginPasswordProc *loginpassword.LoginPasswordProcessor
	textProc          *text.TextProcessor
	cardProc          *card.CardProcessor
}

func NewTerminal(client controller.GophkeeperClient) *Terminal {
	return &Terminal{
		client:            client,
		in:                os.Stdin,
		out:               os.Stdout,
		loginPasswordProc: loginpassword.NewLoginPassword(client),
		textProc:          text.NewText(client),
		cardProc:          card.NewCard(client),
	}
}
