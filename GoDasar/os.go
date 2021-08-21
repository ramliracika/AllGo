package main

import (
	"fmt"
	"os"
)

func main() {
	a := os.Args

	fmt.Println("Argumen : ", a)

	username := os.Getenv("App_username")
	password := os.Getenv("App_password")

	fmt.Println(username)
	fmt.Println(password)

}
