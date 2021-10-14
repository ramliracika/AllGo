package Website

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func AutoEscape(writer http.ResponseWriter, request *http.Request) {
	t.ExecuteTemplate(writer, "xss.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  "<p>This is a paragraf, so remind this </p>",
	})

}

func TestAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	AutoEscape(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}

func TestServerAutoEscape(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(AutoEscape),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func DisabledAutoEscape(writer http.ResponseWriter, request *http.Request) {
	t.ExecuteTemplate(writer, "xss.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape Disabled",
		"Body":  template.HTML("<h1>This will render ass html h1 too </h1>"),
	})
}

func TestDisabledAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recoder := httptest.NewRecorder()

	DisabledAutoEscape(recoder, request)

	body, _ := io.ReadAll(recoder.Result().Body)
	fmt.Println(string(body))
}

func TestServerDisabledAutoEscape(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(DisabledAutoEscape),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TemplateXSS(writer http.ResponseWriter, request *http.Request) {
	t.ExecuteTemplate(writer, "xss.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  template.HTML(request.URL.Query().Get("body")),
	})
}

func TestTemplateXSS(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?body=<p>alert</p>", nil)
	recorder := httptest.NewRecorder()

	TemplateXSS(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestTemplateXSSServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateXSS),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
