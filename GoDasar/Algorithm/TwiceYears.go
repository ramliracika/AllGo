package main

import "fmt"

import "math"

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {

	res := (sonYearsOld * 2) - dadYearsOld

	return int(math.Abs(float64(res)))
}

func main () {

	fmt.Println(TwiceAsOld(29,0))

}

