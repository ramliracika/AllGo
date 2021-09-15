package Website

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	Server := http.Server{
		Addr: "localhost:8080",
	}

	err := Server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
