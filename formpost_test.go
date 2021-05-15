package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)


func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm() //! secara manual parse terlebih dahulu
	if err != nil { //! check jika ada error
		panic(err)
	}

	// request.PostFormValue("first_name") //! secara otomatis tanpa melakukan parsing terlebih dahulu

	first_name := request.PostForm.Get("first_name")
	last_name := request.PostForm.Get("last_name")

	fmt.Fprintf(writer, "Hello %s %s", first_name, last_name)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=Titan&last_name=Yudista")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)

	request.Header.Add("content-type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}


