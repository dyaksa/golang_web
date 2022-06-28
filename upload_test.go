package golang_web

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
	myTemplate.ExecuteTemplate(w, "upload_form.gohtml", nil)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(100 << 20) //upload max upload 100mb
	file, fileheader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	fileDestination, err := os.Create("./resources/" + fileheader.Filename)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}
	name := r.FormValue("name")
	myTemplate.ExecuteTemplate(w, "upload_success", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileheader.Filename,
	})
}

func TestUploadServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
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

//go:embed resources/ssnew.png
var uploadFileTest []byte

func TestUploadedFile(t *testing.T) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "Dyaksa Jauharuddin")

	file, _ := writer.CreateFormFile("file", "Newlogo.png")
	file.Write(uploadFileTest)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	response := httptest.NewRecorder()

	Upload(response, request)

	bodyResponse, _ := io.ReadAll(response.Result().Body)
	fmt.Println(string(bodyResponse))
}
