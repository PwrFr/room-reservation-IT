package repo

import (
	"fmt"

	"github.com/PwrFr/gqlgen/graph/model"
)

// var db = config.NewDBConn()

func (r *RepoDB) GetRoom() ([]*model.Room, error) {
	var room []*model.Room

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

	return room, nil
}

func (r *RepoDB) GetRoomWithRequest() ([]*model.RoomWithRequest, error) {
	var room []*model.RoomWithRequest

	err := r.DB.Model(&room).
		Relation("RoomType").
		Relation("RoomFacility").
		Relation("RoomFacility.Facility").
		Relation("Request").
		// WherePK().
		// Where("room_id = 1").
		Select()
	if err != nil {
		fmt.Println("Err, ", err)
		return nil, err
	}

	return room, nil
}

func (r *RepoDB) InsertRoom(room *model.Room) (*model.Room, error) {
	_, err := r.DB.Model(room).Returning("*").Insert()
	if err != nil {
		return nil, err
	}

	return room, err
}

func (r *RepoDB) UpdateRoom(room_id int, status string) error {
	var room *model.Room
	_, err := r.DB.Model(room).
		Set("room_status = ?", status).
		Where("room_id = ?", room_id).
		Update()
	if err != nil {
		fmt.Println("err", err)
		return err
	}

	return nil
}
