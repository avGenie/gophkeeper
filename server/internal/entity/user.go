package entity

// UserID User id struct
type UserID string

// User User struct. Keeps user id, login and password.
type User struct {
	ID       UserID
	Login    string
	Password string
}
