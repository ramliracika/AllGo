package main

import (
	"errors"
	"fmt"
	"reflect"
)

// function pembanding data

type Product struct {
	Product_id, Stock, Price int
	Name                     string
	Status                   bool
}

func (product *Product) PriceValidation() (int, error) {
	if product.Price <= 1000 && product.Price > 100 {
		return 1, nil
	} else if product.Price < 100 {
		return 2, errors.New("Price lower than 100")
	} else {
		return 3, errors.New("Price more than 1000")
	}
}

func (product *Product) show() {
	fmt.Println("Product_Id :", product.Product_id)
	fmt.Println("Product_Name :", product.Name)
	fmt.Println("Product_Price :", product.Price)
	fmt.Println("Product_Stock :", product.Stock)
	fmt.Println("Is Active :", product.Status)
}

func main() {
	product := Product{123, 10, 150, "Product A", true}

	hasil, err := product.PriceValidation()
	if err != nil {
		fmt.Println("Eror :", err.Error())
	}

	if hasil == 1 {
		product.show()
	}

	v := reflect.ValueOf(product)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fmt.Println(typeOfS.Field(i).Name, v.Field(i).Interface())
	}
}
