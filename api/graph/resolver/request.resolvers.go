package resolver

import (
	"context"
	"time"

	"github.com/PwrFr/gqlgen/graph/model"
)

func (m *mutationResolver) CreateRequest(ctx context.Context, input model.NewRequest) (*model.Request, error) {
	new_req := &model.Request{
		RoomID:          input.RoomID,
		RequestPurpose:  input.RequestPurpose,
		RequestAttendee: input.RequestAttendee,
		RequestStatus:   "pending",
		StartDatetime:   time.Now(),
		EndDatetime:     time.Now(),
		RequestBy:       input.RequestBy,
		RequestDatetime: time.Now(),
	}

	return m.RepoDB.InsertRequest(new_req)
}

func (r *queryResolver) Request(ctx context.Context) ([]*model.Request, error) {
	return r.RepoDB.GetRequest()
}
