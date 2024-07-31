package terminal

const (
	Exiting = "\tExiting..."

	ChoiceBack = "0. Back"
	ChoiceExit = "0. Exit"
	ChoiceReenterChoice = "\tRe-enter your choice!"
	ChoiceEnterChoice = "Enter your choice: "

	UnexpectedError = "\t!!! Unexpected error. Please, contact your administrator. !!!"
	PermissionDenied = "\tPermission denied. Please, login again."
	DataNotFound = "\tRequested data not found."
	DataDeleteWrongName = "\tData with given name not found. Please, try again."
)

type LoginPasswordChoice int

const (
	LoginPassword_Exit LoginPasswordChoice = iota
	LoginPassword_GetByName
	LoginPassword_GetAll
)

type TextChoice int

const (
	Text_Exit TextChoice = iota
	Text_GetByName
	Text_GetAll
)

type CardChoice int

const (
	Card_Exit CardChoice = iota
	Card_GetByName
	Card_GetAll
)

type DeleteChoice int

const (
	Delete_Exit DeleteChoice = iota
	Delete_LoginPassword
	Delete_Text
	Delete_Card
)
