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

func TestPanicHandler(t *testing.T) {

	router := httprouter.New()
	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, err interface{}) {
		fmt.Fprint(writer, "Panic : ", err)
	}
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		panic("Ups")
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)

	response := recoder.Result()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Panic : Ups", string(body))

}
