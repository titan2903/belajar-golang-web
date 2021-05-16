package belajargolangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)


func SimpleHTML(writer http.ResponseWriter, request *http.Request) { //! dynamic template menggunakan text
	templateText := `<html><body>{{.}}</body></html>`
	// t, err := template.New("SIMPLE").Parse(templateText) //!dengan checkiong err
	// if err != nil {
	// 	panic(err)
	// }

	tmp := template.Must(template.New("SIMPLE").Parse(templateText))

	tmp.ExecuteTemplate(writer, "SIMPLE", "HELLO HTML TEMPLATE") //! data "HELLO HTML TEMPLATE" akan di subsitusi ke double kurung kurawal . (titik)
}

func TestSimpleHTML(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)

	result := recorder.Result().Body
	body, _ := io.ReadAll(result)

	fmt.Println(string(body))
}


func SimpleHtmlFile(writer http.ResponseWriter, request *http.Request) { //! dynamic template menggunakan file
	t := template.Must(template.ParseFiles("./template/simple.gohtml"))
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello Go Template HTML File")
}

func TestSimpleHTMLFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHtmlFile(recorder, request)

	result := recorder.Result().Body
	body, _ := io.ReadAll(result)

	fmt.Println(string(body))
}

func SimpleHtmlTemplateDirectory(writer http.ResponseWriter, request *http.Request) { //! dynamic template menggunakan directory
	t := template.Must(template.ParseGlob("./template/*.gohtml"))
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello Go Template HTML File")
}

func TestSimpleHtmlTemplateDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHtmlTemplateDirectory(recorder, request)

	result := recorder.Result().Body
	body, _ := io.ReadAll(result)

	fmt.Println(string(body))
}

/*
! jika template tidak ketemu maka balikkan atau return nya berup string kosong
*/



func TemplateEmbed(writer http.ResponseWriter, request *http.Request) { //! dynamic template menggunakan directory
	t := template.Must(template.ParseFS(templates, "template/*.gohtml"))
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello Go Template HTML File")
}

func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	result := recorder.Result().Body
	body, _ := io.ReadAll(result)

	fmt.Println(string(body))
}
