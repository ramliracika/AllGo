package database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestConnect(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/Golang")

	if err != nil {
		panic(err)
	}

	defer db.Close()
}
