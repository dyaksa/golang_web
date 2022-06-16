package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestHello(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hi?name=dyaksa", nil)
	recorder := httptest.NewRecorder()

	sayHello(recorder, request)

	response := recorder.Result()
	result, _ := io.ReadAll(response.Body)

	strBody := string(result)

	fmt.Println(strBody)
}

func queryMultiple(w http.ResponseWriter, r *http.Request) {
	names := r.URL.Query()
	query := names["firstname"]
	fmt.Println(strings.Join(query, ""))
	firstName := r.URL.Query().Get("firstname")
	lastName := r.URL.Query().Get("lastname")

	if firstName == "" || lastName == "" {
		fmt.Fprint(w, "hello stranger")
	} else {
		fmt.Fprintf(w, "hello %s %s", firstName, lastName)
	}
}

func TestMultipleQuery(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?firstname=dyaksa&lastname=jauharuddin", nil)
	recorder := httptest.NewRecorder()

	queryMultiple(recorder, request)

	response := recorder.Result()

	result, _ := io.ReadAll(response.Body)
	strResult := string(result)
	fmt.Println(strResult)
}
