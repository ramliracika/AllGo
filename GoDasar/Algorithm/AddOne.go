package main

import (
	"fmt"
	"reflect"
)

func Add (int) func(int)int {

	return func(i int) int {
		return i+ Add()
	}

}


func main()  {
	fmt.Println(Add(3))
	fmt.Println(reflect.TypeOf(Add(3)))
}
