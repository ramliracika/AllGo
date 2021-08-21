package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Ramli")
	data.PushBack("Racika")
	data.PushBack("Haryadi")
	data.PushBack("Maulana")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	fmt.Println(" ")

	// data from front to back
	// 	for value := data.Front(); value != nil; value = value.Next() {
	// 		fmt.Println(value.Value)
	// 	}
	// }

	// reverse data
	for Reverse := data.Back(); Reverse != nil; Reverse = Reverse.Prev() {
		fmt.Println(Reverse.Value)
	}
}
