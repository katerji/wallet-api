package model

type TokenAPIResponse struct {
	Data []CoinData `json:"Data"`
}

type CoinData struct {
	CoinInfo  CoinInfo  `json:"CoinInfo"`
	CoinPrice CoinPrice `json:"RAW"`
}
type CoinInfo struct {
	ID     string `json:"Id"`
	Name   string `json:"FullName"`
	Ticker string `json:"Name"`
	Symbol string `json:"ImageURL"`
}

type CoinPrice struct {
	Raw CoinPriceRAW `json:"USD"`
}

type CoinPriceRAW struct {
	Price            float64 `json:"PRICE"`
	ChangePercentage float64 `json:"CHANGEPCT24HOUR"`
}
