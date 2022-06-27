package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateEscape(w http.ResponseWriter, r *http.Request) {
	myTemplate.ExecuteTemplate(w, "escape", map[string]interface{}{
		"Title": "xss scripting",
		"Name":  "Dyaksa Jauhruddin",
		"Body":  template.HTMLEscapeString("<p>Hello World</p>"),
	})
}

func TestTemplateEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	response := httptest.NewRecorder()

	TemplateEscape(response, request)

	body, _ := io.ReadAll(response.Result().Body)
	fmt.Println(string(body))
}

func TestRunningTemplateEscape(t *testing.T) {
	serve := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateEscape),
	}

	err := serve.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TemplateEscapeDisable(w http.ResponseWriter, r *http.Request) {
	myTemplate.ExecuteTemplate(w, "escape", map[string]interface{}{
		"Title": "XSS",
		"Name":  "Dyaksa",
	})
}
