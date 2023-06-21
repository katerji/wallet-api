package model

type Token struct {
	ID               string
	TokenID          string
	Name             string
	Ticker           string
	Symbol           string
	Price            float64
	ChangePercentage float64
}

func (token Token) ToOutput() TokenOutput {
	return TokenOutput(token)
}
