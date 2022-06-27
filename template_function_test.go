package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + myPage.Name + "Good Morning"
}

func TemplateFunction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{ .SayHello "Dyaksa" }}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{
		Name: "Dyaksa Jauharuddin",
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

//global function golang

func TemplateFunctionGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{ len "Dyaksa"}} `))
	t.ExecuteTemplate(w, "FUNCTION", map[string]interface{}{
		"Name": "Dyaksa",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	response := httptest.NewRecorder()

	TemplateFunctionGlobal(response, request)

	body, _ := io.ReadAll(response.Result().Body)
	fmt.Println(string(body))
}

//create template global function
func TemplateFunctionCreateGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t.Parse(`{{ upper .Name }}`)

	t.ExecuteTemplate(w, "FUNCTION", map[string]interface{}{
		"Name": "Dyaksa Jauharuddin",
	})
}

func TestTemplateFunctionCreateGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	response := httptest.NewRecorder()

	TemplateFunctionCreateGlobal(response, request)

	body, _ := io.ReadAll(response.Result().Body)
	fmt.Println(string(body))
}

//template function pipelines
func TemplateFunctionPipeline(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	t.Funcs(map[string]interface{}{
		"sayHello": func(value string) string {
			return "Hello " + value
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t.Parse("{{ sayHello .Name | upper }}")

	t.ExecuteTemplate(w, "FUNCTION", map[string]interface{}{
		"Name": "Dyaksa Jauharuddin",
	})
}

func TestTemplateFunctionPipeline(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	response := httptest.NewRecorder()

	TemplateFunctionPipeline(response, request)

	body, _ := io.ReadAll(response.Result().Body)
	fmt.Println(string(body))
}
