package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Ramli Racika", "Ramli")) //mencari kata Ramli dalam kalimat
	fmt.Println(strings.Contains("Ramli Racika", "Raka"))  // return false

	fmt.Println(strings.Split("Ramli Racika", " ")) // mengubah String menjadi Slice

	fmt.Println(strings.ToUpper("Ramli Racika")) // Menjadi Huruf Besar

	fmt.Println(strings.ToLower("Ramli Racika")) // menjadi huruf kecil

	fmt.Println(strings.Trim("     Ramli Racika    ", " ")) //menghapus spasi atau karakter yg ada di kanan dan kiri

	fmt.Println(strings.ReplaceAll("Ramli Racika", "Ramli", "Racika")) // mengubah kata ramli dalam kalimat menjadi racika

}
