package query

const InsertTokenBaseQuery = "INSERT INTO token (token_id, name, ticker, symbol, price, change_percentage) VALUES "
const FetchTokensQuery = "SELECT id, token_id, name, ticker, symbol, price, change_percentage FROM token LIMIT 1000"
