package main

import "fmt"

func single() {
	fmt.Println("Ini single")
}

func withpm(a string) {
	fmt.Println("Hasil dari parameter :", a)
}

func pmrt(b string) string {
	return b
}

func doublert(c, d int) (int, int) {
	return c, d
}

func named() (ramli string) {
	ramli = "Ramli Racika"
	return ramli
}
func variad(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	single()
	withpm("dengan parameter")
	fmt.Println(pmrt("ini hasilnya"))
	fmt.Println(doublert(10, 10))
	fmt.Println(named())

	//1
	total := variad(10, 10, 10)
	fmt.Println(total)

	//2
	slice := []int{10, 20, 30}
	total = variad(slice...)
	fmt.Println(total)

}
