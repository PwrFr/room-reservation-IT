package model

type Student struct {
	tableName struct{} `pg:"student"`
	AccountID int      `json:"account_id"`
	StudentID int      `json:"student_id"`
	Year      string   `json:"year"`
}
