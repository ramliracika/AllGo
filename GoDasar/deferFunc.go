package main

import "fmt"

// DEFER
// adalah fungsi yg bisa dijadwalkan untuk di eksekusi setelah sebuah function di eksekusi
// akan selalu di eksekusi walau terjadi eror di function yg di eksekusi

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApp(value int) {
	defer logging() // menjalankan function setelah function runApp selesai dijalankan walau ada eror
	fmt.Println("Run App")
	result := 10 / value
	fmt.Println("Result :", result)

}

//==================================================================

func main() {
	runApp(0) // erpr karna 0 dibagi 10
}
