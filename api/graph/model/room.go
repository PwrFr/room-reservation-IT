package model

type NewRoom struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Type     string `json:"type"`
}

type Room struct {
	tableName struct{} `pg:"room"`
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Status    string   `json:"status"`
	Capacity  int      `json:"capacity"`
	Type      string   `json:"type"`
}
