package main

import "fmt"

func Divisors(n int) {
	//Put your code here
	var temp []int
	for i:=1; i <= n; i++ {
		if n % i == 0 {
			temp = append(temp,i)
		}
	}
	fmt.Println(temp)
	fmt.Println(len(temp))
}


func main() {
	Divisors(30)
}

