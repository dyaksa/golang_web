package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	firstname := r.PostForm.Get("firstname")
	lastname := r.PostForm.Get("lastname")

	fmt.Fprintf(w, "%s %s", firstname, lastname)

}

func TestFormPost(t *testing.T) {
	requestbBody := strings.NewReader("firstname=dyaksa&lastname=jauharuddin")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestbBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	strBody := string(body)
	fmt.Println(strBody)
}
