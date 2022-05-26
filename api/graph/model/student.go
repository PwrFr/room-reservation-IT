package model

type Student struct {
	tableName struct{} `pg:"student"`
	AccountID string   `json:"account_id"`
	StudentID string   `json:"student_id"`
	ProgramID string   `json:"program_id"`
	Phone     string   `json:"phone"`
	Year      string   `json:"year"`
}
