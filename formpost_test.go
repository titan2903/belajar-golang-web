package belajargolangweb

import (
	"fmt"
	"net/http"
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

}


