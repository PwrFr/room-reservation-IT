package resolver

import (
	"context"
	"fmt"

	"github.com/PwrFr/gqlgen/graph/model"
	"github.com/PwrFr/gqlgen/graph/resolver/middleware"
)


func (m *mutationResolver) CreateAccount(ctx context.Context, input model.NewAccount) (*model.AccountWithToken, error) {

	new_acc := &model.Account{
		AccountID: input.AccountID,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Role:      "guest",
	}


	acc, _ := m.RepoDB.GetAccountById(input.AccountID)

	//if account is null will create new account
	if acc != nil {
		token, _ := middleware.GenerateJwtToken(acc.Role)

		acc_token := &model.AccountWithToken{AccountID: acc.AccountID,
			FirstName: acc.FirstName,
			LastName:  acc.LastName,
			Email:     acc.Email,
			Role:      acc.Role,
			Token:     &token}

		return acc_token, nil



	}

	isStudent := new_acc.Email[9:] == "it.kmitl.ac.th"
	if isStudent {
		new_acc.Role = "student"
	}
	_, err := m.RepoDB.InsertAccount(new_acc)
	if err != nil {
		return nil, err
	}


	if isStudent {

	

		_, err := m.RepoDB.InsertStudent(new_acc.AccountID, new_acc.Email)
		if err != nil {
			fmt.Println("std_err")
			return nil, err
		}
	}

	token, _ := middleware.GenerateJwtToken(new_acc.Role)
	new_acc_token := &model.AccountWithToken{
		AccountID: new_acc.AccountID,
		FirstName: new_acc.FirstName,
		LastName:  new_acc.LastName,
		Email:     new_acc.Email,
		Role:      new_acc.Role,
		Token:     &token,
	}

	return new_acc_token, nil
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
