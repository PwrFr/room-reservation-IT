package resolver

import (
	"context"
	"errors"

	"github.com/PwrFr/gqlgen/graph/model"
	"github.com/PwrFr/gqlgen/graph/resolver/middleware"
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

func (r *queryResolver) RoomWithRequest(ctx context.Context) ([]*model.RoomWithRequest, error) {
	return r.RepoDB.GetRoomWithRequest()
}

func (m *mutationResolver) UpdateRoom(ctx context.Context, room_id int, status string, token *string) (*string, error) {
	if auth := middleware.RoleValid("staff", token); !auth {
		return nil, errors.New("Forbidden lol")
	}

	err := m.RepoDB.UpdateRoom(room_id, status)
	return nil, err
}
