package handler

import (
	"github.com/sirateek/poker-be/internal/deck"
	"github.com/sirateek/poker-be/internal/room"
	"github.com/sirateek/poker-be/utils"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DeckService    deck.Service
	RoomService    room.Service
	ContextManager utils.ContextManager
}
