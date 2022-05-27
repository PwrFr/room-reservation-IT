package repo

import (
	"fmt"

	"github.com/PwrFr/gqlgen/graph/model"
)

func (r *RepoDB) InsertRequest(request *model.RequestOutput) (*model.RequestOutput, error) {
	_, err := r.DB.Model(request).Returning("*").Insert()
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	return request, nil
}

func (r *RepoDB) GetRequest() ([]*model.Request, error) {
	var request []*model.Request
	stm := `
	SELECT req.*, s.student_id, s.year, 
	a.first_name,  a.last_name, a.email, 
	r.room_name, r.room_status, r.room_capacity 
	FROM request AS req 
	join student AS s on req.request_by = s.account_id 
	join account AS a on s.account_id = a.account_id 
	join room AS r on r.room_id = req.room_id 
`

	_, err := r.DB.Query(&request, stm)
	if err != nil {
		return nil, err
	}
	// x, _ := json.Marshal(request)
	// fmt.Println(string(x))
	return request, nil
}

// func (r *RepoDB) GetRequest() ([]*model.Request, error) {
// 	var request []*model.Request
// 	err := r.DB.Model(&request).Relation("Room").Relation("student").Select()
// 	if err != nil {
// 		return nil, err
// 	}
// 	// x, _ := json.Marshal(request)
// 	// fmt.Println(string(x))
// 	return request, nil
// }

// func (r *RepoDB) GetRequestById(input string) ([]*model.Request, error) {
// 	var request []*model.Request
// 	err := r.DB.Model(&request).Relation("Room").Where("request_by = ?", input).Select()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return request, nil
// }

//for student will show account as staff
func (r *RepoDB) GetRequestById(input string) ([]*model.Request, error) {
	var request []*model.Request
	stm := `
	SELECT req.*, 
	a.first_name,  a.last_name, a.email, 
	r.room_name, r.room_status, r.room_capacity  
	FROM request AS req 
	join account AS a on req.approve_by = a.account_id 
	join room AS r on r.room_id = req.room_id 
	WHERE req.request_by = ?`

	_, err := r.DB.Query(&request, stm, input)
	if err != nil {
		return nil, err
	}
	// x, _ := json.Marshal(request)
	// fmt.Println(string(x))
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
