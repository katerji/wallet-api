package query

const (
	InsertUserQuery     = "INSERT INTO user (email, username, password) VALUES (?, ?, ?)"
	GetUserByEmailQuery = "SELECT id, email, username, password FROM user WHERE email = ?"
)
