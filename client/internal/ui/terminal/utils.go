package terminal

import (
	"bufio"
	"errors"
	"fmt"

	"github.com/avGenie/gophkeeper/client/internal/controller"
	"go.uber.org/zap"
)

func (t *Terminal) scanText() string {
	scanner := bufio.NewScanner(t.in)
	scanner.Scan()

	return scanner.Text()
}

func (t *Terminal) menuGetItems() {
	fmt.Fprintln(t.out, "1. Get by name")
	fmt.Fprintln(t.out, "2. Get all")
	fmt.Fprintln(t.out, ChoiceBack)
	fmt.Fprint(t.out, ChoiceEnterChoice)
	fmt.Fprintln(t.out, "")
}

func (t *Terminal) getProcessError(err error, errMsg string) error {
	if errors.Is(err, controller.ErrUserPermissionDenied) {
		fmt.Fprintln(t.out, PermissionDenied)

		return errExit
	}

	if errors.Is(err, controller.ErrDataNotFound) {
		fmt.Fprintln(t.out, DataNotFound)

		return nil
	}

	zap.S().Error(errMsg, zap.Error(err))
	fmt.Fprintln(t.out, UnexpectedError)

	return errExit
}
