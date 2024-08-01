package terminal

import (
	"fmt"

	"github.com/avGenie/gophkeeper/client/internal/entity"
)

func (t *Terminal) getLoginPassword() error {
	var choice LoginPasswordChoice

	for {
		t.menuGetItems()
		fmt.Scan(&choice)

		switch choice {
		case LoginPassword_GetByName:
			err := t.loginPasswordGetByName()
			if err != nil {
				fmt.Fprintln(t.out, Exiting)

				return err
			}

		case LoginPassword_GetAll:
			err := t.loginPasswordGetAll()
			if err != nil {
				fmt.Fprintln(t.out, Exiting)

				return err
			}

		case LoginPassword_GetExit:
			fmt.Fprintln(t.out, Exiting)

			return nil

		default:
			fmt.Fprintln(t.out, ChoiceReenterChoice)
		}
	}
}

func (t *Terminal) loginPasswordGetByName() error {
	fmt.Fprint(t.out, "Object name: ")
	name := entity.ObjectName(t.scanText())

	data, err := t.loginPasswordProc.GetData(name)
	if err != nil {
		return t.getProcessError(err, "failed to get login-password data object")
	}

	t.printLoginPasswordData(data)

	return nil
}

func (t *Terminal) loginPasswordGetAll() error {
	data, err := t.loginPasswordProc.GetObjects()
	if err != nil {
		return t.getProcessError(err, "failed to get all login-password data")
	}

	t.printLoginPasswordObjects(data)

	return nil
}

func (t *Terminal) getText() error {
	var choice TextChoice

	for {
		t.menuGetItems()
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

		case Text_GetExit:
			fmt.Fprintln(t.out, Exiting)

			return nil

		default:
			fmt.Fprintln(t.out, ChoiceReenterChoice)
		}
	}
}

func (t *Terminal) textGetByName() error {
	fmt.Fprint(t.out, "Object name: ")
	name := entity.ObjectName(t.scanText())

	data, err := t.textProc.GetData(name)
	if err != nil {
		return t.getProcessError(err, "failed to get text data object")
	}

	t.printTextData(data)

	return nil
}

func (t *Terminal) textGetAll() error {
	data, err := t.textProc.GetObjects()
	if err != nil {
		return t.getProcessError(err, "failed to get all text data")
	}

	t.printTextObjects(data)

	return nil
}

func (t *Terminal) getCard() error {
	var choice CardChoice

	for {
		t.menuGetItems()
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

		case Card_GetExit:
			fmt.Fprintln(t.out, Exiting)

			return nil

		default:
			fmt.Fprintln(t.out, ChoiceReenterChoice)
		}
	}
}

func (t *Terminal) cardGetByName() error {
	fmt.Fprint(t.out, "Object name: ")
	name := entity.ObjectName(t.scanText())

	data, err := t.cardProc.GetData(name)
	if err != nil {
		return t.getProcessError(err, "failed to get card data object")
	}

	t.printCardData(data)

	return nil
}

func (t *Terminal) cardGetAll() error {
	data, err := t.cardProc.GetObjects()
	if err != nil {
		return t.getProcessError(err, "failed to get all card data")
	}

	t.printCardObjects(data)

	return nil
}
