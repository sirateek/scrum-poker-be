package handler

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"fmt"

	"github.com/sirateek/poker-be/graph"
	"github.com/sirateek/poker-be/model"
)

// JoinRoom is the resolver for the joinRoom field.
func (r *mutationResolver) JoinRoom(ctx context.Context, id string, passcode string) (*model.Room, error) {
	panic(fmt.Errorf("not implemented: JoinRoom - joinRoom"))
}

// CreateRoom is the resolver for the createRoom field.
func (r *mutationResolver) CreateRoom(_ context.Context, room *model.CreateRoom) (*model.Room, error) {
	return r.RoomService.CreateRoom(room)
}

// GetDeck is the resolver for the getDeck field.
func (r *queryResolver) GetDeck(ctx context.Context, id string) (*model.Deck, error) {
	return r.DeckService.GetDeck(id)
}

// GetAvailableDecks is the resolver for the getAvailableDecks field.
func (r *queryResolver) GetAvailableDecks(ctx context.Context) ([]*model.Deck, error) {
	return r.DeckService.GetAllAvailableDecks(), nil
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
