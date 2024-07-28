package entity

// LoginPassword Contains user login and password
type LoginPassword struct {
	Name     string
	Login    string
	Password string
	Metadata Metadata
}
