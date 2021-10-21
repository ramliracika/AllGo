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

func TestNotFoundHandler(t *testing.T) {

	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Yah Kamu Tersesat")
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)

	response := recoder.Result()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Yah Kamu Tersesat", string(body))

}
