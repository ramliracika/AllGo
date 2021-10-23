package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "456",
		"name":  "monitor",
		"price": 575000,
	}

	bytes, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJsonMapDecode(t *testing.T) {
	jsonstring := `{"id":"456","name":"monitor","price":575000}`
	jsonbytes := []byte(jsonstring)

	var product map[string]interface{}

	err := json.Unmarshal(jsonbytes, &product)
	if err != nil {
		panic(err)
	}

	fmt.Println(product)
	fmt.Println(product["id"])
	fmt.Println(product["name"])
	fmt.Println(product["price"])

}
