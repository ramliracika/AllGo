package main

import "fmt"

//deklarasi struct
type Biodata struct {
	Name, Address string
	Age           int
	IsMarried     bool
}

//membuat struct function

func (biodata Biodata) Intro() {
	fmt.Println("Name :", biodata.Name)
	fmt.Println("Address :", biodata.Address)
	fmt.Println("Age :", biodata.Age)
	fmt.Println("Is Married :", biodata.IsMarried)

}

func main() {
	// contoh pengisian data struct 1
	var Andi Biodata
	Andi.Name = "Andi"
	Andi.Address = "Jakarta"
	Andi.Age = 40
	Andi.IsMarried = true

	// contoh pengisian data struct 2
	Baim := Biodata{
		Name:      "Baim",
		Address:   "Bogor",
		Age:       23,
		IsMarried: false,
	}

	fmt.Println(Andi)
	fmt.Println(" ")
	fmt.Println(Baim)
	fmt.Println(" ")

	// memanggil struct function
	Andi.Intro()
	fmt.Println(" ")
	Baim.Intro()

}
