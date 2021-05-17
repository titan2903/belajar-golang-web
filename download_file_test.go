package belajargolangweb

import (
	"fmt"
	"net/http"
	"testing"
)


func DownloadFile(writer http.ResponseWriter, request *http.Request) {
	file := request.URL.Query().Get("file")

	if file == "" { //! tidak ada file
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "Bad Request")
		return //! memaka berhenti agar tidak di eksekusi program setelahnya
	}


	writer.Header().Add("Content-Disposition", "attachment; filename=\""+file+"\"") //! langsung di minta donwload dan tidak di render
	http.ServeFile(writer, request, "./resources" + file)
}

func TestDownloadFile(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(DownloadFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}