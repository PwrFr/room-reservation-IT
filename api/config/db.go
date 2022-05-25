package config

import (
	"fmt"

	"github.com/go-pg/pg/v10"
)

// func init() {
// 	ConnectDB()
// }

// func ConnectDB() *sql.DB {
// 	pg := fmt.Sprintf(
// 		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 		Env("PG_HOST"),
// 		Env("PG_PORT"),
// 		Env("PG_USER"),
// 		Env("PG_PASSWORD"),
// 		Env("PG_DBNAME"))

// 	db, err := sql.Open("postgres", pg)
// 	if err != nil {
// 		fmt.Println("Pg: failed to connect to database: ", err)
// 		return &sql.DB{}
// 	}

// 	fmt.Println("Pg: connect to database success!!")

// 	return db
// }

var db *pg.DB

func init() {
	db = NewDBConn()

}

func CallDB() *pg.DB {
	return db
}

func NewDBConn() (con *pg.DB) {
	address := fmt.Sprintf("%s:%s", Env("PG_HOST"), Env("PG_PORT"))
	options := &pg.Options{
		User:     Env("PG_USER"),
		Password: Env("PG_PASSWORD"),
		Addr:     address,
		Database: Env("PG_DBNAME"),
		PoolSize: 50,
	}
	con = pg.Connect(options)
	if con == nil {
		fmt.Println("Pg: failed to connect to database: ", con)
	}

	fmt.Println("Pg: connect to database success!!")

	// defer con.Close()

	return
}
