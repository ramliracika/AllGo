package Website

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {

	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	firstName := request.PostForm.Get("First_Name")
	lastName := request.PostForm.Get("Last_Name")

	fmt.Fprintf(writer, "Halo %s %s", firstName, lastName)

}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("First_Name=ABC&Last_Name=123")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	record := httptest.NewRecorder()

	FormPost(record, request)

	response := record.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}
