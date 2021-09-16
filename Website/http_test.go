package Website

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)


func Handler(writer http.ResponseWriter, request *http.Request)  {
	fmt.Fprint(writer,"Hello World")
}

func TestHTTP(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet,"localhost:8080",nil)
	recorder := httptest.NewRecorder()

	Handler(recorder,request)

	response := recorder.Result()
	body,_ := io.ReadAll(recorder.Body)
	ValueBody := string(body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(ValueBody)

}