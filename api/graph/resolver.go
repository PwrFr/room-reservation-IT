package graph

//go:generate go run -mod=mod github.com/99designs/gqlgen generate
import (
	"github.com/PwrFr/gqlgen/graph/model"
	repo "github.com/PwrFr/gqlgen/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	RoomRepo repo.RoomRepo
	rooms    []*model.Room
	users    []*model.User
}
