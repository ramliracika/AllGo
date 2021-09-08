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

	rows, err := db.QueryContext(ctx, script) // to execute query sql without show data
	if err != nil {
		panic(err)
	}

	fmt.Println("Succes Insert Data to Mahasiswa")

	defer rows.Close()

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

}
