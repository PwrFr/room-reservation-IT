package model

type Request struct {
	tableName struct{} `pg:"request"`
	RequestID int      `json:"request_id"`
}
