package database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestMultiType(t *testing.T) {
	db := GetConnet()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT Id,Name,Address,Email,Phone,Birthday,CreateAt,Rating,isMarried From customer;"

	rows, err := db.QueryContext(ctx, script) // to execute query show data
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	fmt.Println("Succes Get Data From Database")

	for rows.Next() {
		var Id, Name, Address string
		var Email sql.NullString
		var Phone int64
		var Birthday sql.NullTime
		var CreateAt time.Time
		var Rating float32
		var isMarried bool

		err = rows.Scan(&Id, &Name, &Address, &Email, &Phone, &Birthday, &CreateAt, &Rating, &isMarried)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id :", Id)
		fmt.Println("Name :", Name)
		fmt.Println("Address :", Address)
		fmt.Println("Email :", Email)
		fmt.Println("Phone :", Phone)
		fmt.Println("Birthday :", Birthday)
		fmt.Println("Create At :", CreateAt)
		fmt.Println("Rating :", Rating)
		fmt.Println("Is Married :", isMarried)
	}
	fmt.Println("Finished Get Data From Database")
}
