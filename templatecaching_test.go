package golang_web

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

//go:embed templates/*.gohtml
var loadTemplate embed.FS

var myTemplate = template.Must(template.ParseFS(loadTemplate, "templates/*.gohtml"))

func TemplateCaching(w http.ResponseWriter, r *http.Request) {
	myTemplate.ExecuteTemplate(w, "testing.gohtml", map[string]interface{}{
		"Name": "Dyaksa",
	})
}

func TestCachingTemplate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	response := httptest.NewRecorder()

	TemplateCaching(response, request)

	body, _ := io.ReadAll(response.Result().Body)
	fmt.Println(string(body))
}
