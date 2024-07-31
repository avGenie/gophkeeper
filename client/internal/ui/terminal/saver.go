package terminal

import (
	"fmt"
	"strconv"

	"github.com/avGenie/gophkeeper/client/internal/entity"
	"github.com/avGenie/gophkeeper/client/internal/ui/validator"
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
	var savedObject entity.LoginPassword

	fmt.Fprintln(t.out, "Input new object data")
	fmt.Fprint(t.out, "Name: ")
	savedObject.Name = entity.ObjectName(t.scanText())

	fmt.Fprint(t.out, "Login: ")
	savedObject.Login = t.scanText()

	fmt.Fprint(t.out, "Password: ")
	savedObject.Password = t.scanText()

	fmt.Fprint(t.out, "Metadata: ")
	savedObject.Metadata.Data = t.scanText()

	err := t.loginPasswordProc.SaveData(savedObject)
	if err != nil {
		return t.saveProcessError(err, "failed to save login-password data object")
	}

	return nil
}

func (t *Terminal) saveText() error {
	var savedObject entity.TextData

	fmt.Fprintln(t.out, "Input new object data")
	fmt.Fprint(t.out, "Name: ")
	savedObject.Name = entity.ObjectName(t.scanText())

	fmt.Fprint(t.out, "Text: ")
	savedObject.Text = t.scanText()

	fmt.Fprint(t.out, "Metadata: ")
	savedObject.Metadata.Data = t.scanText()

	err := t.textProc.SaveData(savedObject)
	if err != nil {
		return t.saveProcessError(err, "failed to save text data object")
	}

	return nil
}

func (t *Terminal) saveCard() error {
	var savedObject entity.CardData

	fmt.Fprintln(t.out, "Input new object data")
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
			fmt.Fprintln(t.out, "Wrong data. Please, enter the number between 1920 and 2090")
			continue
		}

		err = validator.ValidateExpirationMonth(year)
		if err != nil {
			fmt.Fprintln(t.out, "Wrong data. Please, enter the number between 1920 and 2090")
			continue
		}

		savedObject.ExpirationYear = year

		break
	}

	for {
		fmt.Fprint(t.out, "Expiration year: ")
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

	err := t.cardProc.SaveData(savedObject)
	if err != nil {
		return t.saveProcessError(err, "failed to save card data object")
	}

	return nil
}
