package player

import (
	"errors"
	"github.com/sirateek/poker-be/model"
	"github.com/sirateek/poker-be/utils"
)

type playerService struct {
	players map[string]*model.Player
}

type Service interface {
	RegisterPlayer(name string) (*model.Player, error)
	GetPlayer(userID string) (*model.Player, error)
}

var (
	ErrPlayerNotFound = errors.New("player not found")
)

func NewService() Service {
	return &playerService{
		players: map[string]*model.Player{},
	}
}

func (p *playerService) RegisterPlayer(name string) (*model.Player, error) {
	userID := utils.RandStr(30)
	p.players[userID] = &model.Player{
		ID:   userID,
		Name: name,
	}
	return p.players[userID], nil
}

func (p *playerService) GetPlayer(userID string) (*model.Player, error) {
	player, ok := p.players[userID]
	if !ok {
		return nil, ErrPlayerNotFound
	}
	return player, nil
}
