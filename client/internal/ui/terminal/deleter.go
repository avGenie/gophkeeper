package terminal

import (
	"fmt"

	"github.com/avGenie/gophkeeper/client/internal/entity"
	"go.uber.org/zap"
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
	t.printNameMenu()

	name := entity.ObjectName(t.scanText())

	err := t.loginPasswordProc.DeleteData(name)
	if err != nil {
		return t.deleteProcessError(err)
	}

	return nil
}

func (t *Terminal) deleteText() error {
	t.printNameMenu()

	name := entity.ObjectName(t.scanText())

	zap.S().Debug("deleting text", zap.String("name", string(name)))

	err := t.textProc.DeleteData(name)
	if err != nil {
		return t.deleteProcessError(err)
	}

	zap.S().Debug("text successfule deleted")

	return nil
}

func (t *Terminal) deleteCard() error {
	t.printNameMenu()

	name := entity.ObjectName(t.scanText())

	err := t.cardProc.DeleteData(name)
	if err != nil {
		return t.deleteProcessError(err)
	}

	return nil
}
