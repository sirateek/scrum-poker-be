package deck

import (
	"errors"
	"github.com/sirateek/poker-be/model"
)

type deckService struct {
	decks    []*model.Deck
	decksMap map[string]*model.Deck
}

type Service interface {
	GetAllAvailableDecks() []*model.Deck
	GetDeck(id string) (*model.Deck, error)
}

var (
	ErrDeckNotFound = errors.New("deck not found")
)

func NewDeck(decks []*model.Deck) Service {
	decksMap := make(map[string]*model.Deck)
	for _, deck := range decks {
		decksMap[deck.ID] = deck
	}

	return &deckService{
		decks:    decks,
		decksMap: decksMap,
	}
}

func (d *deckService) GetAllAvailableDecks() []*model.Deck {
	return d.decks
}

func (d *deckService) GetDeck(id string) (*model.Deck, error) {
	value, ok := d.decksMap[id]
	if !ok {
		return nil, ErrDeckNotFound
	}
	return value, nil
}
