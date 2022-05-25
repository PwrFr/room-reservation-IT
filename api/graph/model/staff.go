package model

type Staff struct {
	tableName struct{} `pg:"staff"`
	AccountID int      `json:"account_id"`
}
