package database

import (
	"context"
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {
	db := GetConnet()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT NPM,NAMA From Mahasiswa;"

	rows, err := db.QueryContext(ctx, script) // to execute query show data
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	fmt.Println("Succes Get Data From Database")
	for rows.Next() {
		var npm int64
		var nama string

		err = rows.Scan(&npm, &nama)
		if err != nil {
			panic(err)
		}

		fmt.Println("NPM :", npm)
		fmt.Println("NAMA :", nama)

	}
	fmt.Println("Finished Get Data From Database")
}
