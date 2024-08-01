package terminal

import (
	"fmt"
)

func (t *Terminal) save() error {
	var choice SaveChoice

	for {
		t.deleteMenu()
		fmt.Scan(&choice)

		switch choice {
		case Save_LoginPassword:
			err := t.saveLoginPassword()
			if err != nil {
				fmt.Fprintln(t.out, Exiting)

				return err
			}

		case Save_Text:
			err := t.saveText()
			if err != nil {
				fmt.Fprintln(t.out, Exiting)

				return err
			}

		case Save_Card:
			err := t.saveCard()
			if err != nil {
				fmt.Fprintln(t.out, Exiting)

				return err
			}

		case Save_Exit:
			fmt.Fprintln(t.out, Exiting)

			return nil

		default:
			fmt.Fprintln(t.out, ChoiceReenterChoice)
		}
	}
}

func (t *Terminal) saveLoginPassword() error {
	fmt.Fprintln(t.out, "Input new object data")
	savedObject := t.inputLoginPasswordData()

	err := t.loginPasswordProc.SaveData(savedObject)
	if err != nil {
		return t.saveProcessError(err, "failed to save login-password data object")
	}

	return nil
}

func (t *Terminal) saveText() error {
	fmt.Fprintln(t.out, "Input new object data")
	savedObject := t.inputTextData()

	err := t.textProc.SaveData(savedObject)
	if err != nil {
		return t.saveProcessError(err, "failed to save text data object")
	}

	return nil
}

func (t *Terminal) saveCard() error {
	fmt.Fprintln(t.out, "Input new object data")
	savedObject := t.inputCardData()

	err := t.cardProc.SaveData(savedObject)
	if err != nil {
		return t.saveProcessError(err, "failed to save card data object")
	}

	return nil
}
