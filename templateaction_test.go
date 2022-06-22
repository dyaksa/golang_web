package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func HandlerIf(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	t.ExecuteTemplate(w, "if.gohtml", map[string]interface{}{
		"Name":  "Dyaksa",
		"Title": "S.Kom",
	})
}

func TestIfTemplateAction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	HandlerIf(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func CompareTemplate(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/compare.gohtml"))
	t.ExecuteTemplate(w, "compare.gohtml", map[string]interface{}{
		"Title":      "Compare",
		"FinalValue": 80,
	})
}

func TestCompareTemplate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	CompareTemplate(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))
	t.ExecuteTemplate(w, "range.gohtml", map[string]interface{}{
		"Title": "range",
		"Hobbies": []string{
			"Game", "Learning", "Reading",
		},
	})
}

func TestTemplateRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateRange(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/with.gohtml"))
	t.ExecuteTemplate(w, "with.gohtml", map[string]interface{}{
		"Name":  "Dyaksa",
		"Title": "Backend Developer",
		"Address": map[string]interface{}{
			"Street": "Kenep RT 02 RW 03 Mangunjiwan Demak",
			"Number": 20,
		},
	})
}

func TestTemplateWith(t *testing.T) {
	response := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateWith(recorder, response)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
