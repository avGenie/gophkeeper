package entity

// LoginPasswordObjs Array of LoginPassword objects
type LoginPasswordObjects []LoginPassword

// LoginPassword Contains user login and password
type LoginPassword struct {
	Name     ObjectName
	Login    string
	Password string
	Metadata Metadata
}
