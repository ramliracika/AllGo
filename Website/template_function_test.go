package Website

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (mypage MyPage) SayHello(Name string) string {
	return "Hello " + Name + ",My Name Is " + mypage.Name
}

func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("Function").Parse(`{{.SayHello "Budi"}}`))

	t.ExecuteTemplate(writer, "Function", MyPage{
		Name: "Ramli",
	})

}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("function").Parse(`{{len .Name}}`))

	t.ExecuteTemplate(writer, "function", MyPage{
		Name: "Ramli Racika",
	})

}

func TestTemplateGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}

func TemplateAddGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.New("function")
	t.Funcs(map[string]interface{}{
		"upper": func(Name string) string {
			return strings.ToUpper(Name)
		},
	})

	t = template.Must(t.Parse(`{{upper .Name}}`))

	t.ExecuteTemplate(writer, "function", MyPage{
		Name: "Ramli Racika",
	})
}

func TestTemplateAddGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateAddGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}

func TemplateGlobalPipeline(writer http.ResponseWriter, request *http.Request) {
	t := template.New("function")
	t.Funcs(map[string]interface{}{
		"Upper": func(Value string) string {
			return strings.ToUpper(Value)
		},

		"SayHello": func(Name string) string {
			return "Hello " + Name
		},
	})

	t = template.Must(t.Parse(`{{SayHello .Name | Upper}}`))

	t.ExecuteTemplate(writer, "function", MyPage{
		Name: "Ramli Racika",
	})
}

func TestTemplateGlobalPipeline(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateGlobalPipeline(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}
