package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("file")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Bad Request")
		return
	}
	w.Header().Add("Content-Dispotition", "attachment; filename=\""+name+"\"")
	http.ServeFile(w, r, "./resources/"+name)
}

func TestDownloadServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(DownloadFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
