package room

import (
	"github.com/sirateek/poker-be/internal/deck"
	"github.com/sirateek/poker-be/model"
	"github.com/sirupsen/logrus"
	"math/rand"
)

// define the given charset, char only
var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// n is the length of random string we want to generate
func randStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		// randomly select 1 character from given charset
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

type roomService struct {
	rooms       map[string]*model.Room
	deckService deck.Service
}

type Service interface {
	CreateRoom(room *model.CreateRoom) (*model.Room, error)
}

func NewService(deckService deck.Service) Service {
	return &roomService{
		deckService: deckService,
	}
}

func (r *roomService) CreateRoom(room *model.CreateRoom) (*model.Room, error) {
	id := randStr(10)
	deckData, err := r.deckService.GetDeck(room.DeckID)
	if err != nil {
		logrus.Error("Get Deck Error: ", err)
		return nil, err
	}
	r.rooms[id] = &model.Room{
		ID:       id,
		Name:     room.Name,
		Deck:     deckData,
		Passcode: &room.Passcode,
	}

	return r.rooms[id], nil
}
