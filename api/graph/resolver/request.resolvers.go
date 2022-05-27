package resolver

import (
	"context"
	"time"

	"github.com/PwrFr/gqlgen/graph/model"
)

func (m *mutationResolver) CreateRequest(ctx context.Context, input model.NewRequest) (*model.RequestOutput, error) {
	new_req := &model.RequestOutput{
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

func (r *queryResolver) RequestByID(ctx context.Context, accountID string) ([]*model.Request, error) {
	return r.RepoDB.GetRequestById(accountID)
}

func (m *mutationResolver) RemoveRequest(ctx context.Context, req_id int) (*string, error) {
	return m.RepoDB.RemoveRequest(req_id)
}
