package main

import "fmt"

type Handphone struct {
	Nama, Merk string
	Ram, Harga int
}

func Tampil(hp *Handphone) {
	fmt.Println("Nama :", hp.Nama)
	fmt.Println("Merk :", hp.Merk)
	fmt.Println("RAM :", hp.Ram)
	fmt.Println("Harga :", hp.Harga)

}

func main() {
	hp := Handphone{"Note 10", "Xiaomi", 8, 2000}
	hp2 := &hp

	Tampil(&hp)
	fmt.Println(" ")
	Tampil(hp2)

}
