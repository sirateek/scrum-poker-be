package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirateek/poker-be/internal/deck"
	"net/http"
)

type Deck struct {
	DeckService deck.Service
}

func NewDeck(router *gin.RouterGroup, DeckService deck.Service) Deck {
	deckHandler := Deck{
		DeckService: DeckService,
	}
	router.GET("/:id", deckHandler.GetDeck)
	router.GET("/all", deckHandler.GetAvailableDecks)

	return deckHandler
}

// GetDeck is the resolver for the getDeck field.
func (d *Deck) GetDeck(c *gin.Context) {
	id := c.Param("id")
	result, err := d.DeckService.GetDeck(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no id"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// GetAvailableDecks is the resolver for the getAvailableDecks field.
func (d *Deck) GetAvailableDecks(c *gin.Context) {
	result := d.DeckService.GetAllAvailableDecks()
	c.JSON(http.StatusOK, result)
}
