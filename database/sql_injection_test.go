package database

import (
	"context"
	"fmt"
	"testing"
)

func TestInjection(t *testing.T) {
	db := GetConnet()
	defer db.Close()

	ctx := context.Background()

	username := "admin';#" // succes login but username wrong
	password := "admin"

	script := "SELECT username From User WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1 "

	rows, err := db.QueryContext(ctx, script) // to execute query show data
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	if rows.Next() {
		var username string

		err = rows.Scan(&username)
		fmt.Println("Sukses Login")
		if err != nil {
			panic(err)
		}

	} else {
		fmt.Println("Login Gagal")
	}

}

func TestInjectionsafe(t *testing.T) {
	db := GetConnet()
	defer db.Close()

	ctx := context.Background()

	username := "admin';#" // cannot login
	password := "admin"

	script := "SELECT username From User WHERE username = ?  AND password = ?  LIMIT 1 " // sign "?" mark sql need parameter

	rows, err := db.QueryContext(ctx, script, username, password) // add username and password
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	if rows.Next() {
		var username string

		err = rows.Scan(&username)
		fmt.Println("Sukses Login")
		if err != nil {
			panic(err)
		}

	} else {
		fmt.Println("Login Gagal")
	}

}
