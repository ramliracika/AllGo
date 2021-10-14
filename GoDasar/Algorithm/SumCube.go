package main

import (
	"fmt"
	"math"
)

func SumCubes(n int) int {
	var s int
	for i := 1; i <= n; i++ {
		s += int(math.Pow(float64(i),3))

	}
	return s

}

func main()  {
	fmt.Println(SumCubes(123))
}
