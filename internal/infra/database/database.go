package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB  *sql.DB
	err error
)

func OpenDB() {
	DB, err = sql.Open("mysql", "root:root@/aluraflix?parseTime=true")

	if err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}
