package entity

// LoginPasswordObjs Array of LoginPassword objects
type LoginPasswordObjs []LoginPassword

// LoginPassword Contains name, login, password and metadata as data content
type LoginPassword struct {
	Name     string
	Login    string
	Password string
	Metadata Metadata
}
