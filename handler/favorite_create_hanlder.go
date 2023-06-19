package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/katerji/UserAuthKit/model"
	"github.com/katerji/UserAuthKit/service"
)

const FavoriteCreatePath = "/favorite"

type FavoriteCreateRequest struct {
	TokenID string `json:"token_id"`
}

type FavoriteCreateResponse struct {
	Success bool `json:"success"`
}

func FavoriteCreateHandler(c *gin.Context) {
	var request FavoriteCreateRequest
	err := c.BindJSON(&request)
	if err != nil {
		sendBadRequest(c)
		return
	}
	user := c.MustGet("user").(model.User)
	tokenService := service.TokenService{}
	ok := tokenService.CreateFavorite(user.ID, request.TokenID)
	response := FavoriteCreateResponse{ok}
	sendJSONResponse(c, response)
	return
}
