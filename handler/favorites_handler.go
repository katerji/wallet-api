package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/katerji/UserAuthKit/model"
	"github.com/katerji/UserAuthKit/service"
)

const FavoritesPath = "/favorites"

type FavoritesResponse struct {
	Success bool `json:"success"`
}

func FavoritesHandler(c *gin.Context) {
	user := c.MustGet("user").(model.User)
	tokenService := service.TokenService{}
	tokens, err := tokenService.GetUserFavorites(user.ID)
	if err != nil {
		sendInternalErrorResponse(c, "")
		return
	}
	tokenOutputs := []model.TokenOutput{}
	for _, token := range tokens {
		tokenOutputs = append(tokenOutputs, token.ToOutput())
	}
	sendJSONResponse(c, tokenOutputs)
	return
}
