package json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamingEncoder(t *testing.T) {

	writer, _ := os.Create("customer_data.json")

	encoder := json.NewEncoder(writer)

	customer := Customer{
		Name:      "Racika",
		Age:       22,
		IsMarried: false,
		Hobbies:   []string{"gaming", "coding", "chill"},
		Addresses: []Address{
			{
				Street:     "Melati 5",
				City:       "Bogor",
				Postalcode: 168,
				Country:    "Indonesia",
			},
		},
	}

	_ = encoder.Encode(customer)

}

func TestStreamingDecoder(t *testing.T) {
	reader, _ := os.Open("customer_data.json")
	decoder := json.NewDecoder(reader)

	customer := &Customer{}
	decoder.Decode(customer)

	fmt.Println(customer)

}
