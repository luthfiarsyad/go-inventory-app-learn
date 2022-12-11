package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDatabase() *sql.DB {
	con := "root:root@tcp(localhost:3306)/inventory_app?parseTime=True"
	db, err := sql.Open("mysql", con)
	if err != nil {
		panic(err)
	}
	return db
}
