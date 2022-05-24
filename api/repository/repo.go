package repo

import (
	"github.com/PwrFr/gqlgen/graph/model"
	"github.com/go-pg/pg/v10"
)

// var db = config.NewDBConn()

type RepoDB struct {
	DB *pg.DB
}

func (r *RepoDB) GetRoom() ([]*model.Room, error) {
	var room []*model.Room
	err := r.DB.Model(&room).Select()
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (r *RepoDB) InsertRoom(room *model.Room) (*model.Room, error) {
	_, err := r.DB.Model(room).Returning("*").Insert()

	return room, err
}
