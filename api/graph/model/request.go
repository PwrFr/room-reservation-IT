package model

import "time"

type Request struct {
	tableName       struct{}  `pg:"request"`
	RequestID       int       `json:"request_id" pg:"request_id, pk"`
	RoomID          int       `json:"room_id"`
	RequestPurpose  string    `json:"request_purpose"`
	RequestAttendee int       `json:"request_attendee"`
	RequestStatus   string    `json:"request_status"`
	StartDatetime   time.Time `json:"start_datetime"`
	EndDatetime     time.Time `json:"end_datetime"`
	RequestBy       string    `json:"request_by"`
	RequestDatetime time.Time `json:"request_datetime"`
	ApproveBy       string    `json:"approve_by"`
	ApproveDatetime time.Time `json:"approve_datetime"`
	Remark          string    `json:"remark"`
}

type NewRequest struct {
	RoomID          int       `json:"room_id"`
	RequestPurpose  string    `json:"request_purpose"`
	RequestAttendee int       `json:"request_attendee"`
	RequestStatus   string    `json:"request_status"`
	StartDatetime   time.Time `json:"start_datetime"`
	EndDatetime     time.Time `json:"end_datetime"`
	RequestBy       string    `json:"request_by"`
	RequestDatetime time.Time `json:"request_datetime"`
}
