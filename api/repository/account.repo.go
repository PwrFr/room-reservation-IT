package repo

import (
	"fmt"

	"github.com/PwrFr/gqlgen/graph/model"
	"github.com/go-pg/pg/v10"
)

type RepoDB struct {
	DB *pg.DB
}

func (r *RepoDB) GeAccount() ([]*model.Account, error) {
	var account []*model.Account
	err := r.DB.Model(&account).Select()
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (r *RepoDB) InsertAccount(account *model.Account) (*model.Account, error) {
	_, err := r.DB.Model(account).Returning("*").Insert()

	return account, err
}
