package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirateek/poker-be/internal/player"
	"net/http"
)

type Player struct {
	playerService player.Service
}

func NewPlayer(router *gin.RouterGroup, playerService player.Service) Player {
	playerHandler := Player{
		playerService: playerService,
	}

	router.GET("/:id", playerHandler.GetPlayer)

	return playerHandler
}

type RegisterPlayerRequest struct {
	Name string `json:"name"`
}

func (p *Player) RegisterPlayerHandler(c *gin.Context) {
	var request RegisterPlayerRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request body.",
		})
		return
	}

	playerData, err := p.playerService.RegisterPlayer(request.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, playerData)
}

// GetPlayer is the resolver for the getPlayer field.
func (p *Player) GetPlayer(c *gin.Context) {
	id := c.Param("id")
	result, err := p.playerService.GetPlayer(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, result)
}
