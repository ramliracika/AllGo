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

type LogMiddleware struct {
	http.Handler
}

func (middleware LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Menerima Request")
	middleware.Handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {

	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello World!")
	})

	Middleware := LogMiddleware{router}

	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recoder := httptest.NewRecorder()

	Middleware.ServeHTTP(recoder, request)

	response := recoder.Result()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Hello World!", string(body))

}
