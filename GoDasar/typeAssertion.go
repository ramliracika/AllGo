package main

import (
	"fmt"
)

func Assert() interface{} {
	return true

}

func main() {
	result := Assert()

	switch value := result.(type) { //merubah interface kosong menjadi type data yg sesuai
	case string:
		fmt.Println("Value :", value, "is String")
		break
	case int:
		fmt.Println("Value :", value, "is Integer")
		break
	case bool:
		fmt.Println("Value :", value, "is Boolean")
		break
	}

}
