package query

const InsertTokenBaseQuery = "INSERT INTO token (id, name, ticker, symbol, price, change_percentage) VALUES "
const FetchTokensQuery = "SELECT id, name, ticker, symbol, price, change_percentage FROM token LIMIT 1000"
