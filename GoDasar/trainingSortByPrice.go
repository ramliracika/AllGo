package main

import (
	"errors"
	"fmt"
)

// function pembanding harga

type Product struct {
	Product_id, Stock, Price int
	Name                     string
	Status                   bool
}

func SortByPrice(product *Product, product2 *Product) (int, error) {
	if product.Price > product2.Price {
		return 1, nil
	} else if product.Price < product2.Price {
		return 2, nil
	} else {
		return 0, errors.New("Product cannot sorted, because Price is Equals to Other")
	}

}

func Show1(product *Product) {
	fmt.Println("Product_Id :", product.Product_id)
	fmt.Println("Product_Name :", product.Name)
	fmt.Println("Product_Price :", product.Price)
	fmt.Println("Product_Stock :", product.Stock)
	fmt.Println("Is Active :", product.Status)
	fmt.Println(" ")
}

func Show2(product2 *Product) {
	fmt.Println("Product_Id :", product2.Product_id)
	fmt.Println("Product_Name :", product2.Name)
	fmt.Println("Product_Price :", product2.Price)
	fmt.Println("Product_Stock :", product2.Stock)
	fmt.Println("Is Active :", product2.Status)
	fmt.Println(" ")

}

func main() {
	product := Product{123, 10, 1500, "Product A", true}
	product2 := &Product{456, 20, 3000, "Product B", false}

	hasil, err := SortByPrice(&product, product2)
	if err != nil {
		fmt.Println("Eror :", err.Error())
	}

	if hasil == 1 {
		Show1(&product)
		Show2(product2)
	} else if hasil == 2 {
		Show2(product2)
		Show1(&product)
	} else {
		fmt.Println("")
	}
}
