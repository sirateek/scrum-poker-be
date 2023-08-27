package utils

import (
	"github.com/sirateek/poker-be/config"
	"github.com/sirateek/poker-be/model"
)

type DeckConfigParser struct{}

func NewDeckConfigParser() *DeckConfigParser {
	return &DeckConfigParser{}
}

func (d *DeckConfigParser) Parse(deckConfigs []config.Deck) []*model.Deck {
	result := make([]*model.Deck, 0, len(deckConfigs))

	for _, deckConfig := range deckConfigs {
		deck := &model.Deck{
			ID:   deckConfig.ID,
			Name: deckConfig.Name,
		}
		cards := make([]*model.Card, 0, len(deckConfig.Cards))

		for index, card := range deckConfig.Cards {
			cards = append(cards, &model.Card{
				Index:        index,
				DisplayValue: card,
			})
		}

		deck.Cards = cards
		result = append(result, deck)
	}

	return result
}
