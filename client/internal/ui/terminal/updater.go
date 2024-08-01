package terminal

import "fmt"

func (t *Terminal) update() error {
	var choice UpdateChoice

	for {
		t.deleteMenu()
		fmt.Scan(&choice)

		switch choice {
		case Update_LoginPassword:
			err := t.updateLoginPassword()
			if err != nil {
				fmt.Fprintln(t.out, Exiting)

				return err
			}

		case Update_Text:
			err := t.updateText()
			if err != nil {
				fmt.Fprintln(t.out, Exiting)

				return err
			}

		case Update_Card:
			err := t.updateCard()
			if err != nil {
				fmt.Fprintln(t.out, Exiting)

				return err
			}

		case Update_Exit:
			fmt.Fprintln(t.out, Exiting)

			return nil

		default:
			fmt.Fprintln(t.out, ChoiceReenterChoice)
		}
	}
}

func (t *Terminal) updateLoginPassword() error {
	fmt.Fprintln(t.out, "Input new object data")
	savedObject := t.inputLoginPasswordData()

	err := t.loginPasswordProc.UpdateData(savedObject)
	if err != nil {
		return t.updateProcessError(err, "failed to save login-password data object")
	}

	return nil
}

func (t *Terminal) updateText() error {
	fmt.Fprintln(t.out, "Input new object data")
	savedObject := t.inputTextData()

	err := t.textProc.UpdateData(savedObject)
	if err != nil {
		return t.updateProcessError(err, "failed to update text data object")
	}

	return nil
}

func (t *Terminal) updateCard() error {
	fmt.Fprintln(t.out, "Input new object data")
	savedObject := t.inputCardData()

	err := t.cardProc.UpdateData(savedObject)
	if err != nil {
		return t.updateProcessError(err, "failed to update card data object")
	}

	return nil
}
