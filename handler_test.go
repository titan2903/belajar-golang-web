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