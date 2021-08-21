package main

import (
	"errors"
	"fmt"
)

func luasSegitiga(alas, tinggi int) (int, error) {
	if alas == 0 || tinggi == 0 {
		return 0, errors.New("Angka Tidak Boleh Kosong")

	} else {
		result := alas * tinggi / 2
		return result, nil
	}

}

func main() {
	hasil, err := luasSegitiga(10, 0)

	if err == nil {
		fmt.Println("Hasil :", hasil)
	} else {
		fmt.Println("Error : ", err.Error())
	}

}
