package Website

import (
	"fmt"
	"net/http"
	"testing"
)

type MiddleHandler struct {
	Handler http.Handler
}

func (Middleware *MiddleHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before Execute Handler")
	Middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After Execute Handler")
}

type HadlerError struct {
	Handler http.Handler
}

func (HandleE *HadlerError) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi Error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error : %s", err)
		}
	}()
	HandleE.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprint(writer, "Hello Middleware")
	})
	mux.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprint(writer, "Hello Foo")
	})
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		panic("ups")
	})

	logMiddleware := &MiddleHandler{
		Handler: mux,
	}

	errMiddleware := &HadlerError{
		Handler: logMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errMiddleware,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
