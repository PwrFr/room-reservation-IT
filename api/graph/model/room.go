package model

type NewRoom struct {
	RoomName     string `json:"room_name"`
	RoomCapacity int    `json:"room_capacity"`
	TypeID       int    `json:"type_id"`
}

type Room struct {
	tableName    struct{}        `pg:"room"`
	RoomID       int             `json:"room_id" pg:"room_id,pk"`
	RoomName     string          `json:"room_name"`
	RoomStatus   string          `json:"roomstatus"`
	RoomCapacity int             `json:"room_capacity"`
	TypeID       int             `json:"type_id" pg:"type_id"`
	RoomType     *RoomType       `pg:"rel:has-one"`
	RoomFacility []*RoomFacility `pg:"rel:has-many"`
}
type RoomType struct {
	tableName struct{} `pg:"room_type"`
	TypeID    int      `json:"type_id" pg:"type_id,pk"`
	TypeName  string   `json:"type_name"`
}
type RoomFacility struct {
	tableName  struct{}  `pg:"room_facility"`
	FacilityID int       `json:"facility_id" pg:"facility_id,join_fk:room_facility_facility_id_fkey"`
	RoomID     int       `json:"room_id" pg:"join_fk:room_facility_room_id_fkey"`
	Quantity   int       `json:"quantity"`
	Facility   *Facility `pg:"rel:has-one"`
}

type Facility struct {
	tableName    struct{} `pg:"facility"`
	FacilityID   int      `json:"facility_id" pg:"facility_id,pk"`
	FacilityName string   `json:"facility_name"`
}
