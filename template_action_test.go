package belajargolangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TemplateActionIf(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./template/if.gohtml"))
	t.ExecuteTemplate(writer, "if.gohtml", Page{
		Title: "Template Data Struct",
		// Name:  "Titan",
	})
}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)

	result := recorder.Result().Body
	body, _ := io.ReadAll(result)

	fmt.Println(string(body))
}