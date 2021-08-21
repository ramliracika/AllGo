package main

import "fmt"

//deklarasi interface kosong
// tidak ada function didalamnya

// type anything interface {
// 	//return "a"
// 	//retun 5
// 	//return false
// }

func kosong(i int) interface{} { //bisa mengganti interface dengan anything
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Bukan satu dan dua"
	}
}

func main() {
	kosong := kosong(3)
	fmt.Println(kosong)
}
