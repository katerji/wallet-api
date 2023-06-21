package model

type TokenOutput struct {
	ID               string  `json:"id"`
	TokenID          string  `json:"token_id"`
	Name             string  `json:"name"`
	Ticker           string  `json:"ticker"`
	Symbol           string  `json:"symbol"`
	Price            float64 `json:"price"`
	ChangePercentage float64 `json:"change_percentage"`
}
