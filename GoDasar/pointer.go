package main

import "fmt"

type Motor struct {
	nama, merk string
	harga      int
}

type Mobil struct {
	nama, merk string
	harga      int
}

func main() {

	matic1 := Motor{"Beat", "Honda", 16}
	matic2 := &matic1           //matic 2 sebagai pointer matic 1
	var matic3 *Motor = &matic1 // matic 3 sebagai pointer matic 1

	matic2.nama = "Vario" // merubah nama dan affect ke semua data dengan field nama

	matic3 = &Motor{"Aerox", "Yamaha", 30} // membuat data baru, merubah data tanpa affect ke lain

	fmt.Println(" ")

	fmt.Println(matic1) // berubah nama motor menjadi vario
	fmt.Println(matic2) // berubah nama motor menjadi vario

	fmt.Println(matic3) // berubah isi menjadi aerox

	fmt.Println(" ")

	//============================================

	manual1 := Mobil{"Pajero", "Daihatsu", 500} // data awal
	var manual2 *Mobil = &manual1

	*manual2 = Mobil{"Brio", "Honda", 300} // merubah semua data

	fmt.Println("================== ")
	fmt.Println(" ")

	fmt.Println(manual1)
	fmt.Println(manual2)

	fmt.Println(" ")

	fmt.Println("================== ")
	fmt.Println(" ")
	//========================================================

	// create new pointer

	manual3 := new(Mobil)

	manual3.nama = "Avanza"

	fmt.Println(manual3)

}
