package model

type User struct {
	ID    int
	Email string
}

func (user User) ToOutput() UserOutput {
	return UserOutput{
		ID:    user.ID,
		Email: user.Email,
	}
}
