package views

import "net/http"

func loadUrls(){

	http.HandleFunc("/", index)
	http.HandleFunc("/login/", login)
	http.HandleFunc("/success/", success)

}



