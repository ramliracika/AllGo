package database

import (
	"database/sql"
	"time"
)

func GetConnet() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/Golang")

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)                 // set minimum connection
	db.SetMaxOpenConns(50)                 // set max connection
	db.SetConnMaxIdleTime(5 * time.Minute) // set if connection idle will close
	db.SetConnMaxLifetime(1 * time.Hour)   // set if connection time expired will re-new to minimum connection

	return db
}
