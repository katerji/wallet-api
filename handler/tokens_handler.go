package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/katerji/UserAuthKit/model"
	"github.com/katerji/UserAuthKit/service"
)

const TokensPath = "/tokens"

type TokensResponse struct {
	tokens []model.TokenOutput
}

func TokensHandler(c *gin.Context) {
	tokenService := service.TokenService{}
	tokens, err := tokenService.GetTokens()
	if len(tokens) == 0 || err != nil {
		sendInternalErrorResponse(c, "error fetching tokens")
		return
	}
	tokenOutput := []model.TokenOutput{}
	for _, token := range tokens {
		tokenOutput = append(tokenOutput, token.ToOutput())
	}
	sendJSONResponse(c, tokenOutput)
}
