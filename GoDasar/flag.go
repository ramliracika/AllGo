package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("Host", "LocalHost", "Put Your database host")
	username := flag.String("Username", "root", "Put Your database username")
	password := flag.String("Password", "root", "Put Your database password")

	flag.Parse()

	fmt.Println("Host :", *host)
	fmt.Println("Username :", *username)
	fmt.Println("Password :", *password)
}
