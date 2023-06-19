package service

import (
	"fmt"
	"github.com/katerji/UserAuthKit/db"
	"github.com/katerji/UserAuthKit/db/query"
	"github.com/katerji/UserAuthKit/model"
	"strings"
)

type TokenService struct{}

func (TokenService) GetTokens() ([]model.Token, error) {
	rows, err := db.GetDbInstance().FetchRows(query.FetchTokensQuery)
	if err != nil {
		return []model.Token{}, err
	}
	var tokens []model.Token
	for _, row := range rows {
		token := model.TokenFromDB(row)
		tokens = append(tokens, token)
	}
	return tokens, nil
}

func (TokenService) InsertTokens(response model.TokenAPIResponse) {
	if len(response.Data) == 0 {
		return
	}
	insertQuery := query.InsertTokenBaseQuery
	var args []any
	for _, coinData := range response.Data {
		imageUrl := fmt.Sprintf("https://www.cryptocompare.com%s", coinData.CoinInfo.Symbol)
		coinArgs := []any{
			coinData.CoinInfo.ID,
			coinData.CoinInfo.Name,
			coinData.CoinInfo.Ticker,
			imageUrl,
			coinData.CoinPrice.Raw.Price,
			coinData.CoinPrice.Raw.ChangePercentage,
		}
		placeholders := strings.Repeat("?, ", len(coinArgs))
		placeholders = strings.TrimSuffix(placeholders, ", ")
		insertQuery += fmt.Sprintf("(%s), ", placeholders)
		args = append(args, coinArgs...)
	}
	insertQuery = strings.TrimSuffix(insertQuery, ", ")
	insertQuery += " ON DUPLICATE KEY UPDATE name=VALUES(name), ticker=VALUES(ticker), symbol=VALUES(symbol), price=VALUES(price), change_percentage=VALUES(change_percentage)"
	db.GetDbInstance().Exec(insertQuery, args...)
}
