package main

import "fmt"

func FaktorBilangan(n int)  {
	for i:=1; i <= n; i++ {
		if n % i == 0 {
			fmt.Println(i)
		}
	}
}


func main()  {
	FaktorBilangan(20)
}
