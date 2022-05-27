package model

type Account struct {
	tableName struct{} `pg:"account"`
	AccountID string   `json:"account_id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Role      string   `json:"role"`
	Email     string   `json:"email"`
}

type AccountWithToken struct {
	tableName struct{} `pg:"account"`
	AccountID string   `json:"account_id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Role      string   `json:"role"`
	Email     string   `json:"email"`
	Token     *string  `json:"token"`
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
	AccountID string   `json:"account_id" pg:"account_id"`
	StudentID string   `json:"student_id" pg:"student_id"`
	ProgramID int      `json:"program_id"`
	Phone     string   `json:"phone"`
	Year      string   `json:"year"`
}
