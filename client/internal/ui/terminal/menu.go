package terminal

import (
	"fmt"
)

type MenuChoice int

const (
	Menu_Exit MenuChoice = iota
	Menu_GetLoginPassword
	Menu_GetText
	Menu_GetCard
	Menu_Add
	Menu_Delete
	Menu_Update
)

func (t *Terminal) Menu() {
	var choice MenuChoice
	t.menuMenu()

	for {
		fmt.Scan(&choice)

		switch choice {
		case Menu_GetLoginPassword:
			err := t.getLoginPassword()
			if err != nil {
				return
			}

		case Menu_GetText:
			err := t.getText()
			if err != nil {
				return
			}

		case Menu_GetCard:
			err := t.getCard()
			if err != nil {
				return
			}

		case Menu_Delete:
			err := t.delete()
			if err != nil {
				return
			}

		case Menu_Exit:
			fmt.Fprintln(t.out, Exiting)

			return

		default:
			fmt.Fprintln(t.out, ChoiceReenterChoice)
		}

		t.menuMenu()
	}
}

func (t *Terminal) menuMenu() {
	fmt.Fprintln(t.out, "1. Get login-password data")
	fmt.Fprintln(t.out, "2. Get text data")
	fmt.Fprintln(t.out, "3. Get card data")
	fmt.Fprintln(t.out, "4. Add")
	fmt.Fprintln(t.out, "5. Delete")
	fmt.Fprintln(t.out, "6. Update")
	fmt.Fprintln(t.out, ChoiceExit)
	fmt.Fprint(t.out, ChoiceEnterChoice)
	fmt.Fprintln(t.out, "")
}
