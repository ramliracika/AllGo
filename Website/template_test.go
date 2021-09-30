package Website

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func SimpleTemplates(writer http.ResponseWriter, requst *http.Request) {
	simpleText := `<html><body>{{.}}</body></html>`
	t, err := template.New("Simple").Parse(simpleText)
	if err != nil {
		panic(err)
	}

	// t := template.Must(template.New("Simple").Parse(simpleText)) <- simple code

	t.ExecuteTemplate(writer, "Simple", "Hello template HTML")
}

func TestSimpleTemplates(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleTemplates(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateFiles(writer http.ResponseWriter, request *http.Request) {

	t := template.Must(template.ParseFiles("./templates/templates.gohtml"))

	t.ExecuteTemplate(writer, "templates.gohtml", "Hello template HTML File")
}

func TestTemplatesFiles(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFiles(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplatesDirectory(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))

	t.ExecuteTemplate(writer, "templates.gohtml", "Hello templates directory")
}

func TestTemplatesDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplatesDirectory(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

//go:embed templates/templates.gohtml
var templates embed.FS

func TemplatesEmbed(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))

	t.ExecuteTemplate(writer, "templates.gohtml", "Hello templates Embed")
}

func TestTemplatesEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplatesEmbed(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
