package Website

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

type Page struct {
	Title string
	Name  string
}

func TemplatesIf(writer http.ResponseWriter, request *http.Request) {
	x := template.Must(template.ParseFiles("./templates/if.gohtml"))

	x.ExecuteTemplate(writer, "if.gohtml", Page{
		Title: "Templates Data Struct",
		Name:  "Ramli",
	})

}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplatesIf(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}

func TemplateComparator(writer http.ResponseWriter, request *http.Request) {
	x := template.Must(template.ParseFiles("./templates/comparation.gohtml"))

	x.ExecuteTemplate(writer, "comparation.gohtml", map[string]interface{}{
		"Title": "Template Action Operator",
		"Value": 90,
	})

}

func TestTemplateComparator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateComparator(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}

func TemplateRange(writer http.ResponseWriter, request *http.Request) {
	x := template.Must(template.ParseFiles("./templates/range.gohtml"))
	x.ExecuteTemplate(writer, "range.gohtml", map[string]interface{}{
		"Title":   "Template Action Range",
		"Hobbies": []string{"Play Game", "Coding", "Relax"},
	})

}

func TestTemplateRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recoder := httptest.NewRecorder()

	TemplateRange(recoder, request)

	body, _ := io.ReadAll(recoder.Result().Body)
	fmt.Println(string(body))
}

func TemplateWith(writer http.ResponseWriter, request *http.Request) {
	x := template.Must(template.ParseFiles("./templates/address.gohtml"))
	x.ExecuteTemplate(writer, "address.gohtml", map[string]interface{}{
		"Title": "Template Action With",
		"Name":  "Ramli",
		"Address": map[string]interface{}{
			"City":   "Bogor",
			"Street": "Melati V",
		},
	})
}

func TestTemplateWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recoder := httptest.NewRecorder()

	TemplateWith(recoder, request)

	body, _ := io.ReadAll(recoder.Result().Body)
	fmt.Println(string(body))
}
