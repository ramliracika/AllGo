package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	City       string
	Postalcode int
	Country    string
}

type Customer struct {
	Name      string
	Age       int
	Address   string
	IsMarried bool
	Hobbies   []string
	Addresses []Address
}

func TestObjectJson(t *testing.T) {
	customer := Customer{
		Name:      "Racika",
		Age:       22,
		Address:   "Bogor",
		IsMarried: false,
	}

	byte, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(byte))
}
