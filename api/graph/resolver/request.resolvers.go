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
		StartDatetime:   input.StartDatetime,
		EndDatetime:     input.EndDatetime,
		RequestBy:       input.RequestBy,
		RequestDatetime: time.Now().Format("2006-01-02"),
	}

	return m.RepoDB.InsertRequest(new_req)
}

func (m *mutationResolver) UpdateRequest(ctx context.Context, input model.Approve) (*model.ApproveOutput, error) {
	new_req := &model.ApproveOutput{
		RequestID:       input.RequestID,
		RequestStatus:   input.RequestStatus,
		ApproveDatetime: time.Now().Format("2006-01-02"),
		ApproveBy:       input.ApproveBy,
		Remark:          input.Remark,
	}

	if input.RequestStatus == "approved" {
		err := m.RepoDB.UpdateRoom(input.RoomID, "Not available")
		_ = err
	}

	return m.RepoDB.UpdateRequest(new_req)
}

func (r *queryResolver) Request(ctx context.Context) ([]*model.Request, error) {
	return r.RepoDB.GetRequest()
}
