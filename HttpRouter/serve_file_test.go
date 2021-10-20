package httprouter

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed source
var source embed.FS

func TestServeFileHttp(t *testing.T) {

	router := httprouter.New()

	directory, _ := fs.Sub(source, "source")
	router.ServeFiles("/files/*filepath", http.FS(directory))

	request := httptest.NewRequest("GET", "http://localhost:8080/files/Hello.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hello HttpRouter", string(body))
}

func TestServeFileHttpGoodBye(t *testing.T) {

	router := httprouter.New()

	directory, _ := fs.Sub(source, "source")
	router.ServeFiles("/files/*filepath", http.FS(directory))

	request := httptest.NewRequest("GET", "http://localhost:8080/files/GoodBye.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "GoodBye HttpRouter", string(body))
}
