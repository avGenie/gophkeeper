package terminal

import (
	"fmt"
	"strconv"

	"github.com/avGenie/gophkeeper/client/internal/entity"
	"github.com/avGenie/gophkeeper/client/internal/ui/validator"
)

func (t *Terminal) inputLoginPasswordData() entity.LoginPassword {
	var savedObject entity.LoginPassword

	fmt.Fprint(t.out, "Name: ")
	savedObject.Name = entity.ObjectName(t.scanText())

	fmt.Fprint(t.out, "Login: ")
	savedObject.Login = t.scanText()

	fmt.Fprint(t.out, "Password: ")
	savedObject.Password = t.scanText()

	fmt.Fprint(t.out, "Metadata: ")
	savedObject.Metadata.Data = t.scanText()

	return savedObject
}

func (t *Terminal) inputTextData() entity.TextData {
	var savedObject entity.TextData

	fmt.Fprint(t.out, "Name: ")
	savedObject.Name = entity.ObjectName(t.scanText())

	fmt.Fprint(t.out, "Text: ")
	savedObject.Text = t.scanText()

	fmt.Fprint(t.out, "Metadata: ")
	savedObject.Metadata.Data = t.scanText()

	return savedObject
}

func (t *Terminal) inputCardData() entity.CardData {
	var savedObject entity.CardData

	fmt.Fprint(t.out, "Name: ")
	savedObject.Name = entity.ObjectName(t.scanText())

	fmt.Fprint(t.out, "Number: ")
	savedObject.Number = t.scanText()

	for {
		fmt.Fprint(t.out, "Expiration month: ")
		strMonth := t.scanText()

		month, err := strconv.Atoi(strMonth)
		if err != nil {
			fmt.Fprintln(t.out, "Wrong data. Please, enter the number between 1 and 12")
			continue
		}

		err = validator.ValidateExpirationMonth(month)
		if err != nil {
			fmt.Fprintln(t.out, "Wrong data. Please, enter the number between 1 and 12")
			continue
		}

		savedObject.ExpirationMonth = month

		break
	}

	for {
		fmt.Fprint(t.out, "Expiration year: ")
		strMonth := t.scanText()

		year, err := strconv.Atoi(strMonth)
		if err != nil {
			fmt.Fprintln(t.out, "Wrong data. Year should be a number (for example, 2006)")
			continue
		}

		err = validator.ValidateExpirationYear(year)
		if err != nil {
			fmt.Fprintln(t.out, err.Error())
			continue
		}

		savedObject.ExpirationYear = year

		break
	}

	for {
		fmt.Fprint(t.out, "Expiration code: ")
		strCode := t.scanText()

		code, err := strconv.Atoi(strCode)
		if err != nil {
			fmt.Fprintln(t.out, "Wrong data. Please, enter the 3-digit number")
			continue
		}

		err = validator.ValidateCode(code)
		if err != nil {
			fmt.Fprintln(t.out, "Wrong data. Please, enter the 3-digit number")
			continue
		}

		savedObject.Code = code

		break
	}

	fmt.Fprint(t.out, "Cardholder: ")
	savedObject.Cardholder = t.scanText()

	fmt.Fprint(t.out, "Metadata: ")
	savedObject.Metadata.Data = t.scanText()

	return savedObject
}
