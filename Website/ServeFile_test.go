package Website

import (
	_ "embed"
	"fmt"
	"net/http"
	_ "net/http/httptest"
	"testing"
)

func ServeFile(writer http.ResponseWriter, request *http.Request) {

	if request.URL.Query().Get("Name") == "" {
		http.ServeFile(writer, request, "./resources/NotFound.html")
	} else {
		http.ServeFile(writer, request, "./resources/ok.html")
	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

//go:embed resources/ok.html
var oke string

//go:embed resources/NotFound.html
var notFound string

func ServeFileEmbed(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("Name") == "" {
		fmt.Fprint(writer, notFound)
	} else {
		fmt.Fprint(writer, oke)
	}
}

func TestServeFileEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
