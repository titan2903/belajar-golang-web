package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)


func SayHallo(rw http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		fmt.Fprint(rw, "Hello")
	} else {
		fmt.Fprintf(rw, "Hello %s", name)
	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=titan", nil)
	recorder := httptest.NewRecorder()

	SayHallo(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

/*
! Fprint mendukung formatter
! menangkap query parameter di golang web
*/

func MultipleQueryParameter(rw http.ResponseWriter, r *http.Request) { //! tidak ada batasannya
	first_name := r.URL.Query().Get("first_name")
	last_name := r.URL.Query().Get("last_name")

	fmt.Fprintf(rw, "Hello %s %s", first_name, last_name)
}

func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?first_name=titanio&last_name=yudista", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleValueQueryParameter(writer http.ResponseWriter, request *http.Request) { //! mengambil semua values menggunakan query of slice
	query := request.URL.Query()
	names := query["name"]

	fmt.Fprint(writer, strings.Join(names, " "))
}


func TestMultipleValueQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=titanio&name=yudista&name=eko", nil)
	recorder := httptest.NewRecorder()

	MultipleValueQueryParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
