package resolver

import (
	"context"

	"github.com/PwrFr/gqlgen/graph/model"
)

func (m *mutationResolver) CreateRoom(ctx context.Context, input model.NewRoom) (*model.Room, error) {
	new_room := &model.Room{
		Name:     input.Name,
		Status:   "Avalible",
		Capacity: input.Capacity,
		Type:     input.Type,
	}

	return m.RepoDB.InsertRoom(new_room)
}

func (r *queryResolver) Rooms(ctx context.Context) ([]*model.Room, error) {
	// r.RepoDB.DB = config.NewDBConn()
	return r.RepoDB.GetRoom()
}
