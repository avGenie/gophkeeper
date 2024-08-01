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

		case Menu_Add:
			err := t.save()
			if err != nil {
				return
			}

		case Menu_Delete:
			err := t.delete()
			if err != nil {
				return
			}

		case Menu_Update:
			err := t.update()
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

func (t *Terminal) loginMenu() {
	fmt.Fprintln(t.out, "1. Registration")
	fmt.Fprintln(t.out, "2. Authentication")
	fmt.Fprintln(t.out, ChoiceExit)
	fmt.Fprint(t.out, ChoiceEnterChoice)
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

func (t *Terminal) menuGetItems() {
	fmt.Fprintln(t.out, "1. Get by name")
	fmt.Fprintln(t.out, "2. Get all")
	fmt.Fprintln(t.out, ChoiceBack)
	fmt.Fprint(t.out, ChoiceEnterChoice)
	fmt.Fprintln(t.out, "")
}

func (t *Terminal) deleteMenu() {
	fmt.Fprintln(t.out, "1. Login-password data")
	fmt.Fprintln(t.out, "2. Text data")
	fmt.Fprintln(t.out, "3. Card data")
	fmt.Fprintln(t.out, ChoiceExit)
	fmt.Fprint(t.out, ChoiceEnterChoice)
	fmt.Fprintln(t.out, "")
}
