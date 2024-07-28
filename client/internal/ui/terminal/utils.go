package terminal

import "bufio"

func (t *Terminal) scanText() string {
	scanner := bufio.NewScanner(t.in)
	scanner.Scan()

	return scanner.Text()
}
