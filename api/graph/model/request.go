package model

// type Request struct {
// 	tableName       struct{} `pg:"request"`
// 	RequestID       int      `json:"request_id" pg:"request_id, pk"`
// 	RoomID          int      `json:"room_id" pg:"room_id"`
// 	RequestPurpose  string   `json:"request_purpose"`
// 	RequestAttendee int      `json:"request_attendee"`
// 	RequestStatus   string   `json:"request_status"`
// 	StartDatetime   string   `json:"start_datetime"`
// 	EndDatetime     string   `json:"end_datetime"`
// 	RequestBy       string   `json:"request_by" pg:"request_by"`
// 	RequestDatetime string   `json:"request_datetime"`
// 	ApproveBy       string   `json:"approve_by"`
// 	ApproveDatetime string   `json:"approve_datetime"`
// 	Remark          string   `json:"remark"`
// 	Room            *Room    `pg:"rel:has-one"`
// 	Student         *Student `pg:"rel:has-one,join_fk:account_id"`

// }
type Request struct {
	tableName       struct{} `pg:"request"`
	RequestID       int      `json:"request_id" pg:"request_id, pk"`
	RoomID          int      `json:"room_id" pg:"room_id"`
	RequestPurpose  string   `json:"request_purpose"`
	RequestAttendee int      `json:"request_attendee"`
	RequestStatus   string   `json:"request_status"`
	StartDatetime   string   `json:"start_datetime"`
	EndDatetime     string   `json:"end_datetime"`
	RequestBy       string   `json:"request_by" pg:"request_by"`
	RequestDatetime string   `json:"request_datetime"`
	ApproveBy       string   `json:"approve_by"`
	ApproveDatetime string   `json:"approve_datetime"`
	Remark          string   `json:"remark"`
	*Student
	*Account
	RoomName     string `json:"room_name"`
	RoomStatus   string `json:"room_status"`
	RoomCapacity int    `json:"room_capacity"`
	// StudentID       string   `json:"student_id" pg:"student_id,pk"`
	// Phone           string   `json:"phone"`
	// Year            string   `json:"year"`

}

type NewRequest struct {
	RoomID          int    `json:"room_id"`
	RequestPurpose  string `json:"request_purpose"`
	RequestAttendee int    `json:"request_attendee"`
	RequestStatus   string `json:"request_status"`
	StartDatetime   string `json:"start_datetime"`
	EndDatetime     string `json:"end_datetime"`
	RequestBy       string `json:"request_by"`
	RequestDatetime string `json:"request_datetime"`
}

type Approve struct {
	tableName       struct{} `pg:"request"`
	ApproveBy       string   `json:"approve_by"`
	RequestID       int      `json:"request_id"`
	RequestStatus   string   `json:"request_status"`
	ApproveDatetime string   `json:"approve_datetime"`
	Remark          string   `json:"remark"`
	RoomID          int      `json:"room_id"`
}

type ApproveOutput struct {
	tableName       struct{} `pg:"request"`
	ApproveBy       string   `json:"approve_by"`
	RequestID       int      `json:"request_id"`
	RequestStatus   string   `json:"request_status"`
	ApproveDatetime string   `json:"approve_datetime"`
	Remark          string   `json:"remark"`
}
