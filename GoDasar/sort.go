package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)

}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age

}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]

}

func main() {
	users := []User{
		{"Ramli", 23},
		{"Racika", 32},
		{"Haryadi", 11},
		{"Maulana", 21},
	}
	fmt.Println("Before Sort ")
	fmt.Println(users)
	fmt.Println(" ")

	sort.Sort(UserSlice(users))

	fmt.Println("After Sort ")
	fmt.Println(users)
}
