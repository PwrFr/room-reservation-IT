package resolver

import (
	"context"
	"fmt"

	"github.com/PwrFr/gqlgen/graph/model"
)

func (m *mutationResolver) CreateAccount(ctx context.Context, input model.NewAccount) (*model.Account, error) {
	acc, _ := m.RepoDB.GetAccountById(input.AccountID)

	if acc != nil {
		return acc, nil

	}

	//if account(topper) is null will create new account
	new_acc := &model.Account{
		AccountID: input.AccountID,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Role:      "guest",
	}

	is_student := new_acc.Email[9:] == "it.kmitl.ac.th"
	if is_student {
		new_acc.Role = "student"
	}

	_, err := m.RepoDB.InsertAccount(new_acc)
	if err != nil {
		return nil, err
	}

	if is_student {
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
