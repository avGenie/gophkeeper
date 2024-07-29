package terminal

import (
	"bufio"
	"fmt"
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
