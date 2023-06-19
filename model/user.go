package model

type User struct {
	ID       int
	Email    string
	Username string
}

func (user User) ToOutput() UserOutput {
	return UserOutput(user)
}
