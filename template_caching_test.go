package belajargolangweb

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

//go:embed template/*.gohtml
var templates embed.FS

var myTemplates = template.Must(template.ParseFS(templates, "template/*.gohtml")) //! di compile sekali di awal saja

func TemplateCaching(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "simple.gohtml", "hello template caching")
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	result := recorder.Result().Body
	body, _ := io.ReadAll(result)

	fmt.Println(string(body))
}
