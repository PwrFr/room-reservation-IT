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
	//if account is null will create new account
	if acc != nil {
		return acc, nil

	}

	_, err := m.RepoDB.InsertAccount(new_acc)
	if err != nil {
		return nil, err
	}

	if new_acc.Email[9:] == "it.kmitl.ac.th" {
		_, err := m.RepoDB.InsertStudent(new_acc.AccountID, new_acc.Email)
		if err != nil {
			fmt.Println("std_err")
			return nil, err
		}
	}

	return new_acc, nil
}

func (r *queryResolver) Account(ctx context.Context) ([]*model.Account, error) {
	return r.RepoDB.GetAccount()
}

func (r *queryResolver) AccountByID(ctx context.Context, accountID string) (*model.Account, error) {
	return r.RepoDB.GetAccountById(accountID)
}

func (r *queryResolver) AccountStudent(ctx context.Context, accountID string) (*model.AccountStudent, error) {
	return r.RepoDB.GetAccountStudentById(accountID)
}
