package handler

import "github.com/sirateek/poker-be/services"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DeckService services.Deck
}
