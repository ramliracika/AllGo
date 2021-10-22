package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJson(t *testing.T) {
	JsonString := `{"Name":"Racika","Age":22,"Address":"Bogor","IsMarried":false}`
	JsonBytes := []byte(JsonString)

	customer := &Customer{}

	err := json.Unmarshal(JsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Name)
	fmt.Println(customer.Age)
	fmt.Println(customer.Address)
	fmt.Println(customer.IsMarried)

}
