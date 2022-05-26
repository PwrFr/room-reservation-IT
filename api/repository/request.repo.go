package repo

import (
	"fmt"

	"github.com/PwrFr/gqlgen/graph/model"
)

func (r *RepoDB) InsertRequest(request *model.Request) (*model.Request, error) {
	_, err := r.DB.Model(request).Returning("*").Insert()
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	return request, nil
}

func (r *RepoDB) GetRequest() ([]*model.Request, error) {
	var request []*model.Request
	err := r.DB.Model(&request).Relation("Room").Select()
	if err != nil {
		return nil, err
	}
	// x, _ := json.Marshal(request)
	// fmt.Println(string(x))
	return request, nil
}

func (r *RepoDB) GetRequestById(input string) ([]*model.Request, error) {
	var request []*model.Request
	err := r.DB.Model(&request).Relation("Room").Where("request_by = ?", input).Select()
	if err != nil {
		return nil, err
	}

	return request, nil
}

func (r *RepoDB) UpdateRequest(request *model.ApproveOutput) (*model.ApproveOutput, error) {
	_, err := r.DB.Model(request).Where("request_id = ?", request.RequestID).Returning("request_id, request_status, approve_by, approve_datetime").Update()
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	return request, nil
}

func (r *RepoDB) RemoveRequest(request_id int) (*string, error) {
	var req *model.Request
	_, err := r.DB.Model(req).Where("request_id = ?", request_id).Delete()
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	return nil, nil
}
