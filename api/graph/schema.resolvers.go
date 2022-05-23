package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/PwrFr/gqlgen/graph/generated"
	"github.com/PwrFr/gqlgen/graph/model"
)

func (r *mutationResolver) CreateRoom(ctx context.Context, input model.NewRoom) (*model.Room, error) {
	room := &model.Room{
		ID:       fmt.Sprintf("T%d", rand.Int()),
		Name:     input.Name,
		Status:   "Avalible",
		Capacity: input.Capacity,
		Type:     input.Type,
	}
	r.rooms = append(r.rooms, room)
	return room, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		ID:   fmt.Sprintf("T%d", rand.Int()),
		Name: input.Name,
	}
	r.users = append(r.users, user)
	return user, nil
}

func (r *queryResolver) Rooms(ctx context.Context) ([]*model.Room, error) {
	// r.RoomRepo.DB = config.NewDBConn()
	return r.RoomRepo.GetRoom()
}

func (r *queryResolver) User(ctx context.Context) ([]*model.User, error) {
	return r.users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
// func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
// 	todo := &model.Todo{
// 		Text: input.Text,
// 		ID:   fmt.Sprintf("T%d", rand.Int()),
// 		User: &model.User{ID: input.UserID, Name: "User " + input.UserID},
// 	}
// 	r.todos = append(r.todos, todo)
// 	return todo, nil
// }
// func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
// 	return r.todos, nil
// }
