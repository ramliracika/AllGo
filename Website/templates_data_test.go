package Website

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplatesDataInterface(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/data.gohtml"))

	t.ExecuteTemplate(writer, "data.gohtml", map[string]interface{}{
		"Title": "Data Interface",
		"Name":  "Ramli",
		"Address": map[string]interface{}{
			"Street": "Bogor",
		},
	})
}

func TestTestDataInterface(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http:/localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplatesDataInterface(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))

}

type Address struct {
	Street string
}

type Data struct {
	Title   string
	Name    string
	Address Address
}

func TemplatesDataStruct(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/data.gohtml"))

	t.ExecuteTemplate(writer, "data.gohtml", Data{
		Title: "Data Struct",
		Name:  "Racika",
		Address: Address{
			Street: "Bogor",
		},
	})
}

func TestTestDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http:/localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplatesDataStruct(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))

}
