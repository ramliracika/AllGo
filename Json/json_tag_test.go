package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	IsActive bool    `json:"is_active"`
}

func TestJsonTagEncode(t *testing.T) {
	product := Product{

		Id:       321,
		Name:     "Mouse",
		Price:    98500,
		IsActive: true,
	}

	bytes, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

}

func TestJsonTagDecode(t *testing.T) {
	jsonstring := `{"id":321,"name":"Mouse","price":98500,"is_active":true}`
	jsonbyte := []byte(jsonstring)

	product := &Product{}

	err := json.Unmarshal(jsonbyte, product)
	if err != nil {
		panic(err)
	}

	fmt.Println(product)
}
