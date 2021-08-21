package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile("[a-z]@[a-z].[a-z]") // membuat regexp

	// mengecek apakah match dengan string
	fmt.Println(regex.MatchString("i@g.com"))
	fmt.Println(regex.MatchString("a@y.co.id"))
	fmt.Println(regex.MatchString("p.id")) // False

	search := regex.FindAllString("i@g.com adm.com a@y.com", 2)
	fmt.Println(search)
}
