package main

import "fmt"

//  Panic

// function untuk menghentikan program
// biasanya dipanggil ketika terjadi eror pada saat program berjalan
// saat di panggil, program akan berhenti namun defer tetap berjalan

//=======================================================================

// Recovery

// Adalah function untuk melanjutkan program walaupun ada panic
// function recovery harus ditempatkan didalam function defer

//======================================================================

func deferFunc() {
	fmt.Println("Aplicaation Ended")
	message := recover() // menangkap eror dan di tampilkan dalam pesan
	if message != nil {
		fmt.Println("eror with message :", message)
	}

}

func terminateApp(eror bool) {
	defer deferFunc() //memanggil function walau pada function ini terjadi eror
	if eror {
		panic("Aplication ERROR !")
	}

	fmt.Println("Aplication Start")
}

func main() {
	terminateApp(true) // true akan menampilkan panic
}
