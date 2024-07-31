package terminal

import (
	"errors"
	"fmt"

	"github.com/avGenie/gophkeeper/client/internal/controller"
	"go.uber.org/zap"
)

var (
	errExit = fmt.Errorf("exit from app")
)

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

func (t *Terminal) saveProcessError(err error, errMsg string) error {
	if errors.Is(err, controller.ErrUserPermissionDenied) {
		fmt.Fprintln(t.out, PermissionDenied)

		return errExit
	}

	if errors.Is(err, controller.ErrDataDeleteWrongName) {
		fmt.Fprintln(t.out, DataSaveExisted)

		return nil
	}

	if errors.Is(err, controller.ErrDataAlreadyExists) {
		fmt.Fprintln(t.out, DataSaveExisted)

		return nil
	}

	zap.S().Error(errMsg, zap.Error(err))
	fmt.Fprintln(t.out, UnexpectedError)

	return errExit
}
