package query

const (
	InsertUserQuery     = "INSERT INTO user (email, password) VALUES (?, ?)"
	GetUserByEmailQuery = "SELECT id, email, password FROM user WHERE email = ?"
)
