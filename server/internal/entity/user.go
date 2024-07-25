package entity

type UserID string

type User struct {
	ID       UserID
	Login    string
	Password string
}
