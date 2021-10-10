package Website

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.gohtml
var templatess embed.FS

var t = template.Must(template.ParseFS(templatess, "templates/*.gohtml"))

func TemplateCaching(writer http.ResponseWriter, request *http.Request) {
	t.ExecuteTemplate(writer, "templates.gohtml", "Hello Template Caching")
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func BenchmarkCaching(b *testing.B) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()
	for i := 0; i < b.N; i++ {
		TemplateCaching(recorder, request)
	}
}

//

func TemplateWithoutCaching(writer http.ResponseWriter, request *http.Request) {
	var x = template.Must(template.ParseFS(templatess, "templates/*.gohtml"))
	x.ExecuteTemplate(writer, "templates.gohtml", "Hello Template Caching")
}

func TestTemplateWithoutCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateWithoutCaching(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func BenchmarkWithoutCaching(b *testing.B) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()
	for i := 0; i < b.N; i++ {
		TemplateWithoutCaching(recorder, request)
	}
}
