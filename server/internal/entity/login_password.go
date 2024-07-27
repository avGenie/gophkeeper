package entity

// LoginPassword Contains name, login, password and metadata as data content
type LoginPassword struct {
	Name     string
	Login    string
	Password string
	Metadata Metadata
}
