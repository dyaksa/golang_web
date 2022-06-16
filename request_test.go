package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)
		fmt.Fprint(w, r.URL)
	}

	serve := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := serve.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
