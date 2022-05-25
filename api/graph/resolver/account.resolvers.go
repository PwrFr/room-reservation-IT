package resolver

import (
	"context"
	"fmt"

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

	acc, _ := m.RepoDB.GetAccountById(input.AccountID)
	fmt.Println(acc == nil)
	if acc == nil {
		fmt.Println("1")
		return m.RepoDB.InsertAccount(new_acc)
	}
	fmt.Println("2")
	return acc, nil
}

// func (m *mutationResolver) CreateAccount(ctx context.Context, input model.NewAccount) (*model.Account, error) {

// 	new_acc := &model.Account{
// 		AccountID: input.AccountID,
// 		FirstName: input.FirstName,
// 		LastName:  input.LastName,
// 		Email:     input.Email,
// 		Role:      "student",
// 	}

// 	return m.RepoDB.InsertAccount(new_acc)
// }

func (r *queryResolver) Account(ctx context.Context) ([]*model.Account, error) {
	return r.RepoDB.GetAccount()
}

func (r *queryResolver) AccountByID(ctx context.Context, accountID string) (*model.Account, error) {
	return r.RepoDB.GetAccountById(accountID)
}
