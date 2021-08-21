package helper

import "fmt"

func Hi(name string) { // dapat di akses oleh package lain karna awalan Huruf Besar
	fmt.Println("Hi ", name)
}

func sayGoodBye(name string) {
	fmt.Println("GoodBye! ", name)
}
