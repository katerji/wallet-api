package query

const (
	InsertUserQuery     = "INSERT INTO user (email, username, password) VALUES (?, ?, ?)"
	GetUserByEmailQuery = "SELECT id, email, username, password FROM user WHERE email = ?"
	EmailExistsQuery    = "SELECT email FROM user WHERE email = ?"
	UsernameExistsQuery = "SELECT username FROM user WHERE username = ?"
)
