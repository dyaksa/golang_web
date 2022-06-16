package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("content-type")

	fmt.Fprint(w, header)
}

func TestHeaderRequest(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	res, _ := io.ReadAll(response.Body)
	resString := string(res)
	fmt.Println(resString)
}

func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Powered-By", "Dyaksa Jauharuddin")
	fmt.Fprint(w, "hello")
}

func TestHeaderResponse(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	response := recorder.Result()
	res, _ := io.ReadAll(response.Body)
	resString := string(res)
	fmt.Println(resString)

	fmt.Println(response.Header.Get("x-powered-by"))
}
