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

func HandlerTemplate(w http.ResponseWriter, r *http.Request) {
	templateHTML := "<html><body>{{.}}</body></html>"
	// t, err := template.New("SIMPLE").Parse(templateHTML)
	// if err != nil {
	// 	panic(err)
	// }
	t := template.Must(template.New("SIMPLE").Parse(templateHTML))
	t.ExecuteTemplate(w, "SIMPLE", "Hello World")
}

func TestHandlerTemplate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	HandlerTemplate(recorder, request)

	body, _ := io.ReadAll(recorder.Body)
	fmt.Println(string(body))

}

func SimpleHTMLFile(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Hello World ini parse files template")
}

func TestSimpleHTML(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFile(recorder, request)

	body, _ := io.ReadAll(recorder.Body)
	fmt.Println(string(body))
}

func TemplateDirectory(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	t.ExecuteTemplate(w, "testing.gohtml", "Ini Test Directory Template")
}

func TestDirectoryTemplate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDirectory(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

//go:embed templates/*gohtml
var templates embed.FS

func TemplateEmbed(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(w, "simple.gohtml", "Hello Ini Template Embed")
}

func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
