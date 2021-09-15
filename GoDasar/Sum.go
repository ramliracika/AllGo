package main

import "fmt"

func Total(Treshold int) int {
	var arr int
	for i := 1; i <= Treshold; i++ {
		arr += i
	}
	return arr
}

func main() {
	fmt.Println(Total(5))

}
