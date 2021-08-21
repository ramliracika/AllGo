package main

import "fmt"

type Pakaian struct {
	Merk string
}

func (pakaian *Pakaian) daftarPakaian() {
	fmt.Println("Merk :", pakaian.Merk)

}

func main() {
	Jeans := Pakaian{"Levis"}

	Chino := &Jeans
	Chino.Merk = "chino"
	Tshirt := &Jeans

	*Tshirt = Pakaian{"Kaos"}

	Jeans.daftarPakaian()
	Chino.daftarPakaian()
	Tshirt.daftarPakaian()

}
