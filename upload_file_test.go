package belajargolangweb

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)


func UploadFormFile(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "upload.form.gohtml", nil)
}

func Upload(writer http.ResponseWriter, request *http.Request) {
	// request.ParseMultipartForm(100 << 20) //! 100 MB max size file
	file, fileHeader, err := request.FormFile("file") //! mnegambil file dari input datanya
	if err != nil {
		panic(err)
	}

	fileDestination, err := os.Create("./resources" + fileHeader.Filename) //! simpan file ke folder yang di tuju
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(fileDestination, file) //! simpan semua file ke destination filenya
	if err != nil {
		panic(err)
	}

	name := request.PostFormValue("name") //! mengambil yang bukan file
	myTemplates.ExecuteTemplate(writer, "upload.success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestUploadFormFile(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadFormFile)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static",http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

//go:embed resources/evomo.png
var uploadFiletTest []byte

func TestUploadFile(t *testing.T) {
	body := new(bytes.Buffer) //! body kosong

	// ! input data dalam body
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "titanio yudista") //! input field text
	file, _ := writer.CreateFormFile("file", "CONTOHUPLAOD.png") //! input file
	file.Write(uploadFiletTest)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	result := recorder.Result().Body
	bodyResponse, _ := io.ReadAll(result)

	fmt.Println(string(bodyResponse))
}
