package main

import "fmt"

var namamhs = [...]string{
	"Andi",
	"Reza",
	"Budi",
	"Firman",
}
var kelasmhs = [...]string{
	"S8K",
	"R1H",
	"R7K",
	"S7X",
}
var prodimhs = [...]string{
	"Arsitek",
	"Informatika",
	"Fisika",
	"Matematika",
}

var npmmhs = [...]int64{
	201743502427,
	202043501515,
	201164350678,
	201743508905,
}

var nilai = [...]int32{
	67,
	89,
	73,
	92,
}

var nama = namamhs
var kelas = kelasmhs
var prodi = prodimhs
var penilai = nilai

var npm = npmmhs

func mhsbaru() int {
	return 0
}

func mhslama() int {

	return 1
}

func main() {
	maba := mhsbaru()
	mala := mhslama()

	var j int
	var k = len(nama) - 1

	if maba == 2 {
		j = 2

	} else if mala == 0 {
		j = 0
		k = 1

	} else {
		j = 0
		k = -1
		fmt.Println("  ")
		fmt.Println("Maaf tidak ada data yang dapat ditampilkan")
		fmt.Println("  ")

	}

	for i := j; i <= k; i++ {

		fmt.Println("Nama : ", nama[i])
		fmt.Println("NPM : ", npm[i])
		fmt.Println("Kelas : ", kelas[i])
		fmt.Println("Program Studi : ", prodi[i])
		fmt.Println("Total Nilai : ", penilai[i])

		if penilai[i] > 80 {
			fmt.Println("Selamat anda lolos")
		} else {
			fmt.Println("Maaf anda tidak lolos")
		}

		fmt.Println(" ")
	}

}
