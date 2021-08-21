package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Round(1.6))
	fmt.Println(math.Round(1.4))

	fmt.Println(math.Floor(1.6))
	fmt.Println(math.Ceil(1.4))

	fmt.Println(math.Min(10, 20))
	fmt.Println(math.Max(10, 20))

}
