package model

type UserOutput struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func (output UserOutput) ToUser() User {
	return User(output)
}
