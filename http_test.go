package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello world")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	res := recorder.Result()
	body, _ := io.ReadAll(res.Body) //! membaca semua di dalam response body
	bodyString := string(body) //! di convert menjadi string karena body berisi []byte

	fmt.Println(bodyString)
}

/*
! cara membuat unit test di golang web
*/