package database

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)

func TestTransaction(t *testing.T) {
	db := GetConnet()
	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	// do transaction here
	script := "INSERT INTO Comment(email,content) Values (? , ?);"

	for i := 1; i < 10; i++ {
		email := "ramli" + strconv.Itoa(i) + "@gmail.com"
		content := "info yang bagus" + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, content) // to execute query sql without show data
		if err != nil {
			panic(err)
		}

		insertId, err := result.LastInsertId() // to get last insert id auto incement
		if err != nil {
			panic(err)
		}

		fmt.Println("Succes Insert Data to Id", insertId)
	}

	tx.Commit() // -> to commit transaction

	/*err = tx.Rollback()  -> to rollback transaction
	if err != nil {
		panic(err)
	} */
}
