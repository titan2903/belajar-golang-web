package belajargolangweb

import (
	"fmt"
	"net/http"
	"testing"
)


type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before execute Handler")
	//! logic sebelum middleware
	middleware.Handler.ServeHTTP(writer, request) //! dibuat pointer agar tidak duplicte
	//! logic sebelum after
	fmt.Println("After execute Handler")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) { //! menagkap recover error
	defer func ()  {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi Error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error: %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(writer, request)
}


func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request){
		fmt.Println("Handler executed")
		fmt.Fprint(writer, "Hello handler")
	})

	mux.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request){
		fmt.Println("Foo Executed")
		// fmt.Fprint(writer, "Hello Foo")
		panic("ups")
	})

	logMiddleware := &LogMiddleware{
		Handler: mux,
	}

	errorHandler := &ErrorHandler{
		Handler: logMiddleware,
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}