package repo

import (
	"fmt"
	"strconv"
	"time"

	"github.com/PwrFr/gqlgen/graph/model"
	"github.com/go-pg/pg/v10"
)

type RepoDB struct {
	DB *pg.DB
}

func (r *RepoDB) GetAccount() ([]*model.Account, error) {
	var account []*model.Account
	err := r.DB.Model(&account).Select()
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (r *RepoDB) GetAccountById(input string) (*model.Account, error) {
	var account model.Account
	err := r.DB.Model(&account).Where("account_id = ?", input).Select()
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (r *RepoDB) InsertAccount(account *model.Account) (*model.Account, error) {
	_, err := r.DB.Model(account).Returning("*").Insert()

	return account, err
}

func (r *RepoDB) InsertStudent(account_id string, email string) (*model.Student, error) {
	year := time.Now().Year()
	st_year, _ := strconv.Atoi(email[:2])

	student := &model.Student{
		AccountID: account_id,
		StudentID: email[:8],
		Year:      strconv.Itoa(year - 1957 - st_year),
		ProgramID: "1",
	}

	_, err := r.DB.Model(student).Returning("*").Insert()

	return student, err
}
