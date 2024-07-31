package terminal

import (
	"fmt"

	"github.com/avGenie/gophkeeper/client/internal/entity"
)

func (t *Terminal) printLoginPasswordData(data entity.LoginPassword) {
	fmt.Fprintf(t.out, "\t    Name: %s\n", data.Name)
	fmt.Fprintf(t.out, "\t   Login: %s\n", data.Login)
	fmt.Fprintf(t.out, "\tPassword: %s\n", data.Password)
	fmt.Fprintf(t.out, "\tMetadata: %s\n", data.Metadata.Data)
}

func (t *Terminal) printLoginPasswordObjects(data entity.LoginPasswordObjects) {
	for _, object := range data {
		t.printLoginPasswordData(object)
		fmt.Fprintf(t.out, "\n")
	}
}

func (t *Terminal) printTextData(data entity.TextData) {
	fmt.Fprintf(t.out, "\t    Name: %s\n", data.Name)
	fmt.Fprintf(t.out, "\t    Text: %s\n", data.Text)
	fmt.Fprintf(t.out, "\tMetadata: %s\n", data.Metadata.Data)
}

func (t *Terminal) printTextObjects(data entity.TextObjects) {
	for _, object := range data {
		t.printTextData(object)
		fmt.Fprintf(t.out, "\n")
	}
}

func (t *Terminal) printCardData(data entity.CardData) {
	fmt.Fprintf(t.out, "\t            Name: %s\n", data.Name)
	fmt.Fprintf(t.out, "\t          Number: %s\n", data.Number)
	fmt.Fprintf(t.out, "\t            Code: %d\n", data.Code)
	fmt.Fprintf(t.out, "\t      Cardholder: %s\n", data.Cardholder)
	fmt.Fprintf(t.out, "\tExpiration month: %d\n", data.ExpirationMonth)
	fmt.Fprintf(t.out, "\tExpiration month: %d\n", data.ExpirationYear)
	fmt.Fprintf(t.out, "\t        Metadata: %s\n", data.Metadata.Data)
}

func (t *Terminal) printCardObjects(data entity.CardObjects) {
	for _, object := range data {
		t.printCardData(object)
		fmt.Fprintf(t.out, "\n")
	}
}
