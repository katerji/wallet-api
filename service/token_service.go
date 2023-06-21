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
	tokens := []model.Token{}
	rows, err := db.GetDbInstance().FetchRows(query.FetchTokensQuery)
	defer rows.Close()
	if err != nil {
		return tokens, err
	}
	for rows.Next() {
		token := model.Token{}
		rows.Scan(&token.ID, &token.TokenID, &token.Name, &token.Ticker, &token.Symbol, &token.Price, &token.ChangePercentage)
		tokens = append(tokens, token)
	}

	return tokens, nil
}

func (TokenService) GetUserFavorites(userId int) ([]model.Token, error) {
	tokens := []model.Token{}
	rows, err := db.GetDbInstance().FetchRows(query.FetchUserFavorites, userId)
	defer rows.Close()
	if err != nil {
		return tokens, err
	}
	for rows.Next() {
		token := model.Token{}
		rows.Scan(&token.ID, &token.Name, &token.Ticker, &token.Symbol, &token.Price, &token.ChangePercentage)
		tokens = append(tokens, token)
	}

	return tokens, nil
}

func (TokenService) CreateFavorite(userId int, tokenId string) bool {
	return db.GetDbInstance().Exec(query.InsertUserFavorite, userId, tokenId)
}

func (TokenService) DeleteFavorite(userId int, tokenId string) bool {
	return db.GetDbInstance().Exec(query.DeleteUserFavorite, userId, tokenId)
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
