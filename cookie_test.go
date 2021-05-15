package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)


func SetCookie(writer http.ResponseWriter, request *http.Request) { //! menambahkan cookie ke response
	cookie := new(http.Cookie)
	cookie.Name = "X-PZN-Name"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/" //! menggunakan / agar cookie bisa di pakai di semua endpoint atau path

	http.SetCookie(writer, cookie) //! menyimpan cookie ata set cookie agar bisa di simpan di client. Bisa di set sebanyak mungkin yang di inginkan
	fmt.Fprint(writer, "Success crete cookie")
}

func GetCookie(writer http.ResponseWriter, request *http.Request) { //! membaca cookie dari request Http
	cookie, err := request.Cookie("X-PZN-Name")
	if err != nil {
		fmt.Fprint(writer, "No cookie")
	} else {
		name := cookie.Value
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) { //! test membuat cookie
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=titan", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies() //! balikannya []cookie

	for _, cookie := range cookies {
		fmt.Printf("Cookie %s: %s \n", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) { //! mengirim cookie ke server
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-PZN-Name"
	cookie.Value = "titan"
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()
	GetCookie(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)

	if err != nil {
		panic(err)
	}
	
	fmt.Println(string(body))
}