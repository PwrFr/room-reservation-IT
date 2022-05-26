package resolver

import (
	"context"

	"github.com/PwrFr/gqlgen/graph/model"
)

func (m *mutationResolver) CreateRoom(ctx context.Context, input model.NewRoom) (*model.Room, error) {
	new_room := &model.Room{
		RoomName:     input.RoomName,
		RoomStatus:   "Available",
		RoomCapacity: input.RoomCapacity,
		TypeID:       input.TypeID,
	}

	return m.RepoDB.InsertRoom(new_room)
}

func (r *queryResolver) Rooms(ctx context.Context) ([]*model.Room, error) {
	return r.RepoDB.GetRoom()
}
