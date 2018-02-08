package main

import (
	"net/http"
	"testform/views"
)


func main() {

	views.StartUp()

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	server.ListenAndServe()
}