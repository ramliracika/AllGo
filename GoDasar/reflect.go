package main

import (
	"fmt"
	"reflect"
)

type Identity struct {
	Id   string `required:"true"`
	Name string `required:"true"`
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}

	}
	return true
}

func main() {

	Idn := Identity{"123", "Ramli"}

	SampleType := reflect.TypeOf(Idn)

	fmt.Println("Jumlah Field :", SampleType.NumField())
	fmt.Println("Nama Field Pertama :", SampleType.Field(0).Name)
	fmt.Println("Isi Tag Required field Pertama :", SampleType.Field(0).Tag.Get("required")) //mengambil nilai dari tag required field pertama

	Idn.Id = ""
	fmt.Println("Isi dari Validasi 1 :", IsValid(Idn))

	contoh := Identity{"123", "esf"}
	fmt.Println("Isi dari Validasi 2 :", IsValid(contoh))

}
