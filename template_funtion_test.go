package belajargolangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)


type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My name is" + myPage.Name
}

func TemplateFucntion(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Eko"}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Titan",
	})
}

func TestTemplateFucntion(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFucntion(recorder, request)

	result := recorder.Result().Body
	body, _ := io.ReadAll(result)

	fmt.Println(string(body))
}

func TemplateFunctionGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Titan",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	result := recorder.Result().Body
	body, _ := io.ReadAll(result)

	fmt.Println(string(body))
}


func TemplateCreateFunctionGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"Upper": func (value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{ Upper .Name}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Titanio yudista",
	})
}

func TestTemplateCreateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateCreateFunctionGlobal(recorder, request)

	result := recorder.Result().Body
	body, _ := io.ReadAll(result)

	fmt.Println(string(body))
}


func TemplateCreateFunctionGlobalPipeline(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"SayHello" : func (name string) string {
			return "Hello " + name
		},
		"Upper": func (value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{ SayHello .Name | Upper}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Titanio yudista",
	})
}

func TestTemplateCreateFunctionGlobalPipeline(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateCreateFunctionGlobalPipeline(recorder, request)

	result := recorder.Result().Body
	body, _ := io.ReadAll(result)

	fmt.Println(string(body))
}