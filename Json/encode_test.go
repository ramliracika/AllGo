package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func EncodeJson(data interface{}) {
	bytes, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {
	EncodeJson("Racika")
	EncodeJson(2307)
	EncodeJson(true)
	EncodeJson(map[string]interface{}{
		"Hobby":   "Code",
		"Address": "Melati 5 Bogor",
	})

	EncodeJson([]string{"Maulana", "2101", "Senopaty JKT"})
}
