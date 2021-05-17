package belajargolangweb

import (
	"io"
	"net/http"
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