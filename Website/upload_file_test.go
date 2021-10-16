package Website

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadForm(w http.ResponseWriter, r *http.Request) {
	err := t.ExecuteTemplate(w, "upload.gohtml", nil)
	if err != nil {
		panic(err)
	}
}

func Upload(w http.ResponseWriter, r *http.Request) {
	// r.ParseMultipartForm(32 << 20) Default Memory 32MB

	file, fileHeader, err := r.FormFile("file")

	if err != nil {
		panic(err)
	}

	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}

	name := r.PostFormValue("name")
	t.ExecuteTemplate(w, "uploadSucces.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})

}

func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/form", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/4.png
var FileUpload []byte

func TestUnitUpload(t *testing.T) {
	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)
	writer.WriteField("name", "Ramli Racika")
	file, _ := writer.CreateFormFile("file", "upload2.png")
	file.Write(FileUpload)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}
