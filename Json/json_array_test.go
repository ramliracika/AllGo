package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestArrayJson(t *testing.T) {
	customer := &Customer{
		Name:    "Racika",
		Hobbies: []string{"Gaming", "Coding", "Chill"},
	}

	byte, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(byte))
}

func TestDecodeArrayJson(t *testing.T) {
	jsonString := `{"Name":"Racika","Age":0,"Address":"","IsMarried":false,"Hobbies":["Gaming","Coding","Chill"],"Addresses":null}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Hobbies)
}

func TestComplexArrayJson(t *testing.T) {
	customer := &Customer{
		Name: "Haryadi",
		Addresses: []Address{
			{
				Street:     "Jalan Kapling",
				City:       "Bogor",
				Postalcode: 168,
				Country:    "Indonesia",
			},
			{
				Street:     "Jalan Mawar",
				City:       "Jakarta",
				Postalcode: 999,
				Country:    "Indonesia",
			},
		},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

}

func TestDecodeComplexArray(t *testing.T) {
	jsonString := `{"Name":"Haryadi","Age":0,"Address":"","IsMarried":false,"Hobbies":null,"Addresses":[{"Street":"Jalan Kapling","City":"Bogor","Postalcode":168,"Country":"Indonesia"},{"Street":"Jalan Mawar","City":"Jakarta","Postalcode":999,"Country":"Indonesia"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Addresses)

}

func TestOnlyArrayDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalan Kapling","City":"Bogor","Postalcode":168,"Country":"Indonesia"},{"Street":"Jalan Mawar","City":"Jakarta","Postalcode":999,"Country":"Indonesia"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}

	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)

}

func TestOnlArrayJsonEncode(t *testing.T) {
	adresses := []Address{
		{
			Street:     "Jalan Kapling",
			City:       "Bogor",
			Postalcode: 168,
			Country:    "Indonesia",
		},
		{
			Street:     "Jalan Mawar",
			City:       "Jakarta",
			Postalcode: 999,
			Country:    "Indonesia",
		},
	}

	bytes, err := json.Marshal(adresses)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

}
