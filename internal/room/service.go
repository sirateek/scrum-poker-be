package room

import (
	"errors"
	"github.com/sirateek/poker-be/internal/deck"
	"github.com/sirateek/poker-be/internal/player"
	"github.com/sirateek/poker-be/model"
	"github.com/sirateek/poker-be/utils"
	"github.com/sirupsen/logrus"
)

var (
	ErrRoomNotFound  = errors.New("room not found")
	ErrWrongPasscode = errors.New("wrong passcode")
)

type roomService struct {
	rooms         map[string]*model.Room
	deckService   deck.Service
	playerService player.Service
	playerRoomMap map[*model.Player]*model.Room
}

type Service interface {
	CreateRoom(room *model.CreateRoom) (*model.Room, error)
	JoinRoom(userID string, roomID string, passcode string) (result bool, err error)
	GetRoom(roomID string) (*model.Room, error)
}

func NewService(deckService deck.Service, playerService player.Service) Service {
	return &roomService{
		deckService:   deckService,
		playerService: playerService,
		rooms:         map[string]*model.Room{},
		playerRoomMap: make(map[*model.Player]*model.Room),
	}
}

func (r *roomService) CreateRoom(room *model.CreateRoom) (*model.Room, error) {
	id := utils.RandStr(10)
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

func (r *roomService) JoinRoom(userID string, roomID string, passcode string) (result bool, err error) {
	room, ok := r.rooms[roomID]
	if !ok || room == nil {
		return false, ErrRoomNotFound
	}

	if room.Passcode == nil || passcode != *room.Passcode {
		return false, ErrWrongPasscode
	}

	playerData, err := r.playerService.GetPlayer(userID)
	if err != nil {
		return false, err
	}

	room.Players = append(room.Players, playerData)
	return true, nil
}

func (r *roomService) GetRoom(roomID string) (*model.Room, error) {
	room, ok := r.rooms[roomID]
	if !ok {
		return nil, ErrRoomNotFound
	}
	return room, nil
}
