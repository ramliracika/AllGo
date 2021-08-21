package main

import "fmt"

func main() {
	// kemampuan function untuk berinteraksi dengan data-data disekitarnya dalam scope yang sama

	//harap gunakan fitur closure dengan bijak saat membuat aplikasi

	counter := 0
	increment := func() {
		fmt.Println("Increment")
		counter++ // penyebab nilai berubah
		// -->	counter = 5  // deklarasi yang salah
		// -->	counter := 5 //deklarasi yang benar
	}

	increment()
	increment()
	fmt.Println(counter)

}
