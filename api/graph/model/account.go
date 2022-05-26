package model

type Account struct {
	tableName struct{} `pg:"account"`
	AccountID string   `json:"account_id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Role      string   `json:"role"`
	Email     string   `json:"email"`
}

type AccountStudent struct {
	tableName struct{}   `pg:"account"`
	AccountID string     `json:"account_id" pg:"account_id,pk"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Role      string     `json:"role"`
	Email     string     `json:"email"`
	Student   []*Student `pg:"rel:has-many"`
}

type NewAccount struct {
	AccountID string `json:"account_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type Student struct {
	tableName struct{} `pg:"student"`
	AccountID string   `json:"account_id" pg:"account_id,join_fk:fk_account_student"`
	StudentID string   `json:"student_id"`
	ProgramID int      `json:"program_id"`
	Phone     string   `json:"phone"`
	Year      string   `json:"year"`
}
