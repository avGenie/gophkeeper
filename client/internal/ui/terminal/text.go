package terminal

import (
	"errors"
	"fmt"

	"github.com/avGenie/gophkeeper/client/internal/controller"
	"github.com/avGenie/gophkeeper/client/internal/entity"
	"go.uber.org/zap"
)

type TextChoice int

const (
	Text_Exit TextChoice = iota
	Text_GetByName
	Text_GetAll
)

func (t *Terminal) getText() error {
	var choice TextChoice

	for {
		t.menuText()
		fmt.Scan(&choice)

		switch choice {
		case Text_GetByName:
			err := t.textGetByName()
			if err != nil {
				fmt.Fprintln(t.out, Exiting)

				return err
			}

		case Text_GetAll:
			err := t.textGetAll()
			if err != nil {
				fmt.Fprintln(t.out, Exiting)

				return err
			}

		case Text_Exit:
			fmt.Fprintln(t.out, Exiting)

			return nil

		default:
			fmt.Fprintln(t.out, ChoiceReenterChoice)
		}
	}
}

func (t *Terminal) menuText() {
	t.menuGetItems()
}

func (t *Terminal) textGetByName() error {
	fmt.Fprint(t.out, "Object name: ")
	name := entity.ObjectName(t.scanText())

	data, err := t.client.GetTextData(name, t.userToken)
	if err != nil {
		if errors.Is(err, controller.ErrUserPermissionDenied) {
			fmt.Fprintln(t.out, PermissionDenied)

			return errExit
		}

		if errors.Is(err, controller.ErrDataNotFound) {
			fmt.Fprintln(t.out, DataNotFound)

			return nil
		}

		zap.S().Error("failed to get text data object", zap.Error(err))
		fmt.Fprintln(t.out, UnexpectedError)

		return errExit
	}

	t.printTextData(data)

	return nil
}

func (t *Terminal) textGetAll() error {
	data, err := t.client.GetTextObjects(t.userToken)
	if err != nil {
		if errors.Is(err, controller.ErrUserPermissionDenied) {
			fmt.Fprintln(t.out, PermissionDenied)

			return errExit
		}

		if errors.Is(err, controller.ErrDataNotFound) {
			fmt.Fprintln(t.out, DataNotFound)

			return nil
		}

		zap.S().Error("failed to get all text data", zap.Error(err))
		fmt.Fprintln(t.out, UnexpectedError)

		return errExit
	}

	t.printTextObjects(data)

	return nil
}

func (t *Terminal) printTextData(data entity.TextData) {
	fmt.Fprintf(t.out, "\t    Name: %s\n", data.Name)
	fmt.Fprintf(t.out, "\t    Text: %s\n", data.Text)
	fmt.Fprintf(t.out, "\tMetadata: %s\n", data.Metadata.Data)
}

func (t *Terminal) printTextObjects(data entity.TextObjects) {
	for _, object := range data {
		t.printTextData(object)
		fmt.Fprintf(t.out, "\n")
	}
}
