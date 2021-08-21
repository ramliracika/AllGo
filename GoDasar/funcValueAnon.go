package main

import "fmt"

func Hello(name string, filter func(string) string) {
	filtername := filter(name)
	fmt.Println("Hello", filtername)

}
func Fname(name string) string {
	if name == "Anjing" || name == "babi" {
		return "****"
	} else {
		return name
	}

}

//ANON FUNCTION

type Block func(string) bool // deklarasi block sebagai fungsi

func blacklist(name string, block Block) {
	if block(name) {
		fmt.Println("Anda Dilarang Masuk", name)
	} else {
		fmt.Println("Selamat Datang", name)
	}
}

func main() {
	// function value manggil 1
	Hello("Andi", Fname)

	//function value manggil 2
	filter := Fname
	Hello("babi", filter)

	//anon funcion manggil 1
	block := func(name string) bool {
		return name == "Reza"
	}

	blacklist("Rexa", block)
	blacklist("Ramli", block)

	//anon function manggil 2
	blacklist("User", func(name string) bool {
		return name == "Admin"
	})

}
