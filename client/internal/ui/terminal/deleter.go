package terminal

import (
	"errors"
	"fmt"

	"github.com/avGenie/gophkeeper/client/internal/controller"
	"github.com/avGenie/gophkeeper/client/internal/entity"
	"go.uber.org/zap"
)

type DeleteChoice int

const (
	Delete_Exit DeleteChoice = iota
	Delete_LoginPassword
	Delete_Text
	Delete_Card
)

func (t *Terminal) delete() error {
	var choice DeleteChoice

	for {
		t.deleteMenu()
		fmt.Scan(&choice)

		switch choice {
		case Delete_LoginPassword:
			err := t.deleteLoginPassword()
			if err != nil {
				fmt.Fprintln(t.out, Exiting)

				return err
			}

		case Delete_Text:
			err := t.deleteText()
			if err != nil {
				fmt.Fprintln(t.out, Exiting)

				return err
			}

		case Delete_Card:
			err := t.deleteCard()
			if err != nil {
				fmt.Fprintln(t.out, Exiting)

				return err
			}

		case Delete_Exit:
			fmt.Fprintln(t.out, Exiting)

			return nil

		default:
			fmt.Fprintln(t.out, ChoiceReenterChoice)
		}
	}
}

func (t *Terminal) deleteLoginPassword() error {
	t.writeNameMenu()

	name := entity.ObjectName(t.scanText())

	err := t.client.DeleteLoginPasswordData(name, t.userToken)
	if err != nil {
		return t.deleteProcessError(err)
	}

	return nil
}

func (t *Terminal) deleteText() error {
	t.writeNameMenu()

	name := entity.ObjectName(t.scanText())

	zap.S().Debug("deleting text", zap.String("name", string(name)))

	err := t.client.DeleteTextData(name, t.userToken)
	if err != nil {
		return t.deleteProcessError(err)
	}

	zap.S().Debug("text successfule deleted")

	return nil
}

func (t *Terminal) deleteCard() error {
	t.writeNameMenu()

	name := entity.ObjectName(t.scanText())

	err := t.client.DeleteCardData(name, t.userToken)
	if err != nil {
		return t.deleteProcessError(err)
	}

	return nil
}

func (t *Terminal) deleteProcessError(err error) error {
	if errors.Is(err, controller.ErrUserPermissionDenied) {
		fmt.Fprintln(t.out, PermissionDenied)

		return errExit
	}

	if errors.Is(err, controller.ErrDataDeleteWrongName) {
		fmt.Fprintln(t.out, DataDeleteWrongName)

		return nil
	}

	zap.S().Error("failed to delete login-password data object", zap.Error(err))
	fmt.Fprintln(t.out, UnexpectedError)

	return errExit
}

func (t *Terminal) deleteMenu() {
	fmt.Fprintln(t.out, "1. Login-password data")
	fmt.Fprintln(t.out, "2. Text data")
	fmt.Fprintln(t.out, "3. Card data")
	fmt.Fprintln(t.out, ChoiceExit)
	fmt.Fprint(t.out, ChoiceEnterChoice)
	fmt.Fprintln(t.out, "")
}

func (t *Terminal) writeNameMenu() {
	fmt.Fprint(t.out, "Data name: ")
	fmt.Fprintln(t.out, "")
}
