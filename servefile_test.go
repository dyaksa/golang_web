package golang_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.ServeFile(w, r, "./recources/notfound.html")
	} else {
		http.ServeFile(w, r, "./resources/ok.html")
	}
}

func TestServeFile(t *testing.T) {
	serve := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := serve.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/oke.htmk
var ok string

//go::embed resources/notfound.html
var notfound string

func ServeFileEmbed(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, notfound)
	} else {
		fmt.Fprint(w, ok)
	}
}
