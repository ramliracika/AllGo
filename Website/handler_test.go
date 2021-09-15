package Website

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World")
	})

	mux.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Home Page")
	})

	mux.HandleFunc("/login/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Login Page")
	})

	mux.HandleFunc("/login/register/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Register Page")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
