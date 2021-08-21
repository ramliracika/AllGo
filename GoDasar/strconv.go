package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("True") // mengubah string menjadi boolean
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	number, err := strconv.ParseInt("1000000", 10, 64) // mengubah string ke int dengan base decimal dan menjadi int64

	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	value := strconv.FormatInt(10000, 10) // mengubah ke format integer
	fmt.Println(value)

	valueint, _ := strconv.Atoi("50000") // cara tercepat mengubah string ke int
	fmt.Println(valueint)

}
