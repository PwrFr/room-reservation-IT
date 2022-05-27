package resolver

//go:generate go run -mod=mod github.com/99designs/gqlgen generate
import (
	"context"

	repo "github.com/PwrFr/gqlgen/repository"
	"github.com/joho/godotenv"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	RepoDB repo.RepoDB
	// rooms   []*model.Room
	// account []*model.Account
}

func InitEnv() map[string]string {
	_ = godotenv.Load()
	env, _ := godotenv.Read()

	return env
}

//calling env
func Env(val string) string {
	return InitEnv()[val]
}

func (r *queryResolver) Ipwat(ctx context.Context) (*string, error) {
	ip := Env("FRONTEND_IP")
	x := ip + ":  1"
	return &x, nil
}
