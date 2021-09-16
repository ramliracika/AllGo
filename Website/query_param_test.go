package Website

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func QueryParam(writer http.ResponseWriter, request *http.Request)  {
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(writer,"Hello")
	} else {
		fmt.Fprintf(writer,"Hello %s",name)
	}
}

func TestQueryParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet,"localhost:8080/hello?name=Racika",nil)
	recorder := httptest.NewRecorder()

	QueryParam(recorder,request)

	response := recorder.Result()
	body,_:= io.ReadAll(response.Body)
	ValueBody := string(body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(ValueBody)
}

func MultiQuery(writer http.ResponseWriter,request *http.Request)  {
	firstname := request.URL.Query().Get("first_name")
	lastname := request.URL.Query().Get("last_name")

	fmt.Fprintf(writer,"Hello %s %s" ,firstname,lastname)
}

func TestMultiQuery(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet,"localhost:8080/hello?first_name=racika&last_name=ramli",nil)
	recorder := httptest.NewRecorder()

	MultiQuery(recorder,request)

	response := recorder.Result()
	body,_:= io.ReadAll(response.Body)
	ValueBody := string(body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(ValueBody)
}

func MultiValue(writer http.ResponseWriter,request *http.Request)  {
	query := request.URL.Query()
	name := query["name"]

	fmt.Fprintln(writer,strings.Join(name,","))
}


func TestMultiValue(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet,"localhost:8080/hello?name=racika&name=ramli&name=haryadi&name=maulana",nil)
	recorder := httptest.NewRecorder()

	MultiValue(recorder,request)

	response := recorder.Result()
	body,_:= io.ReadAll(response.Body)
	ValueBody := string(body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(ValueBody)
}