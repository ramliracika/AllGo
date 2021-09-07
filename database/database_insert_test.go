package database

import (
	"context"
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	db := GetConnet()
	ctx := context.Background()

	script := "INSERT INTO mahasiswa(NPM,NAMA) VALUES('201743502427','Ramli Racika');"

	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Succes Insert Data to Mahasiswa")

}
