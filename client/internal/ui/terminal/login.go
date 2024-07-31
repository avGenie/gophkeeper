package terminal

import (
	"errors"
	"fmt"

	"github.com/avGenie/gophkeeper/client/internal/controller"
	"github.com/avGenie/gophkeeper/client/internal/entity"
	"github.com/avGenie/gophkeeper/client/internal/ui"
	"go.uber.org/zap"
	"golang.org/x/term"
)

type LoginChoice int

const (
	Login_Exit LoginChoice = iota
	Login_Register
	Login_Auth
)

func (t *Terminal) Login() error {
	var choice LoginChoice
	t.loginMenu()

	for {
		fmt.Scan(&choice)
		switch choice {
		case Login_Register:
			success, err := t.register()
			if err != nil {
				return fmt.Errorf("login register: %w", err)
			}

			if success {
				fmt.Fprintln(t.out, "\n\nUser has been successfuly registered!")
			}

		case Login_Auth:
			success, err := t.authenticate()
			if err != nil {
				return fmt.Errorf("login authenticate: %w", err)
			}

			if success {
				fmt.Fprintln(t.out, "User has been successfuly authenticate!")
				
				return nil
			}

		case Login_Exit:
			fmt.Fprintln(t.out, Exiting)
			return ui.ErrExit

		default:
			fmt.Fprintln(t.out, ChoiceReenterChoice)
		}

		t.loginMenu()
	}
}

func (t *Terminal) loginMenu() {
	fmt.Fprintln(t.out, "1. Registration")
	fmt.Fprintln(t.out, "2. Authentication")
	fmt.Fprintln(t.out, ChoiceExit)
	fmt.Fprint(t.out, ChoiceEnterChoice)
}

func (t *Terminal) enterLoginPassword() (entity.User, error) {
	var loginPass entity.User

	fmt.Fprint(t.out, "Username: ")
	loginPass.Login = t.scanText()

	fmt.Fprint(t.out, "Password: ")
	password, err := term.ReadPassword(int(t.in.Fd()))
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to read password: %w", err)
	}

	loginPass.Password = string(password)

	return loginPass, nil
}

func (t *Terminal) register() (bool, error) {
	fmt.Fprintln(t.out, "=============Registration=============")
	creds, err := t.enterLoginPassword()
	if err != nil {
		zap.S().Error("failed to register user", zap.Error(err))

		return false, fmt.Errorf("couldn't register user: %w", err)
	}

	err = t.client.RegisterUser(creds)
	if err != nil {
		if errors.Is(controller.ErrUserExists, err) {
			fmt.Fprintln(t.out, "User with this credentials exists. Please, try again.")

			return false, nil
		}

		if errors.Is(controller.ErrUserInvalid, err) {
			fmt.Fprintln(t.out, "Invalid input user credentials. Please, try again.")

			return false, nil
		}

		zap.S().Error("failed to register user", zap.Error(err), zap.String("user_login", creds.Login))
		fmt.Fprintln(t.out, UnexpectedError)

		return false, fmt.Errorf("couldn't register user: %w", err)
	}

	return true, nil
}

func (t *Terminal) authenticate() (bool, error) {
	fmt.Fprintln(t.out, "=============Authentication=============")
	creds, err := t.enterLoginPassword()
	if err != nil {
		zap.S().Error("failed to authenticate user", zap.Error(err))

		return false, fmt.Errorf("couldn't authenticate user: %w", err)
	}

	err = t.client.AuthenticateUser(creds)
	if err != nil {
		if errors.Is(controller.ErrUserNotFound, err) {
			fmt.Fprintln(t.out, "User with this credentials not found. Please, try again.")

			return false, nil
		}

		if errors.Is(controller.ErrUserInvalid, err) {
			fmt.Fprintln(t.out, "Invalid input user credentials. Please, try again.")

			return false, nil
		}

		zap.S().Error("failed to authenticate user", zap.Error(err), zap.String("user_login", creds.Login))
		fmt.Fprintln(t.out, UnexpectedError)

		return false, fmt.Errorf("couldn't authenticate user: %w", err)
	}

	return true, nil
}
