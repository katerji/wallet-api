package model

type User struct {
	ID       int
	Email    string
	Username string
}

func (user User) ToOutput() UserOutput {
	return UserOutput(user)
}

func UserFromDB(row map[string]any) User {
	id := row["id"].(int)
	email := row["email"].(string)
	username := row["username"].(string)
	return User{
		ID:       id,
		Email:    email,
		Username: username,
	}
}
