package Website

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(writer http.ResponseWriter,request *http.Request) {
	ContentType := request.Header.Get("Content-Type")
	fmt.Fprint(writer,ContentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost,"http://localhost:8080/",nil)
	request.Header.Add("Content-Type","application/json")

	recorder := httptest.NewRecorder()

	RequestHeader(recorder,request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func ResponseHeader(writer http.ResponseWriter, request *http.Request)  {
	 writer.Header().Add("Powered-By","Racika Ramli")
	fmt.Fprint(writer,"OK")

}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet,"http://localhost:8080/",nil)
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder,request)

	PoweredBy := recorder.Header().Get("powered-by")
	fmt.Println(PoweredBy)
}