package database

import "fmt"

var connection string

func init() {
	fmt.Println("init running")
	connection = "MySQL"

}

func GetDatabase() string {
	return connection

}