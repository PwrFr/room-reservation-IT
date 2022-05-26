package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/PwrFr/gqlgen/graph/generated"
	"github.com/PwrFr/gqlgen/graph/model"
)

func (r *mutationResolver) CreateRoom(ctx context.Context, input model.NewRoom) (*model.Room, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateAccount(ctx context.Context, input model.NewAccount) (*model.Account, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateRequest(ctx context.Context, input model.NewRequest) (*model.Request, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateRequest(ctx context.Context, input model.Approve) (*model.ApproveOutput, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
