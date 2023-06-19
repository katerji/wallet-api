package query

const FetchUserFavorites = "SELECT id, name, ticker, symbol, price, change_percentage FROM token t JOIN favorite f ON t.id = f.token_id WHERE f.user_id = ?"
const InsertUserFavorite = "INSERT INTO favorite (user_id, token_id) VALUES (?, ?)"
const DeleteUserFavorite = "DELETE FROM favorite WHERE user_id = ? AND token_id = ?"
