package model

type Account struct {
	tableName struct{} `pg:"account"`
	AccountID string   `json:"account_id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Role      string   `json:"role"`
	Email     string   `json:"email"`
}

type NewAccount struct {
	AccountID string `json:"account_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
