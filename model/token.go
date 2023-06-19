package model

import "github.com/katerji/UserAuthKit/db"

type Token struct {
	ID               string
	Name             string
	Ticker           string
	Symbol           string
	Price            float64
	ChangePercentage float64
}

func (token Token) ToOutput() TokenOutput {
	return TokenOutput(token)
}

func TokenFromDB(row map[string]any) Token {
	id := db.BytesToString(row["id"])
	name := db.BytesToString(row["name"])
	ticker := db.BytesToString(row["ticker"])
	symbol := db.BytesToString(row["symbol"])
	price := db.StringToFloat64(row["price"])
	changePercentage := db.StringToFloat64(row["change_percentage"])
	return Token{
		ID:               id,
		Name:             name,
		Ticker:           ticker,
		Symbol:           symbol,
		Price:            price,
		ChangePercentage: changePercentage,
	}
}
