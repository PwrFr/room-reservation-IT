package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/PwrFr/gqlgen/graph/generated"
	"github.com/PwrFr/gqlgen/graph/model"
)

func (r *queryResolver) Rooms(ctx context.Context) ([]*model.Room, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Account(ctx context.Context) ([]*model.Account, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AccountByID(ctx context.Context, accountID string) (*model.Account, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AccountStudent(ctx context.Context, accountID string) (*model.AccountStudent, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Request(ctx context.Context) ([]*model.Request, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) RequestByID(ctx context.Context, accountID string) ([]*model.Request, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Ipwat(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
