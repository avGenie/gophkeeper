package entity

type DataRequestType int

const (
	DataRequestLoginPassword DataRequestType = iota
	DataRequestText
	DataRequestCard
	DataRequestInvalid
)

type LoginPasswordObjects []LoginPassword

type LoginPassword struct {
	Name     string
	Login    string
	Password string
	Metadata string
}

type TextObjects []TextData

type TextData struct {
	Name     string
	Text     string
	Metadata string
}

type CardObjects []CardData

type CardData struct {
	Name            string
	Number          string
	ExpirationMonth int
	ExpirationYear  int
	Code            int
	Cardholder      string
	Metadata        string
}
