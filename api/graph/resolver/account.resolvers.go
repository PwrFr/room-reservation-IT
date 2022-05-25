package resolver

import (
	"context"

	"github.com/PwrFr/gqlgen/graph/model"
)

func (m *mutationResolver) CreateAccount(ctx context.Context, input model.NewAccount) (*model.Account, error) {
	new_acc := &model.Account{
		AccountID: input.AccountID,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Role:      "student",
	}

	return m.RepoDB.InsertAccount(new_acc)
}

func (r *queryResolver) Account(ctx context.Context) ([]*model.Account, error) {
	return r.RepoDB.GeAccount()
}
