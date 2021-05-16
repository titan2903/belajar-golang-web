package belajargolangweb

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)


func TestFileServer(t *testing.T) { //! membuat static file server menggunakan golang
	directory := http.Dir("resources")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	
}

//go:embed resources
 var resources embed.FS

 func TestFileServerGoEmbed(t *testing.T) {
	directory, _ := fs.Sub(resources, "resources")
	fileServer := http.FileServer(http.FS(directory)) //! convert resource bawaan golang embed ke file sistem

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
 }