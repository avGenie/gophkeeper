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
			fmt.Fprintln(t.out, "Exitinig...")

			return
		}

		t.menuLoginPassword()

	case LoginPassword_GetAll:

	case LoginPassword_Exit:
		fmt.Fprintln(t.out, "Exitinig...")

		t.Menu()

	default:
		fmt.Fprintln(t.out, "Re-enter your choice!")
		t.menuMenu()
	}
}

func (t *Terminal) menuLoginPassword() {
	fmt.Fprintln(t.out, "1. Get by name")
	fmt.Fprintln(t.out, "2. Get all")
	fmt.Fprintln(t.out, "0. Back")
	fmt.Fprint(t.out, "Enter your choice: ")
}

func (t *Terminal) loginPasswordGetByName() error {
	fmt.Fprint(t.out, "Object name: ")
	name := entity.ObjectName(t.scanText())

	data, err := t.client.GetLoginPasswordUser(name, t.userToken)
	if err != nil {
		if errors.Is(err, controller.ErrUserPermissionDenied) {
			fmt.Fprintln(t.out, "Permission denied. Please, login again.")

			return errExit
		}

		if errors.Is(err, controller.ErrDataNotFound) {
			fmt.Fprintln(t.out, "Requested data not found.")

			return nil
		}

		zap.S().Error("failed to get login-password data", zap.Error(err))
		fmt.Fprintln(t.out, "!!! Unexpected error. Please, contact your administrator. !!!")

		return errExit
	}

	t.printLoginPasswordData(data)

	return nil
}

func (t *Terminal) printLoginPasswordData(data entity.LoginPassword) {
	fmt.Fprintf(t.out, "    Name: %s\n", data.Name)
	fmt.Fprintf(t.out, "   Login: %s\n", data.Login)
	fmt.Fprintf(t.out, "Password: %s\n", data.Password)
	fmt.Fprintf(t.out, "Metadata: %s\n", data.Metadata.Data)
}
