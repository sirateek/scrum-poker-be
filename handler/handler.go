package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirateek/poker-be/internal/player"
	"net/http"
)

type NonGraphHandler struct {
	playerService player.Service
}

type RegisterPlayerRequest struct {
	Name string `json:"name"`
}

func (n *NonGraphHandler) RegisterPlayerHandler(c *gin.Context) {
	var request RegisterPlayerRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request body.",
		})
		return
	}

	playerData, err := n.playerService.RegisterPlayer(request.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, playerData)
}
