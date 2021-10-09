package Website

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	x := template.Must(template.ParseGlob("./templates/*.gohtml"))

	x.ExecuteTemplate(writer, "layout", map[string]interface{}{
		"Title": "Templates Layout",
		"Name":  "Ramli",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
