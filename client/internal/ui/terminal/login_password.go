package terminal

import (
	"errors"
	"fmt"

	"github.com/avGenie/gophkeeper/client/internal/controller"
	"github.com/avGenie/gophkeeper/client/internal/entity"
	"go.uber.org/zap"
)

type LoginPasswordChoice int

const (
	LoginPassword_Exit LoginPasswordChoice = iota
	LoginPassword_GetByName
	LoginPassword_GetAll
)

func (t *Terminal) getLoginPassword() {
	var choice LoginPasswordChoice
	t.menuLoginPassword()

	fmt.Scan(&choice)

	switch choice {
	case LoginPassword_GetByName:
		err := t.loginPasswordGetByName()
		if err != nil {
			fmt.Fprintln(t.out, Exiting)

			return
		}

		t.getLoginPassword()

	case LoginPassword_GetAll:
		err := t.loginPasswordGetAll()
		if err != nil {
			fmt.Fprintln(t.out, Exiting)

			return
		}

		t.getLoginPassword()

	case LoginPassword_Exit:
		fmt.Fprintln(t.out, Exiting)

		t.Menu()

	default:
		fmt.Fprintln(t.out, ChoiceReenterChoice)
		t.getLoginPassword()
	}
}

func (t *Terminal) menuLoginPassword() {
	fmt.Fprintln(t.out, "1. Get by name")
	fmt.Fprintln(t.out, "2. Get all")
	fmt.Fprintln(t.out, ChoiceBack)
	fmt.Fprint(t.out, ChoiceEnterChoice)
}

func (t *Terminal) loginPasswordGetByName() error {
	fmt.Fprint(t.out, "Object name: ")
	name := entity.ObjectName(t.scanText())

	data, err := t.client.GetLoginPasswordUser(name, t.userToken)
	if err != nil {
		if errors.Is(err, controller.ErrUserPermissionDenied) {
			fmt.Fprintln(t.out, PermissionDenied)

			return errExit
		}

		if errors.Is(err, controller.ErrDataNotFound) {
			fmt.Fprintln(t.out, DataNotFound)

			return nil
		}

		zap.S().Error("failed to get login-password data object", zap.Error(err))
		fmt.Fprintln(t.out, UnexpectedError)

		return errExit
	}

	t.printLoginPasswordData(data)

	return nil
}

func (t *Terminal) loginPasswordGetAll() error {
	data, err := t.client.GetLoginPasswordObjects(t.userToken)
	if err != nil {
		if errors.Is(err, controller.ErrUserPermissionDenied) {
			fmt.Fprintln(t.out, PermissionDenied)

			return errExit
		}

		if errors.Is(err, controller.ErrDataNotFound) {
			fmt.Fprintln(t.out, DataNotFound)

			return nil
		}

		zap.S().Error("failed to get all login-password data", zap.Error(err))
		fmt.Fprintln(t.out, UnexpectedError)

		return errExit
	}

	t.printLoginPasswordObjects(data)

	return nil
}

func (t *Terminal) printLoginPasswordData(data entity.LoginPassword) {
	fmt.Fprintf(t.out, "    Name: %s\n", data.Name)
	fmt.Fprintf(t.out, "   Login: %s\n", data.Login)
	fmt.Fprintf(t.out, "Password: %s\n", data.Password)
	fmt.Fprintf(t.out, "Metadata: %s\n", data.Metadata.Data)
}

func (t *Terminal) printLoginPasswordObjects(data entity.LoginPasswordObjects) {
	for _, object := range data {
		t.printLoginPasswordData(object)
		fmt.Fprintf(t.out, "\n")
	}
}
