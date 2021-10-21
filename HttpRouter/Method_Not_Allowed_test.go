package httprouter

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestMethodNotAllowed(t *testing.T) {

	router := httprouter.New()
	router.MethodNotAllowed = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Kamu Dilarang Ya")
	})

	router.POST("/", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(rw, "POST")
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)

	response := recoder.Result()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Kamu Dilarang Ya", string(body))

}
