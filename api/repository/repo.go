package repo

import (
	"github.com/PwrFr/gqlgen/graph/model"
	"github.com/go-pg/pg/v10"
)

// var db = config.NewDBConn()

type RoomRepo struct {
	DB *pg.DB
}

func (r *RoomRepo) GetRoom() ([]*model.Room, error) {
	var room []*model.Room
	err := r.DB.Model(&room).Select()
	if err != nil {
		return nil, err
	}
	return room, nil
}
