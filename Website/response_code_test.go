package Website

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "Data is Empty")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestInvalid(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", nil)
	record := httptest.NewRecorder()

	ResponseCode(record, request)

	response := record.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}

func TestValid(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080?name=racika", nil)
	record := httptest.NewRecorder()

	ResponseCode(record, request)

	response := record.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}
