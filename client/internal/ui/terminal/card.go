package terminal

import (
	"errors"
	"fmt"

	"github.com/avGenie/gophkeeper/client/internal/controller"
	"github.com/avGenie/gophkeeper/client/internal/entity"
	"go.uber.org/zap"
)

type CardChoice int

const (
	Card_Exit CardChoice = iota
	Card_GetByName
	Card_GetAll
)

func (t *Terminal) getCard() error {
	var choice CardChoice
	t.menuCard()

	for {
		fmt.Scan(&choice)

		switch choice {
		case Card_GetByName:
			err := t.cardGetByName()
			if err != nil {
				fmt.Fprintln(t.out, Exiting)

				return err
			}

		case Card_GetAll:
			err := t.cardGetAll()
			if err != nil {
				fmt.Fprintln(t.out, Exiting)

				return err
			}

		case Card_Exit:
			fmt.Fprintln(t.out, Exiting)

			return nil

		default:
			fmt.Fprintln(t.out, ChoiceReenterChoice)
		}

		t.menuCard()
	}
}

func (t *Terminal) menuCard() {
	t.menuGetItems()
}

func (t *Terminal) cardGetByName() error {
	fmt.Fprint(t.out, "Object name: ")
	name := entity.ObjectName(t.scanText())

	data, err := t.client.GetCardData(name, t.userToken)
	if err != nil {
		if errors.Is(err, controller.ErrUserPermissionDenied) {
			fmt.Fprintln(t.out, PermissionDenied)

			return errExit
		}

		if errors.Is(err, controller.ErrDataNotFound) {
			fmt.Fprintln(t.out, DataNotFound)

			return nil
		}

		zap.S().Error("failed to get card data object", zap.Error(err))
		fmt.Fprintln(t.out, UnexpectedError)

		return errExit
	}

	t.printCardData(data)

	return nil
}

func (t *Terminal) cardGetAll() error {
	data, err := t.client.GetCardObjects(t.userToken)
	if err != nil {
		if errors.Is(err, controller.ErrUserPermissionDenied) {
			fmt.Fprintln(t.out, PermissionDenied)

			return errExit
		}

		if errors.Is(err, controller.ErrDataNotFound) {
			fmt.Fprintln(t.out, DataNotFound)

			return nil
		}

		zap.S().Error("failed to get all card data", zap.Error(err))
		fmt.Fprintln(t.out, UnexpectedError)

		return errExit
	}

	t.printCardObjects(data)

	return nil
}

func (t *Terminal) printCardData(data entity.CardData) {
	fmt.Fprintf(t.out, "\t            Name: %s\n", data.Name)
	fmt.Fprintf(t.out, "\t          Number: %s\n", data.Number)
	fmt.Fprintf(t.out, "\t            Code: %d\n", data.Code)
	fmt.Fprintf(t.out, "\t      Cardholder: %s\n", data.Cardholder)
	fmt.Fprintf(t.out, "\tExpiration month: %d\n", data.ExpirationMonth)
	fmt.Fprintf(t.out, "\tExpiration month: %d\n", data.ExpirationYear)
	fmt.Fprintf(t.out, "\t        Metadata: %s\n", data.Metadata.Data)
}

func (t *Terminal) printCardObjects(data entity.CardObjects) {
	for _, object := range data {
		t.printCardData(object)
		fmt.Fprintf(t.out, "\n")
	}
}
