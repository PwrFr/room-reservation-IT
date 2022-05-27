package resolver

//go:generate go run -mod=mod github.com/99designs/gqlgen generate
import (
	"context"
	"os"

	repo "github.com/PwrFr/gqlgen/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	RepoDB repo.RepoDB
	// rooms   []*model.Room
	// account []*model.Account
}

func (r *queryResolver) Ipwat(ctx context.Context) (*string, error) {
	ip := os.Getenv("FRONTEND_HOST")
	return &ip, nil
}
