package belajargolangweb

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)


func ServeFile(writer http.ResponseWriter, request *http.Request) { //! Jika ingin lebih dinamis
	if request.URL.Query().Get("name") != "" {
		http.ServeFile(writer, request, "./resources/ok.html")
	} else {
		http.ServeFile(writer, request, "./resources/notFound.html")
	}
}

func TestServeFileServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err == nil {
		panic(err)
	}
}

//go:embed resources/ok.html
var resourceOk string

//go:embed resources/notFound.html
var resourceNotFound string

func ServeFileGoEmbed(writer http.ResponseWriter, request *http.Request) { //! Jika ingin lebih dinamis dan lebih sederhana lagi dan tidak perlu lagi menentukan file name ada di mana addressnya
	if request.URL.Query().Get("name") != "" {
		fmt.Fprint(writer, resourceOk)
	} else {
		fmt.Fprint(writer, resourceNotFound)
	}
}

func TestServeFileServerEmbed(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServeFileGoEmbed),
	}

	err := server.ListenAndServe()
	if err == nil {
		panic(err)
	}
}