package golang_web

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServe(t *testing.T) {
	directory := http.Dir("./resources")
	fileServe := http.FileServer(directory)

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServe))

	serve := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := serve.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources
var resources embed.FS

func TestFileserveGolangEmbed(t *testing.T) {
	directory, _ := fs.Sub(resources, "resources")
	fileserve := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileserve))

	serve := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := serve.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
