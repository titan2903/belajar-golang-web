package belajargolangweb

import (
	"fmt"
	"net/http"
	"testing"
)


func TestHandler(t *testing.T) { //! menggunakan handler function
	var handler http.HandlerFunc = func (wirter http.ResponseWriter, request *http.Request)  {
		//! Logic web
		fmt.Fprint(wirter, "Hello World")
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServeMux(t *testing.T) { //! ingin membuat multiple endpoint
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hellow world")
	})

	mux.HandleFunc("/hi", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hi")
	})

	mux.HandleFunc("/images/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Image")
	})

	mux.HandleFunc("/images/thumbnails", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Thumbnails")
	})

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/*
! jika endpoint duplicate maka akan ke replace
*/

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, r.Method)
		fmt.Fprint(rw, r.RequestURI)
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}