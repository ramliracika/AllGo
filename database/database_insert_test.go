package database

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)

func TestInsert(t *testing.T) {
	db := GetConnet()
	ctx := context.Background()

	script := "INSERT INTO mahasiswa(NPM,NAMA) VALUES('201743502427','Ramli Racika');"

	_, err := db.ExecContext(ctx, script) // to execute query sql without show data
	if err != nil {
		panic(err)
	}

	fmt.Println("Succes Insert Data to Mahasiswa")

}

func TestMultiInsert(t *testing.T) {
	db := GetConnet()
	ctx := context.Background()

	script := "INSERT INTO customer(Id,Name,Address,Email,Phone,Birthday,Rating,isMarried) VALUES('ID2407','Azcareon','Mars','azcareon@xabre.com','0219876542','22-07-1998','5.0','false'),('ID1121','HayDim','Earth',Null,'0219876542','22-07-1998','4.0','true'),('ID3923','Razor','Pluto',Null,'0219876542',Null,'3.5','false');"

	_, err := db.ExecContext(ctx, script) // to execute query sql without show data
	if err != nil {
		panic(err)
	}

	fmt.Println("Succes Insert Data to Customer")

}

func TestSafeInsert(t *testing.T) {
	db := GetConnet()
	ctx := context.Background()
	username := "admin2"
	password := "admin2"

	script := "INSERT INTO User(username,password) Values (? , ?);"

	_, err := db.ExecContext(ctx, script, username, password) // to execute query sql without show data
	if err != nil {
		panic(err)
	}

	fmt.Println("Succes Insert Data to User")

}

func TestInsertIncement(t *testing.T) {
	db := GetConnet()
	ctx := context.Background()
	email := "ramliracika@gmail.com"
	content := "info yang bagus"

	script := "INSERT INTO Comment(email,content) Values (? , ?);"

	result, err := db.ExecContext(ctx, script, email, content) // to execute query sql without show data
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId() // to get last insert id auto incement
	if err != nil {
		panic(err)
	}

	fmt.Println("Succes Insert Data to Id", insertId)

}

func TestPrepare(t *testing.T) {
	db := GetConnet()
	ctx := context.Background()
	script := "INSERT INTO Comment(email,content) Values (? , ?);"

	statement, err := db.PrepareContext(ctx, script) // to execute query sql without show data
	if err != nil {
		panic(err)
	}
	for i := 1; i < 10; i++ {
		email := "ramliracika" + strconv.Itoa(i) + "@gmail.com"
		content := "info yang bagus" + strconv.Itoa(i)

		statement, err := statement.ExecContext(ctx, email, content) // to execute query sql without show data
		if err != nil {
			panic(err)
		}

		insertId, err := statement.LastInsertId() // to get last insert id auto incement
		if err != nil {
			panic(err)
		}

		fmt.Println("Succes Insert Data to Id", insertId)
	}
	defer statement.Close()
}
