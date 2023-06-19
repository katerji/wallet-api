package crons

import (
	"encoding/json"
	"fmt"
	"github.com/katerji/UserAuthKit/model"
	"github.com/katerji/UserAuthKit/service"
	"io"
	"net/http"
)

const cryptoCompareAPI = "https://min-api.cryptocompare.com/data/top/totaltoptiervolfull?limit=100&tsym=USD"

func FetchAndStoreTokensFromAPI() {
	tokens, ok := fetchTokens()
	if !ok {
		return
	}
	tokenService := service.TokenService{}
	tokenService.InsertTokens(tokens)
}

func fetchTokens() (model.TokenAPIResponse, bool) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", cryptoCompareAPI, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return model.TokenAPIResponse{}, false
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return model.TokenAPIResponse{}, false
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return model.TokenAPIResponse{}, false
	}
	response := &model.TokenAPIResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		fmt.Println(err)
		return model.TokenAPIResponse{}, false
	}
	return *response, true
}
