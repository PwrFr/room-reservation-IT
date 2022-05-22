package graph

//go:generate go run -mod=mod github.com/99designs/gqlgen generate
import "github.com/PwrFr/gqlgen/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	rooms []*model.Room
	users []*model.User
}
