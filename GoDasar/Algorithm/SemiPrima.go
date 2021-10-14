package main

import "fmt"

func SemiPrima(n int)  {


		if n%n == 0 && n%1 == 0 && n% 2 == 0 {
			fmt.Println(n,"YA")
		} else {
			fmt.Println(n,"TIDAK")
		}

}

func main()  {
	SemiPrima(3)
}