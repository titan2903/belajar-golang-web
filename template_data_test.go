package belajargolangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TemplateDataMap(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./template/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", map[string]interface{}{
		"Title": "Template data map",
		"Name": "Titan",
		"Address": map[string]interface{}{
			"Street" : "Jalan Belum Ada",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	result := recorder.Result().Body
	body, _ := io.ReadAll(result)

	fmt.Println(string(body))
}


type Address struct {
	Street string
}
type Page struct {
	Title string
	Name string
	Address Address //! nested struct
}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./template/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", Page{
		Title: "Template Data Struct",
		Name: "Titan",
		Address: Address{
			Street: "Jalan Hijrah",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	result := recorder.Result().Body
	body, _ := io.ReadAll(result)

	fmt.Println(string(body))
}