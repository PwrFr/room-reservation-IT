package repo

import (
	"fmt"

	"github.com/PwrFr/gqlgen/graph/model"
)

// var db = config.NewDBConn()

func (r *RepoDB) GetRoom() ([]*model.Room, error) {
	var room []*model.Room

	// stm := `select * from room as r
	// join room_type as rt on r.type_id = rt.type_id
	// `
	// _, err := r.DB.Query(&room, stm)

	err := r.DB.Model(&room).
		Relation("RoomType").
		Relation("RoomFacility").
		Relation("RoomFacility.Facility").
		// WherePK().
		// Where("room_id = 1").
		Select()
	if err != nil {
		fmt.Println("Err, ", err)
		return nil, err
	}

	// x, _ := json.Marshal(room)
	// fmt.Println(string(x))
	return room, nil
}

func (r *RepoDB) InsertRoom(room *model.Room) (*model.Room, error) {
	_, err := r.DB.Model(room).Returning("*").Insert()
	if err != nil {
		return nil, err
	}

	return room, err
}
