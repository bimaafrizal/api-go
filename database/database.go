package database

import (
	"database/sql"
)

func InitDB() *sql.DB {
	dsn := "root@tcp(localhost:3306)/todos?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return db
}
