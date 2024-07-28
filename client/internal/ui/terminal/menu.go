package terminal

import "fmt"

func (t *Terminal) Menu() {
	t.menuMenu()
}

func (t *Terminal) menuMenu() {
	fmt.Fprintln(t.out, "1. Get login-password data")
	fmt.Fprintln(t.out, "2. Get text data")
	fmt.Fprintln(t.out, "3. Get card data")
	fmt.Fprintln(t.out, "4. Add")
	fmt.Fprintln(t.out, "5. Delete")
	fmt.Fprintln(t.out, "6. Update")
	fmt.Fprintln(t.out, "0. Exit")
	fmt.Fprint(t.out, "Enter your choice: ")
}
