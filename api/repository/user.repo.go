package repo

import (
	"github.com/PwrFr/gqlgen/graph/model"
)

func (r *RepoDB) GetUser() ([]*model.User, error) {
	var room []*model.User
	err := r.DB.Model(&room).Select()
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (r *RepoDB) InsertUser(user *model.User) (*model.User, error) {
	_, err := r.DB.Model(user).Returning("*").Insert()

	return user, err
}
